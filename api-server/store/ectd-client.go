package store

import "go.etcd.io/etcd/clientv3"

type EtcdCtl struct {
	EtcdClient *clientv3.Client
}

func (e *EtcdCtl) Get(key string) {

}

func (e *EtcdCtl) Set(key string, value string) {

}
