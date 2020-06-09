package scheme

import (
	"errors"
	"time"
)

type PodV1 struct {
	ApiVersion string      `yaml:"apiVersion"`
	Kind       string      `yaml:"Kind"`
	Metadata   MetaV1      `yaml:"metadata"`
	Spec       PodSpecV1   `yaml:"spec"`
	Status     PodStatusV1 `yaml:"status,omitempty"`
}

func (p *PodV1) Validate() error {

	return nil
}

type PodSpecV1 struct {
	Containers []ContainerV1 `yaml:"containers"`
}

type PodStatusV1 struct {
	StartTime      time.Time         `yaml:"startTime"`
	Reason         string            `yaml:"reason"`
	Message        string            `yaml:"message"`
	ContainerStaus ContainerStatusV1 `yaml:"ContainerStatus"`
}

type ContainerStatusV1 struct {
	ContainerID string           `yaml:"containerID"`
	Image       string           `yaml:"image"`
	LastState   string           `yaml:"lastState"`
	Name        string           `yaml:"name"`
	Ready       bool             `yaml:"ready"`
	Started     bool             `yaml:"started"`
	State       ContainerStateV1 `yaml:"ContainerState"`
}

type ContainerStateV1 struct {
	Running    string `yaml:"running"`
	Terminated string `yaml:"terminated"`
	Waiting    string `yaml:"waiting"`
}

type ContainerV1 struct {
	Image   string          `yaml:"image"`
	Name    string          `yaml:"name"`
	Ports   ContainerPortV1 `yaml:"ports"`
	Args    []string        `yaml:"args"`
	Command string          `yaml:"command"`
}

type ContainerPortV1 struct {
	ContainerPort int    `yaml:"containerPort"`
	HostPort      int    `yaml:"hostPort"`
	Protocol      string `yaml:"protocol,omitempty"`
	Name          string `yaml:"name,omitempty"`
	HostIP        string `yaml:"hostIP,omitempty"`
}

type MetaV1 struct {
	Name         string    `yaml:"name"`
	Namespace    string    `yaml:"namespace"`
	CreationTime time.Time `yaml:"creationTime"`
	Uid          string    `yaml:"uid"`
}

/* Register Components Model*/

type ComponentType string

const (
	Schedular  ComponentType = "Scheduler"
	Kubelet    ComponentType = "Kubelet"
	Controller ComponentType = "Controller"
	Undefined  ComponentType = "Undefined"
)

func (c ComponentType) IsValid() bool {
	switch c {
	case Schedular, Kubelet, Controller:
		return true
	default:
		return false
	}
	return false
}

func GetComponentType(cmp string) (ComponentType, error) {
	switch cmp {
	case Schedular.String():
		return Schedular, nil
	case Kubelet.String():
		return Kubelet, nil
	case Controller.String():
		return Controller, nil
	default:
		return Undefined, errors.New("Component Not Defined")
	}
}

func (c ComponentType) String() string {
	return string(c)
}

type EventSubscriber struct {
	Name        string        `yaml:"name"`
	CallbackURL string        `yaml:"callbackURL"`
	Type        ComponentType `yaml:"type"`
}
