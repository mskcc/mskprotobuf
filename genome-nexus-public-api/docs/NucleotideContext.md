# NucleotideContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hgvs** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Molecule** | Pointer to **string** |  | [optional] 
**Query** | Pointer to **string** |  | [optional] 
**Seq** | **string** | Nucleotide context sequence | 

## Methods

### NewNucleotideContext

`func NewNucleotideContext(seq string, ) *NucleotideContext`

NewNucleotideContext instantiates a new NucleotideContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNucleotideContextWithDefaults

`func NewNucleotideContextWithDefaults() *NucleotideContext`

NewNucleotideContextWithDefaults instantiates a new NucleotideContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHgvs

`func (o *NucleotideContext) GetHgvs() string`

GetHgvs returns the Hgvs field if non-nil, zero value otherwise.

### GetHgvsOk

`func (o *NucleotideContext) GetHgvsOk() (*string, bool)`

GetHgvsOk returns a tuple with the Hgvs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgvs

`func (o *NucleotideContext) SetHgvs(v string)`

SetHgvs sets Hgvs field to given value.

### HasHgvs

`func (o *NucleotideContext) HasHgvs() bool`

HasHgvs returns a boolean if a field has been set.

### GetId

`func (o *NucleotideContext) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NucleotideContext) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NucleotideContext) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NucleotideContext) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMolecule

`func (o *NucleotideContext) GetMolecule() string`

GetMolecule returns the Molecule field if non-nil, zero value otherwise.

### GetMoleculeOk

`func (o *NucleotideContext) GetMoleculeOk() (*string, bool)`

GetMoleculeOk returns a tuple with the Molecule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMolecule

`func (o *NucleotideContext) SetMolecule(v string)`

SetMolecule sets Molecule field to given value.

### HasMolecule

`func (o *NucleotideContext) HasMolecule() bool`

HasMolecule returns a boolean if a field has been set.

### GetQuery

`func (o *NucleotideContext) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *NucleotideContext) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *NucleotideContext) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *NucleotideContext) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetSeq

`func (o *NucleotideContext) GetSeq() string`

GetSeq returns the Seq field if non-nil, zero value otherwise.

### GetSeqOk

`func (o *NucleotideContext) GetSeqOk() (*string, bool)`

GetSeqOk returns a tuple with the Seq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeq

`func (o *NucleotideContext) SetSeq(v string)`

SetSeq sets Seq field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


