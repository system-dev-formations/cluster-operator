// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/storageos/cluster-operator/pkg/apis/storageos/v1.Job":                    schema_pkg_apis_storageos_v1_Job(ref),
		"github.com/storageos/cluster-operator/pkg/apis/storageos/v1.JobSpec":                schema_pkg_apis_storageos_v1_JobSpec(ref),
		"github.com/storageos/cluster-operator/pkg/apis/storageos/v1.JobStatus":              schema_pkg_apis_storageos_v1_JobStatus(ref),
		"github.com/storageos/cluster-operator/pkg/apis/storageos/v1.NFSServer":              schema_pkg_apis_storageos_v1_NFSServer(ref),
		"github.com/storageos/cluster-operator/pkg/apis/storageos/v1.NFSServerSpec":          schema_pkg_apis_storageos_v1_NFSServerSpec(ref),
		"github.com/storageos/cluster-operator/pkg/apis/storageos/v1.NFSServerStatus":        schema_pkg_apis_storageos_v1_NFSServerStatus(ref),
		"github.com/storageos/cluster-operator/pkg/apis/storageos/v1.StorageOSCluster":       schema_pkg_apis_storageos_v1_StorageOSCluster(ref),
		"github.com/storageos/cluster-operator/pkg/apis/storageos/v1.StorageOSClusterSpec":   schema_pkg_apis_storageos_v1_StorageOSClusterSpec(ref),
		"github.com/storageos/cluster-operator/pkg/apis/storageos/v1.StorageOSClusterStatus": schema_pkg_apis_storageos_v1_StorageOSClusterStatus(ref),
		"github.com/storageos/cluster-operator/pkg/apis/storageos/v1.StorageOSUpgrade":       schema_pkg_apis_storageos_v1_StorageOSUpgrade(ref),
		"github.com/storageos/cluster-operator/pkg/apis/storageos/v1.StorageOSUpgradeSpec":   schema_pkg_apis_storageos_v1_StorageOSUpgradeSpec(ref),
		"github.com/storageos/cluster-operator/pkg/apis/storageos/v1.StorageOSUpgradeStatus": schema_pkg_apis_storageos_v1_StorageOSUpgradeStatus(ref),
	}
}

func schema_pkg_apis_storageos_v1_Job(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Job is the Schema for the jobs API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/storageos/cluster-operator/pkg/apis/storageos/v1.JobSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/storageos/cluster-operator/pkg/apis/storageos/v1.JobStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/storageos/cluster-operator/pkg/apis/storageos/v1.JobSpec", "github.com/storageos/cluster-operator/pkg/apis/storageos/v1.JobStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_storageos_v1_JobSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "JobSpec defines the desired state of Job",
				Properties: map[string]spec.Schema{
					"image": {
						SchemaProps: spec.SchemaProps{
							Description: "Image is the container image to run as the job.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"args": {
						SchemaProps: spec.SchemaProps{
							Description: "Args is an array of strings passed as an argument to the job container.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"mountPath": {
						SchemaProps: spec.SchemaProps{
							Description: "MountPath is the path in the job container where a volume is mounted.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"hostPath": {
						SchemaProps: spec.SchemaProps{
							Description: "HostPath is the path in the host that's mounted into a job container.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"completionWord": {
						SchemaProps: spec.SchemaProps{
							Description: "CompletionWord is the word that's looked for in the pod logs to find out if a DaemonSet Pod has completed its task.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"labelSelector": {
						SchemaProps: spec.SchemaProps{
							Description: "LabelSelector is the label selector for the job Pods.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"nodeSelectorTerms": {
						SchemaProps: spec.SchemaProps{
							Description: "NodeSelectorTerms is the set of placement of the job pods using node affinity requiredDuringSchedulingIgnoredDuringExecution.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.NodeSelectorTerm"),
									},
								},
							},
						},
					},
					"tolerations": {
						SchemaProps: spec.SchemaProps{
							Description: "Tolerations is to set the placement of storageos pods using pod toleration.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.Toleration"),
									},
								},
							},
						},
					},
				},
				Required: []string{"image", "args", "mountPath", "hostPath", "completionWord"},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.NodeSelectorTerm", "k8s.io/api/core/v1.Toleration"},
	}
}

