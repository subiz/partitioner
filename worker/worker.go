package worker

import (
	"context"
	"fmt"
	"github.com/subiz/errors"
	pb "github.com/subiz/header/partitioner"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"hash/fnv"
	"io"
	"strings"
	"sync"
	"time"
)

const (
	// represense blocked state of an partition mean...
	BLOCKED = -1
	NORMAL  = 1

	PartitionKey = "partitionkey"
	RedirectKey  = "partitionredirectkey"
)

// global hashing util, used to hash key to partition number
var ghash = fnv.New32a()

type partition struct {
	worker_id string
	state     int // NORMAL, BLOCKED
}

type Worker struct {
	*sync.Mutex
	id             string
	partitions     []partition
	version        string
	cluster        string
	preparing_term int32
	config         *pb.Configuration
	coor           pb.CoordinatorClient
	onUpdates      []func([]int32)
}

// NewWorker creates and starts a new Worker object
// parameters:
//  host: grpc host address for coordinator or other workers to connect
//    to this worker. E.g: web-1.web:8080
//  cluster: ID of cluster. E.g: web
//  id: ID of the worker, this must be unique for each workers inside
//    cluster. E.g: web-1
//  coordinator: host address of coordinator, used to listen changes
//    and making transaction. E.g: coordinator:8021
func NewWorker(cluster, id, coordinator string) *Worker {
	me := &Worker{Mutex: &sync.Mutex{}, version: "1.0.0", cluster: cluster, id: id}
	for {
		cconn, err := dialGrpc(coordinator)
		if err != nil {
			fmt.Printf("ERR 242SDJFDS while connecting %v. Retry in 2 secs\n", err)
			time.Sleep(2 * time.Second)
			continue
		}
		me.coor = pb.NewCoordinatorClient(cconn)
		break
	}
	me.fetchConfig()
	go me.rebalancePull()
	return me
}

// safe wraps f in side a func to prevent panicking
func safe(f func()) {
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println(errors.New(500, errors.E_unknown).Error())
			}
		}()
		f()
	}()
}

func (me *Worker) rebalancePull() {
	ctx := context.Background()
	for { // replace to rebalance
		safe(func() {
			stream, err := me.coor.Rebalance(ctx, &pb.WorkerID{Cluster: me.cluster, Id: me.id})
			if err != nil {
				fmt.Printf("ERR while joining cluster %v. Retry in 2 secs\n", err)
				time.Sleep(2 * time.Second)
				return
			}
			for {
				conf, err := stream.Recv()
				if err == io.EOF {
					break
				}

				if err != nil {
					fmt.Printf("ERR while rebalancing %v. Retry in 2 secs\n", err)
					time.Sleep(2 * time.Second)
					return
				}

				me.Lock()

				err = me.validateRequest(conf.GetVersion(), conf.GetCluster(), conf.GetTerm())
				if err != nil {
					me.coor.Deny(ctx, &pb.WorkerID{Cluster: me.cluster, Id: me.id, Term: conf.Term})
				} else {
					me.block(conf.GetPartitions())
					me.coor.Accept(ctx, &pb.WorkerID{Cluster: me.cluster, Id: me.id, Term: conf.Term})
				}
				me.fetchConfig()
				me.notifySubscribers()
				me.Unlock()
			}
		})
	}
}

// fetchConfig updates worker partition map to the current configuration of
// coordinator.
// Note: this function may take up to ten of seconds to execute since the
// cluster is in middle of rebalance process. During that time, coordinator
// would be blocking any requests.
func (me *Worker) fetchConfig() {
	var conf *pb.Configuration
	ctx := context.Background()
	fmt.Println("FETCHING CONFIG...")

	for {
		conf, err := me.coor.GetConfig(ctx, &pb.Cluster{Id: me.cluster})
		if err != nil {
			fmt.Printf("ERR #234FOISDOUD config %v. Retry after 2 secs\n", err)
			time.Sleep(2 * time.Second)
			continue
		}

		err = me.validateRequest(conf.GetVersion(), conf.GetCluster(), conf.GetTerm())
		if err != nil {
			fmt.Printf("ERR #234DD4FSDOUD OUTDATED %v. Retry after 2 secs\n", err)
			time.Sleep(2 * time.Second)
			continue
		}

		me.config = conf
		break
	}

	b, _ := me.config.MarshalJSON()
	fmt.Println("FETCHED.", string(b))
	// rebuild partition map using coordinator's configuration
	me.partitions = make([]partition, conf.GetTotalPartitions())
	for workerid, pars := range conf.GetPartitions() {
		for _, p := range pars.GetPartitions() {
			me.partitions[p] = partition{worker_id: workerid, state: NORMAL}
		}
	}
}

