//go:build !ignore_autogenerated

/*
Copyright 2024.

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

// Code generated by controller-gen. DO NOT EDIT.

package config

import ()

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExporterConfigSpec) DeepCopyInto(out *ExporterConfigSpec) {
	*out = *in
	out.Provider = in.Provider
	out.BearerToken = in.BearerToken
	if in.AdditionalVariables != nil {
		in, out := &in.AdditionalVariables, &out.AdditionalVariables
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExporterConfigSpec.
func (in *ExporterConfigSpec) DeepCopy() *ExporterConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ExporterConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExporterScraperConfig) DeepCopyInto(out *ExporterScraperConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExporterScraperConfig.
func (in *ExporterScraperConfig) DeepCopy() *ExporterScraperConfig {
	if in == nil {
		return nil
	}
	out := new(ExporterScraperConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExporterScraperConfigSpec) DeepCopyInto(out *ExporterScraperConfigSpec) {
	*out = *in
	in.ExporterConfig.DeepCopyInto(&out.ExporterConfig)
	out.ScraperConfig = in.ScraperConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExporterScraperConfigSpec.
func (in *ExporterScraperConfigSpec) DeepCopy() *ExporterScraperConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ExporterScraperConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExporterScraperConfigStatus) DeepCopyInto(out *ExporterScraperConfigStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
	out.ActiveExporter = in.ActiveExporter
	out.ConfigMap = in.ConfigMap
	out.Service = in.Service
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExporterScraperConfigStatus.
func (in *ExporterScraperConfigStatus) DeepCopy() *ExporterScraperConfigStatus {
	if in == nil {
		return nil
	}
	out := new(ExporterScraperConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FocusConfig) DeepCopyInto(out *FocusConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FocusConfig.
func (in *FocusConfig) DeepCopy() *FocusConfig {
	if in == nil {
		return nil
	}
	out := new(FocusConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FocusConfigList) DeepCopyInto(out *FocusConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FocusConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FocusConfigList.
func (in *FocusConfigList) DeepCopy() *FocusConfigList {
	if in == nil {
		return nil
	}
	out := new(FocusConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FocusConfigSpec) DeepCopyInto(out *FocusConfigSpec) {
	*out = *in
	in.FocusSpec.DeepCopyInto(&out.FocusSpec)
	out.ScraperConfig = in.ScraperConfig
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FocusConfigSpec.
func (in *FocusConfigSpec) DeepCopy() *FocusConfigSpec {
	if in == nil {
		return nil
	}
	out := new(FocusConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FocusConfigStatus) DeepCopyInto(out *FocusConfigStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FocusConfigStatus.
func (in *FocusConfigStatus) DeepCopy() *FocusConfigStatus {
	if in == nil {
		return nil
	}
	out := new(FocusConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FocusSpec) DeepCopyInto(out *FocusSpec) {
	*out = *in
	out.BilledCost = in.BilledCost.DeepCopy()
	in.BillingPeriodEnd.DeepCopyInto(&out.BillingPeriodEnd)
	in.BillingPeriodStart.DeepCopyInto(&out.BillingPeriodStart)
	in.ChargePeriodEnd.DeepCopyInto(&out.ChargePeriodEnd)
	in.ChargePeriodStart.DeepCopyInto(&out.ChargePeriodStart)
	out.CommitmentDiscountQuantity = in.CommitmentDiscountQuantity.DeepCopy()
	out.ConsumedQuantity = in.ConsumedQuantity.DeepCopy()
	out.ContractedCost = in.ContractedCost.DeepCopy()
	out.ContractedUnitCost = in.ContractedUnitCost.DeepCopy()
	out.EffectiveCost = in.EffectiveCost.DeepCopy()
	out.ListCost = in.ListCost.DeepCopy()
	out.ListUnitPrice = in.ListUnitPrice.DeepCopy()
	out.PricingQuantity = in.PricingQuantity.DeepCopy()
	if in.SkuPriceDetails != nil {
		in, out := &in.SkuPriceDetails, &out.SkuPriceDetails
		*out = make([]TagsType, len(*in))
		copy(*out, *in)
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]TagsType, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FocusSpec.
func (in *FocusSpec) DeepCopy() *FocusSpec {
	if in == nil {
		return nil
	}
	out := new(FocusSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ObjectRef) DeepCopyInto(out *ObjectRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ObjectRef.
func (in *ObjectRef) DeepCopy() *ObjectRef {
	if in == nil {
		return nil
	}
	out := new(ObjectRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScraperConfig) DeepCopyInto(out *ScraperConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScraperConfig.
func (in *ScraperConfig) DeepCopy() *ScraperConfig {
	if in == nil {
		return nil
	}
	out := new(ScraperConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScraperConfigSpec) DeepCopyInto(out *ScraperConfigSpec) {
	*out = *in
	out.ScraperDatabaseConfigRef = in.ScraperDatabaseConfigRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScraperConfigSpec.
func (in *ScraperConfigSpec) DeepCopy() *ScraperConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ScraperConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScraperConfigStatus) DeepCopyInto(out *ScraperConfigStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
	out.ActiveScraper = in.ActiveScraper
	out.ConfigMap = in.ConfigMap
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScraperConfigStatus.
func (in *ScraperConfigStatus) DeepCopy() *ScraperConfigStatus {
	if in == nil {
		return nil
	}
	out := new(ScraperConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretKeySelector) DeepCopyInto(out *SecretKeySelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretKeySelector.
func (in *SecretKeySelector) DeepCopy() *SecretKeySelector {
	if in == nil {
		return nil
	}
	out := new(SecretKeySelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagsType) DeepCopyInto(out *TagsType) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagsType.
func (in *TagsType) DeepCopy() *TagsType {
	if in == nil {
		return nil
	}
	out := new(TagsType)
	in.DeepCopyInto(out)
	return out
}
