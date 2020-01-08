// +build !ignore_autogenerated

// Copyright © 2020 Banzai Cloud
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketWeb) DeepCopyInto(out *BucketWeb) {
	*out = *in
	out.HTTPGracePeriod = in.HTTPGracePeriod
	out.Refresh = in.Refresh
	out.Timeout = in.Timeout
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketWeb.
func (in *BucketWeb) DeepCopy() *BucketWeb {
	if in == nil {
		return nil
	}
	out := new(BucketWeb)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Compactor) DeepCopyInto(out *Compactor) {
	*out = *in
	out.HTTPGracePeriod = in.HTTPGracePeriod
	out.ConsistencyDelay = in.ConsistencyDelay
	out.RetentionResolutionRaw = in.RetentionResolutionRaw
	out.RetentionResolution5m = in.RetentionResolution5m
	out.RetentionResolution1h = in.RetentionResolution1h
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Compactor.
func (in *Compactor) DeepCopy() *Compactor {
	if in == nil {
		return nil
	}
	out := new(Compactor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Local) DeepCopyInto(out *Local) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Local.
func (in *Local) DeepCopy() *Local {
	if in == nil {
		return nil
	}
	out := new(Local)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStore) DeepCopyInto(out *ObjectStore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStore.
func (in *ObjectStore) DeepCopy() *ObjectStore {
	if in == nil {
		return nil
	}
	out := new(ObjectStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObjectStore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStoreList) DeepCopyInto(out *ObjectStoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ObjectStore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStoreList.
func (in *ObjectStoreList) DeepCopy() *ObjectStoreList {
	if in == nil {
		return nil
	}
	out := new(ObjectStoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ObjectStoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStoreSpec) DeepCopyInto(out *ObjectStoreSpec) {
	*out = *in
	in.Config.DeepCopyInto(&out.Config)
	if in.Compactor != nil {
		in, out := &in.Compactor, &out.Compactor
		*out = new(Compactor)
		**out = **in
	}
	if in.BucketWeb != nil {
		in, out := &in.BucketWeb, &out.BucketWeb
		*out = new(BucketWeb)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStoreSpec.
func (in *ObjectStoreSpec) DeepCopy() *ObjectStoreSpec {
	if in == nil {
		return nil
	}
	out := new(ObjectStoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectStoreStatus) DeepCopyInto(out *ObjectStoreStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectStoreStatus.
func (in *ObjectStoreStatus) DeepCopy() *ObjectStoreStatus {
	if in == nil {
		return nil
	}
	out := new(ObjectStoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Remote) DeepCopyInto(out *Remote) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Remote.
func (in *Remote) DeepCopy() *Remote {
	if in == nil {
		return nil
	}
	out := new(Remote)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreGateway) DeepCopyInto(out *StoreGateway) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreGateway.
func (in *StoreGateway) DeepCopy() *StoreGateway {
	if in == nil {
		return nil
	}
	out := new(StoreGateway)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Thanos) DeepCopyInto(out *Thanos) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Thanos.
func (in *Thanos) DeepCopy() *Thanos {
	if in == nil {
		return nil
	}
	out := new(Thanos)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Thanos) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosDiscovery) DeepCopyInto(out *ThanosDiscovery) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosDiscovery.
func (in *ThanosDiscovery) DeepCopy() *ThanosDiscovery {
	if in == nil {
		return nil
	}
	out := new(ThanosDiscovery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosList) DeepCopyInto(out *ThanosList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Thanos, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosList.
func (in *ThanosList) DeepCopy() *ThanosList {
	if in == nil {
		return nil
	}
	out := new(ThanosList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ThanosList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosSpec) DeepCopyInto(out *ThanosSpec) {
	*out = *in
	if in.Remote != nil {
		in, out := &in.Remote, &out.Remote
		*out = new(Remote)
		**out = **in
	}
	if in.ThanosDiscovery != nil {
		in, out := &in.ThanosDiscovery, &out.ThanosDiscovery
		*out = new(ThanosDiscovery)
		**out = **in
	}
	if in.Local != nil {
		in, out := &in.Local, &out.Local
		*out = new(Local)
		**out = **in
	}
	if in.StoreGateway != nil {
		in, out := &in.StoreGateway, &out.StoreGateway
		*out = new(StoreGateway)
		**out = **in
	}
	if in.Rule != nil {
		in, out := &in.Rule, &out.Rule
		*out = new(Rule)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosSpec.
func (in *ThanosSpec) DeepCopy() *ThanosSpec {
	if in == nil {
		return nil
	}
	out := new(ThanosSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ThanosStatus) DeepCopyInto(out *ThanosStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ThanosStatus.
func (in *ThanosStatus) DeepCopy() *ThanosStatus {
	if in == nil {
		return nil
	}
	out := new(ThanosStatus)
	in.DeepCopyInto(out)
	return out
}