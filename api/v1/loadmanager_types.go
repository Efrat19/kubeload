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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type LoadSetup struct {
	MaxLoad     uint64          `json:"maxLoad"`
	InitialLoad uint64          `json:"initialLoad"`
	HatchRate   uint64          `json:"hatchRate"`
	Interval    metav1.Duration `json:"interval"`
}

// LoadManagerSpec defines the desired state of LoadManager
type LoadManagerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Selector  metav1.LabelSelector `json:"selector"`
	LoadSetup LoadSetup            `json:"loadSetup"`
}

// LoadManagerStatus defines the observed state of LoadManager
type LoadManagerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	CurrentLoad uint64 `json:"currentLoad"`
}

// +kubebuilder:object:root=true

// LoadManager is the Schema for the loadmanagers API
type LoadManager struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LoadManagerSpec   `json:"spec,omitempty"`
	Status LoadManagerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadManagerList contains a list of LoadManager
type LoadManagerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadManager `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LoadManager{}, &LoadManagerList{})
}
