/*
Copyright 2021 The Kubernetes Authors.

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

package v1alpha2

import (
	"k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/kops/pkg/apis/kops"
	"k8s.io/kops/pkg/values"
)

// Convert_v1alpha2_CanalNetworkingSpec_To_kops_CanalNetworkingSpec is an autogenerated conversion function.
func Convert_v1alpha2_CanalNetworkingSpec_To_kops_CanalNetworkingSpec(in *CanalNetworkingSpec, out *kops.CanalNetworkingSpec, s conversion.Scope) error {
	if err := autoConvert_v1alpha2_CanalNetworkingSpec_To_kops_CanalNetworkingSpec(in, out, s); err != nil {
		return err
	}
	if in.FlanneldIptablesForwardRules != nil {
		out.FlanneldIptablesForwardRules = values.Bool(!*in.FlanneldIptablesForwardRules)
	}
	return nil
}

// Convert_kops_CanalNetworkingSpec_To_v1alpha2_CanalNetworkingSpec is an autogenerated conversion function.
func Convert_kops_CanalNetworkingSpec_To_v1alpha2_CanalNetworkingSpec(in *kops.CanalNetworkingSpec, out *CanalNetworkingSpec, s conversion.Scope) error {
	if err := autoConvert_kops_CanalNetworkingSpec_To_v1alpha2_CanalNetworkingSpec(in, out, s); err != nil {
		return err
	}
	if in.FlanneldIptablesForwardRules != nil {
		out.FlanneldIptablesForwardRules = values.Bool(!*in.FlanneldIptablesForwardRules)
	}
	return nil
}

func Convert_v1alpha2_CiliumNetworkingSpec_To_kops_CiliumNetworkingSpec(in *CiliumNetworkingSpec, out *kops.CiliumNetworkingSpec, s conversion.Scope) error {
	if err := autoConvert_v1alpha2_CiliumNetworkingSpec_To_kops_CiliumNetworkingSpec(in, out, s); err != nil {
		return err
	}
	if in.InstallIptablesRules != nil {
		out.InstallIptablesRules = values.Bool(!*in.InstallIptablesRules)
	}
	if in.Masquerade != nil {
		out.Masquerade = values.Bool(!*in.Masquerade)
	}
	return nil
}

func Convert_kops_CiliumNetworkingSpec_To_v1alpha2_CiliumNetworkingSpec(in *kops.CiliumNetworkingSpec, out *CiliumNetworkingSpec, s conversion.Scope) error {
	if err := autoConvert_kops_CiliumNetworkingSpec_To_v1alpha2_CiliumNetworkingSpec(in, out, s); err != nil {
		return err
	}
	if in.InstallIptablesRules != nil {
		out.InstallIptablesRules = values.Bool(!*in.InstallIptablesRules)
	}
	if in.Masquerade != nil {
		out.Masquerade = values.Bool(!*in.Masquerade)
	}
	return nil
}

func Convert_v1alpha2_ClusterSpec_To_kops_ClusterSpec(in *ClusterSpec, out *kops.ClusterSpec, s conversion.Scope) error {
	if err := autoConvert_v1alpha2_ClusterSpec_To_kops_ClusterSpec(in, out, s); err != nil {
		return err
	}
	if in.AdditionalPolicies != nil {
		out.AdditionalPolicies = make(map[string]string, len(in.AdditionalPolicies))
		for k, v := range in.AdditionalPolicies {
			if k == "master" {
				k = "control-plane"
			}
			out.AdditionalPolicies[k] = v
		}
	}
	if in.ExternalPolicies != nil {
		policies := make(map[string][]string, len(in.ExternalPolicies))
		for k, v := range in.ExternalPolicies {
			if k == "master" {
				k = "control-plane"
			}
			policies[k] = v
		}
		out.ExternalPolicies = policies
	}
	if in.LegacyNetworking != nil {
		if err := autoConvert_v1alpha2_NetworkingSpec_To_kops_NetworkingSpec(in.LegacyNetworking, &out.Networking, s); err != nil {
			return err
		}
	}
	if in.LegacyAPI != nil {
		if err := autoConvert_v1alpha2_APISpec_To_kops_APISpec(in.LegacyAPI, &out.API, s); err != nil {
			return err
		}
	}
	switch kops.CloudProviderID(in.LegacyCloudProvider) {
	case kops.CloudProviderAWS:
		out.CloudProvider.AWS = &kops.AWSSpec{}
	case kops.CloudProviderAzure:
		out.CloudProvider.Azure = &kops.AzureSpec{}
		if in.CloudConfig != nil && in.CloudConfig.Azure != nil {
			if err := autoConvert_v1alpha2_AzureSpec_To_kops_AzureSpec(in.CloudConfig.Azure, out.CloudProvider.Azure, s); err != nil {
				return err
			}
		}
	case kops.CloudProviderDO:
		out.CloudProvider.DO = &kops.DOSpec{}
	case kops.CloudProviderGCE:
		out.CloudProvider.GCE = &kops.GCESpec{
			Project: in.Project,
		}
	case kops.CloudProviderHetzner:
		out.CloudProvider.Hetzner = &kops.HetznerSpec{}
	case kops.CloudProviderOpenstack:
		out.CloudProvider.Openstack = &kops.OpenstackSpec{}
		if in.CloudConfig != nil && in.CloudConfig.Openstack != nil {
			if err := autoConvert_v1alpha2_OpenstackSpec_To_kops_OpenstackSpec(in.CloudConfig.Openstack, out.CloudProvider.Openstack, s); err != nil {
				return err
			}
		}
	case kops.CloudProviderScaleway:
		out.CloudProvider.Scaleway = &kops.ScalewaySpec{}
	case "":
	default:
		return field.NotSupported(field.NewPath("spec").Child("cloudProvider"), in.LegacyCloudProvider, []string{
			string(kops.CloudProviderGCE),
			string(kops.CloudProviderDO),
			string(kops.CloudProviderAzure),
			string(kops.CloudProviderAWS),
			string(kops.CloudProviderHetzner),
			string(kops.CloudProviderOpenstack),
			string(kops.CloudProviderScaleway),
		})
	}
	if in.CloudConfig != nil {
		if in.CloudConfig.AWSEBSCSIDriver != nil {
			if out.CloudProvider.AWS == nil {
				return field.Forbidden(field.NewPath("spec").Child("cloudConfig", "awsEBSCSIDriver"), "EBS CSI driver supports only AWS")
			}
			out.CloudProvider.AWS.EBSCSIDriver = &kops.EBSCSIDriverSpec{}
			if err := autoConvert_v1alpha2_EBSCSIDriverSpec_To_kops_EBSCSIDriverSpec(in.CloudConfig.AWSEBSCSIDriver, out.CloudProvider.AWS.EBSCSIDriver, s); err != nil {
				return err
			}
		}
		if in.CloudConfig.GCEServiceAccount != "" {
			if out.CloudProvider.GCE == nil {
				return field.Forbidden(field.NewPath("spec").Child("cloudConfig", "gceServiceAccount"), "GCE Service Account supports only GCE")
			}
			out.CloudProvider.GCE.ServiceAccount = in.CloudConfig.GCEServiceAccount
		}
		if in.CloudConfig.DisableSecurityGroupIngress != nil {
			if out.CloudProvider.AWS == nil {
				return field.Forbidden(field.NewPath("spec").Child("cloudConfig", "disableSecurityGroupIngress"), "disableSecurityGroupIngress supports only AWS")
			}
			val := *in.CloudConfig.DisableSecurityGroupIngress
			out.CloudProvider.AWS.DisableSecurityGroupIngress = &val
		}
		if in.CloudConfig.ElbSecurityGroup != nil {
			if out.CloudProvider.AWS == nil {
				return field.Forbidden(field.NewPath("spec").Child("cloudConfig", "elbSecurityGroup"), "elbSecurityGroup supports only AWS")
			}
			val := *in.CloudConfig.ElbSecurityGroup
			out.CloudProvider.AWS.ElbSecurityGroup = &val
		}
		if in.CloudConfig.GCPPDCSIDriver != nil {
			if out.CloudProvider.GCE == nil {
				return field.Forbidden(field.NewPath("spec").Child("cloudConfig", "gcpPDCSIDriver"), "PD CSI driver supports only GCE")
			}
			out.CloudProvider.GCE.PDCSIDriver = &kops.PDCSIDriver{}
			if err := autoConvert_v1alpha2_PDCSIDriver_To_kops_PDCSIDriver(in.CloudConfig.GCPPDCSIDriver, out.CloudProvider.GCE.PDCSIDriver, s); err != nil {
				return err
			}
		}
		if in.CloudConfig.Multizone != nil {
			if out.CloudProvider.GCE == nil {
				return field.Forbidden(field.NewPath("spec").Child("cloudConfig", "multizone"), "multizone supports only GCE")
			}
			out.CloudProvider.GCE.Multizone = in.CloudConfig.Multizone
		}
		if in.CloudConfig.NodeIPFamilies != nil {
			if out.CloudProvider.AWS == nil {
				return field.Forbidden(field.NewPath("spec").Child("cloudConfig", "nodeIPFamilies"), "nodeIPFamilies supports only AWS")
			}
			out.CloudProvider.AWS.NodeIPFamilies = append([]string(nil), in.CloudConfig.NodeIPFamilies...)
		}
		if in.CloudConfig.NodeTags != nil {
			if out.CloudProvider.GCE == nil {
				return field.Forbidden(field.NewPath("spec").Child("cloudConfig", "nodeTags"), "nodeTags supports only GCE")
			}
			out.CloudProvider.GCE.NodeTags = in.CloudConfig.NodeTags
		}
		if in.CloudConfig.NodeInstancePrefix != nil {
			if out.CloudProvider.GCE == nil {
				return field.Forbidden(field.NewPath("spec").Child("cloudConfig", "nodeInstancePrefix"), "nodeInstancePrefix supports only GCE")
			}
			out.CloudProvider.GCE.NodeInstancePrefix = in.CloudConfig.NodeInstancePrefix
		}
		if in.CloudConfig.SpotinstOrientation != nil {
			if out.CloudProvider.AWS == nil {
				return field.Forbidden(field.NewPath("spec").Child("cloudConfig", "spotinstOrientation"), "Spotinst supports only AWS")
			}
			val := *in.CloudConfig.SpotinstOrientation
			out.CloudProvider.AWS.SpotinstOrientation = &val
		}
		if in.CloudConfig.SpotinstProduct != nil {
			if out.CloudProvider.AWS == nil {
				return field.Forbidden(field.NewPath("spec").Child("cloudConfig", "spotinstProduct"), "Spotinst supports only AWS")
			}
			val := *in.CloudConfig.SpotinstProduct
			out.CloudProvider.AWS.SpotinstProduct = &val
		}
	}
	if in.NodeTerminationHandler != nil {
		if out.CloudProvider.AWS == nil {
			return field.Forbidden(field.NewPath("spec").Child("nodeTerminationHandler"), "node termination handler supports only AWS")
		}
		out.CloudProvider.AWS.NodeTerminationHandler = &kops.NodeTerminationHandlerSpec{}
		if err := autoConvert_v1alpha2_NodeTerminationHandlerSpec_To_kops_NodeTerminationHandlerSpec(in.NodeTerminationHandler, out.CloudProvider.AWS.NodeTerminationHandler, s); err != nil {
			return err
		}
	}
	if in.AWSLoadBalancerController != nil {
		if out.CloudProvider.AWS == nil {
			return field.Forbidden(field.NewPath("spec").Child("awsLoadBalancerController"), "AWS Load Balancer Controller supports only AWS")
		}
		out.CloudProvider.AWS.LoadBalancerController = &kops.LoadBalancerControllerSpec{}
		if err := autoConvert_v1alpha2_LoadBalancerControllerSpec_To_kops_LoadBalancerControllerSpec(in.AWSLoadBalancerController, out.CloudProvider.AWS.LoadBalancerController, s); err != nil {
			return err
		}
	}
	if in.WarmPool != nil {
		if out.CloudProvider.AWS == nil {
			return field.Forbidden(field.NewPath("spec", "warmPool"), "warm pool only supported on AWS")
		}
		out.CloudProvider.AWS.WarmPool = &kops.WarmPoolSpec{}
		if err := autoConvert_v1alpha2_WarmPoolSpec_To_kops_WarmPoolSpec(in.WarmPool, out.CloudProvider.AWS.WarmPool, s); err != nil {
			return err
		}
	}
	if in.PodIdentityWebhook != nil {
		if out.CloudProvider.AWS == nil {
			return field.Forbidden(field.NewPath("spec", "podIdentityWebhook"), "pod identity webhook supports only AWS")
		}
		out.CloudProvider.AWS.PodIdentityWebhook = &kops.PodIdentityWebhookSpec{}
		if err := autoConvert_v1alpha2_PodIdentityWebhookSpec_To_kops_PodIdentityWebhookSpec(in.PodIdentityWebhook, out.CloudProvider.AWS.PodIdentityWebhook, s); err != nil {
			return err
		}
	}
	for i, hook := range in.Hooks {
		if hook.Enabled != nil {
			out.Hooks[i].Enabled = values.Bool(!*hook.Enabled)
		}
	}
	if in.Subnets != nil {
		in, out := &in.Subnets, &out.Networking.Subnets
		*out = make([]kops.ClusterSubnetSpec, len(*in))
		for i := range *in {
			if err := Convert_v1alpha2_ClusterSubnetSpec_To_kops_ClusterSubnetSpec(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Networking.Subnets = nil
	}
	out.API.PublicName = in.MasterPublicName
	out.Networking.NetworkCIDR = in.NetworkCIDR
	out.Networking.AdditionalNetworkCIDRs = in.AdditionalNetworkCIDRs
	out.Networking.NetworkID = in.NetworkID
	if in.Topology != nil {
		in, out := &in.Topology, &out.Networking.Topology
		*out = new(kops.TopologySpec)
		if err := Convert_v1alpha2_TopologySpec_To_kops_TopologySpec(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Networking.Topology = nil
	}
	out.Networking.ServiceClusterIPRange = in.ServiceClusterIPRange
	out.Networking.PodCIDR = in.PodCIDR
	out.Networking.NonMasqueradeCIDR = in.NonMasqueradeCIDR
	if in.EgressProxy != nil {
		in, out := &in.EgressProxy, &out.Networking.EgressProxy
		*out = new(kops.EgressProxySpec)
		if err := Convert_v1alpha2_EgressProxySpec_To_kops_EgressProxySpec(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Networking.EgressProxy = nil
	}
	if in.IsolateMasters != nil {
		in, out := &in.IsolateMasters, &out.Networking.IsolateControlPlane
		*out = new(bool)
		**out = **in
	}
	out.API.AdditionalSANs = in.AdditionalSANs
	out.API.Access = in.KubernetesAPIAccess
	if in.TagSubnets != nil {
		out.Networking.TagSubnets = values.Bool(!*in.TagSubnets)
	}
	return nil
}

func Convert_kops_ClusterSpec_To_v1alpha2_ClusterSpec(in *kops.ClusterSpec, out *ClusterSpec, s conversion.Scope) error {
	if err := autoConvert_kops_ClusterSpec_To_v1alpha2_ClusterSpec(in, out, s); err != nil {
		return err
	}
	if in.AdditionalPolicies != nil {
		out.AdditionalPolicies = make(map[string]string, len(in.AdditionalPolicies))
		for k, v := range in.AdditionalPolicies {
			if k == "control-plane" {
				k = "master"
			}
			out.AdditionalPolicies[k] = v
		}
	}
	if in.ExternalPolicies != nil {
		out.ExternalPolicies = make(map[string][]string, len(in.ExternalPolicies))
		for k, v := range in.ExternalPolicies {
			if k == "control-plane" {
				k = "master"
			}
			out.ExternalPolicies[k] = v
		}
	}
	out.LegacyNetworking = &NetworkingSpec{}
	if err := autoConvert_kops_NetworkingSpec_To_v1alpha2_NetworkingSpec(&in.Networking, out.LegacyNetworking, s); err != nil {
		return err
	}
	if out.LegacyNetworking.IsEmpty() {
		out.LegacyNetworking = nil
	}
	out.LegacyAPI = &APISpec{}
	if err := autoConvert_kops_APISpec_To_v1alpha2_APISpec(&in.API, out.LegacyAPI, s); err != nil {
		return err
	}
	if out.API.IsEmpty() {
		out.LegacyAPI = nil
	}
	out.LegacyCloudProvider = string(in.GetCloudProvider())
	switch kops.CloudProviderID(out.LegacyCloudProvider) {
	case kops.CloudProviderAWS:
		aws := in.CloudProvider.AWS
		if aws.DisableSecurityGroupIngress != nil {
			if out.CloudConfig == nil {
				out.CloudConfig = &CloudConfiguration{}
			}
			val := *aws.DisableSecurityGroupIngress
			out.CloudConfig.DisableSecurityGroupIngress = &val
		}
		if aws.EBSCSIDriver != nil {
			if out.CloudConfig == nil {
				out.CloudConfig = &CloudConfiguration{}
			}
			out.CloudConfig.AWSEBSCSIDriver = &EBSCSIDriverSpec{}
			if err := autoConvert_kops_EBSCSIDriverSpec_To_v1alpha2_EBSCSIDriverSpec(aws.EBSCSIDriver, out.CloudConfig.AWSEBSCSIDriver, s); err != nil {
				return err
			}
		}
		if aws.ElbSecurityGroup != nil {
			if out.CloudConfig == nil {
				out.CloudConfig = &CloudConfiguration{}
			}
			val := *aws.ElbSecurityGroup
			out.CloudConfig.ElbSecurityGroup = &val
		}
		if aws.NodeTerminationHandler != nil {
			out.NodeTerminationHandler = &NodeTerminationHandlerSpec{}
			if err := autoConvert_kops_NodeTerminationHandlerSpec_To_v1alpha2_NodeTerminationHandlerSpec(aws.NodeTerminationHandler, out.NodeTerminationHandler, s); err != nil {
				return err
			}
		}
		if aws.LoadBalancerController != nil {
			out.AWSLoadBalancerController = &LoadBalancerControllerSpec{}
			if err := autoConvert_kops_LoadBalancerControllerSpec_To_v1alpha2_LoadBalancerControllerSpec(aws.LoadBalancerController, out.AWSLoadBalancerController, s); err != nil {
				return err
			}
		}
		if aws.NodeIPFamilies != nil {
			if out.CloudConfig == nil {
				out.CloudConfig = &CloudConfiguration{}
			}
			out.CloudConfig.NodeIPFamilies = append([]string(nil), aws.NodeIPFamilies...)
		}
		if aws.SpotinstOrientation != nil {
			if out.CloudConfig == nil {
				out.CloudConfig = &CloudConfiguration{}
			}
			val := *aws.SpotinstOrientation
			out.CloudConfig.SpotinstOrientation = &val
		}
		if aws.SpotinstProduct != nil {
			if out.CloudConfig == nil {
				out.CloudConfig = &CloudConfiguration{}
			}
			val := *aws.SpotinstProduct
			out.CloudConfig.SpotinstProduct = &val
		}
		if aws.WarmPool != nil {
			out.WarmPool = &WarmPoolSpec{}
			if err := autoConvert_kops_WarmPoolSpec_To_v1alpha2_WarmPoolSpec(aws.WarmPool, out.WarmPool, s); err != nil {
				return err
			}
		}
		if aws.PodIdentityWebhook != nil {
			out.PodIdentityWebhook = &PodIdentityWebhookSpec{}
			if err := autoConvert_kops_PodIdentityWebhookSpec_To_v1alpha2_PodIdentityWebhookSpec(aws.PodIdentityWebhook, out.PodIdentityWebhook, s); err != nil {
				return err
			}
		}
	case kops.CloudProviderAzure:
		if out.CloudConfig == nil {
			out.CloudConfig = &CloudConfiguration{}
		}
		if out.CloudConfig.Azure == nil {
			out.CloudConfig.Azure = &AzureSpec{}
		}
		if err := autoConvert_kops_AzureSpec_To_v1alpha2_AzureSpec(in.CloudProvider.Azure, out.CloudConfig.Azure, s); err != nil {
			return err
		}
	case kops.CloudProviderGCE:
		gce := in.CloudProvider.GCE
		out.Project = gce.Project
		if gce.Multizone != nil {
			if out.CloudConfig == nil {
				out.CloudConfig = &CloudConfiguration{}
			}
			out.CloudConfig.Multizone = gce.Multizone
		}
		if gce.NodeTags != nil {
			if out.CloudConfig == nil {
				out.CloudConfig = &CloudConfiguration{}
			}
			out.CloudConfig.NodeTags = gce.NodeTags
		}
		if gce.NodeInstancePrefix != nil {
			if out.CloudConfig == nil {
				out.CloudConfig = &CloudConfiguration{}
			}
			out.CloudConfig.NodeInstancePrefix = gce.NodeInstancePrefix
		}
		if gce.ServiceAccount != "" {
			if out.CloudConfig == nil {
				out.CloudConfig = &CloudConfiguration{}
			}
			out.CloudConfig.GCEServiceAccount = gce.ServiceAccount
		}
		if gce.PDCSIDriver != nil {
			if out.CloudConfig == nil {
				out.CloudConfig = &CloudConfiguration{}
			}
			out.CloudConfig.GCPPDCSIDriver = &PDCSIDriver{}
			if err := autoConvert_kops_PDCSIDriver_To_v1alpha2_PDCSIDriver(gce.PDCSIDriver, out.CloudConfig.GCPPDCSIDriver, s); err != nil {
				return err
			}
		}
	case kops.CloudProviderOpenstack:
		if out.CloudConfig == nil {
			out.CloudConfig = &CloudConfiguration{}
		}
		if out.CloudConfig.Openstack == nil {
			out.CloudConfig.Openstack = &OpenstackSpec{}
		}
		if err := autoConvert_kops_OpenstackSpec_To_v1alpha2_OpenstackSpec(in.CloudProvider.Openstack, out.CloudConfig.Openstack, s); err != nil {
			return err
		}
	}
	if in.Networking.Subnets != nil {
		in, out := &in.Networking.Subnets, &out.Subnets
		*out = make([]ClusterSubnetSpec, len(*in))
		for i := range *in {
			if err := Convert_kops_ClusterSubnetSpec_To_v1alpha2_ClusterSubnetSpec(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Subnets = nil
	}
	out.NetworkCIDR = in.Networking.NetworkCIDR
	out.AdditionalNetworkCIDRs = in.Networking.AdditionalNetworkCIDRs
	out.NetworkID = in.Networking.NetworkID
	if in.Networking.Topology != nil {
		in, out := &in.Networking.Topology, &out.Topology
		*out = new(TopologySpec)
		if err := Convert_kops_TopologySpec_To_v1alpha2_TopologySpec(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Topology = nil
	}
	out.ServiceClusterIPRange = in.Networking.ServiceClusterIPRange
	out.PodCIDR = in.Networking.PodCIDR
	out.NonMasqueradeCIDR = in.Networking.NonMasqueradeCIDR
	if in.Networking.EgressProxy != nil {
		in, out := &in.Networking.EgressProxy, &out.EgressProxy
		*out = new(EgressProxySpec)
		if err := Convert_kops_EgressProxySpec_To_v1alpha2_EgressProxySpec(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.EgressProxy = nil
	}
	if in.Networking.IsolateControlPlane != nil {
		in, out := &in.Networking.IsolateControlPlane, &out.IsolateMasters
		*out = new(bool)
		**out = **in
	}
	if in.Networking.TagSubnets != nil {
		out.TagSubnets = values.Bool(!*in.Networking.TagSubnets)
	}
	for i, hook := range in.Hooks {
		if hook.Enabled != nil {
			out.Hooks[i].Enabled = values.Bool(!*hook.Enabled)
		}
	}
	out.MasterPublicName = in.API.PublicName
	out.AdditionalSANs = in.API.AdditionalSANs
	out.KubernetesAPIAccess = in.API.Access
	return nil
}

func Convert_v1alpha2_ExternalDNSConfig_To_kops_ExternalDNSConfig(in *ExternalDNSConfig, out *kops.ExternalDNSConfig, s conversion.Scope) error {
	if err := autoConvert_v1alpha2_ExternalDNSConfig_To_kops_ExternalDNSConfig(in, out, s); err != nil {
		return err
	}
	if in.Disable {
		out.Provider = kops.ExternalDNSProviderNone
	}
	return nil
}

func Convert_kops_ExternalDNSConfig_To_v1alpha2_ExternalDNSConfig(in *kops.ExternalDNSConfig, out *ExternalDNSConfig, s conversion.Scope) error {
	if err := autoConvert_kops_ExternalDNSConfig_To_v1alpha2_ExternalDNSConfig(in, out, s); err != nil {
		return err
	}
	if in.Provider == kops.ExternalDNSProviderNone {
		out.Disable = true
		out.Provider = ""
	}
	return nil
}

// Convert_v1alpha2_HookSpec_To_kops_HookSpec is an autogenerated conversion function.
func Convert_v1alpha2_HookSpec_To_kops_HookSpec(in *HookSpec, out *kops.HookSpec, s conversion.Scope) error {
	if err := autoConvert_v1alpha2_HookSpec_To_kops_HookSpec(in, out, s); err != nil {
		return err
	}
	if in.Roles != nil {
		for i := range in.Roles {
			if in.Roles[i] == "Master" {
				out.Roles[i] = kops.InstanceGroupRoleControlPlane
			}
		}
	}
	return nil
}

func Convert_kops_HookSpec_To_v1alpha2_HookSpec(in *kops.HookSpec, out *HookSpec, s conversion.Scope) error {
	if err := autoConvert_kops_HookSpec_To_v1alpha2_HookSpec(in, out, s); err != nil {
		return err
	}
	if in.Roles != nil {
		for i := range in.Roles {
			if in.Roles[i] == kops.InstanceGroupRoleControlPlane {
				out.Roles[i] = "Master"
			}
		}
	}
	return nil
}

func Convert_v1alpha2_InstanceGroupSpec_To_kops_InstanceGroupSpec(in *InstanceGroupSpec, out *kops.InstanceGroupSpec, s conversion.Scope) error {
	if err := autoConvert_v1alpha2_InstanceGroupSpec_To_kops_InstanceGroupSpec(in, out, s); err != nil {
		return err
	}
	if in.Role == "Master" {
		out.Role = kops.InstanceGroupRoleControlPlane
	}
	return nil
}

func Convert_kops_InstanceGroupSpec_To_v1alpha2_InstanceGroupSpec(in *kops.InstanceGroupSpec, out *InstanceGroupSpec, s conversion.Scope) error {
	if err := autoConvert_kops_InstanceGroupSpec_To_v1alpha2_InstanceGroupSpec(in, out, s); err != nil {
		return err
	}
	if in.Role == kops.InstanceGroupRoleControlPlane {
		out.Role = "Master"
	}
	return nil
}

func Convert_v1alpha2_TopologySpec_To_kops_TopologySpec(in *TopologySpec, out *kops.TopologySpec, s conversion.Scope) error {
	if err := autoConvert_v1alpha2_TopologySpec_To_kops_TopologySpec(in, out, s); err != nil {
		return err
	}
	if in.LegacyDNS != nil {
		out.DNS = kops.DNSType(in.LegacyDNS.Type)
	}
	return nil
}

func Convert_kops_TopologySpec_To_v1alpha2_TopologySpec(in *kops.TopologySpec, out *TopologySpec, s conversion.Scope) error {
	if err := autoConvert_kops_TopologySpec_To_v1alpha2_TopologySpec(in, out, s); err != nil {
		return err
	}
	if in.DNS != "" {
		out.LegacyDNS = &DNSSpec{Type: DNSType(in.DNS)}
	}
	return nil
}
