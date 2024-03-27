# Gene

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GeneId** | Pointer to **string** | geneid | [optional] 
**Symbol** | Pointer to **string** | symbol | [optional] 

## Methods

### NewGene

`func NewGene() *Gene`

NewGene instantiates a new Gene object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneWithDefaults

`func NewGeneWithDefaults() *Gene`

NewGeneWithDefaults instantiates a new Gene object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGeneId

`func (o *Gene) GetGeneId() string`

GetGeneId returns the GeneId field if non-nil, zero value otherwise.

### GetGeneIdOk

`func (o *Gene) GetGeneIdOk() (*string, bool)`

GetGeneIdOk returns a tuple with the GeneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneId

`func (o *Gene) SetGeneId(v string)`

SetGeneId sets GeneId field to given value.

### HasGeneId

`func (o *Gene) HasGeneId() bool`

HasGeneId returns a boolean if a field has been set.

### GetSymbol

`func (o *Gene) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Gene) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Gene) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Gene) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


