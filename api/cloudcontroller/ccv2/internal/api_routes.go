package internal

import (
	"net/http"

	"github.com/tedsuo/rata"
)

const (
	GetAppInstancesRequest               = "GetAppInstances"
	GetAppStatsRequest                   = "GetAppStats"
	GetAppRequest                        = "GetApp"
	GetRouteAppsRequest                  = "GetRouteApps"
	GetAppsRequest                       = "GetApps"
	PutSecurityGroupSpaceRequest         = "PutSecurityGroupSpace"
	DeleteOrganizationRequest            = "DeleteOrganization"
	DeleteRouteRequest                   = "DeleteRoute"
	DeleteServiceBindingRequest          = "DeleteServiceBinding"
	GetInfoRequest                       = "GetInfo"
	GetJobRequest                        = "GetJob"
	GetOrganizationsRequest              = "GetOrganizations"
	GetQuotaDefinition                   = "GetQuotaDefinition"
	GetPrivateDomainRequest              = "GetPrivateDomain"
	GetOrganizationPrivateDomainsRequest = "GetOrganizationPrivateDomains"
	GetRouteRouteMappingsRequest         = "GetRouteRouteMappings"
	GetAppRoutesRequest                  = "GetAppRoutes"
	GetSpaceRoutesRequest                = "GetSpaceRoutes"
	GetSecurityGroupsRequest             = "GetSecurityGroups"
	GetServiceBindingsRequest            = "GetServiceBindings"
	GetServiceInstancesRequest           = "GetServiceInstances"
	GetSharedDomainRequest               = "GetSharedDomain"
	GetSharedDomainsRequest              = "GetSharedDomains"
	GetSpaceServiceInstancesRequest      = "GetSpaceServiceInstances"
	GetSpacesRequest                     = "GetSpaces"
	GetStackRequest                      = "GetStack"
	PutAppRequest                        = "PutApp"
	GetUsersRequest                      = "GetUsers"
)

// APIRoutes is a list of routes used by the rata library to construct request
// URLs.
// Naming convention: Method + [ParentResourceInstance] + Resource
var APIRoutes = rata.Routes{
	{Path: "/v2/apps", Method: http.MethodGet, Name: GetAppsRequest},
	{Path: "/v2/apps/:app_guid", Method: http.MethodGet, Name: GetAppRequest},
	{Path: "/v2/apps/:app_guid", Method: http.MethodPut, Name: PutAppRequest},
	{Path: "/v2/apps/:app_guid/instances", Method: http.MethodGet, Name: GetAppInstancesRequest},
	{Path: "/v2/apps/:app_guid/routes", Method: http.MethodGet, Name: GetAppRoutesRequest},
	{Path: "/v2/apps/:app_guid/stats", Method: http.MethodGet, Name: GetAppStatsRequest},
	{Path: "/v2/info", Method: http.MethodGet, Name: GetInfoRequest},
	{Path: "/v2/jobs/:job_guid", Method: http.MethodGet, Name: GetJobRequest},
	{Path: "/v2/organizations", Method: http.MethodGet, Name: GetOrganizationsRequest},
	{Path: "/v2/organizations/:organization_guid", Method: http.MethodDelete, Name: DeleteOrganizationRequest},
	{Path: "/v2/organizations/:organization_guid/private_domains", Method: http.MethodGet, Name: GetOrganizationPrivateDomainsRequest},
	{Path: "/v2/private_domains/:private_domain_guid", Method: http.MethodGet, Name: GetPrivateDomainRequest},
	{Path: "/v2/quota_definitions/:organization_quota_guid", Method: http.MethodGet, Name: GetQuotaDefinition},
	{Path: "/v2/routes/:route_guid", Method: http.MethodDelete, Name: DeleteRouteRequest},
	{Path: "/v2/routes/:route_guid/apps", Method: http.MethodGet, Name: GetRouteAppsRequest},
	{Path: "/v2/routes/:route_guid/route_mappings", Method: http.MethodGet, Name: GetRouteRouteMappingsRequest},
	{Path: "/v2/security_groups", Method: http.MethodGet, Name: GetSecurityGroupsRequest},
	{Path: "/v2/security_groups/:security_group_guid/spaces/:space_guid", Method: http.MethodPut, Name: PutSecurityGroupSpaceRequest},
	{Path: "/v2/service_bindings", Method: http.MethodGet, Name: GetServiceBindingsRequest},
	{Path: "/v2/service_bindings/:service_binding_guid", Method: http.MethodDelete, Name: DeleteServiceBindingRequest},
	{Path: "/v2/service_instances", Method: http.MethodGet, Name: GetServiceInstancesRequest},
	{Path: "/v2/shared_domains/:shared_domain_guid", Method: http.MethodGet, Name: GetSharedDomainRequest},
	{Path: "/v2/shared_domains", Method: http.MethodGet, Name: GetSharedDomainsRequest},
	{Path: "/v2/spaces", Method: http.MethodGet, Name: GetSpacesRequest},
	{Path: "/v2/spaces/:guid/service_instances", Method: http.MethodGet, Name: GetSpaceServiceInstancesRequest},
	{Path: "/v2/spaces/:space_guid/routes", Method: http.MethodGet, Name: GetSpaceRoutesRequest},
	{Path: "/v2/stacks/:stack_guid", Method: http.MethodGet, Name: GetStackRequest},
	{Path: "/v2/users", Method: http.MethodPost, Name: GetUsersRequest},
}
