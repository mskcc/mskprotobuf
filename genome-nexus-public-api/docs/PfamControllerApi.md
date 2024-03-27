# \PfamControllerApi

All URIs are relative to *http://www.genomenexus.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchPfamDomainsByAccessionGET**](PfamControllerApi.md#FetchPfamDomainsByAccessionGET) | **Get** /pfam/domain/{pfamAccession} | Retrieves a PFAM domain by a PFAM domain ID
[**FetchPfamDomainsByPfamAccessionPOST**](PfamControllerApi.md#FetchPfamDomainsByPfamAccessionPOST) | **Post** /pfam/domain | Retrieves PFAM domains by PFAM domain accession IDs



## FetchPfamDomainsByAccessionGET

> PfamDomain FetchPfamDomainsByAccessionGET(ctx, pfamAccession).Execute()

Retrieves a PFAM domain by a PFAM domain ID

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
    pfamAccession := "pfamAccession_example" // string | A PFAM domain accession ID. For example PF02827

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PfamControllerApi.FetchPfamDomainsByAccessionGET(context.Background(), pfamAccession).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PfamControllerApi.FetchPfamDomainsByAccessionGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchPfamDomainsByAccessionGET`: PfamDomain
    fmt.Fprintf(os.Stdout, "Response from `PfamControllerApi.FetchPfamDomainsByAccessionGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pfamAccession** | **string** | A PFAM domain accession ID. For example PF02827 | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchPfamDomainsByAccessionGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PfamDomain**](PfamDomain.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchPfamDomainsByPfamAccessionPOST

> []PfamDomain FetchPfamDomainsByPfamAccessionPOST(ctx).PfamAccessions(pfamAccessions).Execute()

Retrieves PFAM domains by PFAM domain accession IDs

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
    pfamAccessions := []string{"Property_example"} // []string | List of PFAM domain accession IDs. For example [\"PF02827\",\"PF00093\",\"PF15276\"]

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PfamControllerApi.FetchPfamDomainsByPfamAccessionPOST(context.Background()).PfamAccessions(pfamAccessions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PfamControllerApi.FetchPfamDomainsByPfamAccessionPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchPfamDomainsByPfamAccessionPOST`: []PfamDomain
    fmt.Fprintf(os.Stdout, "Response from `PfamControllerApi.FetchPfamDomainsByPfamAccessionPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchPfamDomainsByPfamAccessionPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pfamAccessions** | **[]string** | List of PFAM domain accession IDs. For example [\&quot;PF02827\&quot;,\&quot;PF00093\&quot;,\&quot;PF15276\&quot;] | 

### Return type

[**[]PfamDomain**](PfamDomain.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

