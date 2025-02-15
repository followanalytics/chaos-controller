// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2021 Datadog, Inc.

/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"errors"
	"fmt"

	chaosv1beta1 "github.com/DataDog/chaos-controller/api/v1beta1"
	chaostypes "github.com/DataDog/chaos-controller/types"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// RunningTargetSelector finds pods in Running Phase for applying network disruptions to a Kubernetes Cluster
type RunningTargetSelector struct{}

// GetMatchingPods returns a pods list containing all running pods matching the given label selector and namespace
func (r RunningTargetSelector) GetMatchingPods(c client.Client, instance *chaosv1beta1.Disruption) (*corev1.PodList, error) {
	// get parsed selector
	selector, err := getLabelSelectorFromInstance(instance)
	if err != nil {
		return nil, fmt.Errorf("error getting label selector from disruption: %w", err)
	}

	// filter pods based on the label selector and namespace
	pods := &corev1.PodList{}
	listOptions := &client.ListOptions{
		LabelSelector: selector,
		Namespace:     instance.Namespace,
	}

	// fetch pods from label selector
	if err := c.List(context.Background(), pods, listOptions); err != nil {
		return nil, err
	}

	runningPods := &corev1.PodList{}

	for _, pod := range pods.Items {
		// if the disruption is applied on init, we only target pending pods with a running
		// chaos handler init container
		// otherwise, we only target running pods
		if instance.Spec.OnInit {
			hasChaosHandler := false

			// search for a potential running chaos handler init container
			for _, initContainerStatus := range pod.Status.InitContainerStatuses {
				if initContainerStatus.Name == "chaos-handler" && initContainerStatus.State.Running != nil {
					hasChaosHandler = true

					break
				}
			}

			if pod.Status.Phase == corev1.PodPending && hasChaosHandler {
				runningPods.Items = append(runningPods.Items, pod)
			}
		} else if pod.Status.Phase == corev1.PodRunning {
			runningPods.Items = append(runningPods.Items, pod)
		}
	}

	return runningPods, nil
}

// GetMatchingNodes returns a nodes list containing all nodes matching the given label selector
func (r RunningTargetSelector) GetMatchingNodes(c client.Client, instance *chaosv1beta1.Disruption) (*corev1.NodeList, error) {
	// get parsed selector
	selector, err := getLabelSelectorFromInstance(instance)
	if err != nil {
		return nil, fmt.Errorf("error getting label selector from disruption: %w", err)
	}

	// filter nodes based on the label selector
	nodes := &corev1.NodeList{}
	listOptions := &client.ListOptions{
		LabelSelector: selector,
	}

	// fetch nodes from label selector
	if err := c.List(context.Background(), nodes, listOptions); err != nil {
		return nil, err
	}

	runningNodes := &corev1.NodeList{}

	for _, node := range nodes.Items {
		// check if node is ready
		ready := false

		for _, condition := range node.Status.Conditions {
			if condition.Type == corev1.NodeReady && condition.Status == corev1.ConditionTrue {
				ready = true
				break
			}
		}

		if ready {
			runningNodes.Items = append(runningNodes.Items, node)
		}
	}

	return runningNodes, nil
}

// TargetIsHealthy returns an error if the given target is unhealthy or does not exist
func (r RunningTargetSelector) TargetIsHealthy(target string, c client.Client, instance *chaosv1beta1.Disruption) error {
	switch instance.Spec.Level {
	case chaostypes.DisruptionLevelUnspecified, chaostypes.DisruptionLevelPod:
		var p corev1.Pod

		// check if target still exists
		if err := c.Get(context.Background(), types.NamespacedName{Name: target, Namespace: instance.Namespace}, &p); err != nil {
			return err
		}

		// check if pod is running
		if p.Status.Phase != corev1.PodRunning {
			return errors.New("pod is not Running")
		}
	case chaostypes.DisruptionLevelNode:
		var n corev1.Node
		if err := c.Get(context.Background(), client.ObjectKey{Name: target}, &n); err != nil {
			return err
		}

		// check if node is ready
		ready := false

		for _, condition := range n.Status.Conditions {
			if condition.Type == corev1.NodeReady && condition.Status == corev1.ConditionTrue {
				ready = true
				break
			}
		}

		if !ready {
			return errors.New("node is not Ready")
		}
	}

	return nil
}
