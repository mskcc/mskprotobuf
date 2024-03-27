# Vues

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comment** | Pointer to **string** |  | [optional] 
**DefaultEffect** | Pointer to **string** |  | [optional] 
**GenomicLocation** | Pointer to **string** |  | [optional] 
**GenomicLocationDescription** | Pointer to **string** |  | [optional] 
**HugoGeneSymbol** | Pointer to **string** |  | [optional] 
**PubmedIds** | Pointer to **[]int32** |  | [optional] 
**ReferenceText** | Pointer to **string** |  | [optional] 
**RevisedProteinEffect** | Pointer to **string** |  | [optional] 
**TranscriptId** | Pointer to **string** |  | [optional] 
**Variant** | Pointer to **string** |  | [optional] 
**VariantClassification** | Pointer to **string** |  | [optional] 

## Methods

### NewVues

`func NewVues() *Vues`

NewVues instantiates a new Vues object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVuesWithDefaults

`func NewVuesWithDefaults() *Vues`

NewVuesWithDefaults instantiates a new Vues object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComment

`func (o *Vues) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Vues) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Vues) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Vues) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetDefaultEffect

`func (o *Vues) GetDefaultEffect() string`

GetDefaultEffect returns the DefaultEffect field if non-nil, zero value otherwise.

### GetDefaultEffectOk

`func (o *Vues) GetDefaultEffectOk() (*string, bool)`

GetDefaultEffectOk returns a tuple with the DefaultEffect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultEffect

`func (o *Vues) SetDefaultEffect(v string)`

SetDefaultEffect sets DefaultEffect field to given value.

### HasDefaultEffect

`func (o *Vues) HasDefaultEffect() bool`

HasDefaultEffect returns a boolean if a field has been set.

### GetGenomicLocation

`func (o *Vues) GetGenomicLocation() string`

GetGenomicLocation returns the GenomicLocation field if non-nil, zero value otherwise.

### GetGenomicLocationOk

`func (o *Vues) GetGenomicLocationOk() (*string, bool)`

GetGenomicLocationOk returns a tuple with the GenomicLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenomicLocation

`func (o *Vues) SetGenomicLocation(v string)`

SetGenomicLocation sets GenomicLocation field to given value.

### HasGenomicLocation

`func (o *Vues) HasGenomicLocation() bool`

HasGenomicLocation returns a boolean if a field has been set.

### GetGenomicLocationDescription

`func (o *Vues) GetGenomicLocationDescription() string`

GetGenomicLocationDescription returns the GenomicLocationDescription field if non-nil, zero value otherwise.

### GetGenomicLocationDescriptionOk

`func (o *Vues) GetGenomicLocationDescriptionOk() (*string, bool)`

GetGenomicLocationDescriptionOk returns a tuple with the GenomicLocationDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenomicLocationDescription

`func (o *Vues) SetGenomicLocationDescription(v string)`

SetGenomicLocationDescription sets GenomicLocationDescription field to given value.

### HasGenomicLocationDescription

`func (o *Vues) HasGenomicLocationDescription() bool`

HasGenomicLocationDescription returns a boolean if a field has been set.

### GetHugoGeneSymbol

`func (o *Vues) GetHugoGeneSymbol() string`

GetHugoGeneSymbol returns the HugoGeneSymbol field if non-nil, zero value otherwise.

### GetHugoGeneSymbolOk

`func (o *Vues) GetHugoGeneSymbolOk() (*string, bool)`

GetHugoGeneSymbolOk returns a tuple with the HugoGeneSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHugoGeneSymbol

`func (o *Vues) SetHugoGeneSymbol(v string)`

SetHugoGeneSymbol sets HugoGeneSymbol field to given value.

### HasHugoGeneSymbol

`func (o *Vues) HasHugoGeneSymbol() bool`

HasHugoGeneSymbol returns a boolean if a field has been set.

### GetPubmedIds

`func (o *Vues) GetPubmedIds() []int32`

GetPubmedIds returns the PubmedIds field if non-nil, zero value otherwise.

### GetPubmedIdsOk

`func (o *Vues) GetPubmedIdsOk() (*[]int32, bool)`

GetPubmedIdsOk returns a tuple with the PubmedIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPubmedIds

`func (o *Vues) SetPubmedIds(v []int32)`

SetPubmedIds sets PubmedIds field to given value.

### HasPubmedIds

`func (o *Vues) HasPubmedIds() bool`

HasPubmedIds returns a boolean if a field has been set.

### GetReferenceText

`func (o *Vues) GetReferenceText() string`

GetReferenceText returns the ReferenceText field if non-nil, zero value otherwise.

### GetReferenceTextOk

`func (o *Vues) GetReferenceTextOk() (*string, bool)`

GetReferenceTextOk returns a tuple with the ReferenceText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceText

`func (o *Vues) SetReferenceText(v string)`

SetReferenceText sets ReferenceText field to given value.

### HasReferenceText

`func (o *Vues) HasReferenceText() bool`

HasReferenceText returns a boolean if a field has been set.

### GetRevisedProteinEffect

`func (o *Vues) GetRevisedProteinEffect() string`

GetRevisedProteinEffect returns the RevisedProteinEffect field if non-nil, zero value otherwise.

### GetRevisedProteinEffectOk

`func (o *Vues) GetRevisedProteinEffectOk() (*string, bool)`

GetRevisedProteinEffectOk returns a tuple with the RevisedProteinEffect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevisedProteinEffect

`func (o *Vues) SetRevisedProteinEffect(v string)`

SetRevisedProteinEffect sets RevisedProteinEffect field to given value.

### HasRevisedProteinEffect

`func (o *Vues) HasRevisedProteinEffect() bool`

HasRevisedProteinEffect returns a boolean if a field has been set.

### GetTranscriptId

`func (o *Vues) GetTranscriptId() string`

GetTranscriptId returns the TranscriptId field if non-nil, zero value otherwise.

### GetTranscriptIdOk

`func (o *Vues) GetTranscriptIdOk() (*string, bool)`

GetTranscriptIdOk returns a tuple with the TranscriptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptId

`func (o *Vues) SetTranscriptId(v string)`

SetTranscriptId sets TranscriptId field to given value.

### HasTranscriptId

`func (o *Vues) HasTranscriptId() bool`

HasTranscriptId returns a boolean if a field has been set.

### GetVariant

`func (o *Vues) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *Vues) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *Vues) SetVariant(v string)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *Vues) HasVariant() bool`

HasVariant returns a boolean if a field has been set.

### GetVariantClassification

`func (o *Vues) GetVariantClassification() string`

GetVariantClassification returns the VariantClassification field if non-nil, zero value otherwise.

### GetVariantClassificationOk

`func (o *Vues) GetVariantClassificationOk() (*string, bool)`

GetVariantClassificationOk returns a tuple with the VariantClassification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantClassification

`func (o *Vues) SetVariantClassification(v string)`

SetVariantClassification sets VariantClassification field to given value.

### HasVariantClassification

`func (o *Vues) HasVariantClassification() bool`

HasVariantClassification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


