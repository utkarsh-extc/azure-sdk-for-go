// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkeyvault

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// VaultsClient contains the methods for the Vaults group.
// Don't use this type directly, use NewVaultsClient() instead.
type VaultsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVaultsClient creates a new instance of VaultsClient with the specified values.
func NewVaultsClient(con *armcore.Connection, subscriptionID string) VaultsClient {
	return VaultsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client VaultsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// CheckNameAvailability - Checks that the vault name is valid and is not already in use.
func (client VaultsClient) CheckNameAvailability(ctx context.Context, vaultName VaultCheckNameAvailabilityParameters, options *VaultsCheckNameAvailabilityOptions) (CheckNameAvailabilityResultResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, vaultName, options)
	if err != nil {
		return CheckNameAvailabilityResultResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return CheckNameAvailabilityResultResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return CheckNameAvailabilityResultResponse{}, client.checkNameAvailabilityHandleError(resp)
	}
	result, err := client.checkNameAvailabilityHandleResponse(resp)
	if err != nil {
		return CheckNameAvailabilityResultResponse{}, err
	}
	return result, nil
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client VaultsClient) checkNameAvailabilityCreateRequest(ctx context.Context, vaultName VaultCheckNameAvailabilityParameters, options *VaultsCheckNameAvailabilityOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.KeyVault/checkNameAvailability"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-09-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(vaultName)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client VaultsClient) checkNameAvailabilityHandleResponse(resp *azcore.Response) (CheckNameAvailabilityResultResponse, error) {
	result := CheckNameAvailabilityResultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.CheckNameAvailabilityResult)
	return result, err
}

// checkNameAvailabilityHandleError handles the CheckNameAvailability error response.
func (client VaultsClient) checkNameAvailabilityHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginCreateOrUpdate - Create or update a key vault in the specified subscription.
func (client VaultsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vaultName string, parameters VaultCreateOrUpdateParameters, options *VaultsBeginCreateOrUpdateOptions) (VaultPollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, vaultName, parameters, options)
	if err != nil {
		return VaultPollerResponse{}, err
	}
	result := VaultPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VaultsClient.CreateOrUpdate", "", resp, client.createOrUpdateHandleError)
	if err != nil {
		return VaultPollerResponse{}, err
	}
	poller := &vaultPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (VaultResponse, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumeCreateOrUpdate creates a new VaultPoller from the specified resume token.
// token - The value must come from a previous call to VaultPoller.ResumeToken().
func (client VaultsClient) ResumeCreateOrUpdate(token string) (VaultPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VaultsClient.CreateOrUpdate", token, client.createOrUpdateHandleError)
	if err != nil {
		return nil, err
	}
	return &vaultPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// CreateOrUpdate - Create or update a key vault in the specified subscription.
func (client VaultsClient) createOrUpdate(ctx context.Context, resourceGroupName string, vaultName string, parameters VaultCreateOrUpdateParameters, options *VaultsBeginCreateOrUpdateOptions) (*azcore.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, vaultName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.createOrUpdateHandleError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client VaultsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, parameters VaultCreateOrUpdateParameters, options *VaultsBeginCreateOrUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-09-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client VaultsClient) createOrUpdateHandleResponse(resp *azcore.Response) (VaultResponse, error) {
	result := VaultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.Vault)
	return result, err
}

// createOrUpdateHandleError handles the CreateOrUpdate error response.
func (client VaultsClient) createOrUpdateHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Delete - Deletes the specified Azure key vault.
func (client VaultsClient) Delete(ctx context.Context, resourceGroupName string, vaultName string, options *VaultsDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vaultName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp.Response, nil
}

// deleteCreateRequest creates the Delete request.
func (client VaultsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, options *VaultsDeleteOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-09-01")
	req.URL.RawQuery = query.Encode()
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client VaultsClient) deleteHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Get - Gets the specified Azure key vault.
func (client VaultsClient) Get(ctx context.Context, resourceGroupName string, vaultName string, options *VaultsGetOptions) (VaultResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, vaultName, options)
	if err != nil {
		return VaultResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return VaultResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VaultResponse{}, client.getHandleError(resp)
	}
	result, err := client.getHandleResponse(resp)
	if err != nil {
		return VaultResponse{}, err
	}
	return result, nil
}

// getCreateRequest creates the Get request.
func (client VaultsClient) getCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, options *VaultsGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-09-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client VaultsClient) getHandleResponse(resp *azcore.Response) (VaultResponse, error) {
	result := VaultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.Vault)
	return result, err
}

