# \PdbControllerApi

All URIs are relative to *http://www.genomenexus.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchPdbHeaderGET**](PdbControllerApi.md#FetchPdbHeaderGET) | **Get** /pdb/header/{pdbId} | Retrieves PDB header info by a PDB id
[**FetchPdbHeaderPOST**](PdbControllerApi.md#FetchPdbHeaderPOST) | **Post** /pdb/header | Retrieves PDB header info by a PDB id



## FetchPdbHeaderGET

> PdbHeader FetchPdbHeaderGET(ctx, pdbId).Execute()

Retrieves PDB header info by a PDB id

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
    pdbId := "pdbId_example" // string | PDB id, for example 1a37

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PdbControllerApi.FetchPdbHeaderGET(context.Background(), pdbId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PdbControllerApi.FetchPdbHeaderGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchPdbHeaderGET`: PdbHeader
    fmt.Fprintf(os.Stdout, "Response from `PdbControllerApi.FetchPdbHeaderGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pdbId** | **string** | PDB id, for example 1a37 | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchPdbHeaderGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PdbHeader**](PdbHeader.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FetchPdbHeaderPOST

> []PdbHeader FetchPdbHeaderPOST(ctx).PdbIds(pdbIds).Execute()

Retrieves PDB header info by a PDB id

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
    pdbIds := []string{"Property_example"} // []string | List of pdb ids, for example [\"1a37\",\"1a4o\"]

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PdbControllerApi.FetchPdbHeaderPOST(context.Background()).PdbIds(pdbIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PdbControllerApi.FetchPdbHeaderPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FetchPdbHeaderPOST`: []PdbHeader
    fmt.Fprintf(os.Stdout, "Response from `PdbControllerApi.FetchPdbHeaderPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFetchPdbHeaderPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pdbIds** | **[]string** | List of pdb ids, for example [\&quot;1a37\&quot;,\&quot;1a4o\&quot;] | 

### Return type

[**[]PdbHeader**](PdbHeader.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

