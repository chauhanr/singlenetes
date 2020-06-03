package store

import (
	"fmt"

	"github.com/chauhanr/singlenetes/api-server/scheme"
	"go.etcd.io/etcd/clientv3"
)

type EtcdCtl struct {
	EtcdClient *clientv3.Client
}

func (e *EtcdCtl) Put(pod scheme.PodV1) {
	v := pod.ApiVersion
	u := pod.Metadata.Uid
	kind := pod.Kind

	key := generateKey(kind, v, u)
	fmt.Printf("Key: %s\n", key)
}

func generateKey(kind, version, uid string) string {
	key := fmt.Sprintf("/api/%s/%s/%s", version, kind, uid)
	return key
}

func encodeS8Object(v interface{}) string {
	return ""
}
