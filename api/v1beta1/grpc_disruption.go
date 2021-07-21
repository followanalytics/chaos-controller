// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2021 Datadog, Inc.

package v1beta1

import (
	"errors"
	"fmt"
)

var (
	errorMap = map[string]int{
		"OK":                   0,
		"CANCELED_CODE":        1,
		"UNKNOWN_CODE":         2,
		"INVALID_ARGUMENT":     3,
		"DEADLINE_EXCEEDED":    4,
		"NOT_FOUND":            5,
		"ALREADY_EXISTS":       6,
		"PERMISSION_DENIED":    7,
		"RESOURCE_EXHAUSTED":   8,
		"FAILED_PRECONDITION":  9,
		"ABORTED_CODE":         10,
		"OUT_OF_RANGE":         11,
		"UNIMPLEMENTED_CODE":   12,
		"INTERNAL_CODE":        13,
		"UNAVAILABLE_CODE":     14,
		"DATALOSS_CODE":        15,
		"UNAUTHENTICATED_CODE": 16,
	}
)

// GRPCDisruptionSpec represents a gRPC disruption
type GRPCDisruptionSpec []EndpointAlteration

// EndpointAlteration represents an endpoint to disrupt and the corresponding error to return
type EndpointAlteration struct {
	TargetEndpoint string `json:"endpoint,omitempty"`
	// +kubebuilder:validation:Enum=OK;CANCELED_CODE;UNKNOWN_CODE;INVALID_ARGUMENT;DEADLINE_EXCEEDED;NOT_FOUND;ALREADY_EXISTS;PERMISSION_DENIED;RESOURCE_EXHAUSTED;FAILED_PRECONDITION;ABORTED_CODE;OUT_OF_RANGE;UNIMPLEMENTED_CODE;INTERNAL_CODE;UNAVAILABLE_CODE;DATALOSS_CODE;UNAUTHENTICATED_CODE
	ErrorToReturn string `json:"error,omitempty"`
	// +kubebuilder:validation:Enum={}
	OverrideToReturn string `json:"override,omitempty"`
}

// Validate validates that there are no missing hostnames or records for the given grpc disruption spec
func (s GRPCDisruptionSpec) Validate() error {

	if len(s) == 0 {
		return errors.New("the gRPC disruption was selected with no endpoints specified, but endpoints must be specified")
	}

	for _, pair := range s {
		if pair.TargetEndpoint == "" {
			return errors.New("some list items in gRPC disruption are missing endpoints; specify an endpoint for each item in the list")
		} else if (pair.ErrorToReturn != "" && pair.OverrideToReturn != "") || (pair.ErrorToReturn == "" && pair.OverrideToReturn == "") {
			return fmt.Errorf("the gRPC disruption can either return an error or override; specify exactly one for endpoint %s", pair.TargetEndpoint)
		}
	}
	return nil
}

// GenerateArgs generates injection pod arguments for the given spec
func (s GRPCDisruptionSpec) GenerateArgs() []string {
	args := []string{
		"grpc-disruption",
	}
	/*
		hostRecordPairArgs := []string{}

		for _, pair := range s {
			whiteSpaceCleanedIPList := strings.ReplaceAll(pair.Record.Value, " ", "")
			arg := fmt.Sprintf("%s;%s;%s", pair.Hostname, pair.Record.Type, whiteSpaceCleanedIPList)
			hostRecordPairArgs = append(hostRecordPairArgs, arg)
		}

		args = append(args, "--host-record-pairs")

		// Each value passed to --host-record-pairs should be of the form `hostname;type;value`, e.g.
		// `foo.bar.svc.cluster.local;A;10.0.0.0,10.0.0.13`
		args = append(args, strings.Split(strings.Join(hostRecordPairArgs, " --host-record-pairs "), " ")...)
	*/
	return args
}
