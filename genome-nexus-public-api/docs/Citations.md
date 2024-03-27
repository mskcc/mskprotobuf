# Citations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Abstracts** | Pointer to [**[]ArticleAbstract**](ArticleAbstract.md) |  | [optional] 
**Pmids** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCitations

`func NewCitations() *Citations`

NewCitations instantiates a new Citations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCitationsWithDefaults

`func NewCitationsWithDefaults() *Citations`

NewCitationsWithDefaults instantiates a new Citations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbstracts

`func (o *Citations) GetAbstracts() []ArticleAbstract`

GetAbstracts returns the Abstracts field if non-nil, zero value otherwise.

### GetAbstractsOk

`func (o *Citations) GetAbstractsOk() (*[]ArticleAbstract, bool)`

GetAbstractsOk returns a tuple with the Abstracts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstracts

`func (o *Citations) SetAbstracts(v []ArticleAbstract)`

SetAbstracts sets Abstracts field to given value.

### HasAbstracts

`func (o *Citations) HasAbstracts() bool`

HasAbstracts returns a boolean if a field has been set.

### GetPmids

`func (o *Citations) GetPmids() []string`

GetPmids returns the Pmids field if non-nil, zero value otherwise.

### GetPmidsOk

`func (o *Citations) GetPmidsOk() (*[]string, bool)`

GetPmidsOk returns a tuple with the Pmids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmids

`func (o *Citations) SetPmids(v []string)`

SetPmids sets Pmids field to given value.

### HasPmids

`func (o *Citations) HasPmids() bool`

HasPmids returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


