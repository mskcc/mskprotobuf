# UntranslatedRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | UTR Type | 
**Start** | **int32** | Start position of UTR | 
**End** | **int32** | End position of UTR | 
**Strand** | **int32** | Strand UTR is on, -1 for - and 1 for + | 

## Methods

### NewUntranslatedRegion

`func NewUntranslatedRegion(type_ string, start int32, end int32, strand int32, ) *UntranslatedRegion`

NewUntranslatedRegion instantiates a new UntranslatedRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUntranslatedRegionWithDefaults

`func NewUntranslatedRegionWithDefaults() *UntranslatedRegion`

NewUntranslatedRegionWithDefaults instantiates a new UntranslatedRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UntranslatedRegion) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UntranslatedRegion) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UntranslatedRegion) SetType(v string)`

SetType sets Type field to given value.


### GetStart

`func (o *UntranslatedRegion) GetStart() int32`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *UntranslatedRegion) GetStartOk() (*int32, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *UntranslatedRegion) SetStart(v int32)`

SetStart sets Start field to given value.


### GetEnd

`func (o *UntranslatedRegion) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *UntranslatedRegion) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *UntranslatedRegion) SetEnd(v int32)`

SetEnd sets End field to given value.


### GetStrand

`func (o *UntranslatedRegion) GetStrand() int32`

GetStrand returns the Strand field if non-nil, zero value otherwise.

### GetStrandOk

`func (o *UntranslatedRegion) GetStrandOk() (*int32, bool)`

GetStrandOk returns a tuple with the Strand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrand

`func (o *UntranslatedRegion) SetStrand(v int32)`

SetStrand sets Strand field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


