package store

import (
	"github.com/chauhanr/singlenetes/api-server/scheme"
	"github.com/stretchr/testify/mock"
	"go.etcd.io/etcd/clientv3"
)

type MockEtcdCtl struct {
	mock.Mock
}

func (m MockEtcdCtl) AddPod(pod scheme.PodV1) error {
	res := m.Called(pod)
	return res.Error(0)
}

func (m MockEtcdCtl) AddSubscriber(subs scheme.EventSubscriber) error {
	res := m.Called(subs)
	return res.Error(0)
}

func (m MockEtcdCtl) GetPodSubscribers(componentType scheme.ComponentType) ([]scheme.EventSubscriber, error) {
	res := m.Called(componentType)
	return res.Get(0).([]scheme.EventSubscriber), res.Error(1)
}

func (e MockEtcdCtl) GetEtcdClient() *clientv3.Client {
	return nil
}
