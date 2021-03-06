package blockchain

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// LocationsClient is the REST API for Azure Blockchain Service
type LocationsClient struct {
	BaseClient
}

// NewLocationsClient creates an instance of the LocationsClient client.
func NewLocationsClient(subscriptionID string) LocationsClient {
	return NewLocationsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewLocationsClientWithBaseURI creates an instance of the LocationsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewLocationsClientWithBaseURI(baseURI string, subscriptionID string) LocationsClient {
	return LocationsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CheckNameAvailability to check whether a resource name is available.
// Parameters:
// locationName - location Name.
// nameAvailabilityRequest - name availability request payload.
func (client LocationsClient) CheckNameAvailability(ctx context.Context, locationName string, nameAvailabilityRequest *NameAvailabilityRequest) (result NameAvailability, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LocationsClient.CheckNameAvailability")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CheckNameAvailabilityPreparer(ctx, locationName, nameAvailabilityRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.LocationsClient", "CheckNameAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckNameAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "blockchain.LocationsClient", "CheckNameAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckNameAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.LocationsClient", "CheckNameAvailability", resp, "Failure responding to request")
		return
	}

	return
}

// CheckNameAvailabilityPreparer prepares the CheckNameAvailability request.
func (client LocationsClient) CheckNameAvailabilityPreparer(ctx context.Context, locationName string, nameAvailabilityRequest *NameAvailabilityRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"locationName":   autorest.Encode("path", locationName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Blockchain/locations/{locationName}/checkNameAvailability", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if nameAvailabilityRequest != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(nameAvailabilityRequest))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckNameAvailabilitySender sends the CheckNameAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client LocationsClient) CheckNameAvailabilitySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CheckNameAvailabilityResponder handles the response to the CheckNameAvailability request. The method always
// closes the http.Response Body.
func (client LocationsClient) CheckNameAvailabilityResponder(resp *http.Response) (result NameAvailability, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListConsortiums lists the available consortiums for a subscription.
// Parameters:
// locationName - location Name.
func (client LocationsClient) ListConsortiums(ctx context.Context, locationName string) (result ConsortiumCollection, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/LocationsClient.ListConsortiums")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListConsortiumsPreparer(ctx, locationName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.LocationsClient", "ListConsortiums", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListConsortiumsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "blockchain.LocationsClient", "ListConsortiums", resp, "Failure sending request")
		return
	}

	result, err = client.ListConsortiumsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "blockchain.LocationsClient", "ListConsortiums", resp, "Failure responding to request")
		return
	}

	return
}

// ListConsortiumsPreparer prepares the ListConsortiums request.
func (client LocationsClient) ListConsortiumsPreparer(ctx context.Context, locationName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"locationName":   autorest.Encode("path", locationName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-06-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Blockchain/locations/{locationName}/listConsortiums", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListConsortiumsSender sends the ListConsortiums request. The method will close the
// http.Response Body if it receives an error.
func (client LocationsClient) ListConsortiumsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListConsortiumsResponder handles the response to the ListConsortiums request. The method always
// closes the http.Response Body.
func (client LocationsClient) ListConsortiumsResponder(resp *http.Response) (result ConsortiumCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
