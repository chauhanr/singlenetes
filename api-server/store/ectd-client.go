package store

import (
	"bytes"
	"context"
	"fmt"

	"gopkg.in/yaml.v2"

	"github.com/chauhanr/singlenetes/api-server/scheme"
	"go.etcd.io/etcd/clientv3"
)

type EtcdCtl struct {
	Client *clientv3.Client
}

func (e *EtcdCtl) Put(pod scheme.PodV1) error {
	v := pod.ApiVersion
	u := pod.Metadata.Uid
	kind := pod.Kind
	if pod.Metadata.Namespace == "" {
		pod.Metadata.Namespace = "default"
	}

	key := generateKey(kind, v, pod.Metadata.Namespace, u)
	podEncoded, err := encodeS8Object(pod)
	if err != nil {
		return err
	}

	ctx := context.Background()
	_, err = e.Client.Put(ctx, key, podEncoded)

	if err != nil {
		return err
	}
	return nil
}

func generateKey(kind, version, namespace, uid string) string {
	key := fmt.Sprintf("/api/%s/%s/%s/%s", version, kind, namespace, uid)
	return key
}

func encodeS8Object(v interface{}) (string, error) {
	var buf bytes.Buffer

	enc := yaml.NewEncoder(&buf)
	err := enc.Encode(v)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
