//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmobilenetwork_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mobilenetwork/armmobilenetwork/v4"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-02-01/examples/SimGroupDelete.json
func ExampleSimGroupsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmobilenetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSimGroupsClient().BeginDelete(ctx, "testResourceGroupName", "testSimGroup", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-02-01/examples/SimGroupGet.json
func ExampleSimGroupsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmobilenetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSimGroupsClient().Get(ctx, "testResourceGroupName", "testSimGroupName", nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SimGroup = armmobilenetwork.SimGroup{
	// 	Name: to.Ptr("testSimGroup"),
	// 	Type: to.Ptr("Microsoft.MobileNetwork/simGroups"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/simGroups/testSimGroup"),
	// 	SystemData: &armmobilenetwork.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armmobilenetwork.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armmobilenetwork.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Identity: &armmobilenetwork.ManagedServiceIdentity{
	// 		Type: to.Ptr(armmobilenetwork.ManagedServiceIdentityTypeUserAssigned),
	// 		UserAssignedIdentities: map[string]*armmobilenetwork.UserAssignedIdentity{
	// 			"/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testUserAssignedManagedIdentity": &armmobilenetwork.UserAssignedIdentity{
	// 				ClientID: to.Ptr("12345678-abcd-dcba-abcd-0123456789ef"),
	// 				PrincipalID: to.Ptr("12345678-abcd-dcba-abcd-0123456789ef"),
	// 			},
	// 		},
	// 	},
	// 	Properties: &armmobilenetwork.SimGroupPropertiesFormat{
	// 		EncryptionKey: &armmobilenetwork.KeyVaultKey{
	// 			KeyURL: to.Ptr("https://contosovault.vault.azure.net/keys/azureKey"),
	// 		},
	// 		MobileNetwork: &armmobilenetwork.ResourceID{
	// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork"),
	// 		},
	// 		ProvisioningState: to.Ptr(armmobilenetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-02-01/examples/SimGroupCreate.json
func ExampleSimGroupsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmobilenetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := clientFactory.NewSimGroupsClient().BeginCreateOrUpdate(ctx, "rg1", "testSimGroup", armmobilenetwork.SimGroup{
		Location: to.Ptr("eastus"),
		Identity: &armmobilenetwork.ManagedServiceIdentity{
			Type: to.Ptr(armmobilenetwork.ManagedServiceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armmobilenetwork.UserAssignedIdentity{
				"/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testUserAssignedManagedIdentity": {},
			},
		},
		Properties: &armmobilenetwork.SimGroupPropertiesFormat{
			EncryptionKey: &armmobilenetwork.KeyVaultKey{
				KeyURL: to.Ptr("https://contosovault.vault.azure.net/keys/azureKey"),
			},
			MobileNetwork: &armmobilenetwork.ResourceID{
				ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork"),
			},
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SimGroup = armmobilenetwork.SimGroup{
	// 	Name: to.Ptr("testSimGroup"),
	// 	Type: to.Ptr("Microsoft.MobileNetwork/simGroups"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/simGroups/testSimGroup"),
	// 	SystemData: &armmobilenetwork.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armmobilenetwork.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armmobilenetwork.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 	},
	// 	Identity: &armmobilenetwork.ManagedServiceIdentity{
	// 		Type: to.Ptr(armmobilenetwork.ManagedServiceIdentityTypeUserAssigned),
	// 		UserAssignedIdentities: map[string]*armmobilenetwork.UserAssignedIdentity{
	// 			"/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testUserAssignedManagedIdentity": &armmobilenetwork.UserAssignedIdentity{
	// 				ClientID: to.Ptr("12345678-abcd-dcba-abcd-0123456789ef"),
	// 				PrincipalID: to.Ptr("12345678-abcd-dcba-abcd-0123456789ef"),
	// 			},
	// 		},
	// 	},
	// 	Properties: &armmobilenetwork.SimGroupPropertiesFormat{
	// 		EncryptionKey: &armmobilenetwork.KeyVaultKey{
	// 			KeyURL: to.Ptr("https://contosovault.vault.azure.net/keys/azureKey"),
	// 		},
	// 		MobileNetwork: &armmobilenetwork.ResourceID{
	// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork"),
	// 		},
	// 		ProvisioningState: to.Ptr(armmobilenetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-02-01/examples/SimGroupPatch.json
func ExampleSimGroupsClient_UpdateTags() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmobilenetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := clientFactory.NewSimGroupsClient().UpdateTags(ctx, "rg1", "testSimGroup", armmobilenetwork.IdentityAndTagsObject{
		Identity: &armmobilenetwork.ManagedServiceIdentity{
			Type: to.Ptr(armmobilenetwork.ManagedServiceIdentityTypeUserAssigned),
			UserAssignedIdentities: map[string]*armmobilenetwork.UserAssignedIdentity{
				"/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testUserAssignedManagedIdentity": {},
			},
		},
		Tags: map[string]*string{
			"tag1": to.Ptr("value1"),
			"tag2": to.Ptr("value2"),
		},
	}, nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// You could use response here. We use blank identifier for just demo purposes.
	_ = res
	// If the HTTP response code is 200 as defined in example definition, your response structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
	// res.SimGroup = armmobilenetwork.SimGroup{
	// 	Name: to.Ptr("testSimGroup"),
	// 	Type: to.Ptr("Microsoft.MobileNetwork/simGroups"),
	// 	ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/simGroups/testSimGroup"),
	// 	SystemData: &armmobilenetwork.SystemData{
	// 		CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
	// 		CreatedBy: to.Ptr("user1"),
	// 		CreatedByType: to.Ptr(armmobilenetwork.CreatedByTypeUser),
	// 		LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
	// 		LastModifiedBy: to.Ptr("user2"),
	// 		LastModifiedByType: to.Ptr(armmobilenetwork.CreatedByTypeUser),
	// 	},
	// 	Location: to.Ptr("eastus"),
	// 	Tags: map[string]*string{
	// 		"tag1": to.Ptr("value1"),
	// 		"tag2": to.Ptr("value2"),
	// 	},
	// 	Identity: &armmobilenetwork.ManagedServiceIdentity{
	// 		Type: to.Ptr(armmobilenetwork.ManagedServiceIdentityTypeUserAssigned),
	// 		UserAssignedIdentities: map[string]*armmobilenetwork.UserAssignedIdentity{
	// 			"/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testUserAssignedManagedIdentity": &armmobilenetwork.UserAssignedIdentity{
	// 				ClientID: to.Ptr("12345678-abcd-dcba-abcd-0123456789ef"),
	// 				PrincipalID: to.Ptr("12345678-abcd-dcba-abcd-0123456789ef"),
	// 			},
	// 		},
	// 	},
	// 	Properties: &armmobilenetwork.SimGroupPropertiesFormat{
	// 		EncryptionKey: &armmobilenetwork.KeyVaultKey{
	// 			KeyURL: to.Ptr("https://contosovault.vault.azure.net/keys/azureKey"),
	// 		},
	// 		MobileNetwork: &armmobilenetwork.ResourceID{
	// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork"),
	// 		},
	// 		ProvisioningState: to.Ptr(armmobilenetwork.ProvisioningStateSucceeded),
	// 	},
	// }
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-02-01/examples/SimGroupListBySubscription.json
func ExampleSimGroupsClient_NewListBySubscriptionPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmobilenetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSimGroupsClient().NewListBySubscriptionPager(nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.SimGroupListResult = armmobilenetwork.SimGroupListResult{
		// 	Value: []*armmobilenetwork.SimGroup{
		// 		{
		// 			Name: to.Ptr("testSimGroup"),
		// 			Type: to.Ptr("Microsoft.MobileNetwork/simGroups"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/simGroups/testSimGroup"),
		// 			SystemData: &armmobilenetwork.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armmobilenetwork.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armmobilenetwork.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Identity: &armmobilenetwork.ManagedServiceIdentity{
		// 				Type: to.Ptr(armmobilenetwork.ManagedServiceIdentityTypeUserAssigned),
		// 				UserAssignedIdentities: map[string]*armmobilenetwork.UserAssignedIdentity{
		// 					"/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testUserAssignedManagedIdentity": &armmobilenetwork.UserAssignedIdentity{
		// 						ClientID: to.Ptr("12345678-abcd-dcba-abcd-0123456789ef"),
		// 						PrincipalID: to.Ptr("12345678-abcd-dcba-abcd-0123456789ef"),
		// 					},
		// 				},
		// 			},
		// 			Properties: &armmobilenetwork.SimGroupPropertiesFormat{
		// 				EncryptionKey: &armmobilenetwork.KeyVaultKey{
		// 					KeyURL: to.Ptr("https://contosovault.vault.azure.net/keys/azureKey"),
		// 				},
		// 				MobileNetwork: &armmobilenetwork.ResourceID{
		// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork"),
		// 				},
		// 				ProvisioningState: to.Ptr(armmobilenetwork.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/blob/cf5ad1932d00c7d15497705ad6b71171d3d68b1e/specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-02-01/examples/SimGroupListByResourceGroup.json
func ExampleSimGroupsClient_NewListByResourceGroupPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	clientFactory, err := armmobilenetwork.NewClientFactory("<subscription-id>", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := clientFactory.NewSimGroupsClient().NewListByResourceGroupPager("rg1", nil)
	for pager.More() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range page.Value {
			// You could use page here. We use blank identifier for just demo purposes.
			_ = v
		}
		// If the HTTP response code is 200 as defined in example definition, your page structure would look as follows. Please pay attention that all the values in the output are fake values for just demo purposes.
		// page.SimGroupListResult = armmobilenetwork.SimGroupListResult{
		// 	Value: []*armmobilenetwork.SimGroup{
		// 		{
		// 			Name: to.Ptr("testSimGroup"),
		// 			Type: to.Ptr("Microsoft.MobileNetwork/simGroups"),
		// 			ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/simGroups/testSimGroup"),
		// 			SystemData: &armmobilenetwork.SystemData{
		// 				CreatedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-01T17:18:19.123Z"); return t}()),
		// 				CreatedBy: to.Ptr("user1"),
		// 				CreatedByType: to.Ptr(armmobilenetwork.CreatedByTypeUser),
		// 				LastModifiedAt: to.Ptr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2020-01-02T17:18:19.123Z"); return t}()),
		// 				LastModifiedBy: to.Ptr("user2"),
		// 				LastModifiedByType: to.Ptr(armmobilenetwork.CreatedByTypeUser),
		// 			},
		// 			Location: to.Ptr("eastus"),
		// 			Tags: map[string]*string{
		// 			},
		// 			Identity: &armmobilenetwork.ManagedServiceIdentity{
		// 				Type: to.Ptr(armmobilenetwork.ManagedServiceIdentityTypeUserAssigned),
		// 				UserAssignedIdentities: map[string]*armmobilenetwork.UserAssignedIdentity{
		// 					"/subscriptions/00000000-0000-0000-0000-000000000000/resourcegroups/rg1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testUserAssignedManagedIdentity": &armmobilenetwork.UserAssignedIdentity{
		// 						ClientID: to.Ptr("12345678-abcd-dcba-abcd-0123456789ef"),
		// 						PrincipalID: to.Ptr("12345678-abcd-dcba-abcd-0123456789ef"),
		// 					},
		// 				},
		// 			},
		// 			Properties: &armmobilenetwork.SimGroupPropertiesFormat{
		// 				EncryptionKey: &armmobilenetwork.KeyVaultKey{
		// 					KeyURL: to.Ptr("https://contosovault.vault.azure.net/keys/azureKey"),
		// 				},
		// 				MobileNetwork: &armmobilenetwork.ResourceID{
		// 					ID: to.Ptr("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork"),
		// 				},
		// 				ProvisioningState: to.Ptr(armmobilenetwork.ProvisioningStateSucceeded),
		// 			},
		// 	}},
		// }
	}
}
