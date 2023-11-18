/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type NodeConfigurationDefaultObservation struct {

	// CAST AI cluster id
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Id of the node configuration
	ConfigurationID *string `json:"configurationId,omitempty" tf:"configuration_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type NodeConfigurationDefaultParameters struct {

	// CAST AI cluster id
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/crossplane-provider-castai/apis/castai/v1alpha1.EksClusterId
	// +kubebuilder:validation:Optional
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a EksClusterId in castai to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a EksClusterId in castai to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// Id of the node configuration
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/crossplane-provider-castai/apis/castai/v1alpha1.NodeConfiguration
	// +kubebuilder:validation:Optional
	ConfigurationID *string `json:"configurationId,omitempty" tf:"configuration_id,omitempty"`

	// Reference to a NodeConfiguration in castai to populate configurationId.
	// +kubebuilder:validation:Optional
	ConfigurationIDRef *v1.Reference `json:"configurationIdRef,omitempty" tf:"-"`

	// Selector for a NodeConfiguration in castai to populate configurationId.
	// +kubebuilder:validation:Optional
	ConfigurationIDSelector *v1.Selector `json:"configurationIdSelector,omitempty" tf:"-"`
}

// NodeConfigurationDefaultSpec defines the desired state of NodeConfigurationDefault
type NodeConfigurationDefaultSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NodeConfigurationDefaultParameters `json:"forProvider"`
}

// NodeConfigurationDefaultStatus defines the observed state of NodeConfigurationDefault.
type NodeConfigurationDefaultStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NodeConfigurationDefaultObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// NodeConfigurationDefault is the Schema for the NodeConfigurationDefaults API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,castai}
type NodeConfigurationDefault struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NodeConfigurationDefaultSpec   `json:"spec"`
	Status            NodeConfigurationDefaultStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NodeConfigurationDefaultList contains a list of NodeConfigurationDefaults
type NodeConfigurationDefaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeConfigurationDefault `json:"items"`
}

// Repository type metadata.
var (
	NodeConfigurationDefault_Kind             = "NodeConfigurationDefault"
	NodeConfigurationDefault_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NodeConfigurationDefault_Kind}.String()
	NodeConfigurationDefault_KindAPIVersion   = NodeConfigurationDefault_Kind + "." + CRDGroupVersion.String()
	NodeConfigurationDefault_GroupVersionKind = CRDGroupVersion.WithKind(NodeConfigurationDefault_Kind)
)

func init() {
	SchemeBuilder.Register(&NodeConfigurationDefault{}, &NodeConfigurationDefaultList{})
}
