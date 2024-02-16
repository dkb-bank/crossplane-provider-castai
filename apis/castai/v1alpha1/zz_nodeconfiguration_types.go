// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type AksInitParameters struct {

	// (Number) Maximum number of pods that can be run on a node, which affects how many IP addresses you will need for each node. Defaults to 30
	// Maximum number of pods that can be run on a node, which affects how many IP addresses you will need for each node. Defaults to 30
	MaxPodsPerNode *float64 `json:"maxPodsPerNode,omitempty" tf:"max_pods_per_node,omitempty"`

	// ssd, premium-ssd (ultra and premium-ssd-v2 are not supported for os disk)
	// Type of managed os disk attached to the node. (See [disk types](https://learn.microsoft.com/en-us/azure/virtual-machines/disks-types)). One of: standard, standard-ssd, premium-ssd (ultra and premium-ssd-v2 are not supported for os disk)
	OsDiskType *string `json:"osDiskType,omitempty" tf:"os_disk_type,omitempty"`
}

type AksObservation struct {

	// (Number) Maximum number of pods that can be run on a node, which affects how many IP addresses you will need for each node. Defaults to 30
	// Maximum number of pods that can be run on a node, which affects how many IP addresses you will need for each node. Defaults to 30
	MaxPodsPerNode *float64 `json:"maxPodsPerNode,omitempty" tf:"max_pods_per_node,omitempty"`

	// ssd, premium-ssd (ultra and premium-ssd-v2 are not supported for os disk)
	// Type of managed os disk attached to the node. (See [disk types](https://learn.microsoft.com/en-us/azure/virtual-machines/disks-types)). One of: standard, standard-ssd, premium-ssd (ultra and premium-ssd-v2 are not supported for os disk)
	OsDiskType *string `json:"osDiskType,omitempty" tf:"os_disk_type,omitempty"`
}

type AksParameters struct {

	// (Number) Maximum number of pods that can be run on a node, which affects how many IP addresses you will need for each node. Defaults to 30
	// Maximum number of pods that can be run on a node, which affects how many IP addresses you will need for each node. Defaults to 30
	// +kubebuilder:validation:Optional
	MaxPodsPerNode *float64 `json:"maxPodsPerNode,omitempty" tf:"max_pods_per_node,omitempty"`

	// ssd, premium-ssd (ultra and premium-ssd-v2 are not supported for os disk)
	// Type of managed os disk attached to the node. (See [disk types](https://learn.microsoft.com/en-us/azure/virtual-machines/disks-types)). One of: standard, standard-ssd, premium-ssd (ultra and premium-ssd-v2 are not supported for os disk)
	// +kubebuilder:validation:Optional
	OsDiskType *string `json:"osDiskType,omitempty" tf:"os_disk_type,omitempty"`
}

