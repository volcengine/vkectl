package resource
import (
	"net/url"

	"github.com/volcengine/vkectl/pkg/client"

	"github.com/volcengine/vkectl/pkg/model/resource/kitex_gen/addon"
	"github.com/volcengine/vkectl/pkg/model/resource/kitex_gen/baremachine"
	"github.com/volcengine/vkectl/pkg/model/resource/kitex_gen/clb"
	"github.com/volcengine/vkectl/pkg/model/resource/kitex_gen/cluster"
	"github.com/volcengine/vkectl/pkg/model/resource/kitex_gen/event"
	"github.com/volcengine/vkectl/pkg/model/resource/kitex_gen/helper"
	"github.com/volcengine/vkectl/pkg/model/resource/kitex_gen/instance"
	"github.com/volcengine/vkectl/pkg/model/resource/kitex_gen/network"
	"github.com/volcengine/vkectl/pkg/model/resource/kitex_gen/overview"
	"github.com/volcengine/vkectl/pkg/model/resource/kitex_gen/quota"
	"github.com/volcengine/vkectl/pkg/model/resource/kitex_gen/rbac"
	"github.com/volcengine/vkectl/pkg/model/resource/kitex_gen/storage"
	"github.com/volcengine/vkectl/pkg/model/resource/kitex_gen/vpc"
)

// ResourceserviceRawApi is a base client
type ResourceserviceRawApi struct {
	Client *client.Client
}

