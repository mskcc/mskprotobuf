/*
Genome Nexus API

This page shows how to use HTTP requests to access the Genome Nexus API. There are more high level clients available in Python, R, JavaScript, TypeScript and various other languages as well as a command line client to annotate MAF and VCF. See https://docs.genomenexus.org/api.  Aside from programmatic clients there are web based tools to annotate variants, see https://docs.genomenexus.org/tools.   We currently only provide long-term support for the '/annotation' endpoint. The other endpoints might change.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package genome_nexus_public_api

import (
	"encoding/json"
)

// GenomicLocation struct for GenomicLocation
type GenomicLocation struct {
	// Chromosome
	Chromosome string `json:"chromosome"`
	// Start Position
	Start int32 `json:"start"`
	// End Position
	End int32 `json:"end"`
	// Reference Allele
	ReferenceAllele string `json:"referenceAllele"`
	// Variant Allele
	VariantAllele string `json:"variantAllele"`
}

// NewGenomicLocation instantiates a new GenomicLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenomicLocation(chromosome string, start int32, end int32, referenceAllele string, variantAllele string) *GenomicLocation {
	this := GenomicLocation{}
	this.Chromosome = chromosome
	this.Start = start
	this.End = end
	this.ReferenceAllele = referenceAllele
	this.VariantAllele = variantAllele
	return &this
}

// NewGenomicLocationWithDefaults instantiates a new GenomicLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenomicLocationWithDefaults() *GenomicLocation {
	this := GenomicLocation{}
	return &this
}

// GetChromosome returns the Chromosome field value
func (o *GenomicLocation) GetChromosome() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Chromosome
}

// GetChromosomeOk returns a tuple with the Chromosome field value
// and a boolean to check if the value has been set.
func (o *GenomicLocation) GetChromosomeOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Chromosome, true
}

// SetChromosome sets field value
func (o *GenomicLocation) SetChromosome(v string) {
	o.Chromosome = v
}

// GetStart returns the Start field value
func (o *GenomicLocation) GetStart() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *GenomicLocation) GetStartOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *GenomicLocation) SetStart(v int32) {
	o.Start = v
}

// GetEnd returns the End field value
func (o *GenomicLocation) GetEnd() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *GenomicLocation) GetEndOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *GenomicLocation) SetEnd(v int32) {
	o.End = v
}

// GetReferenceAllele returns the ReferenceAllele field value
func (o *GenomicLocation) GetReferenceAllele() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceAllele
}

// GetReferenceAlleleOk returns a tuple with the ReferenceAllele field value
// and a boolean to check if the value has been set.
func (o *GenomicLocation) GetReferenceAlleleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.ReferenceAllele, true
}

// SetReferenceAllele sets field value
func (o *GenomicLocation) SetReferenceAllele(v string) {
	o.ReferenceAllele = v
}

// GetVariantAllele returns the VariantAllele field value
func (o *GenomicLocation) GetVariantAllele() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VariantAllele
}

// GetVariantAlleleOk returns a tuple with the VariantAllele field value
// and a boolean to check if the value has been set.
func (o *GenomicLocation) GetVariantAlleleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.VariantAllele, true
}

// SetVariantAllele sets field value
func (o *GenomicLocation) SetVariantAllele(v string) {
	o.VariantAllele = v
}

func (o GenomicLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["chromosome"] = o.Chromosome
	}
	if true {
		toSerialize["start"] = o.Start
	}
	if true {
		toSerialize["end"] = o.End
	}
	if true {
		toSerialize["referenceAllele"] = o.ReferenceAllele
	}
	if true {
		toSerialize["variantAllele"] = o.VariantAllele
	}
	return json.Marshal(toSerialize)
}

type NullableGenomicLocation struct {
	value *GenomicLocation
	isSet bool
}

func (v NullableGenomicLocation) Get() *GenomicLocation {
	return v.value
}

func (v *NullableGenomicLocation) Set(val *GenomicLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableGenomicLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableGenomicLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenomicLocation(val *GenomicLocation) *NullableGenomicLocation {
	return &NullableGenomicLocation{value: val, isSet: true}
}

func (v NullableGenomicLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenomicLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


