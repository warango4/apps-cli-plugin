//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Catalog) DeepCopyInto(out *Catalog) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.PluginDescriptors != nil {
		in, out := &in.PluginDescriptors, &out.PluginDescriptors
		*out = make([]*PluginDescriptor, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(PluginDescriptor)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.IndexByPath != nil {
		in, out := &in.IndexByPath, &out.IndexByPath
		*out = make(map[string]PluginDescriptor, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.IndexByName != nil {
		in, out := &in.IndexByName, &out.IndexByName
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.StandAlonePlugins != nil {
		in, out := &in.StandAlonePlugins, &out.StandAlonePlugins
		*out = make(PluginAssociation, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ServerPlugins != nil {
		in, out := &in.ServerPlugins, &out.ServerPlugins
		*out = make(map[string]PluginAssociation, len(*in))
		for key, val := range *in {
			var outVal map[string]string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(PluginAssociation, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Catalog.
func (in *Catalog) DeepCopy() *Catalog {
	if in == nil {
		return nil
	}
	out := new(Catalog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Catalog) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CatalogList) DeepCopyInto(out *CatalogList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Catalog, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CatalogList.
func (in *CatalogList) DeepCopy() *CatalogList {
	if in == nil {
		return nil
	}
	out := new(CatalogList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CatalogList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Distro) DeepCopyInto(out *Distro) {
	{
		in := &in
		*out = make(Distro, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Distro.
func (in Distro) DeepCopy() Distro {
	if in == nil {
		return nil
	}
	out := new(Distro)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in PluginAssociation) DeepCopyInto(out *PluginAssociation) {
	{
		in := &in
		*out = make(PluginAssociation, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginAssociation.
func (in PluginAssociation) DeepCopy() PluginAssociation {
	if in == nil {
		return nil
	}
	out := new(PluginAssociation)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginDescriptor) DeepCopyInto(out *PluginDescriptor) {
	*out = *in
	if in.CompletionArgs != nil {
		in, out := &in.CompletionArgs, &out.CompletionArgs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Aliases != nil {
		in, out := &in.Aliases, &out.Aliases
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.PostInstallHook.DeepCopyInto(&out.PostInstallHook)
	if in.DefaultFeatureFlags != nil {
		in, out := &in.DefaultFeatureFlags, &out.DefaultFeatureFlags
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginDescriptor.
func (in *PluginDescriptor) DeepCopy() *PluginDescriptor {
	if in == nil {
		return nil
	}
	out := new(PluginDescriptor)
	in.DeepCopyInto(out)
	return out
}