// getHandleError handles the Get error response.
func (client VaultsClient) getHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// GetDeleted - Gets the deleted Azure key vault.
func (client VaultsClient) GetDeleted(ctx context.Context, vaultName string, location string, options *VaultsGetDeletedOptions) (DeletedVaultResponse, error) {
	req, err := client.getDeletedCreateRequest(ctx, vaultName, location, options)
	if err != nil {
		return DeletedVaultResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return DeletedVaultResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DeletedVaultResponse{}, client.getDeletedHandleError(resp)
	}
	result, err := client.getDeletedHandleResponse(resp)
	if err != nil {
		return DeletedVaultResponse{}, err
	}
	return result, nil
}

// getDeletedCreateRequest creates the GetDeleted request.
func (client VaultsClient) getDeletedCreateRequest(ctx context.Context, vaultName string, location string, options *VaultsGetDeletedOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.KeyVault/locations/{location}/deletedVaults/{vaultName}"
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-09-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getDeletedHandleResponse handles the GetDeleted response.
func (client VaultsClient) getDeletedHandleResponse(resp *azcore.Response) (DeletedVaultResponse, error) {
	result := DeletedVaultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.DeletedVault)
	return result, err
}

// getDeletedHandleError handles the GetDeleted error response.
func (client VaultsClient) getDeletedHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// List - The List operation gets information about the vaults associated with the subscription.
func (client VaultsClient) List(options *VaultsListOptions) ResourceListResultPager {
	return &resourceListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		responder: client.listHandleResponse,
		errorer:   client.listHandleError,
		advancer: func(ctx context.Context, resp ResourceListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.ResourceListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listCreateRequest creates the List request.
func (client VaultsClient) listCreateRequest(ctx context.Context, options *VaultsListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resources"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("$filter", "resourceType eq 'Microsoft.KeyVault/vaults'")
	if options != nil && options.Top != nil {
		query.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	query.Set("api-version", "2015-11-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client VaultsClient) listHandleResponse(resp *azcore.Response) (ResourceListResultResponse, error) {
	result := ResourceListResultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.ResourceListResult)
	return result, err
}

// listHandleError handles the List error response.
func (client VaultsClient) listHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListByResourceGroup - The List operation gets information about the vaults associated with the subscription and within the specified resource group.
func (client VaultsClient) ListByResourceGroup(resourceGroupName string, options *VaultsListByResourceGroupOptions) VaultListResultPager {
	return &vaultListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		responder: client.listByResourceGroupHandleResponse,
		errorer:   client.listByResourceGroupHandleError,
		advancer: func(ctx context.Context, resp VaultListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.VaultListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client VaultsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *VaultsListByResourceGroupOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	if options != nil && options.Top != nil {
		query.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	query.Set("api-version", "2019-09-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client VaultsClient) listByResourceGroupHandleResponse(resp *azcore.Response) (VaultListResultResponse, error) {
	result := VaultListResultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.VaultListResult)
	return result, err
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client VaultsClient) listByResourceGroupHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListBySubscription - The List operation gets information about the vaults associated with the subscription.
func (client VaultsClient) ListBySubscription(options *VaultsListBySubscriptionOptions) VaultListResultPager {
	return &vaultListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		responder: client.listBySubscriptionHandleResponse,
		errorer:   client.listBySubscriptionHandleError,
		advancer: func(ctx context.Context, resp VaultListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.VaultListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client VaultsClient) listBySubscriptionCreateRequest(ctx context.Context, options *VaultsListBySubscriptionOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.KeyVault/vaults"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	if options != nil && options.Top != nil {
		query.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	query.Set("api-version", "2019-09-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client VaultsClient) listBySubscriptionHandleResponse(resp *azcore.Response) (VaultListResultResponse, error) {
	result := VaultListResultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.VaultListResult)
	return result, err
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client VaultsClient) listBySubscriptionHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// ListDeleted - Gets information about the deleted vaults in a subscription.
func (client VaultsClient) ListDeleted(options *VaultsListDeletedOptions) DeletedVaultListResultPager {
	return &deletedVaultListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listDeletedCreateRequest(ctx, options)
		},
		responder: client.listDeletedHandleResponse,
		errorer:   client.listDeletedHandleError,
		advancer: func(ctx context.Context, resp DeletedVaultListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.DeletedVaultListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listDeletedCreateRequest creates the ListDeleted request.
func (client VaultsClient) listDeletedCreateRequest(ctx context.Context, options *VaultsListDeletedOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.KeyVault/deletedVaults"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-09-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listDeletedHandleResponse handles the ListDeleted response.
func (client VaultsClient) listDeletedHandleResponse(resp *azcore.Response) (DeletedVaultListResultResponse, error) {
	result := DeletedVaultListResultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.DeletedVaultListResult)
	return result, err
}

// listDeletedHandleError handles the ListDeleted error response.
func (client VaultsClient) listDeletedHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// BeginPurgeDeleted - Permanently deletes the specified vault. aka Purges the deleted Azure key vault.
func (client VaultsClient) BeginPurgeDeleted(ctx context.Context, vaultName string, location string, options *VaultsBeginPurgeDeletedOptions) (HTTPPollerResponse, error) {
	resp, err := client.purgeDeleted(ctx, vaultName, location, options)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	result := HTTPPollerResponse{
		RawResponse: resp.Response,
	}
	pt, err := armcore.NewPoller("VaultsClient.PurgeDeleted", "", resp, client.purgeDeletedHandleError)
	if err != nil {
		return HTTPPollerResponse{}, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.con.Pipeline(),
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

// ResumePurgeDeleted creates a new HTTPPoller from the specified resume token.
// token - The value must come from a previous call to HTTPPoller.ResumeToken().
func (client VaultsClient) ResumePurgeDeleted(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("VaultsClient.PurgeDeleted", token, client.purgeDeletedHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.con.Pipeline(),
		pt:       pt,
	}, nil
}

// PurgeDeleted - Permanently deletes the specified vault. aka Purges the deleted Azure key vault.
func (client VaultsClient) purgeDeleted(ctx context.Context, vaultName string, location string, options *VaultsBeginPurgeDeletedOptions) (*azcore.Response, error) {
	req, err := client.purgeDeletedCreateRequest(ctx, vaultName, location, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted) {
		return nil, client.purgeDeletedHandleError(resp)
	}
	return resp, nil
}

// purgeDeletedCreateRequest creates the PurgeDeleted request.
func (client VaultsClient) purgeDeletedCreateRequest(ctx context.Context, vaultName string, location string, options *VaultsBeginPurgeDeletedOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.KeyVault/locations/{location}/deletedVaults/{vaultName}/purge"
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-09-01")
	req.URL.RawQuery = query.Encode()
	return req, nil
}

// purgeDeletedHandleError handles the PurgeDeleted error response.
func (client VaultsClient) purgeDeletedHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// Update - Update a key vault in the specified subscription.
func (client VaultsClient) Update(ctx context.Context, resourceGroupName string, vaultName string, parameters VaultPatchParameters, options *VaultsUpdateOptions) (VaultResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, vaultName, parameters, options)
	if err != nil {
		return VaultResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return VaultResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return VaultResponse{}, client.updateHandleError(resp)
	}
	result, err := client.updateHandleResponse(resp)
	if err != nil {
		return VaultResponse{}, err
	}
	return result, nil
}

// updateCreateRequest creates the Update request.
func (client VaultsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, parameters VaultPatchParameters, options *VaultsUpdateOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-09-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateHandleResponse handles the Update response.
func (client VaultsClient) updateHandleResponse(resp *azcore.Response) (VaultResponse, error) {
	result := VaultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.Vault)
	return result, err
}

// updateHandleError handles the Update error response.
func (client VaultsClient) updateHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}

// UpdateAccessPolicy - Update access policies in a key vault in the specified subscription.
func (client VaultsClient) UpdateAccessPolicy(ctx context.Context, resourceGroupName string, vaultName string, operationKind AccessPolicyUpdateKind, parameters VaultAccessPolicyParameters, options *VaultsUpdateAccessPolicyOptions) (VaultAccessPolicyParametersResponse, error) {
	req, err := client.updateAccessPolicyCreateRequest(ctx, resourceGroupName, vaultName, operationKind, parameters, options)
	if err != nil {
		return VaultAccessPolicyParametersResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return VaultAccessPolicyParametersResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return VaultAccessPolicyParametersResponse{}, client.updateAccessPolicyHandleError(resp)
	}
	result, err := client.updateAccessPolicyHandleResponse(resp)
	if err != nil {
		return VaultAccessPolicyParametersResponse{}, err
	}
	return result, nil
}

// updateAccessPolicyCreateRequest creates the UpdateAccessPolicy request.
func (client VaultsClient) updateAccessPolicyCreateRequest(ctx context.Context, resourceGroupName string, vaultName string, operationKind AccessPolicyUpdateKind, parameters VaultAccessPolicyParameters, options *VaultsUpdateAccessPolicyOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}/accessPolicies/{operationKind}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{vaultName}", url.PathEscape(vaultName))
	urlPath = strings.ReplaceAll(urlPath, "{operationKind}", url.PathEscape(string(operationKind)))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-09-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// updateAccessPolicyHandleResponse handles the UpdateAccessPolicy response.
func (client VaultsClient) updateAccessPolicyHandleResponse(resp *azcore.Response) (VaultAccessPolicyParametersResponse, error) {
	result := VaultAccessPolicyParametersResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.VaultAccessPolicyParameters)
	return result, err
}

// updateAccessPolicyHandleError handles the UpdateAccessPolicy error response.
func (client VaultsClient) updateAccessPolicyHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}