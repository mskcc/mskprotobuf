# PfamDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | PFAM domain description | [optional] 
**Name** | **string** | PFAM domain name | 
**PfamAccession** | **string** | PFAM domain accession | 

## Methods

### NewPfamDomain

`func NewPfamDomain(name string, pfamAccession string, ) *PfamDomain`

NewPfamDomain instantiates a new PfamDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPfamDomainWithDefaults

`func NewPfamDomainWithDefaults() *PfamDomain`

NewPfamDomainWithDefaults instantiates a new PfamDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *PfamDomain) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PfamDomain) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PfamDomain) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PfamDomain) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *PfamDomain) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PfamDomain) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PfamDomain) SetName(v string)`

SetName sets Name field to given value.


### GetPfamAccession

`func (o *PfamDomain) GetPfamAccession() string`

GetPfamAccession returns the PfamAccession field if non-nil, zero value otherwise.

### GetPfamAccessionOk

`func (o *PfamDomain) GetPfamAccessionOk() (*string, bool)`

GetPfamAccessionOk returns a tuple with the PfamAccession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfamAccession

`func (o *PfamDomain) SetPfamAccession(v string)`

SetPfamAccession sets PfamAccession field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


