//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2019 Wrangler Sample Controller Authors

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

// Code generated by main. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKSClusterConfig) DeepCopyInto(out *AKSClusterConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKSClusterConfig.
func (in *AKSClusterConfig) DeepCopy() *AKSClusterConfig {
	if in == nil {
		return nil
	}
	out := new(AKSClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AKSClusterConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKSClusterConfigList) DeepCopyInto(out *AKSClusterConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AKSClusterConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKSClusterConfigList.
func (in *AKSClusterConfigList) DeepCopy() *AKSClusterConfigList {
	if in == nil {
		return nil
	}
	out := new(AKSClusterConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AKSClusterConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKSClusterConfigSpec) DeepCopyInto(out *AKSClusterConfigSpec) {
	*out = *in
	if in.BaseURL != nil {
		in, out := &in.BaseURL, &out.BaseURL
		*out = new(string)
		**out = **in
	}
	if in.AuthBaseURL != nil {
		in, out := &in.AuthBaseURL, &out.AuthBaseURL
		*out = new(string)
		**out = **in
	}
	if in.NetworkPlugin != nil {
		in, out := &in.NetworkPlugin, &out.NetworkPlugin
		*out = new(string)
		**out = **in
	}
	if in.VirtualNetworkResourceGroup != nil {
		in, out := &in.VirtualNetworkResourceGroup, &out.VirtualNetworkResourceGroup
		*out = new(string)
		**out = **in
	}
	if in.VirtualNetwork != nil {
		in, out := &in.VirtualNetwork, &out.VirtualNetwork
		*out = new(string)
		**out = **in
	}
	if in.Subnet != nil {
		in, out := &in.Subnet, &out.Subnet
		*out = new(string)
		**out = **in
	}
	if in.NetworkDNSServiceIP != nil {
		in, out := &in.NetworkDNSServiceIP, &out.NetworkDNSServiceIP
		*out = new(string)
		**out = **in
	}
	if in.NetworkServiceCIDR != nil {
		in, out := &in.NetworkServiceCIDR, &out.NetworkServiceCIDR
		*out = new(string)
		**out = **in
	}
	if in.NetworkDockerBridgeCIDR != nil {
		in, out := &in.NetworkDockerBridgeCIDR, &out.NetworkDockerBridgeCIDR
		*out = new(string)
		**out = **in
	}
	if in.NetworkPodCIDR != nil {
		in, out := &in.NetworkPodCIDR, &out.NetworkPodCIDR
		*out = new(string)
		**out = **in
	}
	if in.NodeResourceGroup != nil {
		in, out := &in.NodeResourceGroup, &out.NodeResourceGroup
		*out = new(string)
		**out = **in
	}
	if in.OutboundType != nil {
		in, out := &in.OutboundType, &out.OutboundType
		*out = new(string)
		**out = **in
	}
	if in.LoadBalancerSKU != nil {
		in, out := &in.LoadBalancerSKU, &out.LoadBalancerSKU
		*out = new(string)
		**out = **in
	}
	if in.NetworkPolicy != nil {
		in, out := &in.NetworkPolicy, &out.NetworkPolicy
		*out = new(string)
		**out = **in
	}
	if in.LinuxAdminUsername != nil {
		in, out := &in.LinuxAdminUsername, &out.LinuxAdminUsername
		*out = new(string)
		**out = **in
	}
	if in.LinuxSSHPublicKey != nil {
		in, out := &in.LinuxSSHPublicKey, &out.LinuxSSHPublicKey
		*out = new(string)
		**out = **in
	}
	if in.DNSPrefix != nil {
		in, out := &in.DNSPrefix, &out.DNSPrefix
		*out = new(string)
		**out = **in
	}
	if in.KubernetesVersion != nil {
		in, out := &in.KubernetesVersion, &out.KubernetesVersion
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NodePools != nil {
		in, out := &in.NodePools, &out.NodePools
		*out = make([]AKSNodePool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PrivateCluster != nil {
		in, out := &in.PrivateCluster, &out.PrivateCluster
		*out = new(bool)
		**out = **in
	}
	if in.PrivateDNSZone != nil {
		in, out := &in.PrivateDNSZone, &out.PrivateDNSZone
		*out = new(string)
		**out = **in
	}
	if in.AuthorizedIPRanges != nil {
		in, out := &in.AuthorizedIPRanges, &out.AuthorizedIPRanges
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.HTTPApplicationRouting != nil {
		in, out := &in.HTTPApplicationRouting, &out.HTTPApplicationRouting
		*out = new(bool)
		**out = **in
	}
	if in.Monitoring != nil {
		in, out := &in.Monitoring, &out.Monitoring
		*out = new(bool)
		**out = **in
	}
	if in.LogAnalyticsWorkspaceGroup != nil {
		in, out := &in.LogAnalyticsWorkspaceGroup, &out.LogAnalyticsWorkspaceGroup
		*out = new(string)
		**out = **in
	}
	if in.LogAnalyticsWorkspaceName != nil {
		in, out := &in.LogAnalyticsWorkspaceName, &out.LogAnalyticsWorkspaceName
		*out = new(string)
		**out = **in
	}
	if in.ManagedIdentity != nil {
		in, out := &in.ManagedIdentity, &out.ManagedIdentity
		*out = new(bool)
		**out = **in
	}
	if in.UserAssignedIdentity != nil {
		in, out := &in.UserAssignedIdentity, &out.UserAssignedIdentity
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKSClusterConfigSpec.
func (in *AKSClusterConfigSpec) DeepCopy() *AKSClusterConfigSpec {
	if in == nil {
		return nil
	}
	out := new(AKSClusterConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKSClusterConfigStatus) DeepCopyInto(out *AKSClusterConfigStatus) {
	*out = *in
	if in.RBACEnabled != nil {
		in, out := &in.RBACEnabled, &out.RBACEnabled
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKSClusterConfigStatus.
func (in *AKSClusterConfigStatus) DeepCopy() *AKSClusterConfigStatus {
	if in == nil {
		return nil
	}
	out := new(AKSClusterConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AKSNodePool) DeepCopyInto(out *AKSNodePool) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Count != nil {
		in, out := &in.Count, &out.Count
		*out = new(int32)
		**out = **in
	}
	if in.MaxPods != nil {
		in, out := &in.MaxPods, &out.MaxPods
		*out = new(int32)
		**out = **in
	}
	if in.OsDiskSizeGB != nil {
		in, out := &in.OsDiskSizeGB, &out.OsDiskSizeGB
		*out = new(int32)
		**out = **in
	}
	if in.OrchestratorVersion != nil {
		in, out := &in.OrchestratorVersion, &out.OrchestratorVersion
		*out = new(string)
		**out = **in
	}
	if in.AvailabilityZones != nil {
		in, out := &in.AvailabilityZones, &out.AvailabilityZones
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	if in.MaxSurge != nil {
		in, out := &in.MaxSurge, &out.MaxSurge
		*out = new(string)
		**out = **in
	}
	if in.MaxCount != nil {
		in, out := &in.MaxCount, &out.MaxCount
		*out = new(int32)
		**out = **in
	}
	if in.MinCount != nil {
		in, out := &in.MinCount, &out.MinCount
		*out = new(int32)
		**out = **in
	}
	if in.EnableAutoScaling != nil {
		in, out := &in.EnableAutoScaling, &out.EnableAutoScaling
		*out = new(bool)
		**out = **in
	}
	if in.VnetSubnetID != nil {
		in, out := &in.VnetSubnetID, &out.VnetSubnetID
		*out = new(string)
		**out = **in
	}
	if in.NodeLabels != nil {
		in, out := &in.NodeLabels, &out.NodeLabels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.NodeTaints != nil {
		in, out := &in.NodeTaints, &out.NodeTaints
		*out = new([]string)
		if **in != nil {
			in, out := *in, *out
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AKSNodePool.
func (in *AKSNodePool) DeepCopy() *AKSNodePool {
	if in == nil {
		return nil
	}
	out := new(AKSNodePool)
	in.DeepCopyInto(out)
	return out
}
