# VEPInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cache** | Pointer to [**Version**](Version.md) |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Server** | Pointer to [**Version**](Version.md) |  | [optional] 

## Methods

### NewVEPInfo

`func NewVEPInfo() *VEPInfo`

NewVEPInfo instantiates a new VEPInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVEPInfoWithDefaults

`func NewVEPInfoWithDefaults() *VEPInfo`

NewVEPInfoWithDefaults instantiates a new VEPInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCache

`func (o *VEPInfo) GetCache() Version`

GetCache returns the Cache field if non-nil, zero value otherwise.

### GetCacheOk

`func (o *VEPInfo) GetCacheOk() (*Version, bool)`

GetCacheOk returns a tuple with the Cache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCache

`func (o *VEPInfo) SetCache(v Version)`

SetCache sets Cache field to given value.

### HasCache

`func (o *VEPInfo) HasCache() bool`

HasCache returns a boolean if a field has been set.

### GetComment

`func (o *VEPInfo) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *VEPInfo) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *VEPInfo) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *VEPInfo) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetServer

`func (o *VEPInfo) GetServer() Version`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *VEPInfo) GetServerOk() (*Version, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *VEPInfo) SetServer(v Version)`

SetServer sets Server field to given value.

### HasServer

`func (o *VEPInfo) HasServer() bool`

HasServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


