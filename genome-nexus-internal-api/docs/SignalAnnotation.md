# SignalAnnotation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotation** | Pointer to [**[]SignalMutation**](SignalMutation.md) |  | [optional] 
**License** | Pointer to **string** |  | [optional] 

## Methods

### NewSignalAnnotation

`func NewSignalAnnotation() *SignalAnnotation`

NewSignalAnnotation instantiates a new SignalAnnotation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignalAnnotationWithDefaults

`func NewSignalAnnotationWithDefaults() *SignalAnnotation`

NewSignalAnnotationWithDefaults instantiates a new SignalAnnotation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotation

`func (o *SignalAnnotation) GetAnnotation() []SignalMutation`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *SignalAnnotation) GetAnnotationOk() (*[]SignalMutation, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotation

`func (o *SignalAnnotation) SetAnnotation(v []SignalMutation)`

SetAnnotation sets Annotation field to given value.

### HasAnnotation

`func (o *SignalAnnotation) HasAnnotation() bool`

HasAnnotation returns a boolean if a field has been set.

### GetLicense

`func (o *SignalAnnotation) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *SignalAnnotation) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *SignalAnnotation) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *SignalAnnotation) HasLicense() bool`

HasLicense returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


