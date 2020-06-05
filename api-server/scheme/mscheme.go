package scheme

import "time"

/**
This is the master scheme that will hold master copy of the configurations
This is what will be persisted to the etcd and will be transformed based on
which version is being used.
*/

type Validator interface {
	Validate() error
}

type Pod struct {
	ApiVersion string    `yaml:"apiVersion"`
	Kind       string    `yaml:"Kind"`
	Metadata   Meta      `yaml:"metadata"`
	Spec       Spec      `yaml:"spec"`
	PodStatus  PodStatus `yaml:"status,omitempty"`
}

func (p *Pod) Validate() error {
	/*validate the pod master data*/
	return nil
}

type Spec struct {
	Containers []Container `yaml:"containers"`
}

type PodStatus struct {
	StartTime      time.Time         `yaml:"startTime"`
	Reason         string            `yaml:"reason"`
	Message        string            `yaml:"message"`
	ContainerStaus ContainerStatusV1 `yaml:"ContainerStatus"`
}

type ContainerStatus struct {
	ContainerID string         `yaml:"containerID"`
	Image       string         `yaml:"image"`
	LastState   string         `yaml:"lastState"`
	Name        string         `yaml:"name"`
	Ready       bool           `yaml:"ready"`
	Started     bool           `yaml:"started"`
	State       ContainerState `yaml:"ContainerState"`
}

type Container struct {
	Image   string        `yaml:"image"`
	Name    string        `yaml:"name"`
	Ports   ContainerPort `yaml:"ports"`
	Args    []string      `yaml:"args,omitempty"`
	Command string        `yaml:"command,omitempty"`
}

type ContainerState struct {
	Running    string `yaml:"running"`
	Terminated string `yaml:"terminated"`
	Waiting    string `yaml:"waiting"`
}

type ContainerPort struct {
	ContainerPort int    `yaml:"containerPort"`
	HostPort      int    `yaml:"hostPort"`
	Protocol      string `yaml:"protocol,omitempty"`
	Name          string `yaml:"name,omitempty"`
	HostIP        string `yaml:"hostIP,omitempty"`
}

type Meta struct {
	Name         string    `yaml:"name"`
	Namespace    string    `yaml:"namespace"`
	CreationTime time.Time `yaml:"creationTime"`
	Uid          string    `yaml:"uid"`
}

func (p *Pod) Transform(pod PodV1) {
	p.ApiVersion = pod.ApiVersion
	p.Kind = pod.Kind
	p.Spec = Spec{}
}
