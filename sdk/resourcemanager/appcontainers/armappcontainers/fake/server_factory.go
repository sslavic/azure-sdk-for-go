//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// ServerFactory is a fake server for instances of the armappcontainers.ClientFactory type.
type ServerFactory struct {
	AppResiliencyServer                                AppResiliencyServer
	AvailableWorkloadProfilesServer                    AvailableWorkloadProfilesServer
	BillingMetersServer                                BillingMetersServer
	BuildAuthTokenServer                               BuildAuthTokenServer
	BuildersServer                                     BuildersServer
	BuildsByBuilderResourceServer                      BuildsByBuilderResourceServer
	BuildsServer                                       BuildsServer
	CertificatesServer                                 CertificatesServer
	ConnectedEnvironmentsCertificatesServer            ConnectedEnvironmentsCertificatesServer
	ConnectedEnvironmentsServer                        ConnectedEnvironmentsServer
	ConnectedEnvironmentsDaprComponentsServer          ConnectedEnvironmentsDaprComponentsServer
	ConnectedEnvironmentsStoragesServer                ConnectedEnvironmentsStoragesServer
	ContainerAppsAPIServer                             ContainerAppsAPIServer
	ContainerAppsAuthConfigsServer                     ContainerAppsAuthConfigsServer
	ContainerAppsBuildsByContainerAppServer            ContainerAppsBuildsByContainerAppServer
	ContainerAppsBuildsServer                          ContainerAppsBuildsServer
	ContainerAppsServer                                ContainerAppsServer
	ContainerAppsDiagnosticsServer                     ContainerAppsDiagnosticsServer
	ContainerAppsPatchesServer                         ContainerAppsPatchesServer
	ContainerAppsRevisionReplicasServer                ContainerAppsRevisionReplicasServer
	ContainerAppsRevisionsServer                       ContainerAppsRevisionsServer
	ContainerAppsSessionPoolsServer                    ContainerAppsSessionPoolsServer
	ContainerAppsSourceControlsServer                  ContainerAppsSourceControlsServer
	DaprComponentResiliencyPoliciesServer              DaprComponentResiliencyPoliciesServer
	DaprComponentsServer                               DaprComponentsServer
	DaprSubscriptionsServer                            DaprSubscriptionsServer
	DotNetComponentsServer                             DotNetComponentsServer
	FunctionsExtensionServer                           FunctionsExtensionServer
	JavaComponentsServer                               JavaComponentsServer
	JobsServer                                         JobsServer
	JobsExecutionsServer                               JobsExecutionsServer
	LogicAppsServer                                    LogicAppsServer
	ManagedCertificatesServer                          ManagedCertificatesServer
	ManagedEnvironmentDiagnosticsServer                ManagedEnvironmentDiagnosticsServer
	ManagedEnvironmentPrivateEndpointConnectionsServer ManagedEnvironmentPrivateEndpointConnectionsServer
	ManagedEnvironmentPrivateLinkResourcesServer       ManagedEnvironmentPrivateLinkResourcesServer
	ManagedEnvironmentUsagesServer                     ManagedEnvironmentUsagesServer
	ManagedEnvironmentsServer                          ManagedEnvironmentsServer
	ManagedEnvironmentsDiagnosticsServer               ManagedEnvironmentsDiagnosticsServer
	ManagedEnvironmentsStoragesServer                  ManagedEnvironmentsStoragesServer
	NamespacesServer                                   NamespacesServer
	OperationsServer                                   OperationsServer
	UsagesServer                                       UsagesServer
}

// NewServerFactoryTransport creates a new instance of ServerFactoryTransport with the provided implementation.
// The returned ServerFactoryTransport instance is connected to an instance of armappcontainers.ClientFactory via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewServerFactoryTransport(srv *ServerFactory) *ServerFactoryTransport {
	return &ServerFactoryTransport{
		srv: srv,
	}
}

