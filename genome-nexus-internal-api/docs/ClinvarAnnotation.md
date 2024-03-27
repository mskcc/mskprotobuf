# ClinvarAnnotation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Annotation** | Pointer to [**Clinvar**](Clinvar.md) |  | [optional] 

## Methods

### NewClinvarAnnotation

`func NewClinvarAnnotation() *ClinvarAnnotation`

NewClinvarAnnotation instantiates a new ClinvarAnnotation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClinvarAnnotationWithDefaults

`func NewClinvarAnnotationWithDefaults() *ClinvarAnnotation`

NewClinvarAnnotationWithDefaults instantiates a new ClinvarAnnotation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnotation

`func (o *ClinvarAnnotation) GetAnnotation() Clinvar`

GetAnnotation returns the Annotation field if non-nil, zero value otherwise.

### GetAnnotationOk

`func (o *ClinvarAnnotation) GetAnnotationOk() (*Clinvar, bool)`

GetAnnotationOk returns a tuple with the Annotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotation

`func (o *ClinvarAnnotation) SetAnnotation(v Clinvar)`

SetAnnotation sets Annotation field to given value.

### HasAnnotation

`func (o *ClinvarAnnotation) HasAnnotation() bool`

HasAnnotation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


