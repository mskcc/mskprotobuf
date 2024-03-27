# NucleotideContextAnnotation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotation** | Pointer to [**NucleotideContext**](NucleotideContext.md) |  | [optional] 
**License** | Pointer to **string** |  | [optional] 

## Methods

### NewNucleotideContextAnnotation

`func NewNucleotideContextAnnotation() *NucleotideContextAnnotation`

NewNucleotideContextAnnotation instantiates a new NucleotideContextAnnotation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNucleotideContextAnnotationWithDefaults

`func NewNucleotideContextAnnotationWithDefaults() *NucleotideContextAnnotation`

NewNucleotideContextAnnotationWithDefaults instantiates a new NucleotideContextAnnotation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotation

`func (o *NucleotideContextAnnotation) GetAnnotation() NucleotideContext`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *NucleotideContextAnnotation) GetAnnotationOk() (*NucleotideContext, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotation

`func (o *NucleotideContextAnnotation) SetAnnotation(v NucleotideContext)`

SetAnnotation sets Annotation field to given value.

### HasAnnotation

`func (o *NucleotideContextAnnotation) HasAnnotation() bool`

HasAnnotation returns a boolean if a field has been set.

### GetLicense

`func (o *NucleotideContextAnnotation) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *NucleotideContextAnnotation) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *NucleotideContextAnnotation) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *NucleotideContextAnnotation) HasLicense() bool`

HasLicense returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


