package store

import (
	"bytes"
	"context"
	"errors"
	"fmt"

	"gopkg.in/yaml.v2"

	"github.com/chauhanr/singlenetes/api-server/scheme"
	"go.etcd.io/etcd/clientv3"
)

const (
	SUBSCRIBER_POD_KEY_PREFIX = "/consumer/pod"
	POD_KEY_PREFIX            = "/api/v1/Pod"
)

type EtcdCtl struct {
	Client *clientv3.Client
}

func (e *EtcdCtl) AddPod(pod scheme.PodV1) error {
	v := pod.ApiVersion
	u := pod.Metadata.Uid
	kind := pod.Kind
	if pod.Metadata.Namespace == "" {
		pod.Metadata.Namespace = "default"
	}

	key := generatePodKey(kind, v, pod.Metadata.Namespace, u)
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

/*
   This add subscriber method will to check for key duplication of the subscriber
   if a duplicate key is provided by the name of the component the key will be overriden
   So it is the responsibility of the sender to ensure that the key are not duplicate and
   if they are then operation / update is intended.
*/

func (e *EtcdCtl) AddSubscriber(subs scheme.EventSubscriber) error {
	componentType := subs.Type
	name := subs.Name

	if name == "" || !componentType.IsValid() {
		return errors.New("Invalid Subsriber Information name missing or component invalid")
	}

	key := generateSubscriberKey(componentType, name)
	subsEncoded, err := encodeS8Object(subs)
	if err != nil {
		return err
	}

	ctx := context.Background()
	_, err = e.Client.Put(ctx, key, subsEncoded)
	if err != nil {
		return err
	}

	return nil
}

func (e *EtcdCtl) GetPodSubscribers(componentType scheme.ComponentType) ([]scheme.EventSubscriber, error) {
	key := fmt.Sprintf("%s/%s", SUBSCRIBER_POD_KEY_PREFIX, componentType)
	ctx := context.Background()
	gRes, err := e.Client.Get(ctx, key, clientv3.WithPrefix())
	subs := []scheme.EventSubscriber{}
	for _, kv := range gRes.Kvs {
		sub := scheme.EventSubscriber{}
		err = decodeS8Object(kv.Value, &sub)
		if err != nil {
			return subs, err
		}
		subs = append(subs, sub)
	}
	return subs, nil
}

func generateSubscriberKey(componentType scheme.ComponentType, name string) string {
	key := fmt.Sprintf("%s/%s/%s", SUBSCRIBER_POD_KEY_PREFIX, componentType, name)
	return key
}

func generatePodKey(kind, version, namespace, uid string) string {
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

func decodeS8Object(value []byte, o interface{}) error {
	err := yaml.Unmarshal(value, o)
	if err != nil {
		return err
	}
	return nil
}