// ServerFactoryTransport connects instances of armappcontainers.ClientFactory to instances of ServerFactory.
// Don't use this type directly, use NewServerFactoryTransport instead.
type ServerFactoryTransport struct {
	srv                                                  *ServerFactory
	trMu                                                 sync.Mutex
	trAppResiliencyServer                                *AppResiliencyServerTransport
	trAvailableWorkloadProfilesServer                    *AvailableWorkloadProfilesServerTransport
	trBillingMetersServer                                *BillingMetersServerTransport
	trBuildAuthTokenServer                               *BuildAuthTokenServerTransport
	trBuildersServer                                     *BuildersServerTransport
	trBuildsByBuilderResourceServer                      *BuildsByBuilderResourceServerTransport
	trBuildsServer                                       *BuildsServerTransport
	trCertificatesServer                                 *CertificatesServerTransport
	trConnectedEnvironmentsCertificatesServer            *ConnectedEnvironmentsCertificatesServerTransport
	trConnectedEnvironmentsServer                        *ConnectedEnvironmentsServerTransport
	trConnectedEnvironmentsDaprComponentsServer          *ConnectedEnvironmentsDaprComponentsServerTransport
	trConnectedEnvironmentsStoragesServer                *ConnectedEnvironmentsStoragesServerTransport
	trContainerAppsAPIServer                             *ContainerAppsAPIServerTransport
	trContainerAppsAuthConfigsServer                     *ContainerAppsAuthConfigsServerTransport
	trContainerAppsBuildsByContainerAppServer            *ContainerAppsBuildsByContainerAppServerTransport
	trContainerAppsBuildsServer                          *ContainerAppsBuildsServerTransport
	trContainerAppsServer                                *ContainerAppsServerTransport
	trContainerAppsDiagnosticsServer                     *ContainerAppsDiagnosticsServerTransport
	trContainerAppsPatchesServer                         *ContainerAppsPatchesServerTransport
	trContainerAppsRevisionReplicasServer                *ContainerAppsRevisionReplicasServerTransport
	trContainerAppsRevisionsServer                       *ContainerAppsRevisionsServerTransport
	trContainerAppsSessionPoolsServer                    *ContainerAppsSessionPoolsServerTransport
	trContainerAppsSourceControlsServer                  *ContainerAppsSourceControlsServerTransport
	trDaprComponentResiliencyPoliciesServer              *DaprComponentResiliencyPoliciesServerTransport
	trDaprComponentsServer                               *DaprComponentsServerTransport
	trDaprSubscriptionsServer                            *DaprSubscriptionsServerTransport
	trDotNetComponentsServer                             *DotNetComponentsServerTransport
	trFunctionsExtensionServer                           *FunctionsExtensionServerTransport
	trJavaComponentsServer                               *JavaComponentsServerTransport
	trJobsServer                                         *JobsServerTransport
	trJobsExecutionsServer                               *JobsExecutionsServerTransport
	trLogicAppsServer                                    *LogicAppsServerTransport
	trManagedCertificatesServer                          *ManagedCertificatesServerTransport
	trManagedEnvironmentDiagnosticsServer                *ManagedEnvironmentDiagnosticsServerTransport
	trManagedEnvironmentPrivateEndpointConnectionsServer *ManagedEnvironmentPrivateEndpointConnectionsServerTransport
	trManagedEnvironmentPrivateLinkResourcesServer       *ManagedEnvironmentPrivateLinkResourcesServerTransport
	trManagedEnvironmentUsagesServer                     *ManagedEnvironmentUsagesServerTransport
	trManagedEnvironmentsServer                          *ManagedEnvironmentsServerTransport
	trManagedEnvironmentsDiagnosticsServer               *ManagedEnvironmentsDiagnosticsServerTransport
	trManagedEnvironmentsStoragesServer                  *ManagedEnvironmentsStoragesServerTransport
	trNamespacesServer                                   *NamespacesServerTransport
	trOperationsServer                                   *OperationsServerTransport
	trUsagesServer                                       *UsagesServerTransport
}

