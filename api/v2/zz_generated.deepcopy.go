//go:build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v2

import (
	"encoding/json"
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackRestCondition) DeepCopyInto(out *BackRestCondition) {
	*out = *in
	if in.FailureCause != nil {
		in, out := &in.FailureCause, &out.FailureCause
		*out = make([]FailureCause, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackRestCondition.
func (in *BackRestCondition) DeepCopy() *BackRestCondition {
	if in == nil {
		return nil
	}
	out := new(BackRestCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackRestSidecar) DeepCopyInto(out *BackRestSidecar) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackRestSidecar.
func (in *BackRestSidecar) DeepCopy() *BackRestSidecar {
	if in == nil {
		return nil
	}
	out := new(BackRestSidecar)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackRestStatus) DeepCopyInto(out *BackRestStatus) {
	*out = *in
	if in.Condition != nil {
		in, out := &in.Condition, &out.Condition
		*out = new(BackRestCondition)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackRestStatus.
func (in *BackRestStatus) DeepCopy() *BackRestStatus {
	if in == nil {
		return nil
	}
	out := new(BackRestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraBackup) DeepCopyInto(out *CassandraBackup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraBackup.
func (in *CassandraBackup) DeepCopy() *CassandraBackup {
	if in == nil {
		return nil
	}
	out := new(CassandraBackup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CassandraBackup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraBackupList) DeepCopyInto(out *CassandraBackupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CassandraBackup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraBackupList.
func (in *CassandraBackupList) DeepCopy() *CassandraBackupList {
	if in == nil {
		return nil
	}
	out := new(CassandraBackupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CassandraBackupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraBackupSpec) DeepCopyInto(out *CassandraBackupSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraBackupSpec.
func (in *CassandraBackupSpec) DeepCopy() *CassandraBackupSpec {
	if in == nil {
		return nil
	}
	out := new(CassandraBackupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraCluster) DeepCopyInto(out *CassandraCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraCluster.
func (in *CassandraCluster) DeepCopy() *CassandraCluster {
	if in == nil {
		return nil
	}
	out := new(CassandraCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CassandraCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraClusterList) DeepCopyInto(out *CassandraClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CassandraCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraClusterList.
func (in *CassandraClusterList) DeepCopy() *CassandraClusterList {
	if in == nil {
		return nil
	}
	out := new(CassandraClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CassandraClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraClusterSpec) DeepCopyInto(out *CassandraClusterSpec) {
	*out = *in
	if in.ReadOnlyRootFilesystem != nil {
		in, out := &in.ReadOnlyRootFilesystem, &out.ReadOnlyRootFilesystem
		*out = new(bool)
		**out = **in
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Pod != nil {
		in, out := &in.Pod, &out.Pod
		*out = new(PodPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(ServicePolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.StorageConfigs != nil {
		in, out := &in.StorageConfigs, &out.StorageConfigs
		*out = make([]StorageConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SidecarConfigs != nil {
		in, out := &in.SidecarConfigs, &out.SidecarConfigs
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	out.ImagePullSecret = in.ImagePullSecret
	out.ImageJolokiaSecret = in.ImageJolokiaSecret
	in.Topology.DeepCopyInto(&out.Topology)
	if in.LivenessInitialDelaySeconds != nil {
		in, out := &in.LivenessInitialDelaySeconds, &out.LivenessInitialDelaySeconds
		*out = new(int32)
		**out = **in
	}
	if in.LivenessHealthCheckTimeout != nil {
		in, out := &in.LivenessHealthCheckTimeout, &out.LivenessHealthCheckTimeout
		*out = new(int32)
		**out = **in
	}
	if in.LivenessHealthCheckPeriod != nil {
		in, out := &in.LivenessHealthCheckPeriod, &out.LivenessHealthCheckPeriod
		*out = new(int32)
		**out = **in
	}
	if in.LivenessFailureThreshold != nil {
		in, out := &in.LivenessFailureThreshold, &out.LivenessFailureThreshold
		*out = new(int32)
		**out = **in
	}
	if in.LivenessSuccessThreshold != nil {
		in, out := &in.LivenessSuccessThreshold, &out.LivenessSuccessThreshold
		*out = new(int32)
		**out = **in
	}
	if in.ReadinessInitialDelaySeconds != nil {
		in, out := &in.ReadinessInitialDelaySeconds, &out.ReadinessInitialDelaySeconds
		*out = new(int32)
		**out = **in
	}
	if in.ReadinessHealthCheckTimeout != nil {
		in, out := &in.ReadinessHealthCheckTimeout, &out.ReadinessHealthCheckTimeout
		*out = new(int32)
		**out = **in
	}
	if in.ReadinessHealthCheckPeriod != nil {
		in, out := &in.ReadinessHealthCheckPeriod, &out.ReadinessHealthCheckPeriod
		*out = new(int32)
		**out = **in
	}
	if in.ReadinessFailureThreshold != nil {
		in, out := &in.ReadinessFailureThreshold, &out.ReadinessFailureThreshold
		*out = new(int32)
		**out = **in
	}
	if in.ReadinessSuccessThreshold != nil {
		in, out := &in.ReadinessSuccessThreshold, &out.ReadinessSuccessThreshold
		*out = new(int32)
		**out = **in
	}
	if in.ShareProcessNamespace != nil {
		in, out := &in.ShareProcessNamespace, &out.ShareProcessNamespace
		*out = new(bool)
		**out = **in
	}
	if in.BackRestSidecar != nil {
		in, out := &in.BackRestSidecar, &out.BackRestSidecar
		*out = new(BackRestSidecar)
		(*in).DeepCopyInto(*out)
	}
	if in.JMXConfiguration != nil {
		in, out := &in.JMXConfiguration, &out.JMXConfiguration
		*out = new(JMXConfiguration)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraClusterSpec.
func (in *CassandraClusterSpec) DeepCopy() *CassandraClusterSpec {
	if in == nil {
		return nil
	}
	out := new(CassandraClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraClusterStatus) DeepCopyInto(out *CassandraClusterStatus) {
	*out = *in
	if in.SeedList != nil {
		in, out := &in.SeedList, &out.SeedList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CassandraNodesStatus != nil {
		in, out := &in.CassandraNodesStatus, &out.CassandraNodesStatus
		*out = make(map[string]CassandraNodeStatus, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.CassandraRackStatus != nil {
		in, out := &in.CassandraRackStatus, &out.CassandraRackStatus
		*out = make(map[string]*CassandraRackStatus, len(*in))
		for key, val := range *in {
			var outVal *CassandraRackStatus
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(CassandraRackStatus)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraClusterStatus.
func (in *CassandraClusterStatus) DeepCopy() *CassandraClusterStatus {
	if in == nil {
		return nil
	}
	out := new(CassandraClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraLastAction) DeepCopyInto(out *CassandraLastAction) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = (*in).DeepCopy()
	}
	if in.UpdatedNodes != nil {
		in, out := &in.UpdatedNodes, &out.UpdatedNodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraLastAction.
func (in *CassandraLastAction) DeepCopy() *CassandraLastAction {
	if in == nil {
		return nil
	}
	out := new(CassandraLastAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraNodeStatus) DeepCopyInto(out *CassandraNodeStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraNodeStatus.
func (in *CassandraNodeStatus) DeepCopy() *CassandraNodeStatus {
	if in == nil {
		return nil
	}
	out := new(CassandraNodeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraRackStatus) DeepCopyInto(out *CassandraRackStatus) {
	*out = *in
	in.CassandraLastAction.DeepCopyInto(&out.CassandraLastAction)
	in.PodLastOperation.DeepCopyInto(&out.PodLastOperation)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraRackStatus.
func (in *CassandraRackStatus) DeepCopy() *CassandraRackStatus {
	if in == nil {
		return nil
	}
	out := new(CassandraRackStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraRestore) DeepCopyInto(out *CassandraRestore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraRestore.
func (in *CassandraRestore) DeepCopy() *CassandraRestore {
	if in == nil {
		return nil
	}
	out := new(CassandraRestore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CassandraRestore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraRestoreList) DeepCopyInto(out *CassandraRestoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CassandraRestore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraRestoreList.
func (in *CassandraRestoreList) DeepCopy() *CassandraRestoreList {
	if in == nil {
		return nil
	}
	out := new(CassandraRestoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CassandraRestoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CassandraRestoreSpec) DeepCopyInto(out *CassandraRestoreSpec) {
	*out = *in
	if in.ConcurrentConnection != nil {
		in, out := &in.ConcurrentConnection, &out.ConcurrentConnection
		*out = new(int32)
		**out = **in
	}
	if in.Rename != nil {
		in, out := &in.Rename, &out.Rename
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CassandraRestoreSpec.
func (in *CassandraRestoreSpec) DeepCopy() *CassandraRestoreSpec {
	if in == nil {
		return nil
	}
	out := new(CassandraRestoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterStateInfo) DeepCopyInto(out *ClusterStateInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterStateInfo.
func (in *ClusterStateInfo) DeepCopy() *ClusterStateInfo {
	if in == nil {
		return nil
	}
	out := new(ClusterStateInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DC) DeepCopyInto(out *DC) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	if in.Rack != nil {
		in, out := &in.Rack, &out.Rack
		*out = make(RackSlice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodesPerRacks != nil {
		in, out := &in.NodesPerRacks, &out.NodesPerRacks
		*out = new(int32)
		**out = **in
	}
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DC.
func (in *DC) DeepCopy() *DC {
	if in == nil {
		return nil
	}
	out := new(DC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in DCSlice) DeepCopyInto(out *DCSlice) {
	{
		in := &in
		*out = make(DCSlice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DCSlice.
func (in DCSlice) DeepCopy() DCSlice {
	if in == nil {
		return nil
	}
	out := new(DCSlice)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FailureCause) DeepCopyInto(out *FailureCause) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FailureCause.
func (in *FailureCause) DeepCopy() *FailureCause {
	if in == nil {
		return nil
	}
	out := new(FailureCause)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JMXConfiguration) DeepCopyInto(out *JMXConfiguration) {
	*out = *in
	if in.JMXRemote != nil {
		in, out := &in.JMXRemote, &out.JMXRemote
		*out = new(bool)
		**out = **in
	}
	if in.JXMRemoteSSL != nil {
		in, out := &in.JXMRemoteSSL, &out.JXMRemoteSSL
		*out = new(bool)
		**out = **in
	}
	if in.JMXRemoteAuthenticate != nil {
		in, out := &in.JMXRemoteAuthenticate, &out.JMXRemoteAuthenticate
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JMXConfiguration.
func (in *JMXConfiguration) DeepCopy() *JMXConfiguration {
	if in == nil {
		return nil
	}
	out := new(JMXConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodLastOperation) DeepCopyInto(out *PodLastOperation) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = (*in).DeepCopy()
	}
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PodsOK != nil {
		in, out := &in.PodsOK, &out.PodsOK
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PodsKO != nil {
		in, out := &in.PodsKO, &out.PodsKO
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodLastOperation.
func (in *PodLastOperation) DeepCopy() *PodLastOperation {
	if in == nil {
		return nil
	}
	out := new(PodLastOperation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodPolicy) DeepCopyInto(out *PodPolicy) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodPolicy.
func (in *PodPolicy) DeepCopy() *PodPolicy {
	if in == nil {
		return nil
	}
	out := new(PodPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rack) DeepCopyInto(out *Rack) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rack.
func (in *Rack) DeepCopy() *Rack {
	if in == nil {
		return nil
	}
	out := new(Rack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in RackSlice) DeepCopyInto(out *RackSlice) {
	{
		in := &in
		*out = make(RackSlice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RackSlice.
func (in RackSlice) DeepCopy() RackSlice {
	if in == nil {
		return nil
	}
	out := new(RackSlice)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicePolicy) DeepCopyInto(out *ServicePolicy) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicePolicy.
func (in *ServicePolicy) DeepCopy() *ServicePolicy {
	if in == nil {
		return nil
	}
	out := new(ServicePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageConfig) DeepCopyInto(out *StorageConfig) {
	*out = *in
	if in.PVCSpec != nil {
		in, out := &in.PVCSpec, &out.PVCSpec
		*out = new(v1.PersistentVolumeClaimSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageConfig.
func (in *StorageConfig) DeepCopy() *StorageConfig {
	if in == nil {
		return nil
	}
	out := new(StorageConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Topology) DeepCopyInto(out *Topology) {
	*out = *in
	if in.DC != nil {
		in, out := &in.DC, &out.DC
		*out = make(DCSlice, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Topology.
func (in *Topology) DeepCopy() *Topology {
	if in == nil {
		return nil
	}
	out := new(Topology)
	in.DeepCopyInto(out)
	return out
}
