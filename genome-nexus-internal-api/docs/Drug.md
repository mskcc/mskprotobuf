# Drug

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DrugName** | Pointer to **string** |  | [optional] 
**NcitCode** | Pointer to **string** |  | [optional] 
**Synonyms** | Pointer to **[]string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 

## Methods

### NewDrug

`func NewDrug() *Drug`

NewDrug instantiates a new Drug object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDrugWithDefaults

`func NewDrugWithDefaults() *Drug`

NewDrugWithDefaults instantiates a new Drug object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDrugName

`func (o *Drug) GetDrugName() string`

GetDrugName returns the DrugName field if non-nil, zero value otherwise.

### GetDrugNameOk

`func (o *Drug) GetDrugNameOk() (*string, bool)`

GetDrugNameOk returns a tuple with the DrugName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugName

`func (o *Drug) SetDrugName(v string)`

SetDrugName sets DrugName field to given value.

### HasDrugName

`func (o *Drug) HasDrugName() bool`

HasDrugName returns a boolean if a field has been set.

### GetNcitCode

`func (o *Drug) GetNcitCode() string`

GetNcitCode returns the NcitCode field if non-nil, zero value otherwise.

### GetNcitCodeOk

`func (o *Drug) GetNcitCodeOk() (*string, bool)`

GetNcitCodeOk returns a tuple with the NcitCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNcitCode

`func (o *Drug) SetNcitCode(v string)`

SetNcitCode sets NcitCode field to given value.

### HasNcitCode

`func (o *Drug) HasNcitCode() bool`

HasNcitCode returns a boolean if a field has been set.

### GetSynonyms

`func (o *Drug) GetSynonyms() []string`

GetSynonyms returns the Synonyms field if non-nil, zero value otherwise.

### GetSynonymsOk

`func (o *Drug) GetSynonymsOk() (*[]string, bool)`

GetSynonymsOk returns a tuple with the Synonyms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSynonyms

`func (o *Drug) SetSynonyms(v []string)`

SetSynonyms sets Synonyms field to given value.

### HasSynonyms

`func (o *Drug) HasSynonyms() bool`

HasSynonyms returns a boolean if a field has been set.

### GetUuid

`func (o *Drug) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Drug) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Drug) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Drug) HasUuid() bool`

HasUuid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


