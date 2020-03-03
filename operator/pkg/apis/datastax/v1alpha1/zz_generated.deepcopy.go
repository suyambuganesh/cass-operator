// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	json "encoding/json"

	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DseDatacenter) DeepCopyInto(out *DseDatacenter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DseDatacenter.
func (in *DseDatacenter) DeepCopy() *DseDatacenter {
	if in == nil {
		return nil
	}
	out := new(DseDatacenter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DseDatacenter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DseDatacenterList) DeepCopyInto(out *DseDatacenterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DseDatacenter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DseDatacenterList.
func (in *DseDatacenterList) DeepCopy() *DseDatacenterList {
	if in == nil {
		return nil
	}
	out := new(DseDatacenterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DseDatacenterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DseDatacenterSpec) DeepCopyInto(out *DseDatacenterSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = make(json.RawMessage, len(*in))
		copy(*out, *in)
	}
	in.ManagementApiAuth.DeepCopyInto(&out.ManagementApiAuth)
	in.Resources.DeepCopyInto(&out.Resources)
	if in.Racks != nil {
		in, out := &in.Racks, &out.Racks
		*out = make([]Rack, len(*in))
		copy(*out, *in)
	}
	if in.StorageClaim != nil {
		in, out := &in.StorageClaim, &out.StorageClaim
		*out = new(StorageClaim)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DseDatacenterSpec.
func (in *DseDatacenterSpec) DeepCopy() *DseDatacenterSpec {
	if in == nil {
		return nil
	}
	out := new(DseDatacenterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DseDatacenterStatus) DeepCopyInto(out *DseDatacenterStatus) {
	*out = *in
	in.SuperUserUpserted.DeepCopyInto(&out.SuperUserUpserted)
	in.LastDseNodeStarted.DeepCopyInto(&out.LastDseNodeStarted)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DseDatacenterStatus.
func (in *DseDatacenterStatus) DeepCopy() *DseDatacenterStatus {
	if in == nil {
		return nil
	}
	out := new(DseDatacenterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementApiAuthConfig) DeepCopyInto(out *ManagementApiAuthConfig) {
	*out = *in
	if in.Insecure != nil {
		in, out := &in.Insecure, &out.Insecure
		*out = new(ManagementApiAuthInsecureConfig)
		**out = **in
	}
	if in.Manual != nil {
		in, out := &in.Manual, &out.Manual
		*out = new(ManagementApiAuthManualConfig)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementApiAuthConfig.
func (in *ManagementApiAuthConfig) DeepCopy() *ManagementApiAuthConfig {
	if in == nil {
		return nil
	}
	out := new(ManagementApiAuthConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementApiAuthInsecureConfig) DeepCopyInto(out *ManagementApiAuthInsecureConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementApiAuthInsecureConfig.
func (in *ManagementApiAuthInsecureConfig) DeepCopy() *ManagementApiAuthInsecureConfig {
	if in == nil {
		return nil
	}
	out := new(ManagementApiAuthInsecureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagementApiAuthManualConfig) DeepCopyInto(out *ManagementApiAuthManualConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagementApiAuthManualConfig.
func (in *ManagementApiAuthManualConfig) DeepCopy() *ManagementApiAuthManualConfig {
	if in == nil {
		return nil
	}
	out := new(ManagementApiAuthManualConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rack) DeepCopyInto(out *Rack) {
	*out = *in
	return
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
func (in *StorageClaim) DeepCopyInto(out *StorageClaim) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageClaim.
func (in *StorageClaim) DeepCopy() *StorageClaim {
	if in == nil {
		return nil
	}
	out := new(StorageClaim)
	in.DeepCopyInto(out)
	return out
}