type EksInitParameters struct {

	// (String) IP address to use for DNS queries within the cluster
	// IP address to use for DNS queries within the cluster
	DNSClusterIP *string `json:"dnsClusterIp,omitempty" tf:"dns_cluster_ip,omitempty"`

	// (Number) Allow configure the IMDSv2 hop limit, the default is 2
	// Allow configure the IMDSv2 hop limit, the default is 2
	ImdsHopLimit *float64 `json:"imdsHopLimit,omitempty" tf:"imds_hop_limit,omitempty"`

	// (Boolean) When the value is true both IMDSv1 and IMDSv2 are enabled. Setting the value to false disables permanently IMDSv1 and might affect legacy workloads running on the node created with this configuration. The default is true if the flag isn't provided
	// When the value is true both IMDSv1 and IMDSv2 are enabled. Setting the value to false disables permanently IMDSv1 and might affect legacy workloads running on the node created with this configuration. The default is true if the flag isn't provided
	ImdsV1 *bool `json:"imdsV1,omitempty" tf:"imds_v1,omitempty"`

	// (String) Cluster's instance profile ARN used for CAST provisioned nodes
	// Cluster's instance profile ARN used for CAST provisioned nodes
	InstanceProfileArn *string `json:"instanceProfileArn,omitempty" tf:"instance_profile_arn,omitempty"`

	// (String) AWS key pair ID to be used for CAST provisioned nodes. Has priority over ssh_public_key
	// AWS key pair ID to be used for CAST provisioned nodes. Has priority over ssh_public_key
	KeyPairID *string `json:"keyPairId,omitempty" tf:"key_pair_id,omitempty"`

	// (List of String) Cluster's security groups configuration for CAST provisioned nodes
	// Cluster's security groups configuration for CAST provisioned nodes
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// (Number) AWS EBS volume IOPS to be used for CAST provisioned nodes
	// AWS EBS volume IOPS to be used for CAST provisioned nodes
	VolumeIops *float64 `json:"volumeIops,omitempty" tf:"volume_iops,omitempty"`

	// (String) AWS KMS key ARN for encrypting EBS volume attached to the node
	// AWS KMS key ARN for encrypting EBS volume attached to the node
	VolumeKMSKeyArn *string `json:"volumeKmsKeyArn,omitempty" tf:"volume_kms_key_arn,omitempty"`

	// (Number) AWS EBS volume throughput in MiB/s to be used for CAST provisioned nodes
	// AWS EBS volume throughput in MiB/s to be used for CAST provisioned nodes
	VolumeThroughput *float64 `json:"volumeThroughput,omitempty" tf:"volume_throughput,omitempty"`

	// (String) AWS EBS volume type to be used for CAST provisioned nodes. One of: gp3, io1, io2
	// AWS EBS volume type to be used for CAST provisioned nodes. One of: gp3, io1, io2
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type EksObservation struct {

	// (String) IP address to use for DNS queries within the cluster
	// IP address to use for DNS queries within the cluster
	DNSClusterIP *string `json:"dnsClusterIp,omitempty" tf:"dns_cluster_ip,omitempty"`

	// (Number) Allow configure the IMDSv2 hop limit, the default is 2
	// Allow configure the IMDSv2 hop limit, the default is 2
	ImdsHopLimit *float64 `json:"imdsHopLimit,omitempty" tf:"imds_hop_limit,omitempty"`

	// (Boolean) When the value is true both IMDSv1 and IMDSv2 are enabled. Setting the value to false disables permanently IMDSv1 and might affect legacy workloads running on the node created with this configuration. The default is true if the flag isn't provided
	// When the value is true both IMDSv1 and IMDSv2 are enabled. Setting the value to false disables permanently IMDSv1 and might affect legacy workloads running on the node created with this configuration. The default is true if the flag isn't provided
	ImdsV1 *bool `json:"imdsV1,omitempty" tf:"imds_v1,omitempty"`

	// (String) Cluster's instance profile ARN used for CAST provisioned nodes
	// Cluster's instance profile ARN used for CAST provisioned nodes
	InstanceProfileArn *string `json:"instanceProfileArn,omitempty" tf:"instance_profile_arn,omitempty"`

	// (String) AWS key pair ID to be used for CAST provisioned nodes. Has priority over ssh_public_key
	// AWS key pair ID to be used for CAST provisioned nodes. Has priority over ssh_public_key
	KeyPairID *string `json:"keyPairId,omitempty" tf:"key_pair_id,omitempty"`

	// (List of String) Cluster's security groups configuration for CAST provisioned nodes
	// Cluster's security groups configuration for CAST provisioned nodes
	SecurityGroups []*string `json:"securityGroups,omitempty" tf:"security_groups,omitempty"`

	// (Number) AWS EBS volume IOPS to be used for CAST provisioned nodes
	// AWS EBS volume IOPS to be used for CAST provisioned nodes
	VolumeIops *float64 `json:"volumeIops,omitempty" tf:"volume_iops,omitempty"`

	// (String) AWS KMS key ARN for encrypting EBS volume attached to the node
	// AWS KMS key ARN for encrypting EBS volume attached to the node
	VolumeKMSKeyArn *string `json:"volumeKmsKeyArn,omitempty" tf:"volume_kms_key_arn,omitempty"`

	// (Number) AWS EBS volume throughput in MiB/s to be used for CAST provisioned nodes
	// AWS EBS volume throughput in MiB/s to be used for CAST provisioned nodes
	VolumeThroughput *float64 `json:"volumeThroughput,omitempty" tf:"volume_throughput,omitempty"`

	// (String) AWS EBS volume type to be used for CAST provisioned nodes. One of: gp3, io1, io2
	// AWS EBS volume type to be used for CAST provisioned nodes. One of: gp3, io1, io2
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type EksParameters struct {

	// (String) IP address to use for DNS queries within the cluster
	// IP address to use for DNS queries within the cluster
	// +kubebuilder:validation:Optional
	DNSClusterIP *string `json:"dnsClusterIp,omitempty" tf:"dns_cluster_ip,omitempty"`

	// (Number) Allow configure the IMDSv2 hop limit, the default is 2
	// Allow configure the IMDSv2 hop limit, the default is 2
	// +kubebuilder:validation:Optional
	ImdsHopLimit *float64 `json:"imdsHopLimit,omitempty" tf:"imds_hop_limit,omitempty"`

	// (Boolean) When the value is true both IMDSv1 and IMDSv2 are enabled. Setting the value to false disables permanently IMDSv1 and might affect legacy workloads running on the node created with this configuration. The default is true if the flag isn't provided
	// When the value is true both IMDSv1 and IMDSv2 are enabled. Setting the value to false disables permanently IMDSv1 and might affect legacy workloads running on the node created with this configuration. The default is true if the flag isn't provided
	// +kubebuilder:validation:Optional
	ImdsV1 *bool `json:"imdsV1,omitempty" tf:"imds_v1,omitempty"`

	// (String) Cluster's instance profile ARN used for CAST provisioned nodes
	// Cluster's instance profile ARN used for CAST provisioned nodes
	// +kubebuilder:validation:Optional
	InstanceProfileArn *string `json:"instanceProfileArn" tf:"instance_profile_arn,omitempty"`

	// (String) AWS key pair ID to be used for CAST provisioned nodes. Has priority over ssh_public_key
	// AWS key pair ID to be used for CAST provisioned nodes. Has priority over ssh_public_key
	// +kubebuilder:validation:Optional
	KeyPairID *string `json:"keyPairId,omitempty" tf:"key_pair_id,omitempty"`

	// (List of String) Cluster's security groups configuration for CAST provisioned nodes
	// Cluster's security groups configuration for CAST provisioned nodes
	// +kubebuilder:validation:Optional
	SecurityGroups []*string `json:"securityGroups" tf:"security_groups,omitempty"`

	// (Number) AWS EBS volume IOPS to be used for CAST provisioned nodes
	// AWS EBS volume IOPS to be used for CAST provisioned nodes
	// +kubebuilder:validation:Optional
	VolumeIops *float64 `json:"volumeIops,omitempty" tf:"volume_iops,omitempty"`

	// (String) AWS KMS key ARN for encrypting EBS volume attached to the node
	// AWS KMS key ARN for encrypting EBS volume attached to the node
	// +kubebuilder:validation:Optional
	VolumeKMSKeyArn *string `json:"volumeKmsKeyArn,omitempty" tf:"volume_kms_key_arn,omitempty"`

	// (Number) AWS EBS volume throughput in MiB/s to be used for CAST provisioned nodes
	// AWS EBS volume throughput in MiB/s to be used for CAST provisioned nodes
	// +kubebuilder:validation:Optional
	VolumeThroughput *float64 `json:"volumeThroughput,omitempty" tf:"volume_throughput,omitempty"`

	// (String) AWS EBS volume type to be used for CAST provisioned nodes. One of: gp3, io1, io2
	// AWS EBS volume type to be used for CAST provisioned nodes. One of: gp3, io1, io2
	// +kubebuilder:validation:Optional
	VolumeType *string `json:"volumeType,omitempty" tf:"volume_type,omitempty"`
}

type GkeInitParameters struct {

	// standard, pd-balanced, pd-ssd, pd-extreme
	// Type of boot disk attached to the node. (See [disk types](https://cloud.google.com/compute/docs/disks#pdspecs)). One of: pd-standard, pd-balanced, pd-ssd, pd-extreme
	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	// (Number) Maximum number of pods that can be run on a node, which affects how many IP addresses you will need for each node. Defaults to 30
	// Maximum number of pods that can be run on a node, which affects how many IP addresses you will need for each node. Defaults to 110
	MaxPodsPerNode *float64 `json:"maxPodsPerNode,omitempty" tf:"max_pods_per_node,omitempty"`

	// (List of String) Network tags to be added on a VM. (See network tags)
	// Network tags to be added on a VM. (See [network tags](https://cloud.google.com/vpc/docs/add-remove-network-tags))
	NetworkTags []*string `json:"networkTags,omitempty" tf:"network_tags,omitempty"`
}

type GkeObservation struct {

	// standard, pd-balanced, pd-ssd, pd-extreme
	// Type of boot disk attached to the node. (See [disk types](https://cloud.google.com/compute/docs/disks#pdspecs)). One of: pd-standard, pd-balanced, pd-ssd, pd-extreme
	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	// (Number) Maximum number of pods that can be run on a node, which affects how many IP addresses you will need for each node. Defaults to 30
	// Maximum number of pods that can be run on a node, which affects how many IP addresses you will need for each node. Defaults to 110
	MaxPodsPerNode *float64 `json:"maxPodsPerNode,omitempty" tf:"max_pods_per_node,omitempty"`

	// (List of String) Network tags to be added on a VM. (See network tags)
	// Network tags to be added on a VM. (See [network tags](https://cloud.google.com/vpc/docs/add-remove-network-tags))
	NetworkTags []*string `json:"networkTags,omitempty" tf:"network_tags,omitempty"`
}

type GkeParameters struct {

	// standard, pd-balanced, pd-ssd, pd-extreme
	// Type of boot disk attached to the node. (See [disk types](https://cloud.google.com/compute/docs/disks#pdspecs)). One of: pd-standard, pd-balanced, pd-ssd, pd-extreme
	// +kubebuilder:validation:Optional
	DiskType *string `json:"diskType,omitempty" tf:"disk_type,omitempty"`

	// (Number) Maximum number of pods that can be run on a node, which affects how many IP addresses you will need for each node. Defaults to 30
	// Maximum number of pods that can be run on a node, which affects how many IP addresses you will need for each node. Defaults to 110
	// +kubebuilder:validation:Optional
	MaxPodsPerNode *float64 `json:"maxPodsPerNode,omitempty" tf:"max_pods_per_node,omitempty"`

	// (List of String) Network tags to be added on a VM. (See network tags)
	// Network tags to be added on a VM. (See [network tags](https://cloud.google.com/vpc/docs/add-remove-network-tags))
	// +kubebuilder:validation:Optional
	NetworkTags []*string `json:"networkTags,omitempty" tf:"network_tags,omitempty"`
}

type KopsInitParameters struct {

	// (String) AWS key pair ID to be used for CAST provisioned nodes. Has priority over ssh_public_key
	// AWS key pair ID to be used for provisioned nodes. Has priority over sshPublicKey
	KeyPairID *string `json:"keyPairId,omitempty" tf:"key_pair_id,omitempty"`
}

type KopsObservation struct {

	// (String) AWS key pair ID to be used for CAST provisioned nodes. Has priority over ssh_public_key
	// AWS key pair ID to be used for provisioned nodes. Has priority over sshPublicKey
	KeyPairID *string `json:"keyPairId,omitempty" tf:"key_pair_id,omitempty"`
}

type KopsParameters struct {

	// (String) AWS key pair ID to be used for CAST provisioned nodes. Has priority over ssh_public_key
	// AWS key pair ID to be used for provisioned nodes. Has priority over sshPublicKey
	// +kubebuilder:validation:Optional
	KeyPairID *string `json:"keyPairId,omitempty" tf:"key_pair_id,omitempty"`
}

type NodeConfigurationInitParameters struct {

	// (Block List, Max: 1) (see below for nested schema)
	Aks []AksInitParameters `json:"aks,omitempty" tf:"aks,omitempty"`

	// (String) CAST AI cluster id
	// CAST AI cluster id
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/crossplane-provider-castai/apis/castai/v1alpha1.EksClusterId
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// Reference to a EksClusterId in castai to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDRef *v1.Reference `json:"clusterIdRef,omitempty" tf:"-"`

	// Selector for a EksClusterId in castai to populate clusterId.
	// +kubebuilder:validation:Optional
	ClusterIDSelector *v1.Selector `json:"clusterIdSelector,omitempty" tf:"-"`

	// (String) Optional container runtime to be used by kubelet. Applicable for EKS only.  Supported values include: dockerd, containerd
	// Optional container runtime to be used by kubelet. Applicable for EKS only.  Supported values include: `dockerd`, `containerd`
	ContainerRuntime *string `json:"containerRuntime,omitempty" tf:"container_runtime,omitempty"`

	// (Number) Disk to CPU ratio. Sets the number of GiBs to be added for every CPU on the node. Defaults to 0
	// Disk to CPU ratio. Sets the number of GiBs to be added for every CPU on the node. Defaults to 0
	DiskCPURatio *float64 `json:"diskCpuRatio,omitempty" tf:"disk_cpu_ratio,omitempty"`

	// (String) Optional docker daemon configuration properties in JSON format. Provide only properties that you want to override. Applicable for EKS only. Available values
	// Optional docker daemon configuration properties in JSON format. Provide only properties that you want to override. Applicable for EKS only. [Available values](https://docs.docker.com/engine/reference/commandline/dockerd/#daemon-configuration-file)
	DockerConfig *string `json:"dockerConfig,omitempty" tf:"docker_config,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	Eks []EksInitParameters `json:"eks,omitempty" tf:"eks,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	Gke []GkeInitParameters `json:"gke,omitempty" tf:"gke,omitempty"`

	// (String) Image to be used while provisioning the node. If nothing is provided will be resolved to latest available image based on Kubernetes version if possible
	// Image to be used while provisioning the node. If nothing is provided will be resolved to latest available image based on Kubernetes version if possible
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// (String) Init script to be run on your instance at launch. Should not contain any sensitive data. Value should be base64 encoded
	// Init script to be run on your instance at launch. Should not contain any sensitive data. Value should be base64 encoded
	InitScript *string `json:"initScript,omitempty" tf:"init_script,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	Kops []KopsInitParameters `json:"kops,omitempty" tf:"kops,omitempty"`

	// (String) Optional kubelet configuration properties in JSON format. Provide only properties that you want to override. Applicable for EKS only. Available values
	// Optional kubelet configuration properties in JSON format. Provide only properties that you want to override. Applicable for EKS only. [Available values](https://kubernetes.io/docs/reference/config-api/kubelet-config.v1beta1/)
	KubeletConfig *string `json:"kubeletConfig,omitempty" tf:"kubelet_config,omitempty"`

	// (Number) Minimal disk size in GiB. Defaults to 100, min 30, max 1000
	// Minimal disk size in GiB. Defaults to 100, min 30, max 1000
	MinDiskSize *float64 `json:"minDiskSize,omitempty" tf:"min_disk_size,omitempty"`

	// (String) Name of the node configuration
	// Name of the node configuration
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) SSH public key to be used for provisioned nodes
	// SSH public key to be used for provisioned nodes
	SSHPublicKey *string `json:"sshPublicKey,omitempty" tf:"ssh_public_key,omitempty"`

	// (List of String) Subnet ids to be used for provisioned nodes
	// Subnet ids to be used for provisioned nodes
	Subnets []*string `json:"subnets,omitempty" tf:"subnets,omitempty"`

	// (Map of String) Tags to be added on cloud instances for provisioned nodes
	// Tags to be added on cloud instances for provisioned nodes
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type NodeConfigurationObservation struct {

	// (Block List, Max: 1) (see below for nested schema)
	Aks []AksObservation `json:"aks,omitempty" tf:"aks,omitempty"`

	// (String) CAST AI cluster id
	// CAST AI cluster id
	ClusterID *string `json:"clusterId,omitempty" tf:"cluster_id,omitempty"`

	// (String) Optional container runtime to be used by kubelet. Applicable for EKS only.  Supported values include: dockerd, containerd
	// Optional container runtime to be used by kubelet. Applicable for EKS only.  Supported values include: `dockerd`, `containerd`
	ContainerRuntime *string `json:"containerRuntime,omitempty" tf:"container_runtime,omitempty"`

	// (Number) Disk to CPU ratio. Sets the number of GiBs to be added for every CPU on the node. Defaults to 0
	// Disk to CPU ratio. Sets the number of GiBs to be added for every CPU on the node. Defaults to 0
	DiskCPURatio *float64 `json:"diskCpuRatio,omitempty" tf:"disk_cpu_ratio,omitempty"`

	// (String) Optional docker daemon configuration properties in JSON format. Provide only properties that you want to override. Applicable for EKS only. Available values
	// Optional docker daemon configuration properties in JSON format. Provide only properties that you want to override. Applicable for EKS only. [Available values](https://docs.docker.com/engine/reference/commandline/dockerd/#daemon-configuration-file)
	DockerConfig *string `json:"dockerConfig,omitempty" tf:"docker_config,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	Eks []EksObservation `json:"eks,omitempty" tf:"eks,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	Gke []GkeObservation `json:"gke,omitempty" tf:"gke,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Image to be used while provisioning the node. If nothing is provided will be resolved to latest available image based on Kubernetes version if possible
	// Image to be used while provisioning the node. If nothing is provided will be resolved to latest available image based on Kubernetes version if possible
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// (String) Init script to be run on your instance at launch. Should not contain any sensitive data. Value should be base64 encoded
	// Init script to be run on your instance at launch. Should not contain any sensitive data. Value should be base64 encoded
	InitScript *string `json:"initScript,omitempty" tf:"init_script,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	Kops []KopsObservation `json:"kops,omitempty" tf:"kops,omitempty"`

	// (String) Optional kubelet configuration properties in JSON format. Provide only properties that you want to override. Applicable for EKS only. Available values
	// Optional kubelet configuration properties in JSON format. Provide only properties that you want to override. Applicable for EKS only. [Available values](https://kubernetes.io/docs/reference/config-api/kubelet-config.v1beta1/)
	KubeletConfig *string `json:"kubeletConfig,omitempty" tf:"kubelet_config,omitempty"`

	// (Number) Minimal disk size in GiB. Defaults to 100, min 30, max 1000
	// Minimal disk size in GiB. Defaults to 100, min 30, max 1000
	MinDiskSize *float64 `json:"minDiskSize,omitempty" tf:"min_disk_size,omitempty"`

	// (String) Name of the node configuration
	// Name of the node configuration
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) SSH public key to be used for provisioned nodes
	// SSH public key to be used for provisioned nodes
	SSHPublicKey *string `json:"sshPublicKey,omitempty" tf:"ssh_public_key,omitempty"`

	// (List of String) Subnet ids to be used for provisioned nodes
	// Subnet ids to be used for provisioned nodes
	Subnets []*string `json:"subnets,omitempty" tf:"subnets,omitempty"`

	// (Map of String) Tags to be added on cloud instances for provisioned nodes
	// Tags to be added on cloud instances for provisioned nodes
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type NodeConfigurationParameters struct {

	// (Block List, Max: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Aks []AksParameters `json:"aks,omitempty" tf:"aks,omitempty"`

	// (String) CAST AI cluster id
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

	// (String) Optional container runtime to be used by kubelet. Applicable for EKS only.  Supported values include: dockerd, containerd
	// Optional container runtime to be used by kubelet. Applicable for EKS only.  Supported values include: `dockerd`, `containerd`
	// +kubebuilder:validation:Optional
	ContainerRuntime *string `json:"containerRuntime,omitempty" tf:"container_runtime,omitempty"`

	// (Number) Disk to CPU ratio. Sets the number of GiBs to be added for every CPU on the node. Defaults to 0
	// Disk to CPU ratio. Sets the number of GiBs to be added for every CPU on the node. Defaults to 0
	// +kubebuilder:validation:Optional
	DiskCPURatio *float64 `json:"diskCpuRatio,omitempty" tf:"disk_cpu_ratio,omitempty"`

	// (String) Optional docker daemon configuration properties in JSON format. Provide only properties that you want to override. Applicable for EKS only. Available values
	// Optional docker daemon configuration properties in JSON format. Provide only properties that you want to override. Applicable for EKS only. [Available values](https://docs.docker.com/engine/reference/commandline/dockerd/#daemon-configuration-file)
	// +kubebuilder:validation:Optional
	DockerConfig *string `json:"dockerConfig,omitempty" tf:"docker_config,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Eks []EksParameters `json:"eks,omitempty" tf:"eks,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Gke []GkeParameters `json:"gke,omitempty" tf:"gke,omitempty"`

	// (String) Image to be used while provisioning the node. If nothing is provided will be resolved to latest available image based on Kubernetes version if possible
	// Image to be used while provisioning the node. If nothing is provided will be resolved to latest available image based on Kubernetes version if possible
	// +kubebuilder:validation:Optional
	Image *string `json:"image,omitempty" tf:"image,omitempty"`

	// (String) Init script to be run on your instance at launch. Should not contain any sensitive data. Value should be base64 encoded
	// Init script to be run on your instance at launch. Should not contain any sensitive data. Value should be base64 encoded
	// +kubebuilder:validation:Optional
	InitScript *string `json:"initScript,omitempty" tf:"init_script,omitempty"`

	// (Block List, Max: 1) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Kops []KopsParameters `json:"kops,omitempty" tf:"kops,omitempty"`

	// (String) Optional kubelet configuration properties in JSON format. Provide only properties that you want to override. Applicable for EKS only. Available values
	// Optional kubelet configuration properties in JSON format. Provide only properties that you want to override. Applicable for EKS only. [Available values](https://kubernetes.io/docs/reference/config-api/kubelet-config.v1beta1/)
	// +kubebuilder:validation:Optional
	KubeletConfig *string `json:"kubeletConfig,omitempty" tf:"kubelet_config,omitempty"`

	// (Number) Minimal disk size in GiB. Defaults to 100, min 30, max 1000
	// Minimal disk size in GiB. Defaults to 100, min 30, max 1000
	// +kubebuilder:validation:Optional
	MinDiskSize *float64 `json:"minDiskSize,omitempty" tf:"min_disk_size,omitempty"`

	// (String) Name of the node configuration
	// Name of the node configuration
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// (String) SSH public key to be used for provisioned nodes
	// SSH public key to be used for provisioned nodes
	// +kubebuilder:validation:Optional
	SSHPublicKey *string `json:"sshPublicKey,omitempty" tf:"ssh_public_key,omitempty"`

	// (List of String) Subnet ids to be used for provisioned nodes
	// Subnet ids to be used for provisioned nodes
	// +kubebuilder:validation:Optional
	Subnets []*string `json:"subnets,omitempty" tf:"subnets,omitempty"`

	// (Map of String) Tags to be added on cloud instances for provisioned nodes
	// Tags to be added on cloud instances for provisioned nodes
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// NodeConfigurationSpec defines the desired state of NodeConfiguration
type NodeConfigurationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     NodeConfigurationParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider NodeConfigurationInitParameters `json:"initProvider,omitempty"`
}

// NodeConfigurationStatus defines the observed state of NodeConfiguration.
type NodeConfigurationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        NodeConfigurationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// NodeConfiguration is the Schema for the NodeConfigurations API. Create node configuration for given cluster. Node configuration reference https://docs.cast.ai/docs/node-configuration
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,castai}
type NodeConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.subnets) || (has(self.initProvider) && has(self.initProvider.subnets))",message="spec.forProvider.subnets is a required parameter"
	Spec   NodeConfigurationSpec   `json:"spec"`
	Status NodeConfigurationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// NodeConfigurationList contains a list of NodeConfigurations
type NodeConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeConfiguration `json:"items"`
}

// Repository type metadata.
var (
	NodeConfiguration_Kind             = "NodeConfiguration"
	NodeConfiguration_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: NodeConfiguration_Kind}.String()
	NodeConfiguration_KindAPIVersion   = NodeConfiguration_Kind + "." + CRDGroupVersion.String()
	NodeConfiguration_GroupVersionKind = CRDGroupVersion.WithKind(NodeConfiguration_Kind)
)

func init() {
	SchemeBuilder.Register(&NodeConfiguration{}, &NodeConfigurationList{})
}
