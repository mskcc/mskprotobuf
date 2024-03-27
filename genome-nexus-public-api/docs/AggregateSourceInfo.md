# AggregateSourceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnnotationSourcesInfo** | Pointer to [**[]SourceVersionInfo**](SourceVersionInfo.md) |  | [optional] 
**GenomeNexus** | Pointer to [**GenomeNexusInfo**](GenomeNexusInfo.md) |  | [optional] 
**Vep** | Pointer to [**VEPInfo**](VEPInfo.md) |  | [optional] 

## Methods

### NewAggregateSourceInfo

`func NewAggregateSourceInfo() *AggregateSourceInfo`

NewAggregateSourceInfo instantiates a new AggregateSourceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAggregateSourceInfoWithDefaults

`func NewAggregateSourceInfoWithDefaults() *AggregateSourceInfo`

NewAggregateSourceInfoWithDefaults instantiates a new AggregateSourceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotationSourcesInfo

`func (o *AggregateSourceInfo) GetAnnotationSourcesInfo() []SourceVersionInfo`

GetAnnotationSourcesInfo returns the AnnotationSourcesInfo field if non-nil, zero value otherwise.

### GetAnnotationSourcesInfoOk

`func (o *AggregateSourceInfo) GetAnnotationSourcesInfoOk() (*[]SourceVersionInfo, bool)`

GetAnnotationSourcesInfoOk returns a tuple with the AnnotationSourcesInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotationSourcesInfo

`func (o *AggregateSourceInfo) SetAnnotationSourcesInfo(v []SourceVersionInfo)`

SetAnnotationSourcesInfo sets AnnotationSourcesInfo field to given value.

### HasAnnotationSourcesInfo

`func (o *AggregateSourceInfo) HasAnnotationSourcesInfo() bool`

HasAnnotationSourcesInfo returns a boolean if a field has been set.

### GetGenomeNexus

`func (o *AggregateSourceInfo) GetGenomeNexus() GenomeNexusInfo`

GetGenomeNexus returns the GenomeNexus field if non-nil, zero value otherwise.

### GetGenomeNexusOk

`func (o *AggregateSourceInfo) GetGenomeNexusOk() (*GenomeNexusInfo, bool)`

GetGenomeNexusOk returns a tuple with the GenomeNexus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenomeNexus

`func (o *AggregateSourceInfo) SetGenomeNexus(v GenomeNexusInfo)`

SetGenomeNexus sets GenomeNexus field to given value.

### HasGenomeNexus

`func (o *AggregateSourceInfo) HasGenomeNexus() bool`

HasGenomeNexus returns a boolean if a field has been set.

### GetVep

`func (o *AggregateSourceInfo) GetVep() VEPInfo`

GetVep returns the Vep field if non-nil, zero value otherwise.

### GetVepOk

`func (o *AggregateSourceInfo) GetVepOk() (*VEPInfo, bool)`

GetVepOk returns a tuple with the Vep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVep

`func (o *AggregateSourceInfo) SetVep(v VEPInfo)`

SetVep sets Vep field to given value.

### HasVep

`func (o *AggregateSourceInfo) HasVep() bool`

HasVep returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


