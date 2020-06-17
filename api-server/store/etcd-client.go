package store

import (
	"context"
	"errors"
	"fmt"

	"github.com/chauhanr/singlenetes/api-server/scheme"
	"github.com/chauhanr/singlenetes/api-server/util"
	"go.etcd.io/etcd/clientv3"
)

const (
	SUBSCRIBER_POD_KEY_PREFIX = "/consumer/pod"
	POD_KEY_PREFIX            = "/api/v1/Pod"
)

type EtcdCtl interface {
	GetEtcdClient() *clientv3.Client
	AddPod(scheme.PodV1) error
	AddSubscriber(scheme.EventSubscriber) error
	GetPodSubscribers(scheme.ComponentType) ([]scheme.EventSubscriber, error)
}

type EtcdCtlImpl struct {
	Client *clientv3.Client
}

func (e EtcdCtlImpl) GetEtcdClient() *clientv3.Client {
	return e.Client
}

func (e EtcdCtlImpl) AddPod(pod scheme.PodV1) error {
	v := pod.ApiVersion
	u := pod.Metadata.Uid
	kind := pod.Kind
	if pod.Metadata.Namespace == "" {
		pod.Metadata.Namespace = "default"
	}

	key := generatePodKey(kind, v, pod.Metadata.Namespace, u)
	podEncoded, err := util.EncodeS8Object(pod)
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

func (e EtcdCtlImpl) AddSubscriber(subs scheme.EventSubscriber) error {
	componentType := subs.Type
	name := subs.Name

	if name == "" || !componentType.IsValid() {
		return errors.New("Invalid Subsriber Information name missing or component invalid")
	}

	key := generateSubscriberKey(componentType, name)
	subsEncoded, err := util.EncodeS8Object(subs)
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

func (e EtcdCtlImpl) GetPodSubscribers(componentType scheme.ComponentType) ([]scheme.EventSubscriber, error) {
	key := fmt.Sprintf("%s/%s", SUBSCRIBER_POD_KEY_PREFIX, componentType)
	ctx := context.Background()
	gRes, err := e.Client.Get(ctx, key, clientv3.WithPrefix())
	subs := []scheme.EventSubscriber{}
	for _, kv := range gRes.Kvs {
		sub := scheme.EventSubscriber{}
		err = util.DecodeS8Object(kv.Value, &sub)
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
