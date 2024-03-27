# PdbHeader

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Compound** | Pointer to **map[string]interface{}** |  | [optional] 
**PdbId** | **string** | PDB id | 
**Source** | Pointer to **map[string]interface{}** |  | [optional] 
**Title** | **string** | PDB description | 

## Methods

### NewPdbHeader

`func NewPdbHeader(pdbId string, title string, ) *PdbHeader`

NewPdbHeader instantiates a new PdbHeader object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPdbHeaderWithDefaults

`func NewPdbHeaderWithDefaults() *PdbHeader`

NewPdbHeaderWithDefaults instantiates a new PdbHeader object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompound

`func (o *PdbHeader) GetCompound() map[string]interface{}`

GetCompound returns the Compound field if non-nil, zero value otherwise.

### GetCompoundOk

`func (o *PdbHeader) GetCompoundOk() (*map[string]interface{}, bool)`

GetCompoundOk returns a tuple with the Compound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompound

`func (o *PdbHeader) SetCompound(v map[string]interface{})`

SetCompound sets Compound field to given value.

### HasCompound

`func (o *PdbHeader) HasCompound() bool`

HasCompound returns a boolean if a field has been set.

### GetPdbId

`func (o *PdbHeader) GetPdbId() string`

GetPdbId returns the PdbId field if non-nil, zero value otherwise.

### GetPdbIdOk

`func (o *PdbHeader) GetPdbIdOk() (*string, bool)`

GetPdbIdOk returns a tuple with the PdbId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdbId

`func (o *PdbHeader) SetPdbId(v string)`

SetPdbId sets PdbId field to given value.


### GetSource

`func (o *PdbHeader) GetSource() map[string]interface{}`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PdbHeader) GetSourceOk() (*map[string]interface{}, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PdbHeader) SetSource(v map[string]interface{})`

SetSource sets Source field to given value.

### HasSource

`func (o *PdbHeader) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTitle

`func (o *PdbHeader) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PdbHeader) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PdbHeader) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


