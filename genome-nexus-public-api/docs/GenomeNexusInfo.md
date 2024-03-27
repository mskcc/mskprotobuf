# GenomeNexusInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Database** | Pointer to [**Version**](Version.md) |  | [optional] 
**Server** | Pointer to [**Version**](Version.md) |  | [optional] 

## Methods

### NewGenomeNexusInfo

`func NewGenomeNexusInfo() *GenomeNexusInfo`

NewGenomeNexusInfo instantiates a new GenomeNexusInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenomeNexusInfoWithDefaults

`func NewGenomeNexusInfoWithDefaults() *GenomeNexusInfo`

NewGenomeNexusInfoWithDefaults instantiates a new GenomeNexusInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatabase

`func (o *GenomeNexusInfo) GetDatabase() Version`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *GenomeNexusInfo) GetDatabaseOk() (*Version, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *GenomeNexusInfo) SetDatabase(v Version)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *GenomeNexusInfo) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetServer

`func (o *GenomeNexusInfo) GetServer() Version`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *GenomeNexusInfo) GetServerOk() (*Version, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *GenomeNexusInfo) SetServer(v Version)`

SetServer sets Server field to given value.

### HasServer

`func (o *GenomeNexusInfo) HasServer() bool`

HasServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