// NewAPIClient 生成一个客户端
func NewAPIClient(ak, sk, host, service, region string) *ResourceserviceRawApi {
	c := client.NewBaseClient()
	c.ServiceInfo = client.NewServiceInfo()
	c.ServiceInfo.Host = host
	c.ServiceInfo.Credentials.AccessKeyID = ak
	c.ServiceInfo.Credentials.SecretAccessKey = sk
	c.ServiceInfo.Credentials.Service = service
	c.ServiceInfo.Credentials.Region = region
	c.SdkVersion = client.DefaultSdkVersion

	return &ResourceserviceRawApi{Client: c}
}
func (p *ResourceserviceRawApi) GetClusterOverview(body string, query url.Values) (r *overview.GetClusterOverviewResponse, statusCode int, err error) {
	action := "GetClusterOverview"
	r = &overview.GetClusterOverviewResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) GetKubeConfig(body string, query url.Values) (r *cluster.GetKubeConfigResponse, statusCode int, err error) {
	action := "GetKubeConfig"
	r = &cluster.GetKubeConfigResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) RevokeKubeConfig(body string, query url.Values) (r *cluster.RevokeKubeConfigResponse, statusCode int, err error) {
	action := "RevokeKubeConfig"
	r = &cluster.RevokeKubeConfigResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListKubeConfig(body string, query url.Values) (r *cluster.ListKubeConfigResponse, statusCode int, err error) {
	action := "ListKubeConfig"
	r = &cluster.ListKubeConfigResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) CreateCluster(body string, query url.Values) (r *cluster.CreateClusterResponse, statusCode int, err error) {
	action := "CreateCluster"
	r = &cluster.CreateClusterResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) RegisterCluster(body string, query url.Values) (r *cluster.RegisterClusterResponse, statusCode int, err error) {
	action := "RegisterCluster"
	r = &cluster.RegisterClusterResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) UpdateCluster(body string, query url.Values) (r *cluster.UpdateClusterResponse, statusCode int, err error) {
	action := "UpdateCluster"
	r = &cluster.UpdateClusterResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) UpgradeCluster(body string, query url.Values) (r *cluster.UpgradeClusterResponse, statusCode int, err error) {
	action := "UpgradeCluster"
	r = &cluster.UpgradeClusterResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) DeleteCluster(body string, query url.Values) (r *cluster.DeleteClusterResponse, statusCode int, err error) {
	action := "DeleteCluster"
	r = &cluster.DeleteClusterResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) GetCluster(body string, query url.Values) (r *cluster.GetClusterResponse, statusCode int, err error) {
	action := "GetCluster"
	r = &cluster.GetClusterResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListCluster(body string, query url.Values) (r *cluster.ListClusterResponse, statusCode int, err error) {
	action := "ListCluster"
	r = &cluster.ListClusterResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) GetClusterDeployProgress(body string, query url.Values) (r *cluster.GetClusterDeployProgressResponse, statusCode int, err error) {
	action := "GetClusterDeployProgress"
	r = &cluster.GetClusterDeployProgressResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListClusterKubernetesVersion(body string, query url.Values) (r *cluster.ListClusterKubernetesVersionResponse, statusCode int, err error) {
	action := "ListClusterKubernetesVersion"
	r = &cluster.ListClusterKubernetesVersionResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListClusterNetworkCidr(body string, query url.Values) (r *cluster.ListClusterNetworkCidrResponse, statusCode int, err error) {
	action := "ListClusterNetworkCidr"
	r = &cluster.ListClusterNetworkCidrResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListSupportGpuModel(body string, query url.Values) (r *cluster.ListSupportGpuModelResponse, statusCode int, err error) {
	action := "ListSupportGpuModel"
	r = &cluster.ListSupportGpuModelResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListClusterNode(body string, query url.Values) (r *cluster.ListClusterNodeResponse, statusCode int, err error) {
	action := "ListClusterNode"
	r = &cluster.ListClusterNodeResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) AddClusterNode(body string, query url.Values) (r *cluster.AddClusterNodeResponse, statusCode int, err error) {
	action := "AddClusterNode"
	r = &cluster.AddClusterNodeResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) GetClusterNode(body string, query url.Values) (r *cluster.GetClusterNodeResponse, statusCode int, err error) {
	action := "GetClusterNode"
	r = &cluster.GetClusterNodeResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) DeleteClusterNode(body string, query url.Values) (r *cluster.DeleteClusterNodeResponse, statusCode int, err error) {
	action := "DeleteClusterNode"
	r = &cluster.DeleteClusterNodeResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListClusterNodeLabel(body string, query url.Values) (r *cluster.ListClusterNodeLabelResponse, statusCode int, err error) {
	action := "ListClusterNodeLabel"
	r = &cluster.ListClusterNodeLabelResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) UpdateClusterNode(body string, query url.Values) (r *cluster.UpdateClusterNodeResponse, statusCode int, err error) {
	action := "UpdateClusterNode"
	r = &cluster.UpdateClusterNodeResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) GetAutoScalingRule(body string, query url.Values) (r *cluster.GetAutoScalingRuleResponse, statusCode int, err error) {
	action := "GetAutoScalingRule"
	r = &cluster.GetAutoScalingRuleResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) UpdateAutoScalingRule(body string, query url.Values) (r *cluster.UpdateAutoScalingRuleResponse, statusCode int, err error) {
	action := "UpdateAutoScalingRule"
	r = &cluster.UpdateAutoScalingRuleResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) NodePoolScaleUp(body string, query url.Values) (r *cluster.NodePoolScaleUpResponse, statusCode int, err error) {
	action := "NodePoolScaleUp"
	r = &cluster.NodePoolScaleUpResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) NodePoolScaleDown(body string, query url.Values) (r *cluster.NodePoolScaleDownResponse, statusCode int, err error) {
	action := "NodePoolScaleDown"
	r = &cluster.NodePoolScaleDownResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListNodePool(body string, query url.Values) (r *cluster.ListNodePoolResponse, statusCode int, err error) {
	action := "ListNodePool"
	r = &cluster.ListNodePoolResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) CreateNodePool(body string, query url.Values) (r *cluster.CreateNodePoolResponse, statusCode int, err error) {
	action := "CreateNodePool"
	r = &cluster.CreateNodePoolResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) GetNodePool(body string, query url.Values) (r *cluster.GetNodePoolResponse, statusCode int, err error) {
	action := "GetNodePool"
	r = &cluster.GetNodePoolResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) UpdateNodePool(body string, query url.Values) (r *cluster.UpdateNodePoolResponse, statusCode int, err error) {
	action := "UpdateNodePool"
	r = &cluster.UpdateNodePoolResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) DeleteNodePool(body string, query url.Values) (r *cluster.DeleteNodePoolResponse, statusCode int, err error) {
	action := "DeleteNodePool"
	r = &cluster.DeleteNodePoolResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListNodePoolScalingRecords(body string, query url.Values) (r *cluster.ListNodePoolScalingRecordResponse, statusCode int, err error) {
	action := "ListNodePoolScalingRecords"
	r = &cluster.ListNodePoolScalingRecordResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListNodePoolNode(body string, query url.Values) (r *cluster.ListNodePoolNodeResponse, statusCode int, err error) {
	action := "ListNodePoolNode"
	r = &cluster.ListNodePoolNodeResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListNamespace(body string, query url.Values) (r *cluster.ListNamespaceResponse, statusCode int, err error) {
	action := "ListNamespace"
	r = &cluster.ListNamespaceResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) CreateNamespace(body string, query url.Values) (r *cluster.CreateNamespaceResponse, statusCode int, err error) {
	action := "CreateNamespace"
	r = &cluster.CreateNamespaceResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) GetNamespace(body string, query url.Values) (r *cluster.GetNamespaceResponse, statusCode int, err error) {
	action := "GetNamespace"
	r = &cluster.GetNamespaceResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) UpdateNamespace(body string, query url.Values) (r *cluster.UpdateNamespaceResponse, statusCode int, err error) {
	action := "UpdateNamespace"
	r = &cluster.UpdateNamespaceResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) UpdateNamespaceResource(body string, query url.Values) (r *cluster.UpdateNamespaceResourceResponse, statusCode int, err error) {
	action := "UpdateNamespaceResource"
	r = &cluster.UpdateNamespaceResourceResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) GetNamespaceResource(body string, query url.Values) (r *cluster.GetNamespaceResourceResponse, statusCode int, err error) {
	action := "GetNamespaceResource"
	r = &cluster.GetNamespaceResourceResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) DeleteNamespace(body string, query url.Values) (r *cluster.DeleteNamespaceResponse, statusCode int, err error) {
	action := "DeleteNamespace"
	r = &cluster.DeleteNamespaceResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListResourceQuota(body string, query url.Values) (r *cluster.ListResourceQuotaResponse, statusCode int, err error) {
	action := "ListResourceQuota"
	r = &cluster.ListResourceQuotaResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListIngressTypes(body string, query url.Values) (r *network.ListIngressTypeResponse, statusCode int, err error) {
	action := "ListIngressTypes"
	r = &network.ListIngressTypeResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListIngress(body string, query url.Values) (r *network.ListIngressResponse, statusCode int, err error) {
	action := "ListIngress"
	r = &network.ListIngressResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) CreateIngress(body string, query url.Values) (r *network.CreateIngressResponse, statusCode int, err error) {
	action := "CreateIngress"
	r = &network.CreateIngressResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) GetIngress(body string, query url.Values) (r *network.GetIngressResponse, statusCode int, err error) {
	action := "GetIngress"
	r = &network.GetIngressResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) UpdateIngress(body string, query url.Values) (r *network.UpdateIngressResponse, statusCode int, err error) {
	action := "UpdateIngress"
	r = &network.UpdateIngressResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) DeleteIngress(body string, query url.Values) (r *network.DeleteIngressResponse, statusCode int, err error) {
	action := "DeleteIngress"
	r = &network.DeleteIngressResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) GetSecretDomains(body string, query url.Values) (r *network.GetSecretDomainsResponse, statusCode int, err error) {
	action := "GetSecretDomains"
	r = &network.GetSecretDomainsResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListStorageClass(body string, query url.Values) (r *storage.ListStorageClassResponse, statusCode int, err error) {
	action := "ListStorageClass"
	r = &storage.ListStorageClassResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) CreateStorageClass(body string, query url.Values) (r *storage.CreateStorageClassResponse, statusCode int, err error) {
	action := "CreateStorageClass"
	r = &storage.CreateStorageClassResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) GetStorageClass(body string, query url.Values) (r *storage.GetStorageClassResponse, statusCode int, err error) {
	action := "GetStorageClass"
	r = &storage.GetStorageClassResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) DeleteStorageClass(body string, query url.Values) (r *storage.DeleteStorageClassResponse, statusCode int, err error) {
	action := "DeleteStorageClass"
	r = &storage.DeleteStorageClassResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListPersistentVolume(body string, query url.Values) (r *storage.ListPersistentVolumeResponse, statusCode int, err error) {
	action := "ListPersistentVolume"
	r = &storage.ListPersistentVolumeResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) CreatePersistentVolume(body string, query url.Values) (r *storage.CreatePersistentVolumeResponse, statusCode int, err error) {
	action := "CreatePersistentVolume"
	r = &storage.CreatePersistentVolumeResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) GetPersistentVolume(body string, query url.Values) (r *storage.GetPersistentVolumeResponse, statusCode int, err error) {
	action := "GetPersistentVolume"
	r = &storage.GetPersistentVolumeResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) DeletePersistentVolume(body string, query url.Values) (r *storage.DeletePersistentVolumeResponse, statusCode int, err error) {
	action := "DeletePersistentVolume"
	r = &storage.DeletePersistentVolumeResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListPersistentVolumeClaim(body string, query url.Values) (r *storage.ListPersistentVolumeClaimResponse, statusCode int, err error) {
	action := "ListPersistentVolumeClaim"
	r = &storage.ListPersistentVolumeClaimResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) CreatePersistentVolumeClaim(body string, query url.Values) (r *storage.CreatePersistentVolumeClaimResponse, statusCode int, err error) {
	action := "CreatePersistentVolumeClaim"
	r = &storage.CreatePersistentVolumeClaimResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) GetPersistentVolumeClaim(body string, query url.Values) (r *storage.GetPersistentVolumeClaimResponse, statusCode int, err error) {
	action := "GetPersistentVolumeClaim"
	r = &storage.GetPersistentVolumeClaimResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) DeletePersistentVolumeClaim(body string, query url.Values) (r *storage.DeletePersistentVolumeClaimResponse, statusCode int, err error) {
	action := "DeletePersistentVolumeClaim"
	r = &storage.DeletePersistentVolumeClaimResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListCephCluster(body string, query url.Values) (r *storage.ListCephClusterResponse, statusCode int, err error) {
	action := "ListCephCluster"
	r = &storage.ListCephClusterResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListEvent(body string, query url.Values) (r *event.ListEventResponse, statusCode int, err error) {
	action := "ListEvent"
	r = &event.ListEventResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListUserRbacs(body string, query url.Values) (r *rbac.ListUserRbacResponse, statusCode int, err error) {
	action := "ListUserRbacs"
	r = &rbac.ListUserRbacResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) CreateUserRbac(body string, query url.Values) (r *rbac.CreateUserRbacResponse, statusCode int, err error) {
	action := "CreateUserRbac"
	r = &rbac.CreateUserRbacResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) UpdateUserRbac(body string, query url.Values) (r *rbac.UpdateUserRbacResponse, statusCode int, err error) {
	action := "UpdateUserRbac"
	r = &rbac.UpdateUserRbacResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) DeleteUserRbac(body string, query url.Values) (r *rbac.DeleteUserRbacResponse, statusCode int, err error) {
	action := "DeleteUserRbac"
	r = &rbac.DeleteUserRbacResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListCustomRoles(body string, query url.Values) (r *rbac.ListCustomRolesResponse, statusCode int, err error) {
	action := "ListCustomRoles"
	r = &rbac.ListCustomRolesResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListSupportedAddons(body string, query url.Values) (r *addon.ListSupportedAddonResponse, statusCode int, err error) {
	action := "ListSupportedAddons"
	r = &addon.ListSupportedAddonResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListAddons(body string, query url.Values) (r *addon.ListAddonResponse, statusCode int, err error) {
	action := "ListAddons"
	r = &addon.ListAddonResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) InstallAddon(body string, query url.Values) (r *addon.InstallAddonResponse, statusCode int, err error) {
	action := "InstallAddon"
	r = &addon.InstallAddonResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) UninstallAddon(body string, query url.Values) (r *addon.UninstallAddonResponse, statusCode int, err error) {
	action := "UninstallAddon"
	r = &addon.UninstallAddonResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) UpgradeAddon(body string, query url.Values) (r *addon.UpgradeAddonResponse, statusCode int, err error) {
	action := "UpgradeAddon"
	r = &addon.UpgradeAddonResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) GetAddon(body string, query url.Values) (r *addon.GetAddonResponse, statusCode int, err error) {
	action := "GetAddon"
	r = &addon.GetAddonResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) CheckResourceExist(body string, query url.Values) (r *helper.CheckResourceExistResponse, statusCode int, err error) {
	action := "CheckResourceExist"
	r = &helper.CheckResourceExistResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListInstances(body string, query url.Values) (r *instance.ListInstancesResponse, statusCode int, err error) {
	action := "ListInstances"
	r = &instance.ListInstancesResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListFlavors(body string, query url.Values) (r *instance.ListFlavorsResponse, statusCode int, err error) {
	action := "ListFlavors"
	r = &instance.ListFlavorsResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) GetInstanceConsole(body string, query url.Values) (r *instance.GetInstanceConsoleResponse, statusCode int, err error) {
	action := "GetInstanceConsole"
	r = &instance.GetInstanceConsoleResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListZones(body string, query url.Values) (r *instance.ListZonesResponse, statusCode int, err error) {
	action := "ListZones"
	r = &instance.ListZonesResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListVolumes(body string, query url.Values) (r *instance.ListVolumesResponse, statusCode int, err error) {
	action := "ListVolumes"
	r = &instance.ListVolumesResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListSubnets(body string, query url.Values) (r *vpc.ListSubnetsResponse, statusCode int, err error) {
	action := "ListSubnets"
	r = &vpc.ListSubnetsResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListElasticIPPools(body string, query url.Values) (r *vpc.ListElasticIPPoolsResponse, statusCode int, err error) {
	action := "ListElasticIPPools"
	r = &vpc.ListElasticIPPoolsResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListVpcs(body string, query url.Values) (r *vpc.ListVpcsResponse, statusCode int, err error) {
	action := "ListVpcs"
	r = &vpc.ListVpcsResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListSecurityGroups(body string, query url.Values) (r *vpc.ListSecurityGroupsResponse, statusCode int, err error) {
	action := "ListSecurityGroups"
	r = &vpc.ListSecurityGroupsResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListClbs(body string, query url.Values) (r *clb.ListClbResponse, statusCode int, err error) {
	action := "ListClbs"
	r = &clb.ListClbResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListQuotas(body string, query url.Values) (r *quota.ListQuotasResponse, statusCode int, err error) {
	action := "ListQuotas"
	r = &quota.ListQuotasResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) GetQuota(body string, query url.Values) (r *quota.GetQuotaResponse, statusCode int, err error) {
	action := "GetQuota"
	r = &quota.GetQuotaResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) CreateBareMachine(body string, query url.Values) (r *baremachine.CreateBareMachineResponse, statusCode int, err error) {
	action := "CreateBareMachine"
	r = &baremachine.CreateBareMachineResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) UpdateBareMachine(body string, query url.Values) (r *baremachine.UpdateBareMachineResponse, statusCode int, err error) {
	action := "UpdateBareMachine"
	r = &baremachine.UpdateBareMachineResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) DeleteBareMachine(body string, query url.Values) (r *baremachine.DeleteBareMachineResponse, statusCode int, err error) {
	action := "DeleteBareMachine"
	r = &baremachine.DeleteBareMachineResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) GetBareMachine(body string, query url.Values) (r *baremachine.GetBareMachineResponse, statusCode int, err error) {
	action := "GetBareMachine"
	r = &baremachine.GetBareMachineResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListBareMachine(body string, query url.Values) (r *baremachine.ListBareMachineResponse, statusCode int, err error) {
	action := "ListBareMachine"
	r = &baremachine.ListBareMachineResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) GetBareMachineSshPublicKey(body string, query url.Values) (r *baremachine.GetBareMachineSshPublicKeyResponse, statusCode int, err error) {
	action := "GetBareMachineSshPublicKey"
	r = &baremachine.GetBareMachineSshPublicKeyResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) ListBareMachineExpectOs(body string, query url.Values) (r *baremachine.ListBareMachineExpectOsResponse, statusCode int, err error) {
	action := "ListBareMachineExpectOs"
	r = &baremachine.ListBareMachineExpectOsResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) GetBareMachineImportExcelTemplate(body string, query url.Values) (r *baremachine.GetBareMachineImportExcelTemplateResponse, statusCode int, err error) {
	action := "GetBareMachineImportExcelTemplate"
	r = &baremachine.GetBareMachineImportExcelTemplateResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) CheckCidrConflict(body string, query url.Values) (r *cluster.CheckCidrConflictResponse, statusCode int, err error) {
	action := "CheckCidrConflict"
	r = &cluster.CheckCidrConflictResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *ResourceserviceRawApi) RecommendCidr(body string, query url.Values) (r *cluster.RecommendCidrResponse, statusCode int, err error) {
	action := "RecommendCidr"
	r = &cluster.RecommendCidrResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}
