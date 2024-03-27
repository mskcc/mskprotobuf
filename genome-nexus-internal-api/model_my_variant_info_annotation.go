/*
Genome Nexus API

This page shows how to use HTTP requests to access the Genome Nexus API. There are more high level clients available in Python, R, JavaScript, TypeScript and various other languages as well as a command line client to annotate MAF and VCF. See https://docs.genomenexus.org/api.  Aside from programmatic clients there are web based tools to annotate variants, see https://docs.genomenexus.org/tools.   We currently only provide long-term support for the '/annotation' endpoint. The other endpoints might change.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package genome_nexus_internal_api

import (
	"encoding/json"
)

// MyVariantInfoAnnotation struct for MyVariantInfoAnnotation
type MyVariantInfoAnnotation struct {
	Annotation *MyVariantInfo `json:"annotation,omitempty"`
	License *string `json:"license,omitempty"`
}

// NewMyVariantInfoAnnotation instantiates a new MyVariantInfoAnnotation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMyVariantInfoAnnotation() *MyVariantInfoAnnotation {
	this := MyVariantInfoAnnotation{}
	return &this
}

// NewMyVariantInfoAnnotationWithDefaults instantiates a new MyVariantInfoAnnotation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMyVariantInfoAnnotationWithDefaults() *MyVariantInfoAnnotation {
	this := MyVariantInfoAnnotation{}
	return &this
}

// GetAnnotation returns the Annotation field value if set, zero value otherwise.
func (o *MyVariantInfoAnnotation) GetAnnotation() MyVariantInfo {
	if o == nil || isNil(o.Annotation) {
		var ret MyVariantInfo
		return ret
	}
	return *o.Annotation
}

// GetAnnotationOk returns a tuple with the Annotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyVariantInfoAnnotation) GetAnnotationOk() (*MyVariantInfo, bool) {
	if o == nil || isNil(o.Annotation) {
    return nil, false
	}
	return o.Annotation, true
}

// HasAnnotation returns a boolean if a field has been set.
func (o *MyVariantInfoAnnotation) HasAnnotation() bool {
	if o != nil && !isNil(o.Annotation) {
		return true
	}

	return false
}

// SetAnnotation gets a reference to the given MyVariantInfo and assigns it to the Annotation field.
func (o *MyVariantInfoAnnotation) SetAnnotation(v MyVariantInfo) {
	o.Annotation = &v
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *MyVariantInfoAnnotation) GetLicense() string {
	if o == nil || isNil(o.License) {
		var ret string
		return ret
	}
	return *o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MyVariantInfoAnnotation) GetLicenseOk() (*string, bool) {
	if o == nil || isNil(o.License) {
    return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *MyVariantInfoAnnotation) HasLicense() bool {
	if o != nil && !isNil(o.License) {
		return true
	}

	return false
}

// SetLicense gets a reference to the given string and assigns it to the License field.
func (o *MyVariantInfoAnnotation) SetLicense(v string) {
	o.License = &v
}

func (o MyVariantInfoAnnotation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Annotation) {
		toSerialize["annotation"] = o.Annotation
	}
	if !isNil(o.License) {
		toSerialize["license"] = o.License
	}
	return json.Marshal(toSerialize)
}

type NullableMyVariantInfoAnnotation struct {
	value *MyVariantInfoAnnotation
	isSet bool
}

func (v NullableMyVariantInfoAnnotation) Get() *MyVariantInfoAnnotation {
	return v.value
}

func (v *NullableMyVariantInfoAnnotation) Set(val *MyVariantInfoAnnotation) {
	v.value = val
	v.isSet = true
}

func (v NullableMyVariantInfoAnnotation) IsSet() bool {
	return v.isSet
}

func (v *NullableMyVariantInfoAnnotation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMyVariantInfoAnnotation(val *MyVariantInfoAnnotation) *NullableMyVariantInfoAnnotation {
	return &NullableMyVariantInfoAnnotation{value: val, isSet: true}
}

func (v NullableMyVariantInfoAnnotation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMyVariantInfoAnnotation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