// Do implements the policy.Transporter interface for ServerFactoryTransport.
func (s *ServerFactoryTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	client := method[:strings.Index(method, ".")]
	var resp *http.Response
	var err error

	switch client {
	case "AppResiliencyClient":
		initServer(s, &s.trAppResiliencyServer, func() *AppResiliencyServerTransport {
			return NewAppResiliencyServerTransport(&s.srv.AppResiliencyServer)
		})
		resp, err = s.trAppResiliencyServer.Do(req)
	case "AvailableWorkloadProfilesClient":
		initServer(s, &s.trAvailableWorkloadProfilesServer, func() *AvailableWorkloadProfilesServerTransport {
			return NewAvailableWorkloadProfilesServerTransport(&s.srv.AvailableWorkloadProfilesServer)
		})
		resp, err = s.trAvailableWorkloadProfilesServer.Do(req)
	case "BillingMetersClient":
		initServer(s, &s.trBillingMetersServer, func() *BillingMetersServerTransport {
			return NewBillingMetersServerTransport(&s.srv.BillingMetersServer)
		})
		resp, err = s.trBillingMetersServer.Do(req)
	case "BuildAuthTokenClient":
		initServer(s, &s.trBuildAuthTokenServer, func() *BuildAuthTokenServerTransport {
			return NewBuildAuthTokenServerTransport(&s.srv.BuildAuthTokenServer)
		})
		resp, err = s.trBuildAuthTokenServer.Do(req)
	case "BuildersClient":
		initServer(s, &s.trBuildersServer, func() *BuildersServerTransport { return NewBuildersServerTransport(&s.srv.BuildersServer) })
		resp, err = s.trBuildersServer.Do(req)
	case "BuildsByBuilderResourceClient":
		initServer(s, &s.trBuildsByBuilderResourceServer, func() *BuildsByBuilderResourceServerTransport {
			return NewBuildsByBuilderResourceServerTransport(&s.srv.BuildsByBuilderResourceServer)
		})
		resp, err = s.trBuildsByBuilderResourceServer.Do(req)
	case "BuildsClient":
		initServer(s, &s.trBuildsServer, func() *BuildsServerTransport { return NewBuildsServerTransport(&s.srv.BuildsServer) })
		resp, err = s.trBuildsServer.Do(req)
	case "CertificatesClient":
		initServer(s, &s.trCertificatesServer, func() *CertificatesServerTransport { return NewCertificatesServerTransport(&s.srv.CertificatesServer) })
		resp, err = s.trCertificatesServer.Do(req)
	case "ConnectedEnvironmentsCertificatesClient":
		initServer(s, &s.trConnectedEnvironmentsCertificatesServer, func() *ConnectedEnvironmentsCertificatesServerTransport {
			return NewConnectedEnvironmentsCertificatesServerTransport(&s.srv.ConnectedEnvironmentsCertificatesServer)
		})
		resp, err = s.trConnectedEnvironmentsCertificatesServer.Do(req)
	case "ConnectedEnvironmentsClient":
		initServer(s, &s.trConnectedEnvironmentsServer, func() *ConnectedEnvironmentsServerTransport {
			return NewConnectedEnvironmentsServerTransport(&s.srv.ConnectedEnvironmentsServer)
		})
		resp, err = s.trConnectedEnvironmentsServer.Do(req)
	case "ConnectedEnvironmentsDaprComponentsClient":
		initServer(s, &s.trConnectedEnvironmentsDaprComponentsServer, func() *ConnectedEnvironmentsDaprComponentsServerTransport {
			return NewConnectedEnvironmentsDaprComponentsServerTransport(&s.srv.ConnectedEnvironmentsDaprComponentsServer)
		})
		resp, err = s.trConnectedEnvironmentsDaprComponentsServer.Do(req)
	case "ConnectedEnvironmentsStoragesClient":
		initServer(s, &s.trConnectedEnvironmentsStoragesServer, func() *ConnectedEnvironmentsStoragesServerTransport {
			return NewConnectedEnvironmentsStoragesServerTransport(&s.srv.ConnectedEnvironmentsStoragesServer)
		})
		resp, err = s.trConnectedEnvironmentsStoragesServer.Do(req)
	case "ContainerAppsAPIClient":
		initServer(s, &s.trContainerAppsAPIServer, func() *ContainerAppsAPIServerTransport {
			return NewContainerAppsAPIServerTransport(&s.srv.ContainerAppsAPIServer)
		})
		resp, err = s.trContainerAppsAPIServer.Do(req)
	case "ContainerAppsAuthConfigsClient":
		initServer(s, &s.trContainerAppsAuthConfigsServer, func() *ContainerAppsAuthConfigsServerTransport {
			return NewContainerAppsAuthConfigsServerTransport(&s.srv.ContainerAppsAuthConfigsServer)
		})
		resp, err = s.trContainerAppsAuthConfigsServer.Do(req)
	case "ContainerAppsBuildsByContainerAppClient":
		initServer(s, &s.trContainerAppsBuildsByContainerAppServer, func() *ContainerAppsBuildsByContainerAppServerTransport {
			return NewContainerAppsBuildsByContainerAppServerTransport(&s.srv.ContainerAppsBuildsByContainerAppServer)
		})
		resp, err = s.trContainerAppsBuildsByContainerAppServer.Do(req)
	case "ContainerAppsBuildsClient":
		initServer(s, &s.trContainerAppsBuildsServer, func() *ContainerAppsBuildsServerTransport {
			return NewContainerAppsBuildsServerTransport(&s.srv.ContainerAppsBuildsServer)
		})
		resp, err = s.trContainerAppsBuildsServer.Do(req)
	case "ContainerAppsClient":
		initServer(s, &s.trContainerAppsServer, func() *ContainerAppsServerTransport {
			return NewContainerAppsServerTransport(&s.srv.ContainerAppsServer)
		})
		resp, err = s.trContainerAppsServer.Do(req)
	case "ContainerAppsDiagnosticsClient":
		initServer(s, &s.trContainerAppsDiagnosticsServer, func() *ContainerAppsDiagnosticsServerTransport {
			return NewContainerAppsDiagnosticsServerTransport(&s.srv.ContainerAppsDiagnosticsServer)
		})
		resp, err = s.trContainerAppsDiagnosticsServer.Do(req)
	case "ContainerAppsPatchesClient":
		initServer(s, &s.trContainerAppsPatchesServer, func() *ContainerAppsPatchesServerTransport {
			return NewContainerAppsPatchesServerTransport(&s.srv.ContainerAppsPatchesServer)
		})
		resp, err = s.trContainerAppsPatchesServer.Do(req)
	case "ContainerAppsRevisionReplicasClient":
		initServer(s, &s.trContainerAppsRevisionReplicasServer, func() *ContainerAppsRevisionReplicasServerTransport {
			return NewContainerAppsRevisionReplicasServerTransport(&s.srv.ContainerAppsRevisionReplicasServer)
		})
		resp, err = s.trContainerAppsRevisionReplicasServer.Do(req)
	case "ContainerAppsRevisionsClient":
		initServer(s, &s.trContainerAppsRevisionsServer, func() *ContainerAppsRevisionsServerTransport {
			return NewContainerAppsRevisionsServerTransport(&s.srv.ContainerAppsRevisionsServer)
		})
		resp, err = s.trContainerAppsRevisionsServer.Do(req)
	case "ContainerAppsSessionPoolsClient":
		initServer(s, &s.trContainerAppsSessionPoolsServer, func() *ContainerAppsSessionPoolsServerTransport {
			return NewContainerAppsSessionPoolsServerTransport(&s.srv.ContainerAppsSessionPoolsServer)
		})
		resp, err = s.trContainerAppsSessionPoolsServer.Do(req)
	case "ContainerAppsSourceControlsClient":
		initServer(s, &s.trContainerAppsSourceControlsServer, func() *ContainerAppsSourceControlsServerTransport {
			return NewContainerAppsSourceControlsServerTransport(&s.srv.ContainerAppsSourceControlsServer)
		})
		resp, err = s.trContainerAppsSourceControlsServer.Do(req)
	case "DaprComponentResiliencyPoliciesClient":
		initServer(s, &s.trDaprComponentResiliencyPoliciesServer, func() *DaprComponentResiliencyPoliciesServerTransport {
			return NewDaprComponentResiliencyPoliciesServerTransport(&s.srv.DaprComponentResiliencyPoliciesServer)
		})
		resp, err = s.trDaprComponentResiliencyPoliciesServer.Do(req)
	case "DaprComponentsClient":
		initServer(s, &s.trDaprComponentsServer, func() *DaprComponentsServerTransport {
			return NewDaprComponentsServerTransport(&s.srv.DaprComponentsServer)
		})
		resp, err = s.trDaprComponentsServer.Do(req)
	case "DaprSubscriptionsClient":
		initServer(s, &s.trDaprSubscriptionsServer, func() *DaprSubscriptionsServerTransport {
			return NewDaprSubscriptionsServerTransport(&s.srv.DaprSubscriptionsServer)
		})
		resp, err = s.trDaprSubscriptionsServer.Do(req)
	case "DotNetComponentsClient":
		initServer(s, &s.trDotNetComponentsServer, func() *DotNetComponentsServerTransport {
			return NewDotNetComponentsServerTransport(&s.srv.DotNetComponentsServer)
		})
		resp, err = s.trDotNetComponentsServer.Do(req)
	case "FunctionsExtensionClient":
		initServer(s, &s.trFunctionsExtensionServer, func() *FunctionsExtensionServerTransport {
			return NewFunctionsExtensionServerTransport(&s.srv.FunctionsExtensionServer)
		})
		resp, err = s.trFunctionsExtensionServer.Do(req)
	case "JavaComponentsClient":
		initServer(s, &s.trJavaComponentsServer, func() *JavaComponentsServerTransport {
			return NewJavaComponentsServerTransport(&s.srv.JavaComponentsServer)
		})
		resp, err = s.trJavaComponentsServer.Do(req)
	case "JobsClient":
		initServer(s, &s.trJobsServer, func() *JobsServerTransport { return NewJobsServerTransport(&s.srv.JobsServer) })
		resp, err = s.trJobsServer.Do(req)
	case "JobsExecutionsClient":
		initServer(s, &s.trJobsExecutionsServer, func() *JobsExecutionsServerTransport {
			return NewJobsExecutionsServerTransport(&s.srv.JobsExecutionsServer)
		})
		resp, err = s.trJobsExecutionsServer.Do(req)
	case "LogicAppsClient":
		initServer(s, &s.trLogicAppsServer, func() *LogicAppsServerTransport { return NewLogicAppsServerTransport(&s.srv.LogicAppsServer) })
		resp, err = s.trLogicAppsServer.Do(req)
	case "ManagedCertificatesClient":
		initServer(s, &s.trManagedCertificatesServer, func() *ManagedCertificatesServerTransport {
			return NewManagedCertificatesServerTransport(&s.srv.ManagedCertificatesServer)
		})
		resp, err = s.trManagedCertificatesServer.Do(req)
	case "ManagedEnvironmentDiagnosticsClient":
		initServer(s, &s.trManagedEnvironmentDiagnosticsServer, func() *ManagedEnvironmentDiagnosticsServerTransport {
			return NewManagedEnvironmentDiagnosticsServerTransport(&s.srv.ManagedEnvironmentDiagnosticsServer)
		})
		resp, err = s.trManagedEnvironmentDiagnosticsServer.Do(req)
	case "ManagedEnvironmentPrivateEndpointConnectionsClient":
		initServer(s, &s.trManagedEnvironmentPrivateEndpointConnectionsServer, func() *ManagedEnvironmentPrivateEndpointConnectionsServerTransport {
			return NewManagedEnvironmentPrivateEndpointConnectionsServerTransport(&s.srv.ManagedEnvironmentPrivateEndpointConnectionsServer)
		})
		resp, err = s.trManagedEnvironmentPrivateEndpointConnectionsServer.Do(req)
	case "ManagedEnvironmentPrivateLinkResourcesClient":
		initServer(s, &s.trManagedEnvironmentPrivateLinkResourcesServer, func() *ManagedEnvironmentPrivateLinkResourcesServerTransport {
			return NewManagedEnvironmentPrivateLinkResourcesServerTransport(&s.srv.ManagedEnvironmentPrivateLinkResourcesServer)
		})
		resp, err = s.trManagedEnvironmentPrivateLinkResourcesServer.Do(req)
	case "ManagedEnvironmentUsagesClient":
		initServer(s, &s.trManagedEnvironmentUsagesServer, func() *ManagedEnvironmentUsagesServerTransport {
			return NewManagedEnvironmentUsagesServerTransport(&s.srv.ManagedEnvironmentUsagesServer)
		})
		resp, err = s.trManagedEnvironmentUsagesServer.Do(req)
	case "ManagedEnvironmentsClient":
		initServer(s, &s.trManagedEnvironmentsServer, func() *ManagedEnvironmentsServerTransport {
			return NewManagedEnvironmentsServerTransport(&s.srv.ManagedEnvironmentsServer)
		})
		resp, err = s.trManagedEnvironmentsServer.Do(req)
	case "ManagedEnvironmentsDiagnosticsClient":
		initServer(s, &s.trManagedEnvironmentsDiagnosticsServer, func() *ManagedEnvironmentsDiagnosticsServerTransport {
			return NewManagedEnvironmentsDiagnosticsServerTransport(&s.srv.ManagedEnvironmentsDiagnosticsServer)
		})
		resp, err = s.trManagedEnvironmentsDiagnosticsServer.Do(req)
	case "ManagedEnvironmentsStoragesClient":
		initServer(s, &s.trManagedEnvironmentsStoragesServer, func() *ManagedEnvironmentsStoragesServerTransport {
			return NewManagedEnvironmentsStoragesServerTransport(&s.srv.ManagedEnvironmentsStoragesServer)
		})
		resp, err = s.trManagedEnvironmentsStoragesServer.Do(req)
	case "NamespacesClient":
		initServer(s, &s.trNamespacesServer, func() *NamespacesServerTransport { return NewNamespacesServerTransport(&s.srv.NamespacesServer) })
		resp, err = s.trNamespacesServer.Do(req)
	case "OperationsClient":
		initServer(s, &s.trOperationsServer, func() *OperationsServerTransport { return NewOperationsServerTransport(&s.srv.OperationsServer) })
		resp, err = s.trOperationsServer.Do(req)
	case "UsagesClient":
		initServer(s, &s.trUsagesServer, func() *UsagesServerTransport { return NewUsagesServerTransport(&s.srv.UsagesServer) })
		resp, err = s.trUsagesServer.Do(req)
	default:
		err = fmt.Errorf("unhandled client %s", client)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func initServer[T any](s *ServerFactoryTransport, dst **T, src func() *T) {
	s.trMu.Lock()
	if *dst == nil {
		*dst = src()
	}
	s.trMu.Unlock()
}
