/*
Copyright 2021-2022 the original author or authors.

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

package v1

import (
	corev1 "k8s.io/api/core/v1"
)

// +die
type _ = corev1.ObjectReference

// +die
type _ = corev1.LocalObjectReference

// +die
type _ = corev1.TypedLocalObjectReference

// +die
type _ = corev1.TypedObjectReference

// +die
type _ = corev1.SecretReference

// +die
type _ = corev1.TopologySelectorTerm

func (d *TopologySelectorTermDie) MatchLabelExpressionsDie(requirements ...*TopologySelectorLabelRequirementDie) *TopologySelectorTermDie {
	return d.DieStamp(func(r *corev1.TopologySelectorTerm) {
		r.MatchLabelExpressions = make([]corev1.TopologySelectorLabelRequirement, len(requirements))
		for i := range requirements {
			r.MatchLabelExpressions[i] = requirements[i].DieRelease()
		}
	})
}

// +die
type _ = corev1.TopologySelectorLabelRequirement
