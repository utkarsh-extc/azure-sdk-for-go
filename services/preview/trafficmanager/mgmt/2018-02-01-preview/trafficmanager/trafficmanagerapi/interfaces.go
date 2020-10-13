package trafficmanagerapi

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
	"github.com/Azure/azure-sdk-for-go/services/preview/trafficmanager/mgmt/2018-02-01-preview/trafficmanager"
)

// EndpointsClientAPI contains the set of methods on the EndpointsClient type.
type EndpointsClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, profileName string, endpointType string, endpointName string, parameters trafficmanager.Endpoint) (result trafficmanager.Endpoint, err error)
	Delete(ctx context.Context, resourceGroupName string, profileName string, endpointType string, endpointName string) (result trafficmanager.DeleteOperationResult, err error)
	Get(ctx context.Context, resourceGroupName string, profileName string, endpointType string, endpointName string) (result trafficmanager.Endpoint, err error)
	Update(ctx context.Context, resourceGroupName string, profileName string, endpointType string, endpointName string, parameters trafficmanager.Endpoint) (result trafficmanager.Endpoint, err error)
}

var _ EndpointsClientAPI = (*trafficmanager.EndpointsClient)(nil)

// ProfilesClientAPI contains the set of methods on the ProfilesClient type.
type ProfilesClientAPI interface {
	CheckTrafficManagerRelativeDNSNameAvailability(ctx context.Context, parameters trafficmanager.CheckTrafficManagerRelativeDNSNameAvailabilityParameters) (result trafficmanager.NameAvailability, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, profileName string, parameters trafficmanager.Profile) (result trafficmanager.Profile, err error)
	Delete(ctx context.Context, resourceGroupName string, profileName string) (result trafficmanager.DeleteOperationResult, err error)
	Get(ctx context.Context, resourceGroupName string, profileName string) (result trafficmanager.Profile, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result trafficmanager.ProfileListResult, err error)
	ListBySubscription(ctx context.Context) (result trafficmanager.ProfileListResult, err error)
	Update(ctx context.Context, resourceGroupName string, profileName string, parameters trafficmanager.Profile) (result trafficmanager.Profile, err error)
}

var _ ProfilesClientAPI = (*trafficmanager.ProfilesClient)(nil)

// GeographicHierarchiesClientAPI contains the set of methods on the GeographicHierarchiesClient type.
type GeographicHierarchiesClientAPI interface {
	GetDefault(ctx context.Context) (result trafficmanager.GeographicHierarchy, err error)
}

var _ GeographicHierarchiesClientAPI = (*trafficmanager.GeographicHierarchiesClient)(nil)

// HeatMapClientAPI contains the set of methods on the HeatMapClient type.
type HeatMapClientAPI interface {
	Get(ctx context.Context, resourceGroupName string, profileName string, topLeft []float64, botRight []float64) (result trafficmanager.HeatMapModel, err error)
}

var _ HeatMapClientAPI = (*trafficmanager.HeatMapClient)(nil)

// UserMetricsKeysClientAPI contains the set of methods on the UserMetricsKeysClient type.
type UserMetricsKeysClientAPI interface {
	CreateOrUpdate(ctx context.Context) (result trafficmanager.UserMetricsKeyModel, err error)
	Delete(ctx context.Context) (result trafficmanager.DeleteOperationResult, err error)
	Get(ctx context.Context) (result trafficmanager.UserMetricsKeyModel, err error)
}

var _ UserMetricsKeysClientAPI = (*trafficmanager.UserMetricsKeysClient)(nil)