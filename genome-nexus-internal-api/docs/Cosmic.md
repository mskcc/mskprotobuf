# Cosmic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alt** | Pointer to **string** | alt | [optional] 
**Chrom** | Pointer to **string** | chrom | [optional] 
**CosmicId** | Pointer to **string** | cosmic_id | [optional] 
**Hg19** | Pointer to [**Hg19**](Hg19.md) |  | [optional] 
**License** | Pointer to **string** | _license | [optional] 
**MutFreq** | Pointer to **float64** | mut_freq | [optional] 
**MutNt** | Pointer to **string** | mut_nt | [optional] 
**Ref** | Pointer to **string** | ref | [optional] 
**TumorSite** | Pointer to **string** | tumor_site | [optional] 

## Methods

### NewCosmic

`func NewCosmic() *Cosmic`

NewCosmic instantiates a new Cosmic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCosmicWithDefaults

`func NewCosmicWithDefaults() *Cosmic`

NewCosmicWithDefaults instantiates a new Cosmic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlt

`func (o *Cosmic) GetAlt() string`

GetAlt returns the Alt field if non-nil, zero value otherwise.

### GetAltOk

`func (o *Cosmic) GetAltOk() (*string, bool)`

GetAltOk returns a tuple with the Alt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlt

`func (o *Cosmic) SetAlt(v string)`

SetAlt sets Alt field to given value.

### HasAlt

`func (o *Cosmic) HasAlt() bool`

HasAlt returns a boolean if a field has been set.

### GetChrom

`func (o *Cosmic) GetChrom() string`

GetChrom returns the Chrom field if non-nil, zero value otherwise.

### GetChromOk

`func (o *Cosmic) GetChromOk() (*string, bool)`

GetChromOk returns a tuple with the Chrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChrom

`func (o *Cosmic) SetChrom(v string)`

SetChrom sets Chrom field to given value.

### HasChrom

`func (o *Cosmic) HasChrom() bool`

HasChrom returns a boolean if a field has been set.

### GetCosmicId

`func (o *Cosmic) GetCosmicId() string`

GetCosmicId returns the CosmicId field if non-nil, zero value otherwise.

### GetCosmicIdOk

`func (o *Cosmic) GetCosmicIdOk() (*string, bool)`

GetCosmicIdOk returns a tuple with the CosmicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCosmicId

`func (o *Cosmic) SetCosmicId(v string)`

SetCosmicId sets CosmicId field to given value.

### HasCosmicId

`func (o *Cosmic) HasCosmicId() bool`

HasCosmicId returns a boolean if a field has been set.

### GetHg19

`func (o *Cosmic) GetHg19() Hg19`

GetHg19 returns the Hg19 field if non-nil, zero value otherwise.

### GetHg19Ok

`func (o *Cosmic) GetHg19Ok() (*Hg19, bool)`

GetHg19Ok returns a tuple with the Hg19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHg19

`func (o *Cosmic) SetHg19(v Hg19)`

SetHg19 sets Hg19 field to given value.

### HasHg19

`func (o *Cosmic) HasHg19() bool`

HasHg19 returns a boolean if a field has been set.

### GetLicense

`func (o *Cosmic) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *Cosmic) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *Cosmic) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *Cosmic) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetMutFreq

`func (o *Cosmic) GetMutFreq() float64`

GetMutFreq returns the MutFreq field if non-nil, zero value otherwise.

### GetMutFreqOk

`func (o *Cosmic) GetMutFreqOk() (*float64, bool)`

GetMutFreqOk returns a tuple with the MutFreq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutFreq

`func (o *Cosmic) SetMutFreq(v float64)`

SetMutFreq sets MutFreq field to given value.

### HasMutFreq

`func (o *Cosmic) HasMutFreq() bool`

HasMutFreq returns a boolean if a field has been set.

### GetMutNt

`func (o *Cosmic) GetMutNt() string`

GetMutNt returns the MutNt field if non-nil, zero value otherwise.

### GetMutNtOk

`func (o *Cosmic) GetMutNtOk() (*string, bool)`

GetMutNtOk returns a tuple with the MutNt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutNt

`func (o *Cosmic) SetMutNt(v string)`

SetMutNt sets MutNt field to given value.

### HasMutNt

`func (o *Cosmic) HasMutNt() bool`

HasMutNt returns a boolean if a field has been set.

### GetRef

`func (o *Cosmic) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Cosmic) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Cosmic) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Cosmic) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetTumorSite

`func (o *Cosmic) GetTumorSite() string`

GetTumorSite returns the TumorSite field if non-nil, zero value otherwise.

### GetTumorSiteOk

`func (o *Cosmic) GetTumorSiteOk() (*string, bool)`

GetTumorSiteOk returns a tuple with the TumorSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTumorSite

`func (o *Cosmic) SetTumorSite(v string)`

SetTumorSite sets TumorSite field to given value.

### HasTumorSite

`func (o *Cosmic) HasTumorSite() bool`

HasTumorSite returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


