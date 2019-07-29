// +build !ignore_autogenerated

/*
Copyright 2019 The Kubernetes Authors.

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

package cloud

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Config) DeepCopyInto(out *Config) {
	*out = *in
	in.Global.DeepCopyInto(&out.Global)
	if in.VCenter != nil {
		in, out := &in.VCenter, &out.VCenter
		*out = make(map[string]*VCenterConfig, len(*in))
		for key, val := range *in {
			var outVal *VCenterConfig
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(VCenterConfig)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	out.Network = in.Network
	out.Disk = in.Disk
	out.Workspace = in.Workspace
	out.Labels = in.Labels
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Config.
func (in *Config) DeepCopy() *Config {
	if in == nil {
		return nil
	}
	out := new(Config)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskConfig) DeepCopyInto(out *DiskConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskConfig.
func (in *DiskConfig) DeepCopy() *DiskConfig {
	if in == nil {
		return nil
	}
	out := new(DiskConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalConfig) DeepCopyInto(out *GlobalConfig) {
	*out = *in
	if in.APIDisable != nil {
		in, out := &in.APIDisable, &out.APIDisable
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalConfig.
func (in *GlobalConfig) DeepCopy() *GlobalConfig {
	if in == nil {
		return nil
	}
	out := new(GlobalConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LabelConfig) DeepCopyInto(out *LabelConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LabelConfig.
func (in *LabelConfig) DeepCopy() *LabelConfig {
	if in == nil {
		return nil
	}
	out := new(LabelConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworConfig) DeepCopyInto(out *NetworConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworConfig.
func (in *NetworConfig) DeepCopy() *NetworConfig {
	if in == nil {
		return nil
	}
	out := new(NetworConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnmarshalINIOptions) DeepCopyInto(out *UnmarshalINIOptions) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnmarshalINIOptions.
func (in *UnmarshalINIOptions) DeepCopy() *UnmarshalINIOptions {
	if in == nil {
		return nil
	}
	out := new(UnmarshalINIOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VCenterConfig) DeepCopyInto(out *VCenterConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VCenterConfig.
func (in *VCenterConfig) DeepCopy() *VCenterConfig {
	if in == nil {
		return nil
	}
	out := new(VCenterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkspaceConfig) DeepCopyInto(out *WorkspaceConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkspaceConfig.
func (in *WorkspaceConfig) DeepCopy() *WorkspaceConfig {
	if in == nil {
		return nil
	}
	out := new(WorkspaceConfig)
	in.DeepCopyInto(out)
	return out
}