// OnUpdate registers a subscriber to workers.
// The subscriber get call synchronously when worker's partitions is changed
// worker will pass list of its new partitions through subscriber parameter
func (me *Worker) OnUpdate(f func([]int32)) {
	me.Lock()
	defer me.Unlock()
	me.onUpdates = append(me.onUpdates, f)
}

// notifySubscriber calls subscribed functions to notify them that worker's
// partitions has changed. This must done synchronously since order of
// changes might be critical to some subscribers.
func (me *Worker) notifySubscribers() {
	ourpars := make([]int32, 0)
	for p, par := range me.partitions {
		if par.worker_id == me.id {
			ourpars = append(ourpars, int32(p))
		}
	}
	wg := &sync.WaitGroup{}
	wg.Add(len(me.onUpdates))
	for _, f := range me.onUpdates {
		safe(func() {
			defer wg.Done()
			f(ourpars)
		})
	}
	wg.Wait()
}

// validateRequest makes sure version, cluster and term are upto date and
//   is sent to correct cluster
func (me *Worker) validateRequest(version, cluster string, term int32) error {
	if version != me.version {
		return errors.New(400, errors.E_invalid_partition_version,
			"only support version "+me.version)
	}

	if cluster != me.cluster {
		return errors.New(400, errors.E_invalid_partition_cluster,
			"cluster should be "+me.cluster+" not "+cluster)
	}

	if term < me.config.GetTerm() {
		return errors.New(400, errors.E_invalid_partition_term,
			"term should be %d, not %d", me.config.GetTerm(), term)
	}
	return nil
}

// GetConfig is a GRPC handler, it returns current configuration of the cluster
func (me *Worker) GetConfig(ctx context.Context, cluster *pb.Cluster) (*pb.Configuration, error) {
	me.Lock()
	for me.config != nil {
		me.Unlock()
		time.Sleep(1 * time.Second)
	}
	config := me.config
	me.Unlock()
	return config, nil
}

// block blocks outdated partitions
func (me *Worker) block(partitions map[string]*pb.WorkerPartitions) {
	for workerid, pars := range partitions {
		for _, p := range pars.GetPartitions() {
			if workerid != me.partitions[p].worker_id {
				me.partitions[p].state = BLOCKED
			}
		}
	}
}

// CreateIntercept makes a new GRPC server interceptor
func (me *Worker) CreateIntercept() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, in interface{}, sinfo *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (out interface{}, err error) {
		md, _ := metadata.FromIncomingContext(ctx)
		pkey := strings.Join(md[PartitionKey], "")

		if pkey == "" {
			return handler(ctx, in)
		}

		ghash.Write([]byte(pkey))
		parindex := ghash.Sum32() % uint32(len(me.partitions)) // 1024
		ghash.Reset()

		me.Lock()
		par := me.partitions[parindex]
		me.Unlock()

		// block
		for par.state == BLOCKED {
			time.Sleep(2 * time.Second)
			me.Lock()
			par = me.partitions[parindex]
			me.Unlock()
		}

		if par.worker_id == me.id { // correct parittion
			return handler(ctx, in)
		}

		header, _ := metadata.FromOutgoingContext(ctx)
		if header == nil {
			header = metadata.New(nil)
		}
		header.Append(RedirectKey, par.worker_id)
		grpc.SendHeader(ctx, header)
		return nil, errors.New(400, errors.E_wrong_partition_host,
			"should be %s, not %s", par.worker_id, me.id)
	}
}

// dialGrpc makes a GRPC connection to service
func dialGrpc(service string) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())
	opts = append(opts, grpc.WithTimeout(5*time.Second))
	return grpc.Dial(service, opts...)
}
