package scheme

/**
This is the master scheme that will hold master copy of the configurations
This is what will be persisted to the etcd and will be transformed based on
which version is being used.
*/

type Pod struct {
	ApiVersion string    `yaml:"apiVerion"`
	Kind       string    `yaml:"Kind"`
	Metadata   string    `yaml:"metadata"`
	Spec       Spec      `yaml:"spec"`
	PodStatus  PodStatus `yaml:"status,omitempty"`
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
	Args    []string      `yaml:"args"`
	Command string        `yaml:"command"`
}

type ContainerState struct {
	Running    string `yaml:"running"`
	Terminated string `yaml:"terminated"`
	Waiting    string `yaml:"waiting"`
}

type ContainerPort struct {
	ContainerPort int    `yaml:"containerPort"`
	HostPort      int    `yaml:"hostPort"`
	Protocol      string `yaml:"protocol"`
	Name          string `yaml:"name"`
	HostIP        string `yaml:"hostIP"`
}
