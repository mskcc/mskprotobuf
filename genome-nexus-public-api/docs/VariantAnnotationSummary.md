# VariantAnnotationSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssemblyName** | Pointer to **string** | Assembly name | [optional] 
**CanonicalTranscriptId** | Pointer to **string** | Canonical transcript id | [optional] 
**GenomicLocation** | Pointer to [**GenomicLocation**](GenomicLocation.md) |  | [optional] 
**StrandSign** | Pointer to **string** | Strand (- or +) | [optional] 
**TranscriptConsequenceSummaries** | [**[]TranscriptConsequenceSummary**](TranscriptConsequenceSummary.md) | All transcript consequence summaries | 
**TranscriptConsequenceSummary** | [**TranscriptConsequenceSummary**](TranscriptConsequenceSummary.md) |  | 
**TranscriptConsequences** | [**[]TranscriptConsequenceSummary**](TranscriptConsequenceSummary.md) | (Deprecated) Transcript consequence summaries (list of one when using annotation/, multiple when using annotation/summary/ | 
**Variant** | **string** | Variant key | 
**VariantType** | Pointer to **string** | Variant type | [optional] 
**Vues** | Pointer to [**Vues**](Vues.md) |  | [optional] 

## Methods

### NewVariantAnnotationSummary

`func NewVariantAnnotationSummary(transcriptConsequenceSummaries []TranscriptConsequenceSummary, transcriptConsequenceSummary TranscriptConsequenceSummary, transcriptConsequences []TranscriptConsequenceSummary, variant string, ) *VariantAnnotationSummary`

NewVariantAnnotationSummary instantiates a new VariantAnnotationSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariantAnnotationSummaryWithDefaults

`func NewVariantAnnotationSummaryWithDefaults() *VariantAnnotationSummary`

NewVariantAnnotationSummaryWithDefaults instantiates a new VariantAnnotationSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssemblyName

`func (o *VariantAnnotationSummary) GetAssemblyName() string`

GetAssemblyName returns the AssemblyName field if non-nil, zero value otherwise.

### GetAssemblyNameOk

`func (o *VariantAnnotationSummary) GetAssemblyNameOk() (*string, bool)`

GetAssemblyNameOk returns a tuple with the AssemblyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssemblyName

`func (o *VariantAnnotationSummary) SetAssemblyName(v string)`

SetAssemblyName sets AssemblyName field to given value.

### HasAssemblyName

`func (o *VariantAnnotationSummary) HasAssemblyName() bool`

HasAssemblyName returns a boolean if a field has been set.

### GetCanonicalTranscriptId

`func (o *VariantAnnotationSummary) GetCanonicalTranscriptId() string`

GetCanonicalTranscriptId returns the CanonicalTranscriptId field if non-nil, zero value otherwise.

### GetCanonicalTranscriptIdOk

`func (o *VariantAnnotationSummary) GetCanonicalTranscriptIdOk() (*string, bool)`

GetCanonicalTranscriptIdOk returns a tuple with the CanonicalTranscriptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalTranscriptId

`func (o *VariantAnnotationSummary) SetCanonicalTranscriptId(v string)`

SetCanonicalTranscriptId sets CanonicalTranscriptId field to given value.

### HasCanonicalTranscriptId

`func (o *VariantAnnotationSummary) HasCanonicalTranscriptId() bool`

HasCanonicalTranscriptId returns a boolean if a field has been set.

### GetGenomicLocation

`func (o *VariantAnnotationSummary) GetGenomicLocation() GenomicLocation`

GetGenomicLocation returns the GenomicLocation field if non-nil, zero value otherwise.

### GetGenomicLocationOk

`func (o *VariantAnnotationSummary) GetGenomicLocationOk() (*GenomicLocation, bool)`

GetGenomicLocationOk returns a tuple with the GenomicLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenomicLocation

`func (o *VariantAnnotationSummary) SetGenomicLocation(v GenomicLocation)`

SetGenomicLocation sets GenomicLocation field to given value.

### HasGenomicLocation

`func (o *VariantAnnotationSummary) HasGenomicLocation() bool`

HasGenomicLocation returns a boolean if a field has been set.

### GetStrandSign

`func (o *VariantAnnotationSummary) GetStrandSign() string`

GetStrandSign returns the StrandSign field if non-nil, zero value otherwise.

### GetStrandSignOk

`func (o *VariantAnnotationSummary) GetStrandSignOk() (*string, bool)`

GetStrandSignOk returns a tuple with the StrandSign field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrandSign

`func (o *VariantAnnotationSummary) SetStrandSign(v string)`

SetStrandSign sets StrandSign field to given value.

### HasStrandSign

`func (o *VariantAnnotationSummary) HasStrandSign() bool`

HasStrandSign returns a boolean if a field has been set.

### GetTranscriptConsequenceSummaries

`func (o *VariantAnnotationSummary) GetTranscriptConsequenceSummaries() []TranscriptConsequenceSummary`

GetTranscriptConsequenceSummaries returns the TranscriptConsequenceSummaries field if non-nil, zero value otherwise.

### GetTranscriptConsequenceSummariesOk

`func (o *VariantAnnotationSummary) GetTranscriptConsequenceSummariesOk() (*[]TranscriptConsequenceSummary, bool)`

GetTranscriptConsequenceSummariesOk returns a tuple with the TranscriptConsequenceSummaries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptConsequenceSummaries

`func (o *VariantAnnotationSummary) SetTranscriptConsequenceSummaries(v []TranscriptConsequenceSummary)`

SetTranscriptConsequenceSummaries sets TranscriptConsequenceSummaries field to given value.


### GetTranscriptConsequenceSummary

`func (o *VariantAnnotationSummary) GetTranscriptConsequenceSummary() TranscriptConsequenceSummary`

GetTranscriptConsequenceSummary returns the TranscriptConsequenceSummary field if non-nil, zero value otherwise.

### GetTranscriptConsequenceSummaryOk

`func (o *VariantAnnotationSummary) GetTranscriptConsequenceSummaryOk() (*TranscriptConsequenceSummary, bool)`

GetTranscriptConsequenceSummaryOk returns a tuple with the TranscriptConsequenceSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptConsequenceSummary

`func (o *VariantAnnotationSummary) SetTranscriptConsequenceSummary(v TranscriptConsequenceSummary)`

SetTranscriptConsequenceSummary sets TranscriptConsequenceSummary field to given value.


### GetTranscriptConsequences

`func (o *VariantAnnotationSummary) GetTranscriptConsequences() []TranscriptConsequenceSummary`

GetTranscriptConsequences returns the TranscriptConsequences field if non-nil, zero value otherwise.

### GetTranscriptConsequencesOk

`func (o *VariantAnnotationSummary) GetTranscriptConsequencesOk() (*[]TranscriptConsequenceSummary, bool)`

GetTranscriptConsequencesOk returns a tuple with the TranscriptConsequences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptConsequences

`func (o *VariantAnnotationSummary) SetTranscriptConsequences(v []TranscriptConsequenceSummary)`

SetTranscriptConsequences sets TranscriptConsequences field to given value.


### GetVariant

`func (o *VariantAnnotationSummary) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *VariantAnnotationSummary) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *VariantAnnotationSummary) SetVariant(v string)`

SetVariant sets Variant field to given value.


### GetVariantType

`func (o *VariantAnnotationSummary) GetVariantType() string`

GetVariantType returns the VariantType field if non-nil, zero value otherwise.

### GetVariantTypeOk

`func (o *VariantAnnotationSummary) GetVariantTypeOk() (*string, bool)`

GetVariantTypeOk returns a tuple with the VariantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantType

`func (o *VariantAnnotationSummary) SetVariantType(v string)`

SetVariantType sets VariantType field to given value.

### HasVariantType

`func (o *VariantAnnotationSummary) HasVariantType() bool`

HasVariantType returns a boolean if a field has been set.

### GetVues

`func (o *VariantAnnotationSummary) GetVues() Vues`

GetVues returns the Vues field if non-nil, zero value otherwise.

### GetVuesOk

`func (o *VariantAnnotationSummary) GetVuesOk() (*Vues, bool)`

GetVuesOk returns a tuple with the Vues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVues

`func (o *VariantAnnotationSummary) SetVues(v Vues)`

SetVues sets Vues field to given value.

### HasVues

`func (o *VariantAnnotationSummary) HasVues() bool`

HasVues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


