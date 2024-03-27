# MutationAssessorAnnotation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotation** | Pointer to [**MutationAssessor**](MutationAssessor.md) |  | [optional] 
**License** | Pointer to **string** |  | [optional] 

## Methods

### NewMutationAssessorAnnotation

`func NewMutationAssessorAnnotation() *MutationAssessorAnnotation`

NewMutationAssessorAnnotation instantiates a new MutationAssessorAnnotation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutationAssessorAnnotationWithDefaults

`func NewMutationAssessorAnnotationWithDefaults() *MutationAssessorAnnotation`

NewMutationAssessorAnnotationWithDefaults instantiates a new MutationAssessorAnnotation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotation

`func (o *MutationAssessorAnnotation) GetAnnotation() MutationAssessor`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *MutationAssessorAnnotation) GetAnnotationOk() (*MutationAssessor, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotation

`func (o *MutationAssessorAnnotation) SetAnnotation(v MutationAssessor)`

SetAnnotation sets Annotation field to given value.

### HasAnnotation

`func (o *MutationAssessorAnnotation) HasAnnotation() bool`

HasAnnotation returns a boolean if a field has been set.

### GetLicense

`func (o *MutationAssessorAnnotation) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *MutationAssessorAnnotation) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *MutationAssessorAnnotation) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *MutationAssessorAnnotation) HasLicense() bool`

HasLicense returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


