# IndicatorQueryTreatment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Abstracts** | Pointer to [**[]ArticleAbstract**](ArticleAbstract.md) |  | [optional] 
**Alterations** | Pointer to **[]string** |  | [optional] 
**ApprovedIndications** | Pointer to **[]string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Drugs** | Pointer to [**[]Drug**](Drug.md) |  | [optional] 
**FdaApproved** | Pointer to **bool** |  | [optional] 
**Level** | Pointer to **string** |  | [optional] 
**LevelAssociatedCancerType** | Pointer to [**TumorType**](TumorType.md) |  | [optional] 
**Pmids** | Pointer to **[]string** |  | [optional] 

## Methods

### NewIndicatorQueryTreatment

`func NewIndicatorQueryTreatment() *IndicatorQueryTreatment`

NewIndicatorQueryTreatment instantiates a new IndicatorQueryTreatment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIndicatorQueryTreatmentWithDefaults

`func NewIndicatorQueryTreatmentWithDefaults() *IndicatorQueryTreatment`

NewIndicatorQueryTreatmentWithDefaults instantiates a new IndicatorQueryTreatment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbstracts

`func (o *IndicatorQueryTreatment) GetAbstracts() []ArticleAbstract`

GetAbstracts returns the Abstracts field if non-nil, zero value otherwise.

### GetAbstractsOk

`func (o *IndicatorQueryTreatment) GetAbstractsOk() (*[]ArticleAbstract, bool)`

GetAbstractsOk returns a tuple with the Abstracts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbstracts

`func (o *IndicatorQueryTreatment) SetAbstracts(v []ArticleAbstract)`

SetAbstracts sets Abstracts field to given value.

### HasAbstracts

`func (o *IndicatorQueryTreatment) HasAbstracts() bool`

HasAbstracts returns a boolean if a field has been set.

### GetAlterations

`func (o *IndicatorQueryTreatment) GetAlterations() []string`

GetAlterations returns the Alterations field if non-nil, zero value otherwise.

### GetAlterationsOk

`func (o *IndicatorQueryTreatment) GetAlterationsOk() (*[]string, bool)`

GetAlterationsOk returns a tuple with the Alterations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlterations

`func (o *IndicatorQueryTreatment) SetAlterations(v []string)`

SetAlterations sets Alterations field to given value.

### HasAlterations

`func (o *IndicatorQueryTreatment) HasAlterations() bool`

HasAlterations returns a boolean if a field has been set.

### GetApprovedIndications

`func (o *IndicatorQueryTreatment) GetApprovedIndications() []string`

GetApprovedIndications returns the ApprovedIndications field if non-nil, zero value otherwise.

### GetApprovedIndicationsOk

`func (o *IndicatorQueryTreatment) GetApprovedIndicationsOk() (*[]string, bool)`

GetApprovedIndicationsOk returns a tuple with the ApprovedIndications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedIndications

`func (o *IndicatorQueryTreatment) SetApprovedIndications(v []string)`

SetApprovedIndications sets ApprovedIndications field to given value.

### HasApprovedIndications

`func (o *IndicatorQueryTreatment) HasApprovedIndications() bool`

HasApprovedIndications returns a boolean if a field has been set.

### GetDescription

`func (o *IndicatorQueryTreatment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IndicatorQueryTreatment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IndicatorQueryTreatment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IndicatorQueryTreatment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDrugs

`func (o *IndicatorQueryTreatment) GetDrugs() []Drug`

GetDrugs returns the Drugs field if non-nil, zero value otherwise.

### GetDrugsOk

`func (o *IndicatorQueryTreatment) GetDrugsOk() (*[]Drug, bool)`

GetDrugsOk returns a tuple with the Drugs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrugs

`func (o *IndicatorQueryTreatment) SetDrugs(v []Drug)`

SetDrugs sets Drugs field to given value.

### HasDrugs

`func (o *IndicatorQueryTreatment) HasDrugs() bool`

HasDrugs returns a boolean if a field has been set.

### GetFdaApproved

`func (o *IndicatorQueryTreatment) GetFdaApproved() bool`

GetFdaApproved returns the FdaApproved field if non-nil, zero value otherwise.

### GetFdaApprovedOk

`func (o *IndicatorQueryTreatment) GetFdaApprovedOk() (*bool, bool)`

GetFdaApprovedOk returns a tuple with the FdaApproved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFdaApproved

`func (o *IndicatorQueryTreatment) SetFdaApproved(v bool)`

SetFdaApproved sets FdaApproved field to given value.

### HasFdaApproved

`func (o *IndicatorQueryTreatment) HasFdaApproved() bool`

HasFdaApproved returns a boolean if a field has been set.

### GetLevel

`func (o *IndicatorQueryTreatment) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *IndicatorQueryTreatment) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *IndicatorQueryTreatment) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *IndicatorQueryTreatment) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetLevelAssociatedCancerType

`func (o *IndicatorQueryTreatment) GetLevelAssociatedCancerType() TumorType`

GetLevelAssociatedCancerType returns the LevelAssociatedCancerType field if non-nil, zero value otherwise.

### GetLevelAssociatedCancerTypeOk

`func (o *IndicatorQueryTreatment) GetLevelAssociatedCancerTypeOk() (*TumorType, bool)`

GetLevelAssociatedCancerTypeOk returns a tuple with the LevelAssociatedCancerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevelAssociatedCancerType

`func (o *IndicatorQueryTreatment) SetLevelAssociatedCancerType(v TumorType)`

SetLevelAssociatedCancerType sets LevelAssociatedCancerType field to given value.

### HasLevelAssociatedCancerType

`func (o *IndicatorQueryTreatment) HasLevelAssociatedCancerType() bool`

HasLevelAssociatedCancerType returns a boolean if a field has been set.

### GetPmids

`func (o *IndicatorQueryTreatment) GetPmids() []string`

GetPmids returns the Pmids field if non-nil, zero value otherwise.

### GetPmidsOk

`func (o *IndicatorQueryTreatment) GetPmidsOk() (*[]string, bool)`

GetPmidsOk returns a tuple with the Pmids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPmids

`func (o *IndicatorQueryTreatment) SetPmids(v []string)`

SetPmids sets Pmids field to given value.

### HasPmids

`func (o *IndicatorQueryTreatment) HasPmids() bool`

HasPmids returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


