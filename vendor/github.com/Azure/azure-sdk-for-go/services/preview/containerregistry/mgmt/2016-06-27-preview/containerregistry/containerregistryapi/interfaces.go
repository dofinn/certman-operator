package containerregistryapi

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
	"github.com/Azure/azure-sdk-for-go/services/preview/containerregistry/mgmt/2016-06-27-preview/containerregistry"
	"github.com/Azure/go-autorest/autorest"
)

// RegistriesClientAPI contains the set of methods on the RegistriesClient type.
type RegistriesClientAPI interface {
	CheckNameAvailability(ctx context.Context, registryNameCheckRequest containerregistry.RegistryNameCheckRequest) (result containerregistry.RegistryNameStatus, err error)
	CreateOrUpdate(ctx context.Context, resourceGroupName string, registryName string, registry containerregistry.Registry) (result containerregistry.Registry, err error)
	Delete(ctx context.Context, resourceGroupName string, registryName string) (result autorest.Response, err error)
	GetCredentials(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.RegistryCredentials, err error)
	GetProperties(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.Registry, err error)
	List(ctx context.Context) (result containerregistry.RegistryListResultPage, err error)
	ListComplete(ctx context.Context) (result containerregistry.RegistryListResultIterator, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result containerregistry.RegistryListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result containerregistry.RegistryListResultIterator, err error)
	RegenerateCredentials(ctx context.Context, resourceGroupName string, registryName string) (result containerregistry.RegistryCredentials, err error)
	Update(ctx context.Context, resourceGroupName string, registryName string, registryUpdateParameters containerregistry.RegistryUpdateParameters) (result containerregistry.Registry, err error)
}

var _ RegistriesClientAPI = (*containerregistry.RegistriesClient)(nil)
