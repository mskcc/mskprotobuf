/*
Genome Nexus API

Testing InfoControllerApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package genome_nexus_internal_api

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_genome_nexus_internal_api_InfoControllerApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test InfoControllerApiService FetchVersionGET", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.InfoControllerApi.FetchVersionGET(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}
