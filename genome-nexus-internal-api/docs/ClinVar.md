# ClinVar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlleleId** | Pointer to **int32** | allele_id | [optional] 
**Alt** | Pointer to **string** | alt | [optional] 
**Chrom** | Pointer to **string** | chrom | [optional] 
**Cytogenic** | Pointer to **string** | cytogenic | [optional] 
**Gene** | Pointer to [**Gene**](Gene.md) |  | [optional] 
**Hg19** | Pointer to [**Hg19**](Hg19.md) |  | [optional] 
**Hg38** | Pointer to [**Hg38**](Hg38.md) |  | [optional] 
**Hgvs** | Pointer to [**Hgvs**](Hgvs.md) |  | [optional] 
**License** | Pointer to **string** | license | [optional] 
**Rcv** | Pointer to [**[]Rcv**](Rcv.md) |  | [optional] 
**VariantId** | Pointer to **int32** | variant_id | [optional] 

## Methods

### NewClinVar

`func NewClinVar() *ClinVar`

NewClinVar instantiates a new ClinVar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClinVarWithDefaults

`func NewClinVarWithDefaults() *ClinVar`

NewClinVarWithDefaults instantiates a new ClinVar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlleleId

`func (o *ClinVar) GetAlleleId() int32`

GetAlleleId returns the AlleleId field if non-nil, zero value otherwise.

### GetAlleleIdOk

`func (o *ClinVar) GetAlleleIdOk() (*int32, bool)`

GetAlleleIdOk returns a tuple with the AlleleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlleleId

`func (o *ClinVar) SetAlleleId(v int32)`

SetAlleleId sets AlleleId field to given value.

### HasAlleleId

`func (o *ClinVar) HasAlleleId() bool`

HasAlleleId returns a boolean if a field has been set.

### GetAlt

`func (o *ClinVar) GetAlt() string`

GetAlt returns the Alt field if non-nil, zero value otherwise.

### GetAltOk

`func (o *ClinVar) GetAltOk() (*string, bool)`

GetAltOk returns a tuple with the Alt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlt

`func (o *ClinVar) SetAlt(v string)`

SetAlt sets Alt field to given value.

### HasAlt

`func (o *ClinVar) HasAlt() bool`

HasAlt returns a boolean if a field has been set.

### GetChrom

`func (o *ClinVar) GetChrom() string`

GetChrom returns the Chrom field if non-nil, zero value otherwise.

### GetChromOk

`func (o *ClinVar) GetChromOk() (*string, bool)`

GetChromOk returns a tuple with the Chrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChrom

`func (o *ClinVar) SetChrom(v string)`

SetChrom sets Chrom field to given value.

### HasChrom

`func (o *ClinVar) HasChrom() bool`

HasChrom returns a boolean if a field has been set.

### GetCytogenic

`func (o *ClinVar) GetCytogenic() string`

GetCytogenic returns the Cytogenic field if non-nil, zero value otherwise.

### GetCytogenicOk

`func (o *ClinVar) GetCytogenicOk() (*string, bool)`

GetCytogenicOk returns a tuple with the Cytogenic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCytogenic

`func (o *ClinVar) SetCytogenic(v string)`

SetCytogenic sets Cytogenic field to given value.

### HasCytogenic

`func (o *ClinVar) HasCytogenic() bool`

HasCytogenic returns a boolean if a field has been set.

### GetGene

`func (o *ClinVar) GetGene() Gene`

GetGene returns the Gene field if non-nil, zero value otherwise.

### GetGeneOk

`func (o *ClinVar) GetGeneOk() (*Gene, bool)`

GetGeneOk returns a tuple with the Gene field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGene

`func (o *ClinVar) SetGene(v Gene)`

SetGene sets Gene field to given value.

### HasGene

`func (o *ClinVar) HasGene() bool`

HasGene returns a boolean if a field has been set.

### GetHg19

`func (o *ClinVar) GetHg19() Hg19`

GetHg19 returns the Hg19 field if non-nil, zero value otherwise.

### GetHg19Ok

`func (o *ClinVar) GetHg19Ok() (*Hg19, bool)`

GetHg19Ok returns a tuple with the Hg19 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHg19

`func (o *ClinVar) SetHg19(v Hg19)`

SetHg19 sets Hg19 field to given value.

### HasHg19

`func (o *ClinVar) HasHg19() bool`

HasHg19 returns a boolean if a field has been set.

### GetHg38

`func (o *ClinVar) GetHg38() Hg38`

GetHg38 returns the Hg38 field if non-nil, zero value otherwise.

### GetHg38Ok

`func (o *ClinVar) GetHg38Ok() (*Hg38, bool)`

GetHg38Ok returns a tuple with the Hg38 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHg38

`func (o *ClinVar) SetHg38(v Hg38)`

SetHg38 sets Hg38 field to given value.

### HasHg38

`func (o *ClinVar) HasHg38() bool`

HasHg38 returns a boolean if a field has been set.

### GetHgvs

`func (o *ClinVar) GetHgvs() Hgvs`

GetHgvs returns the Hgvs field if non-nil, zero value otherwise.

### GetHgvsOk

`func (o *ClinVar) GetHgvsOk() (*Hgvs, bool)`

GetHgvsOk returns a tuple with the Hgvs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHgvs

`func (o *ClinVar) SetHgvs(v Hgvs)`

SetHgvs sets Hgvs field to given value.

### HasHgvs

`func (o *ClinVar) HasHgvs() bool`

HasHgvs returns a boolean if a field has been set.

### GetLicense

`func (o *ClinVar) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *ClinVar) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *ClinVar) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *ClinVar) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetRcv

`func (o *ClinVar) GetRcv() []Rcv`

GetRcv returns the Rcv field if non-nil, zero value otherwise.

### GetRcvOk

`func (o *ClinVar) GetRcvOk() (*[]Rcv, bool)`

GetRcvOk returns a tuple with the Rcv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcv

`func (o *ClinVar) SetRcv(v []Rcv)`

SetRcv sets Rcv field to given value.

### HasRcv

`func (o *ClinVar) HasRcv() bool`

HasRcv returns a boolean if a field has been set.

### GetVariantId

`func (o *ClinVar) GetVariantId() int32`

GetVariantId returns the VariantId field if non-nil, zero value otherwise.

### GetVariantIdOk

`func (o *ClinVar) GetVariantIdOk() (*int32, bool)`

GetVariantIdOk returns a tuple with the VariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantId

`func (o *ClinVar) SetVariantId(v int32)`

SetVariantId sets VariantId field to given value.

### HasVariantId

`func (o *ClinVar) HasVariantId() bool`

HasVariantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


