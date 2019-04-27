package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func (in *Tuned) DeepCopyInto(out *Tuned) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}
func (in *Tuned) DeepCopy() *Tuned {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(Tuned)
	in.DeepCopyInto(out)
	return out
}
func (in *Tuned) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *TunedList) DeepCopyInto(out *TunedList) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Tuned, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *TunedList) DeepCopy() *TunedList {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(TunedList)
	in.DeepCopyInto(out)
	return out
}
func (in *TunedList) DeepCopyObject() runtime.Object {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
func (in *TunedMatch) DeepCopyInto(out *TunedMatch) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Label != nil {
		in, out := &in.Label, &out.Label
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Match != nil {
		in, out := &in.Match, &out.Match
		*out = make([]TunedMatch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *TunedMatch) DeepCopy() *TunedMatch {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(TunedMatch)
	in.DeepCopyInto(out)
	return out
}
func (in *TunedProfile) DeepCopyInto(out *TunedProfile) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Data != nil {
		in, out := &in.Data, &out.Data
		*out = new(string)
		**out = **in
	}
	return
}
func (in *TunedProfile) DeepCopy() *TunedProfile {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(TunedProfile)
	in.DeepCopyInto(out)
	return out
}
func (in *TunedRecommend) DeepCopyInto(out *TunedRecommend) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Profile != nil {
		in, out := &in.Profile, &out.Profile
		*out = new(string)
		**out = **in
	}
	if in.Priority != nil {
		in, out := &in.Priority, &out.Priority
		*out = new(uint64)
		**out = **in
	}
	if in.Match != nil {
		in, out := &in.Match, &out.Match
		*out = make([]TunedMatch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *TunedRecommend) DeepCopy() *TunedRecommend {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(TunedRecommend)
	in.DeepCopyInto(out)
	return out
}
func (in *TunedSpec) DeepCopyInto(out *TunedSpec) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	if in.Profile != nil {
		in, out := &in.Profile, &out.Profile
		*out = make([]TunedProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Recommend != nil {
		in, out := &in.Recommend, &out.Recommend
		*out = make([]TunedRecommend, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}
func (in *TunedSpec) DeepCopy() *TunedSpec {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(TunedSpec)
	in.DeepCopyInto(out)
	return out
}
func (in *TunedStatus) DeepCopyInto(out *TunedStatus) {
	_logClusterCodePath()
	defer _logClusterCodePath()
	*out = *in
	return
}
func (in *TunedStatus) DeepCopy() *TunedStatus {
	_logClusterCodePath()
	defer _logClusterCodePath()
	if in == nil {
		return nil
	}
	out := new(TunedStatus)
	in.DeepCopyInto(out)
	return out
}
