# Dbsnp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to **string** | class | [optional] 
**AlleleOrigin** | Pointer to **string** | allele_origin | [optional] 
**Alleles** | Pointer to [**[]Alleles**](Alleles.md) | alleles | [optional] 
**Alt** | Pointer to **string** | alt | [optional] 
**Chrom** | Pointer to **string** | chrom | [optional] 
**DbsnpBuild** | Pointer to **int32** | dbsnp_build | [optional] 
**Flags** | Pointer to **[]string** | flags | [optional] 
**Hg19** | Pointer to [**Hg19**](Hg19.md) |  | [optional] 
**License** | Pointer to **string** | _license | [optional] 
**Ref** | Pointer to **string** | ref | [optional] 
**Rsid** | Pointer to **string** | rsid | [optional] 
**Validated** | Pointer to **bool** | validated | [optional] 
**VarSubtype** | Pointer to **string** | var_subtype | [optional] 
**Vartype** | Pointer to **string** | vartype | [optional] 

## Methods

### NewDbsnp

`func NewDbsnp() *Dbsnp`

NewDbsnp instantiates a new Dbsnp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDbsnpWithDefaults

`func NewDbsnpWithDefaults() *Dbsnp`

NewDbsnpWithDefaults instantiates a new Dbsnp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *Dbsnp) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *Dbsnp) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *Dbsnp) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *Dbsnp) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetAlleleOrigin

`func (o *Dbsnp) GetAlleleOrigin() string`

GetAlleleOrigin returns the AlleleOrigin field if non-nil, zero value otherwise.

### GetAlleleOriginOk

`func (o *Dbsnp) GetAlleleOriginOk() (*string, bool)`

GetAlleleOriginOk returns a tuple with the AlleleOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlleleOrigin

`func (o *Dbsnp) SetAlleleOrigin(v string)`

SetAlleleOrigin sets AlleleOrigin field to given value.

### HasAlleleOrigin

`func (o *Dbsnp) HasAlleleOrigin() bool`

HasAlleleOrigin returns a boolean if a field has been set.

### GetAlleles

`func (o *Dbsnp) GetAlleles() []Alleles`

GetAlleles returns the Alleles field if non-nil, zero value otherwise.

### GetAllelesOk

`func (o *Dbsnp) GetAllelesOk() (*[]Alleles, bool)`

GetAllelesOk returns a tuple with the Alleles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlleles

`func (o *Dbsnp) SetAlleles(v []Alleles)`

SetAlleles sets Alleles field to given value.

### HasAlleles

`func (o *Dbsnp) HasAlleles() bool`

HasAlleles returns a boolean if a field has been set.

### GetAlt

`func (o *Dbsnp) GetAlt() string`

GetAlt returns the Alt field if non-nil, zero value otherwise.

### GetAltOk

`func (o *Dbsnp) GetAltOk() (*string, bool)`

GetAltOk returns a tuple with the Alt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlt

`func (o *Dbsnp) SetAlt(v string)`

SetAlt sets Alt field to given value.

### HasAlt

`func (o *Dbsnp) HasAlt() bool`

HasAlt returns a boolean if a field has been set.

### GetChrom

`func (o *Dbsnp) GetChrom() string`

GetChrom returns the Chrom field if non-nil, zero value otherwise.

### GetChromOk

`func (o *Dbsnp) GetChromOk() (*string, bool)`

GetChromOk returns a tuple with the Chrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChrom

`func (o *Dbsnp) SetChrom(v string)`

SetChrom sets Chrom field to given value.

### HasChrom

`func (o *Dbsnp) HasChrom() bool`

HasChrom returns a boolean if a field has been set.

### GetDbsnpBuild

`func (o *Dbsnp) GetDbsnpBuild() int32`

GetDbsnpBuild returns the DbsnpBuild field if non-nil, zero value otherwise.

### GetDbsnpBuildOk

`func (o *Dbsnp) GetDbsnpBuildOk() (*int32, bool)`

GetDbsnpBuildOk returns a tuple with the DbsnpBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbsnpBuild

`func (o *Dbsnp) SetDbsnpBuild(v int32)`

SetDbsnpBuild sets DbsnpBuild field to given value.

### HasDbsnpBuild

`func (o *Dbsnp) HasDbsnpBuild() bool`

HasDbsnpBuild returns a boolean if a field has been set.

### GetFlags

`func (o *Dbsnp) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *Dbsnp) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *Dbsnp) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *Dbsnp) HasFlags() bool`

HasFlags returns a boolean if a field has been set.

### GetHg19

`func (o *Dbsnp) GetHg19() Hg19`

GetHg19 returns the Hg19 field if non-nil, zero value otherwise.

### GetHg19Ok

`func (o *Dbsnp) GetHg19Ok() (*Hg19, bool)`

GetHg19Ok returns a tuple with the Hg19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHg19

`func (o *Dbsnp) SetHg19(v Hg19)`

SetHg19 sets Hg19 field to given value.

### HasHg19

`func (o *Dbsnp) HasHg19() bool`

HasHg19 returns a boolean if a field has been set.

### GetLicense

`func (o *Dbsnp) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *Dbsnp) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *Dbsnp) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *Dbsnp) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetRef

`func (o *Dbsnp) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *Dbsnp) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *Dbsnp) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *Dbsnp) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetRsid

`func (o *Dbsnp) GetRsid() string`

GetRsid returns the Rsid field if non-nil, zero value otherwise.

### GetRsidOk

`func (o *Dbsnp) GetRsidOk() (*string, bool)`

GetRsidOk returns a tuple with the Rsid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRsid

`func (o *Dbsnp) SetRsid(v string)`

SetRsid sets Rsid field to given value.

### HasRsid

`func (o *Dbsnp) HasRsid() bool`

HasRsid returns a boolean if a field has been set.

### GetValidated

`func (o *Dbsnp) GetValidated() bool`

GetValidated returns the Validated field if non-nil, zero value otherwise.

### GetValidatedOk

`func (o *Dbsnp) GetValidatedOk() (*bool, bool)`

GetValidatedOk returns a tuple with the Validated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidated

`func (o *Dbsnp) SetValidated(v bool)`

SetValidated sets Validated field to given value.

### HasValidated

`func (o *Dbsnp) HasValidated() bool`

HasValidated returns a boolean if a field has been set.

### GetVarSubtype

`func (o *Dbsnp) GetVarSubtype() string`

GetVarSubtype returns the VarSubtype field if non-nil, zero value otherwise.

### GetVarSubtypeOk

`func (o *Dbsnp) GetVarSubtypeOk() (*string, bool)`

GetVarSubtypeOk returns a tuple with the VarSubtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVarSubtype

`func (o *Dbsnp) SetVarSubtype(v string)`

SetVarSubtype sets VarSubtype field to given value.

### HasVarSubtype

`func (o *Dbsnp) HasVarSubtype() bool`

HasVarSubtype returns a boolean if a field has been set.

### GetVartype

`func (o *Dbsnp) GetVartype() string`

GetVartype returns the Vartype field if non-nil, zero value otherwise.

### GetVartypeOk

`func (o *Dbsnp) GetVartypeOk() (*string, bool)`

GetVartypeOk returns a tuple with the Vartype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVartype

`func (o *Dbsnp) SetVartype(v string)`

SetVartype sets Vartype field to given value.

### HasVartype

`func (o *Dbsnp) HasVartype() bool`

HasVartype returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


