package main

import (
//	"encoding/json"
		"github.com/subiz/errors"
	pb "github.com/subiz/partitioner/header"
	"testing"
//	"time"
)

func TestJoin(t *testing.T) {
	var cluster = "cluster1"
	db := NewDBMock()
	wc := NewWCMock()
	coor := NewCoordinator(cluster, db, wc)

	wc.RegisterWorker("worker1", func(conf *pb.Configuration) error {
		return nil
	})
	var err error
	err = coor.Join(&pb.WorkerRequest{
		Version: VERSION,
		Term:    0,
		Cluster: cluster,
		Id:      "worker1",
		Host:    "worker1:8080",
	})
	if err != nil {
		t.Fatal(err)
	}

	conf := coor.GetConfig()
	if conf.GetTerm() != 1 && conf.GetNextTerm() != 2 {
		t.Errorf("wrong term")
	}
	if len(conf.Workers) != 1 && conf.Workers["worker1"] == nil {
		t.Errorf("wrong worker")
	}
}

func TestJoinFailed(t *testing.T) {
	var cluster = "cluster1"
	db := NewDBMock()
	wc := NewWCMock()
	coor := NewCoordinator(cluster, db, wc)

	wc.RegisterWorker("worker1", func(conf *pb.Configuration) error {
		return errors.New(400, errors.E_error_from_partition_peer)
	})
	var err error

	err = coor.Join(&pb.WorkerRequest{
		Version: VERSION,
		Term:    0,
		Cluster: cluster,
		Id:      "worker1",
		Host:    "worker1:8080",
	})
	if err == nil {
		t.Fatal("shoule be error, got nil")
	}

	err = coor.Join(&pb.WorkerRequest{
		Version: VERSION,
		Term:    0,
		Cluster: cluster,
		Id:      "worker1",
		Host:    "worker1:8080",
	})
	if err == nil {
		t.Fatal("shoule be error, got nil")
	}

	conf := coor.GetConfig()
	if conf.GetTerm() != 0 && conf.GetNextTerm() != 1 {
		t.Errorf("wrong term")
	}
	if len(conf.Workers) != 0 {
		t.Errorf("wrong worker")
	}
}

type WCMock struct {
	workers map[string]func(*pb.Configuration) error
}

func NewWCMock() *WCMock {
	return &WCMock{
		workers:make(map[string]func(*pb.Configuration) error),
	}
}

func (me *WCMock) RegisterWorker(workerid string, f func(*pb.Configuration) error) {
	me.workers[workerid] = f
}
func (me *WCMock) Prepare(cluster, workerid string, conf *pb.Configuration) error {
	return me.workers[workerid](conf)
}