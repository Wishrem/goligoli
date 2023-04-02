package etcd

import (
	"context"
	"fmt"

	"github.com/wishrem/goligoli/pkg/conf"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
)

var etcdClient *clientv3.Client

func EtcdRegister(addr string, name string) error {
	etcdClient, err := clientv3.NewFromURL(conf.App.Etcd.URL)
	if err != nil {
		return err
	}

	em, err := endpoints.NewManager(etcdClient, name)
	if err != nil {
		return err
	}

	lease, _ := etcdClient.Grant(context.TODO(), 10)

	err = em.AddEndpoint(context.TODO(), fmt.Sprintf("%s/%s", name, addr), endpoints.Endpoint{Addr: addr}, clientv3.WithLease(lease.ID))
	if err != nil {
		return err
	}

	alive, err := etcdClient.KeepAlive(context.TODO(), lease.ID)
	if err != nil {
		return err
	}

	go func() {
		for {
			<-alive
		}
	}()

	return nil
}

func EtcdUnRegister(addr string, name string) error {
	if etcdClient != nil {
		em, err := endpoints.NewManager(etcdClient, name)
		if err != nil {
			return err
		}
		err = em.DeleteEndpoint(context.TODO(), fmt.Sprintf("%s/%s", name, addr))
		if err != nil {
			return err
		}
	}
	return nil
}
