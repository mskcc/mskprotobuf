# TumorType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Children** | Pointer to [**map[string]TumorType**](TumorType.md) |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Level** | Pointer to **int32** |  | [optional] 
**MainType** | Pointer to [**MainType**](MainType.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Parent** | Pointer to **string** |  | [optional] 
**Tissue** | Pointer to **string** |  | [optional] 
**TumorForm** | Pointer to **string** |  | [optional] 

## Methods

### NewTumorType

`func NewTumorType() *TumorType`

NewTumorType instantiates a new TumorType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTumorTypeWithDefaults

`func NewTumorTypeWithDefaults() *TumorType`

NewTumorTypeWithDefaults instantiates a new TumorType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildren

`func (o *TumorType) GetChildren() map[string]TumorType`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *TumorType) GetChildrenOk() (*map[string]TumorType, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *TumorType) SetChildren(v map[string]TumorType)`

SetChildren sets Children field to given value.

### HasChildren

`func (o *TumorType) HasChildren() bool`

HasChildren returns a boolean if a field has been set.

### GetCode

`func (o *TumorType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TumorType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TumorType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *TumorType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetColor

`func (o *TumorType) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *TumorType) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *TumorType) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *TumorType) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetId

`func (o *TumorType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TumorType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TumorType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TumorType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLevel

`func (o *TumorType) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *TumorType) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *TumorType) SetLevel(v int32)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *TumorType) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetMainType

`func (o *TumorType) GetMainType() MainType`

GetMainType returns the MainType field if non-nil, zero value otherwise.

### GetMainTypeOk

`func (o *TumorType) GetMainTypeOk() (*MainType, bool)`

GetMainTypeOk returns a tuple with the MainType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainType

`func (o *TumorType) SetMainType(v MainType)`

SetMainType sets MainType field to given value.

### HasMainType

`func (o *TumorType) HasMainType() bool`

HasMainType returns a boolean if a field has been set.

### GetName

`func (o *TumorType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TumorType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TumorType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TumorType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParent

`func (o *TumorType) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *TumorType) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *TumorType) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *TumorType) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetTissue

`func (o *TumorType) GetTissue() string`

GetTissue returns the Tissue field if non-nil, zero value otherwise.

### GetTissueOk

`func (o *TumorType) GetTissueOk() (*string, bool)`

GetTissueOk returns a tuple with the Tissue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTissue

`func (o *TumorType) SetTissue(v string)`

SetTissue sets Tissue field to given value.

### HasTissue

`func (o *TumorType) HasTissue() bool`

HasTissue returns a boolean if a field has been set.

### GetTumorForm

`func (o *TumorType) GetTumorForm() string`

GetTumorForm returns the TumorForm field if non-nil, zero value otherwise.

### GetTumorFormOk

`func (o *TumorType) GetTumorFormOk() (*string, bool)`

GetTumorFormOk returns a tuple with the TumorForm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTumorForm

`func (o *TumorType) SetTumorForm(v string)`

SetTumorForm sets TumorForm field to given value.

### HasTumorForm

`func (o *TumorType) HasTumorForm() bool`

HasTumorForm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


