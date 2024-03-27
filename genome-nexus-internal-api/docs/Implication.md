# Implication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alterations** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**LevelOfEvidence** | Pointer to **string** |  | [optional] 
**TumorType** | Pointer to [**TumorType**](TumorType.md) |  | [optional] 

## Methods

### NewImplication

`func NewImplication() *Implication`

NewImplication instantiates a new Implication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewImplicationWithDefaults

`func NewImplicationWithDefaults() *Implication`

NewImplicationWithDefaults instantiates a new Implication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlterations

`func (o *Implication) GetAlterations() []string`

GetAlterations returns the Alterations field if non-nil, zero value otherwise.

### GetAlterationsOk

`func (o *Implication) GetAlterationsOk() (*[]string, bool)`

GetAlterationsOk returns a tuple with the Alterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlterations

`func (o *Implication) SetAlterations(v []string)`

SetAlterations sets Alterations field to given value.

### HasAlterations

`func (o *Implication) HasAlterations() bool`

HasAlterations returns a boolean if a field has been set.

### GetDescription

`func (o *Implication) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Implication) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Implication) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Implication) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLevelOfEvidence

`func (o *Implication) GetLevelOfEvidence() string`

GetLevelOfEvidence returns the LevelOfEvidence field if non-nil, zero value otherwise.

### GetLevelOfEvidenceOk

`func (o *Implication) GetLevelOfEvidenceOk() (*string, bool)`

GetLevelOfEvidenceOk returns a tuple with the LevelOfEvidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelOfEvidence

`func (o *Implication) SetLevelOfEvidence(v string)`

SetLevelOfEvidence sets LevelOfEvidence field to given value.

### HasLevelOfEvidence

`func (o *Implication) HasLevelOfEvidence() bool`

HasLevelOfEvidence returns a boolean if a field has been set.

### GetTumorType

`func (o *Implication) GetTumorType() TumorType`

GetTumorType returns the TumorType field if non-nil, zero value otherwise.

### GetTumorTypeOk

`func (o *Implication) GetTumorTypeOk() (*TumorType, bool)`

GetTumorTypeOk returns a tuple with the TumorType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTumorType

`func (o *Implication) SetTumorType(v TumorType)`

SetTumorType sets TumorType field to given value.

### HasTumorType

`func (o *Implication) HasTumorType() bool`

HasTumorType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


