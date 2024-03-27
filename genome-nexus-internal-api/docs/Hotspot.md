# Hotspot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HugoSymbol** | Pointer to **string** | Hugo gene symbol | [optional] 
**InframeCount** | Pointer to **int32** | Inframe mutation count | [optional] 
**MissenseCount** | Pointer to **int32** | Missense mutation count | [optional] 
**Residue** | Pointer to **string** | Hotspot residue | [optional] 
**SpliceCount** | Pointer to **int32** | Splice mutation count | [optional] 
**TranscriptId** | Pointer to **string** | Ensembl Transcript Id | [optional] 
**TruncatingCount** | Pointer to **int32** | Truncation mutation count | [optional] 
**TumorCount** | Pointer to **int32** | Tumor count | [optional] 
**Type** | Pointer to **string** | Hotspot type | [optional] 

## Methods

### NewHotspot

`func NewHotspot() *Hotspot`

NewHotspot instantiates a new Hotspot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHotspotWithDefaults

`func NewHotspotWithDefaults() *Hotspot`

NewHotspotWithDefaults instantiates a new Hotspot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHugoSymbol

`func (o *Hotspot) GetHugoSymbol() string`

GetHugoSymbol returns the HugoSymbol field if non-nil, zero value otherwise.

### GetHugoSymbolOk

`func (o *Hotspot) GetHugoSymbolOk() (*string, bool)`

GetHugoSymbolOk returns a tuple with the HugoSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHugoSymbol

`func (o *Hotspot) SetHugoSymbol(v string)`

SetHugoSymbol sets HugoSymbol field to given value.

### HasHugoSymbol

`func (o *Hotspot) HasHugoSymbol() bool`

HasHugoSymbol returns a boolean if a field has been set.

### GetInframeCount

`func (o *Hotspot) GetInframeCount() int32`

GetInframeCount returns the InframeCount field if non-nil, zero value otherwise.

### GetInframeCountOk

`func (o *Hotspot) GetInframeCountOk() (*int32, bool)`

GetInframeCountOk returns a tuple with the InframeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInframeCount

`func (o *Hotspot) SetInframeCount(v int32)`

SetInframeCount sets InframeCount field to given value.

### HasInframeCount

`func (o *Hotspot) HasInframeCount() bool`

HasInframeCount returns a boolean if a field has been set.

### GetMissenseCount

`func (o *Hotspot) GetMissenseCount() int32`

GetMissenseCount returns the MissenseCount field if non-nil, zero value otherwise.

### GetMissenseCountOk

`func (o *Hotspot) GetMissenseCountOk() (*int32, bool)`

GetMissenseCountOk returns a tuple with the MissenseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissenseCount

`func (o *Hotspot) SetMissenseCount(v int32)`

SetMissenseCount sets MissenseCount field to given value.

### HasMissenseCount

`func (o *Hotspot) HasMissenseCount() bool`

HasMissenseCount returns a boolean if a field has been set.

### GetResidue

`func (o *Hotspot) GetResidue() string`

GetResidue returns the Residue field if non-nil, zero value otherwise.

### GetResidueOk

`func (o *Hotspot) GetResidueOk() (*string, bool)`

GetResidueOk returns a tuple with the Residue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResidue

`func (o *Hotspot) SetResidue(v string)`

SetResidue sets Residue field to given value.

### HasResidue

`func (o *Hotspot) HasResidue() bool`

HasResidue returns a boolean if a field has been set.

### GetSpliceCount

`func (o *Hotspot) GetSpliceCount() int32`

GetSpliceCount returns the SpliceCount field if non-nil, zero value otherwise.

### GetSpliceCountOk

`func (o *Hotspot) GetSpliceCountOk() (*int32, bool)`

GetSpliceCountOk returns a tuple with the SpliceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpliceCount

`func (o *Hotspot) SetSpliceCount(v int32)`

SetSpliceCount sets SpliceCount field to given value.

### HasSpliceCount

`func (o *Hotspot) HasSpliceCount() bool`

HasSpliceCount returns a boolean if a field has been set.

### GetTranscriptId

`func (o *Hotspot) GetTranscriptId() string`

GetTranscriptId returns the TranscriptId field if non-nil, zero value otherwise.

### GetTranscriptIdOk

`func (o *Hotspot) GetTranscriptIdOk() (*string, bool)`

GetTranscriptIdOk returns a tuple with the TranscriptId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscriptId

`func (o *Hotspot) SetTranscriptId(v string)`

SetTranscriptId sets TranscriptId field to given value.

### HasTranscriptId

`func (o *Hotspot) HasTranscriptId() bool`

HasTranscriptId returns a boolean if a field has been set.

### GetTruncatingCount

`func (o *Hotspot) GetTruncatingCount() int32`

GetTruncatingCount returns the TruncatingCount field if non-nil, zero value otherwise.

### GetTruncatingCountOk

`func (o *Hotspot) GetTruncatingCountOk() (*int32, bool)`

GetTruncatingCountOk returns a tuple with the TruncatingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncatingCount

`func (o *Hotspot) SetTruncatingCount(v int32)`

SetTruncatingCount sets TruncatingCount field to given value.

### HasTruncatingCount

`func (o *Hotspot) HasTruncatingCount() bool`

HasTruncatingCount returns a boolean if a field has been set.

### GetTumorCount

`func (o *Hotspot) GetTumorCount() int32`

GetTumorCount returns the TumorCount field if non-nil, zero value otherwise.

### GetTumorCountOk

`func (o *Hotspot) GetTumorCountOk() (*int32, bool)`

GetTumorCountOk returns a tuple with the TumorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTumorCount

`func (o *Hotspot) SetTumorCount(v int32)`

SetTumorCount sets TumorCount field to given value.

### HasTumorCount

`func (o *Hotspot) HasTumorCount() bool`

HasTumorCount returns a boolean if a field has been set.

### GetType

`func (o *Hotspot) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Hotspot) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Hotspot) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Hotspot) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


