// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2021 Datadog, Inc.

package v1beta1

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	// FlowEgress is the string representation of network disruptions applied to outgoing packets
	FlowEgress = "egress"
	// FlowIngress is the string representation of network disruptions applied to incoming packets
	FlowIngress = "ingress"
)

// NetworkDisruptionSpec represents a network disruption injection
// +ddmark:validation:LinkedFields={Delay,DelayJitter}
type NetworkDisruptionSpec struct {
	// +nullable
	Hosts []NetworkDisruptionHostSpec `json:"hosts,omitempty"`
	// +nullable
	Services []NetworkDisruptionServiceSpec `json:"services,omitempty"`
	// +ddmark:validation:Enum=ingress;egress
	Flow string `json:"flow,omitempty"`
	// +ddmark:validation:Minimum=0
	// +ddmark:validation:Maximum=100
	Drop int `json:"drop,omitempty"`
	// +ddmark:validation:Minimum=0
	// +ddmark:validation:Maximum=100
	Duplicate int `json:"duplicate,omitempty"`
	// +ddmark:validation:Minimum=0
	// +ddmark:validation:Maximum=100
	Corrupt int `json:"corrupt,omitempty"`
	// +ddmark:validation:Minimum=0
	// +ddmark:validation:Maximum=59999
	Delay uint `json:"delay,omitempty"`
	// +ddmark:validation:Minimum=0
	// +ddmark:validation:Maximum=100
	DelayJitter uint `json:"delayJitter,omitempty"`
	// +ddmark:validation:Minimum=0
	BandwidthLimit int `json:"bandwidthLimit,omitempty"`
	// +ddmark:validation:Minimum=0
	// +ddmark:validation:Maximum=65535
	// +nullable
	DeprecatedPort *int `json:"port,omitempty"`
}

type NetworkDisruptionHostSpec struct {
	Host string `json:"host,omitempty"`
	// +ddmark:validation:Minimum=0
	// +ddmark:validation:Maximum=65535
	Port int `json:"port,omitempty"`
	// +ddmark:validation:Enum=tcp;udp;""
	Protocol string `json:"protocol,omitempty"`
}

type NetworkDisruptionServiceSpec struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

// Validate validates args for the given disruption
func (s *NetworkDisruptionSpec) Validate() error {
	// check that at least one network disruption is set
	if s.BandwidthLimit == 0 &&
		s.Drop == 0 &&
		s.Delay == 0 &&
		s.Corrupt == 0 &&
		s.Duplicate == 0 {
		return errors.New("the network disruption was selected, but no disruption type was specified. Please set at least one of: drop, delay, bandwidthLimit, corrupt, or duplicate. No injection will occur")
	}

	// ensure spec filters on something if ingress mode is enabled
	if s.Flow == FlowIngress {
		if len(s.Hosts) == 0 && len(s.Services) == 0 {
			return errors.New("the network disruption has ingress flow enabled but no hosts or services are provided, which is required for it to work")
		}
	}

	if k8sClient != nil {
		err := validateServices(k8sClient, s.Services)
		if err != nil {
			return err
		}
	}

	// ensure deprecated fields are not used
	if s.DeprecatedPort != nil {
		return fmt.Errorf("the port specification at the network disruption level is deprecated; apply to network disruption hosts instead")
	}

	return nil
}

// GenerateArgs generates injection or cleanup pod arguments for the given spec
func (s *NetworkDisruptionSpec) GenerateArgs() []string {
	args := []string{
		"network-disruption",
		"--corrupt",
		strconv.Itoa(s.Corrupt),
		"--drop",
		strconv.Itoa(s.Drop),
		"--duplicate",
		strconv.Itoa(s.Duplicate),
		"--delay",
		strconv.Itoa(int(s.Delay)),
		"--delay-jitter",
		strconv.Itoa(int(s.DelayJitter)),
		"--bandwidth-limit",
		strconv.Itoa(s.BandwidthLimit),
	}

	// append hosts
	for _, host := range s.Hosts {
		args = append(args, "--hosts", fmt.Sprintf("%s;%d;%s", host.Host, host.Port, host.Protocol))
	}

	// append services
	for _, service := range s.Services {
		args = append(args, "--services", fmt.Sprintf("%s;%s", service.Name, service.Namespace))
	}

	// append flow
	if s.Flow != "" {
		args = append(args, "--flow", s.Flow)
	}

	return args
}
