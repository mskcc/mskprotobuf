# HotspotAnnotation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotation** | Pointer to [**[][]Hotspot**]([]Hotspot.md) |  | [optional] 
**License** | Pointer to **string** |  | [optional] 

## Methods

### NewHotspotAnnotation

`func NewHotspotAnnotation() *HotspotAnnotation`

NewHotspotAnnotation instantiates a new HotspotAnnotation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHotspotAnnotationWithDefaults

`func NewHotspotAnnotationWithDefaults() *HotspotAnnotation`

NewHotspotAnnotationWithDefaults instantiates a new HotspotAnnotation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotation

`func (o *HotspotAnnotation) GetAnnotation() [][]Hotspot`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *HotspotAnnotation) GetAnnotationOk() (*[][]Hotspot, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotation

`func (o *HotspotAnnotation) SetAnnotation(v [][]Hotspot)`

SetAnnotation sets Annotation field to given value.

### HasAnnotation

`func (o *HotspotAnnotation) HasAnnotation() bool`

HasAnnotation returns a boolean if a field has been set.

### GetLicense

`func (o *HotspotAnnotation) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *HotspotAnnotation) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *HotspotAnnotation) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *HotspotAnnotation) HasLicense() bool`

HasLicense returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


