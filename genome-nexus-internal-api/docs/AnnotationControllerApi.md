# \AnnotationControllerApi

All URIs are relative to *http://www.genomenexus.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchVariantAnnotationByGenomicLocationGET**](AnnotationControllerApi.md#FetchVariantAnnotationByGenomicLocationGET) | **Get** /annotation/genomic/{genomicLocation} | Retrieves VEP annotation for the provided genomic location
[**FetchVariantAnnotationByGenomicLocationPOST**](AnnotationControllerApi.md#FetchVariantAnnotationByGenomicLocationPOST) | **Post** /annotation/genomic | Retrieves VEP annotation for the provided list of genomic locations
[**FetchVariantAnnotationByIdGET**](AnnotationControllerApi.md#FetchVariantAnnotationByIdGET) | **Get** /annotation/dbsnp/{variantId} | Retrieves VEP annotation for the give dbSNP id
[**FetchVariantAnnotationByIdPOST**](AnnotationControllerApi.md#FetchVariantAnnotationByIdPOST) | **Post** /annotation/dbsnp/ | Retrieves VEP annotation for the provided list of dbSNP ids
[**FetchVariantAnnotationGET**](AnnotationControllerApi.md#FetchVariantAnnotationGET) | **Get** /annotation/{variant} | Retrieves VEP annotation for the provided variant
[**FetchVariantAnnotationPOST**](AnnotationControllerApi.md#FetchVariantAnnotationPOST) | **Post** /annotation | Retrieves VEP annotation for the provided list of variants



## FetchVariantAnnotationByGenomicLocationGET

> VariantAnnotation FetchVariantAnnotationByGenomicLocationGET(ctx, genomicLocation).IsoformOverrideSource(isoformOverrideSource).Token(token).Fields(fields).Execute()

Retrieves VEP annotation for the provided genomic location

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    genomicLocation := "genomicLocation_example" // string | A genomic location. For example 7,140453136,140453136,A,T
    isoformOverrideSource := "isoformOverrideSource_example" // string | Isoform override source. For example uniprot (optional)
    token := "token_example" // string | Map of tokens. For example {\"source1\":\"put-your-token1-here\",\"source2\":\"put-your-token2-here\"} (optional)
    fields := []string{"Inner_example"} // []string | Comma separated list of fields to include (case-sensitive!). For example: hotspots (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnnotationControllerApi.FetchVariantAnnotationByGenomicLocationGET(context.Background(), genomicLocation).IsoformOverrideSource(isoformOverrideSource).Token(token).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnnotationControllerApi.FetchVariantAnnotationByGenomicLocationGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchVariantAnnotationByGenomicLocationGET`: VariantAnnotation
    fmt.Fprintf(os.Stdout, "Response from `AnnotationControllerApi.FetchVariantAnnotationByGenomicLocationGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**genomicLocation** | **string** | A genomic location. For example 7,140453136,140453136,A,T | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchVariantAnnotationByGenomicLocationGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isoformOverrideSource** | **string** | Isoform override source. For example uniprot | 
 **token** | **string** | Map of tokens. For example {\&quot;source1\&quot;:\&quot;put-your-token1-here\&quot;,\&quot;source2\&quot;:\&quot;put-your-token2-here\&quot;} | 
 **fields** | **[]string** | Comma separated list of fields to include (case-sensitive!). For example: hotspots | 

### Return type

[**VariantAnnotation**](VariantAnnotation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchVariantAnnotationByGenomicLocationPOST

> []VariantAnnotation FetchVariantAnnotationByGenomicLocationPOST(ctx).GenomicLocations(genomicLocations).IsoformOverrideSource(isoformOverrideSource).Token(token).Fields(fields).Execute()

Retrieves VEP annotation for the provided list of genomic locations

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    genomicLocations := []openapiclient.GenomicLocation{*openapiclient.NewGenomicLocation("Chromosome_example", int32(123), int32(123), "ReferenceAllele_example", "VariantAllele_example")} // []GenomicLocation | List of Genomic Locations
    isoformOverrideSource := "isoformOverrideSource_example" // string | Isoform override source. For example uniprot (optional)
    token := "token_example" // string | Map of tokens. For example {\"source1\":\"put-your-token1-here\",\"source2\":\"put-your-token2-here\"} (optional)
    fields := []string{"Inner_example"} // []string | Comma separated list of fields to include (case-sensitive!). For example: hotspots (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnnotationControllerApi.FetchVariantAnnotationByGenomicLocationPOST(context.Background()).GenomicLocations(genomicLocations).IsoformOverrideSource(isoformOverrideSource).Token(token).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnnotationControllerApi.FetchVariantAnnotationByGenomicLocationPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchVariantAnnotationByGenomicLocationPOST`: []VariantAnnotation
    fmt.Fprintf(os.Stdout, "Response from `AnnotationControllerApi.FetchVariantAnnotationByGenomicLocationPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchVariantAnnotationByGenomicLocationPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **genomicLocations** | [**[]GenomicLocation**](GenomicLocation.md) | List of Genomic Locations | 
 **isoformOverrideSource** | **string** | Isoform override source. For example uniprot | 
 **token** | **string** | Map of tokens. For example {\&quot;source1\&quot;:\&quot;put-your-token1-here\&quot;,\&quot;source2\&quot;:\&quot;put-your-token2-here\&quot;} | 
 **fields** | **[]string** | Comma separated list of fields to include (case-sensitive!). For example: hotspots | 

### Return type

[**[]VariantAnnotation**](VariantAnnotation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchVariantAnnotationByIdGET

> VariantAnnotation FetchVariantAnnotationByIdGET(ctx, variantId).IsoformOverrideSource(isoformOverrideSource).Token(token).Fields(fields).Execute()

Retrieves VEP annotation for the give dbSNP id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    variantId := "variantId_example" // string | dbSNP id. For example rs116035550.
    isoformOverrideSource := "isoformOverrideSource_example" // string | Isoform override source. For example uniprot (optional)
    token := "token_example" // string | Map of tokens. For example {\"source1\":\"put-your-token1-here\",\"source2\":\"put-your-token2-here\"} (optional)
    fields := []string{"Inner_example"} // []string | Comma separated list of fields to include (case-sensitive!). For example: annotation_summary (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnnotationControllerApi.FetchVariantAnnotationByIdGET(context.Background(), variantId).IsoformOverrideSource(isoformOverrideSource).Token(token).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnnotationControllerApi.FetchVariantAnnotationByIdGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchVariantAnnotationByIdGET`: VariantAnnotation
    fmt.Fprintf(os.Stdout, "Response from `AnnotationControllerApi.FetchVariantAnnotationByIdGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**variantId** | **string** | dbSNP id. For example rs116035550. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchVariantAnnotationByIdGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isoformOverrideSource** | **string** | Isoform override source. For example uniprot | 
 **token** | **string** | Map of tokens. For example {\&quot;source1\&quot;:\&quot;put-your-token1-here\&quot;,\&quot;source2\&quot;:\&quot;put-your-token2-here\&quot;} | 
 **fields** | **[]string** | Comma separated list of fields to include (case-sensitive!). For example: annotation_summary | 

### Return type

[**VariantAnnotation**](VariantAnnotation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchVariantAnnotationByIdPOST

> []VariantAnnotation FetchVariantAnnotationByIdPOST(ctx).VariantIds(variantIds).IsoformOverrideSource(isoformOverrideSource).Token(token).Fields(fields).Execute()

Retrieves VEP annotation for the provided list of dbSNP ids

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    variantIds := []string{"Property_example"} // []string | List of variant IDs. For example [\"rs116035550\"]
    isoformOverrideSource := "isoformOverrideSource_example" // string | Isoform override source. For example uniprot (optional)
    token := "token_example" // string | Map of tokens. For example {\"source1\":\"put-your-token1-here\",\"source2\":\"put-your-token2-here\"} (optional)
    fields := []string{"Inner_example"} // []string | Comma separated list of fields to include (case-sensitive!). For example: annotation_summary (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnnotationControllerApi.FetchVariantAnnotationByIdPOST(context.Background()).VariantIds(variantIds).IsoformOverrideSource(isoformOverrideSource).Token(token).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnnotationControllerApi.FetchVariantAnnotationByIdPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchVariantAnnotationByIdPOST`: []VariantAnnotation
    fmt.Fprintf(os.Stdout, "Response from `AnnotationControllerApi.FetchVariantAnnotationByIdPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchVariantAnnotationByIdPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variantIds** | **[]string** | List of variant IDs. For example [\&quot;rs116035550\&quot;] | 
 **isoformOverrideSource** | **string** | Isoform override source. For example uniprot | 
 **token** | **string** | Map of tokens. For example {\&quot;source1\&quot;:\&quot;put-your-token1-here\&quot;,\&quot;source2\&quot;:\&quot;put-your-token2-here\&quot;} | 
 **fields** | **[]string** | Comma separated list of fields to include (case-sensitive!). For example: annotation_summary | 

### Return type

[**[]VariantAnnotation**](VariantAnnotation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchVariantAnnotationGET

> VariantAnnotation FetchVariantAnnotationGET(ctx, variant).IsoformOverrideSource(isoformOverrideSource).Token(token).Fields(fields).Execute()

Retrieves VEP annotation for the provided variant

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    variant := "variant_example" // string | Variant. For example 17:g.41242962_41242963insGA
    isoformOverrideSource := "isoformOverrideSource_example" // string | Isoform override source. For example uniprot (optional)
    token := "token_example" // string | Map of tokens. For example {\"source1\":\"put-your-token1-here\",\"source2\":\"put-your-token2-here\"} (optional)
    fields := []string{"Inner_example"} // []string | Comma separated list of fields to include (case-sensitive!). For example: hotspots (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnnotationControllerApi.FetchVariantAnnotationGET(context.Background(), variant).IsoformOverrideSource(isoformOverrideSource).Token(token).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnnotationControllerApi.FetchVariantAnnotationGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchVariantAnnotationGET`: VariantAnnotation
    fmt.Fprintf(os.Stdout, "Response from `AnnotationControllerApi.FetchVariantAnnotationGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**variant** | **string** | Variant. For example 17:g.41242962_41242963insGA | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchVariantAnnotationGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isoformOverrideSource** | **string** | Isoform override source. For example uniprot | 
 **token** | **string** | Map of tokens. For example {\&quot;source1\&quot;:\&quot;put-your-token1-here\&quot;,\&quot;source2\&quot;:\&quot;put-your-token2-here\&quot;} | 
 **fields** | **[]string** | Comma separated list of fields to include (case-sensitive!). For example: hotspots | 

### Return type

[**VariantAnnotation**](VariantAnnotation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchVariantAnnotationPOST

> []VariantAnnotation FetchVariantAnnotationPOST(ctx).Variants(variants).IsoformOverrideSource(isoformOverrideSource).Token(token).Fields(fields).Execute()

Retrieves VEP annotation for the provided list of variants

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    variants := []string{"Property_example"} // []string | List of variants. For example [\"X:g.66937331T>A\",\"17:g.41242962_41242963insGA\"] (GRCh37) or [\"1:g.182712A>C\", \"2:g.265023C>T\", \"3:g.319781del\", \"19:g.110753dup\", \"1:g.1385015_1387562del\"] (GRCh38)
    isoformOverrideSource := "isoformOverrideSource_example" // string | Isoform override source. For example uniprot (optional)
    token := "token_example" // string | Map of tokens. For example {\"source1\":\"put-your-token1-here\",\"source2\":\"put-your-token2-here\"} (optional)
    fields := []string{"Inner_example"} // []string | Comma separated list of fields to include (case-sensitive!). For example: hotspots (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnnotationControllerApi.FetchVariantAnnotationPOST(context.Background()).Variants(variants).IsoformOverrideSource(isoformOverrideSource).Token(token).Fields(fields).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnnotationControllerApi.FetchVariantAnnotationPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchVariantAnnotationPOST`: []VariantAnnotation
    fmt.Fprintf(os.Stdout, "Response from `AnnotationControllerApi.FetchVariantAnnotationPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchVariantAnnotationPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variants** | **[]string** | List of variants. For example [\&quot;X:g.66937331T&gt;A\&quot;,\&quot;17:g.41242962_41242963insGA\&quot;] (GRCh37) or [\&quot;1:g.182712A&gt;C\&quot;, \&quot;2:g.265023C&gt;T\&quot;, \&quot;3:g.319781del\&quot;, \&quot;19:g.110753dup\&quot;, \&quot;1:g.1385015_1387562del\&quot;] (GRCh38) | 
 **isoformOverrideSource** | **string** | Isoform override source. For example uniprot | 
 **token** | **string** | Map of tokens. For example {\&quot;source1\&quot;:\&quot;put-your-token1-here\&quot;,\&quot;source2\&quot;:\&quot;put-your-token2-here\&quot;} | 
 **fields** | **[]string** | Comma separated list of fields to include (case-sensitive!). For example: hotspots | 

### Return type

[**[]VariantAnnotation**](VariantAnnotation.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

