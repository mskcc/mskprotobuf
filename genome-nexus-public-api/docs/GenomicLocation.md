# GenomicLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chromosome** | **string** | Chromosome | 
**Start** | **int32** | Start Position | 
**End** | **int32** | End Position | 
**ReferenceAllele** | **string** | Reference Allele | 
**VariantAllele** | **string** | Variant Allele | 

## Methods

### NewGenomicLocation

`func NewGenomicLocation(chromosome string, start int32, end int32, referenceAllele string, variantAllele string, ) *GenomicLocation`

NewGenomicLocation instantiates a new GenomicLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenomicLocationWithDefaults

`func NewGenomicLocationWithDefaults() *GenomicLocation`

NewGenomicLocationWithDefaults instantiates a new GenomicLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChromosome

`func (o *GenomicLocation) GetChromosome() string`

GetChromosome returns the Chromosome field if non-nil, zero value otherwise.

### GetChromosomeOk

`func (o *GenomicLocation) GetChromosomeOk() (*string, bool)`

GetChromosomeOk returns a tuple with the Chromosome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChromosome

`func (o *GenomicLocation) SetChromosome(v string)`

SetChromosome sets Chromosome field to given value.


### GetStart

`func (o *GenomicLocation) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *GenomicLocation) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *GenomicLocation) SetStart(v int32)`

SetStart sets Start field to given value.


### GetEnd

`func (o *GenomicLocation) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *GenomicLocation) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *GenomicLocation) SetEnd(v int32)`

SetEnd sets End field to given value.


### GetReferenceAllele

`func (o *GenomicLocation) GetReferenceAllele() string`

GetReferenceAllele returns the ReferenceAllele field if non-nil, zero value otherwise.

### GetReferenceAlleleOk

`func (o *GenomicLocation) GetReferenceAlleleOk() (*string, bool)`

GetReferenceAlleleOk returns a tuple with the ReferenceAllele field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceAllele

`func (o *GenomicLocation) SetReferenceAllele(v string)`

SetReferenceAllele sets ReferenceAllele field to given value.


### GetVariantAllele

`func (o *GenomicLocation) GetVariantAllele() string`

GetVariantAllele returns the VariantAllele field if non-nil, zero value otherwise.

### GetVariantAlleleOk

`func (o *GenomicLocation) GetVariantAlleleOk() (*string, bool)`

GetVariantAlleleOk returns a tuple with the VariantAllele field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantAllele

`func (o *GenomicLocation) SetVariantAllele(v string)`

SetVariantAllele sets VariantAllele field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


