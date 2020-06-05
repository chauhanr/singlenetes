package scheme

import "time"

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