func schema_pkg_apis_storageos_v1_JobStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "JobStatus defines the observed state of Job",
				Properties: map[string]spec.Schema{
					"completed": {
						SchemaProps: spec.SchemaProps{
							Description: "Completed indicates the complete status of job.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_storageos_v1_NFSServer(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NFSServer is the Schema for the nfsservers API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/storageos/cluster-operator/pkg/apis/storageos/v1.NFSServerSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/storageos/cluster-operator/pkg/apis/storageos/v1.NFSServerStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/storageos/cluster-operator/pkg/apis/storageos/v1.NFSServerSpec", "github.com/storageos/cluster-operator/pkg/apis/storageos/v1.NFSServerStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_storageos_v1_NFSServerSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NFSServerSpec defines the desired state of NFSServer",
				Properties: map[string]spec.Schema{
					"nfsContainer": {
						SchemaProps: spec.SchemaProps{
							Description: "NFSContainer is the container image to use for the NFS server.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"storageClassName": {
						SchemaProps: spec.SchemaProps{
							Description: "StorageClassName is the name of the StorageClass used by the NFS volume.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"tolerations": {
						SchemaProps: spec.SchemaProps{
							Description: "Tolerations is to set the placement of NFS server pods using pod toleration.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.Toleration"),
									},
								},
							},
						},
					},
					"resources": {
						SchemaProps: spec.SchemaProps{
							Description: "Resources represents the minimum resources required",
							Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
					"annotations": {
						SchemaProps: spec.SchemaProps{
							Description: "The annotations-related configuration to add/set on each Pod related object.",
							Type:        []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"export": {
						SchemaProps: spec.SchemaProps{
							Description: "The parameters to configure the NFS export",
							Ref:         ref("github.com/storageos/cluster-operator/pkg/apis/storageos/v1.ExportSpec"),
						},
					},
					"persistentVolumeReclaimPolicy": {
						SchemaProps: spec.SchemaProps{
							Description: "Reclamation policy for the persistent volume shared to the user's pod.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"mountOptions": {
						SchemaProps: spec.SchemaProps{
							Description: "PV mount options. Not validated - mount of the PVs will simply fail if one is invalid.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"persistentVolumeClaim": {
						SchemaProps: spec.SchemaProps{
							Description: "PersistentVolumeClaim is the PVC source of the PVC to be used with the NFS Server. If not specified, a new PVC is provisioned and used.",
							Ref:         ref("k8s.io/api/core/v1.PersistentVolumeClaimVolumeSource"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/storageos/cluster-operator/pkg/apis/storageos/v1.ExportSpec", "k8s.io/api/core/v1.PersistentVolumeClaimVolumeSource", "k8s.io/api/core/v1.ResourceRequirements", "k8s.io/api/core/v1.Toleration"},
	}
}

func schema_pkg_apis_storageos_v1_NFSServerStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "NFSServerStatus defines the observed state of NFSServer",
				Properties: map[string]spec.Schema{
					"remoteTarget": {
						SchemaProps: spec.SchemaProps{
							Description: "RemoteTarget is the connection string that clients can use to access the shared filesystem.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"phase": {
						SchemaProps: spec.SchemaProps{
							Description: "Phase is a simple, high-level summary of where the NFS Server is in its lifecycle. Phase will be set to Ready when the NFS Server is ready for use.  It is intended to be similar to the PodStatus Phase described at: https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.14/#podstatus-v1-core\n\nThere are five possible phase values:\n  - Pending: The NFS Server has been accepted by the Kubernetes system,\n    but one or more of the components has not been created. This includes\n    time before being scheduled as well as time spent downloading images\n    over the network, which could take a while.\n  - Running: The NFS Server has been bound to a node, and all of the\n    dependencies have been created.\n  - Succeeded: All NFS Server dependencies have terminated in success,\n    and will not be restarted.\n  - Failed: All NFS Server dependencies in the pod have terminated, and\n    at least one container has terminated in failure. The container\n    either exited with non-zero status or was terminated by the system.\n  - Unknown: For some reason the state of the NFS Server could not be\n    obtained, typically due to an error in communicating with the host of\n    the pod.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"accessModes": {
						SchemaProps: spec.SchemaProps{
							Description: "AccessModes is the access modes supported by the NFS server.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_storageos_v1_StorageOSCluster(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StorageOSCluster is the Schema for the storageosclusters API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/storageos/cluster-operator/pkg/apis/storageos/v1.StorageOSClusterSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/storageos/cluster-operator/pkg/apis/storageos/v1.StorageOSClusterStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/storageos/cluster-operator/pkg/apis/storageos/v1.StorageOSClusterSpec", "github.com/storageos/cluster-operator/pkg/apis/storageos/v1.StorageOSClusterStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_storageos_v1_StorageOSClusterSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StorageOSClusterSpec defines the desired state of StorageOSCluster",
				Properties: map[string]spec.Schema{
					"join": {
						SchemaProps: spec.SchemaProps{
							Description: "Join is the join token used for service discovery.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"csi": {
						SchemaProps: spec.SchemaProps{
							Description: "CSI defines the configurations for CSI.",
							Ref:         ref("github.com/storageos/cluster-operator/pkg/apis/storageos/v1.StorageOSClusterCSI"),
						},
					},
					"namespace": {
						SchemaProps: spec.SchemaProps{
							Description: "Namespace is the kubernetes Namespace where storageos resources are provisioned.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"storageClassName": {
						SchemaProps: spec.SchemaProps{
							Description: "StorageClassName is the name of default StorageClass created for StorageOS volumes.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"service": {
						SchemaProps: spec.SchemaProps{
							Description: "Service is the Service configuration for the cluster nodes.",
							Ref:         ref("github.com/storageos/cluster-operator/pkg/apis/storageos/v1.StorageOSClusterService"),
						},
					},
					"secretRefName": {
						SchemaProps: spec.SchemaProps{
							Description: "SecretRefName is the name of the secret object that contains all the sensitive cluster configurations.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"secretRefNamespace": {
						SchemaProps: spec.SchemaProps{
							Description: "SecretRefNamespace is the namespace of the secret reference.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"sharedDir": {
						SchemaProps: spec.SchemaProps{
							Description: "SharedDir is the shared directory to be used when the kubelet is running in a container. Typically: \"/var/lib/kubelet/plugins/kubernetes.io~storageos\". If not set, defaults will be used.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"ingress": {
						SchemaProps: spec.SchemaProps{
							Description: "Ingress defines the ingress configurations used in the cluster.",
							Ref:         ref("github.com/storageos/cluster-operator/pkg/apis/storageos/v1.StorageOSClusterIngress"),
						},
					},
					"images": {
						SchemaProps: spec.SchemaProps{
							Description: "Images defines the various container images used in the cluster.",
							Ref:         ref("github.com/storageos/cluster-operator/pkg/apis/storageos/v1.ContainerImages"),
						},
					},
					"kvBackend": {
						SchemaProps: spec.SchemaProps{
							Description: "KVBackend defines the key-value store backend used in the cluster.",
							Ref:         ref("github.com/storageos/cluster-operator/pkg/apis/storageos/v1.StorageOSClusterKVBackend"),
						},
					},
					"pause": {
						SchemaProps: spec.SchemaProps{
							Description: "Pause is to pause the operator for the cluster.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"debug": {
						SchemaProps: spec.SchemaProps{
							Description: "Debug is to set debug mode of the cluster.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"nodeSelectorTerms": {
						SchemaProps: spec.SchemaProps{
							Description: "NodeSelectorTerms is to set the placement of storageos pods using node affinity requiredDuringSchedulingIgnoredDuringExecution.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.NodeSelectorTerm"),
									},
								},
							},
						},
					},
					"tolerations": {
						SchemaProps: spec.SchemaProps{
							Description: "Tolerations is to set the placement of storageos pods using pod toleration.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.Toleration"),
									},
								},
							},
						},
					},
					"resources": {
						SchemaProps: spec.SchemaProps{
							Description: "Resources is to set the resource requirements of the storageos containers.",
							Ref:         ref("k8s.io/api/core/v1.ResourceRequirements"),
						},
					},
					"disableFencing": {
						SchemaProps: spec.SchemaProps{
							Description: "Disable Pod Fencing.  With StatefulSets, Pods are only re-scheduled if the Pod has been marked as killed.  In practice this means that failover of a StatefulSet pod is a manual operation.\n\nBy enabling Pod Fencing and setting the `storageos.com/fenced=true` label on a Pod, StorageOS will enable automated Pod failover (by killing the application Pod on the failed node) if the following conditions exist:\n\n- Pod fencing has not been explicitly disabled. - StorageOS has determined that the node the Pod is running on is\n  offline.  StorageOS uses Gossip and TCP checks and will retry for 30\n  seconds.  At this point all volumes on the failed node are marked\n  offline (irrespective of whether fencing is enabled) and volume\n  failover starts.\n- The Pod has the label `storageos.com/fenced=true` set. - The Pod has at least one StorageOS volume attached. - Each StorageOS volume has at least 1 healthy replica.\n\nWhen Pod Fencing is disabled, StorageOS will not perform any interaction with Kubernetes when it detects that a node has gone offline. Additionally, the Kubernetes permissions required for Fencing will not be added to the StorageOS role.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"disableTelemetry": {
						SchemaProps: spec.SchemaProps{
							Description: "Disable Telemetry.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"disableTCMU": {
						SchemaProps: spec.SchemaProps{
							Description: "Disable TCMU can be set to true to disable the TCMU storage driver.  This is required when there are multiple storage systems running on the same node and you wish to avoid conflicts.  Only one TCMU-based storage system can run on a node at a time.\n\nDisabling TCMU will degrade performance.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"forceTCMU": {
						SchemaProps: spec.SchemaProps{
							Description: "Force TCMU can be set to true to ensure that TCMU is enabled or cause StorageOS to abort startup.\n\nAt startup, StorageOS will automatically fallback to non-TCMU mode if another TCMU-based storage system is running on the node.  Since non-TCMU will degrade performance, this may not always be desired.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
					"tlsEtcdSecretRefName": {
						SchemaProps: spec.SchemaProps{
							Description: "TLSEtcdSecretRefName is the name of the secret object that contains the etcd TLS certs. This secret is shared with etcd, therefore it's not part of the main storageos secret.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"tlsEtcdSecretRefNamespace": {
						SchemaProps: spec.SchemaProps{
							Description: "TLSEtcdSecretRefNamespace is the namespace of the etcd TLS secret object.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"k8sDistro": {
						SchemaProps: spec.SchemaProps{
							Description: "K8sDistro is the name of the Kubernetes distribution where the operator is being deployed.  It should be in the format: `name[-1.0]`, where the version is optional and should only be appended if known.  Suitable names include: `openshift`, `rancher`, `aks`, `gke`, `eks`, or the deployment method if using upstream directly, e.g `minishift` or `kubeadm`.\n\nSetting k8sDistro is optional, and will be used to simplify cluster configuration by setting appropriate defaults for the distribution.  The distribution information will also be included in the product telemetry (if enabled), to help focus development efforts.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"disableScheduler": {
						SchemaProps: spec.SchemaProps{
							Description: "Disable StorageOS scheduler extender.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
				Required: []string{"secretRefName", "secretRefNamespace"},
			},
		},
		Dependencies: []string{
			"github.com/storageos/cluster-operator/pkg/apis/storageos/v1.ContainerImages", "github.com/storageos/cluster-operator/pkg/apis/storageos/v1.StorageOSClusterCSI", "github.com/storageos/cluster-operator/pkg/apis/storageos/v1.StorageOSClusterIngress", "github.com/storageos/cluster-operator/pkg/apis/storageos/v1.StorageOSClusterKVBackend", "github.com/storageos/cluster-operator/pkg/apis/storageos/v1.StorageOSClusterService", "k8s.io/api/core/v1.NodeSelectorTerm", "k8s.io/api/core/v1.ResourceRequirements", "k8s.io/api/core/v1.Toleration"},
	}
}

func schema_pkg_apis_storageos_v1_StorageOSClusterStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StorageOSClusterStatus defines the observed state of StorageOSCluster",
				Properties: map[string]spec.Schema{
					"phase": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"nodeHealthStatus": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"object"},
							AdditionalProperties: &spec.SchemaOrBool{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/storageos/cluster-operator/pkg/apis/storageos/v1.NodeHealth"),
									},
								},
							},
						},
					},
					"nodes": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
					"ready": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"members": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/storageos/cluster-operator/pkg/apis/storageos/v1.MembersStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/storageos/cluster-operator/pkg/apis/storageos/v1.MembersStatus", "github.com/storageos/cluster-operator/pkg/apis/storageos/v1.NodeHealth"},
	}
}

func schema_pkg_apis_storageos_v1_StorageOSUpgrade(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StorageOSUpgrade is the Schema for the storageosupgrades API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/storageos/cluster-operator/pkg/apis/storageos/v1.StorageOSUpgradeSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/storageos/cluster-operator/pkg/apis/storageos/v1.StorageOSUpgradeStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/storageos/cluster-operator/pkg/apis/storageos/v1.StorageOSUpgradeSpec", "github.com/storageos/cluster-operator/pkg/apis/storageos/v1.StorageOSUpgradeStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_storageos_v1_StorageOSUpgradeSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StorageOSUpgradeSpec defines the desired state of StorageOSUpgrade",
				Properties: map[string]spec.Schema{
					"newImage": {
						SchemaProps: spec.SchemaProps{
							Description: "NewImage is the new StorageOS node container image.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"newImage"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_storageos_v1_StorageOSUpgradeStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "StorageOSUpgradeStatus defines the observed state of StorageOSUpgrade",
				Properties: map[string]spec.Schema{
					"completed": {
						SchemaProps: spec.SchemaProps{
							Description: "Completed is the status of upgrade process.",
							Type:        []string{"boolean"},
							Format:      "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}
