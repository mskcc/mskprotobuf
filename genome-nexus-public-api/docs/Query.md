# Query

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alteration** | Pointer to **string** |  | [optional] 
**AlterationType** | Pointer to **string** |  | [optional] 
**Consequence** | Pointer to **string** |  | [optional] 
**EntrezGeneId** | Pointer to **int32** |  | [optional] 
**Hgvs** | Pointer to **string** |  | [optional] 
**HugoSymbol** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**ProteinEnd** | Pointer to **int32** |  | [optional] 
**ProteinStart** | Pointer to **int32** |  | [optional] 
**SvType** | Pointer to **string** |  | [optional] 
**TumorType** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewQuery

`func NewQuery() *Query`

NewQuery instantiates a new Query object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueryWithDefaults

`func NewQueryWithDefaults() *Query`

NewQueryWithDefaults instantiates a new Query object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlteration

`func (o *Query) GetAlteration() string`

GetAlteration returns the Alteration field if non-nil, zero value otherwise.

### GetAlterationOk

`func (o *Query) GetAlterationOk() (*string, bool)`

GetAlterationOk returns a tuple with the Alteration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlteration

`func (o *Query) SetAlteration(v string)`

SetAlteration sets Alteration field to given value.

### HasAlteration

`func (o *Query) HasAlteration() bool`

HasAlteration returns a boolean if a field has been set.

### GetAlterationType

`func (o *Query) GetAlterationType() string`

GetAlterationType returns the AlterationType field if non-nil, zero value otherwise.

### GetAlterationTypeOk

`func (o *Query) GetAlterationTypeOk() (*string, bool)`

GetAlterationTypeOk returns a tuple with the AlterationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlterationType

`func (o *Query) SetAlterationType(v string)`

SetAlterationType sets AlterationType field to given value.

### HasAlterationType

`func (o *Query) HasAlterationType() bool`

HasAlterationType returns a boolean if a field has been set.

### GetConsequence

`func (o *Query) GetConsequence() string`

GetConsequence returns the Consequence field if non-nil, zero value otherwise.

### GetConsequenceOk

`func (o *Query) GetConsequenceOk() (*string, bool)`

GetConsequenceOk returns a tuple with the Consequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsequence

`func (o *Query) SetConsequence(v string)`

SetConsequence sets Consequence field to given value.

### HasConsequence

`func (o *Query) HasConsequence() bool`

HasConsequence returns a boolean if a field has been set.

### GetEntrezGeneId

`func (o *Query) GetEntrezGeneId() int32`

GetEntrezGeneId returns the EntrezGeneId field if non-nil, zero value otherwise.

### GetEntrezGeneIdOk

`func (o *Query) GetEntrezGeneIdOk() (*int32, bool)`

GetEntrezGeneIdOk returns a tuple with the EntrezGeneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrezGeneId

`func (o *Query) SetEntrezGeneId(v int32)`

SetEntrezGeneId sets EntrezGeneId field to given value.

### HasEntrezGeneId

`func (o *Query) HasEntrezGeneId() bool`

HasEntrezGeneId returns a boolean if a field has been set.

### GetHgvs

`func (o *Query) GetHgvs() string`

GetHgvs returns the Hgvs field if non-nil, zero value otherwise.

### GetHgvsOk

`func (o *Query) GetHgvsOk() (*string, bool)`

GetHgvsOk returns a tuple with the Hgvs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgvs

`func (o *Query) SetHgvs(v string)`

SetHgvs sets Hgvs field to given value.

### HasHgvs

`func (o *Query) HasHgvs() bool`

HasHgvs returns a boolean if a field has been set.

### GetHugoSymbol

`func (o *Query) GetHugoSymbol() string`

GetHugoSymbol returns the HugoSymbol field if non-nil, zero value otherwise.

### GetHugoSymbolOk

`func (o *Query) GetHugoSymbolOk() (*string, bool)`

GetHugoSymbolOk returns a tuple with the HugoSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHugoSymbol

`func (o *Query) SetHugoSymbol(v string)`

SetHugoSymbol sets HugoSymbol field to given value.

### HasHugoSymbol

`func (o *Query) HasHugoSymbol() bool`

HasHugoSymbol returns a boolean if a field has been set.

### GetId

`func (o *Query) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Query) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Query) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Query) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProteinEnd

`func (o *Query) GetProteinEnd() int32`

GetProteinEnd returns the ProteinEnd field if non-nil, zero value otherwise.

### GetProteinEndOk

`func (o *Query) GetProteinEndOk() (*int32, bool)`

GetProteinEndOk returns a tuple with the ProteinEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProteinEnd

`func (o *Query) SetProteinEnd(v int32)`

SetProteinEnd sets ProteinEnd field to given value.

### HasProteinEnd

`func (o *Query) HasProteinEnd() bool`

HasProteinEnd returns a boolean if a field has been set.

### GetProteinStart

`func (o *Query) GetProteinStart() int32`

GetProteinStart returns the ProteinStart field if non-nil, zero value otherwise.

### GetProteinStartOk

`func (o *Query) GetProteinStartOk() (*int32, bool)`

GetProteinStartOk returns a tuple with the ProteinStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProteinStart

`func (o *Query) SetProteinStart(v int32)`

SetProteinStart sets ProteinStart field to given value.

### HasProteinStart

`func (o *Query) HasProteinStart() bool`

HasProteinStart returns a boolean if a field has been set.

### GetSvType

`func (o *Query) GetSvType() string`

GetSvType returns the SvType field if non-nil, zero value otherwise.

### GetSvTypeOk

`func (o *Query) GetSvTypeOk() (*string, bool)`

GetSvTypeOk returns a tuple with the SvType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSvType

`func (o *Query) SetSvType(v string)`

SetSvType sets SvType field to given value.

### HasSvType

`func (o *Query) HasSvType() bool`

HasSvType returns a boolean if a field has been set.

### GetTumorType

`func (o *Query) GetTumorType() string`

GetTumorType returns the TumorType field if non-nil, zero value otherwise.

### GetTumorTypeOk

`func (o *Query) GetTumorTypeOk() (*string, bool)`

GetTumorTypeOk returns a tuple with the TumorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTumorType

`func (o *Query) SetTumorType(v string)`

SetTumorType sets TumorType field to given value.

### HasTumorType

`func (o *Query) HasTumorType() bool`

HasTumorType returns a boolean if a field has been set.

### GetType

`func (o *Query) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Query) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Query) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Query) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


