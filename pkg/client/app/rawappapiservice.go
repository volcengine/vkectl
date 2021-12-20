package app
import (
	"net/url"

	"github.com/volcengine/vkectl/pkg/client"

	"github.com/volcengine/vkectl/pkg/model/app/kitex_gen/paastob/vke/appserver"
)

// AppapiserviceRawApi is a base client
type AppapiserviceRawApi struct {
	Client *client.Client
}

// NewAPIClient 生成一个客户端
func NewAPIClient(ak, sk, host, service, region string) *AppapiserviceRawApi {
	c := client.NewBaseClient()
	c.ServiceInfo = client.NewServiceInfo()
	c.ServiceInfo.Host = host
	c.ServiceInfo.Credentials.AccessKeyID = ak
	c.ServiceInfo.Credentials.SecretAccessKey = sk
	c.ServiceInfo.Credentials.Service = service
	c.ServiceInfo.Credentials.Region = region
	c.SdkVersion = client.DefaultSdkVersion

	return &AppapiserviceRawApi{Client: c}
}
func (p *AppapiserviceRawApi) GetConfigMap(body string, query url.Values) (r *appserver.GetConfigMapResponse, statusCode int, err error) {
	action := "GetConfigMap"
	r = &appserver.GetConfigMapResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) ListConfigMaps(body string, query url.Values) (r *appserver.ListConfigMapsResponse, statusCode int, err error) {
	action := "ListConfigMaps"
	r = &appserver.ListConfigMapsResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) CreateConfigMap(body string, query url.Values) (r *appserver.GetConfigMapResponse, statusCode int, err error) {
	action := "CreateConfigMap"
	r = &appserver.GetConfigMapResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) UpdateConfigMap(body string, query url.Values) (r *appserver.GetConfigMapResponse, statusCode int, err error) {
	action := "UpdateConfigMap"
	r = &appserver.GetConfigMapResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) DeleteConfigMap(body string, query url.Values) (r *appserver.EmptyResponse, statusCode int, err error) {
	action := "DeleteConfigMap"
	r = &appserver.EmptyResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) GetService(body string, query url.Values) (r *appserver.GetServiceResponse, statusCode int, err error) {
	action := "GetService"
	r = &appserver.GetServiceResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) ListServices(body string, query url.Values) (r *appserver.ListServicesResponse, statusCode int, err error) {
	action := "ListServices"
	r = &appserver.ListServicesResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) CreateService(body string, query url.Values) (r *appserver.GetServiceResponse, statusCode int, err error) {
	action := "CreateService"
	r = &appserver.GetServiceResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) UpdateService(body string, query url.Values) (r *appserver.GetServiceResponse, statusCode int, err error) {
	action := "UpdateService"
	r = &appserver.GetServiceResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) DeleteService(body string, query url.Values) (r *appserver.EmptyResponse, statusCode int, err error) {
	action := "DeleteService"
	r = &appserver.EmptyResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) PatchWorkloadImage(body string, query url.Values) (r *appserver.EmptyResponse, statusCode int, err error) {
	action := "PatchWorkloadImage"
	r = &appserver.EmptyResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) PatchWorkloadReplicas(body string, query url.Values) (r *appserver.EmptyResponse, statusCode int, err error) {
	action := "PatchWorkloadReplicas"
	r = &appserver.EmptyResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) RestartWorkload(body string, query url.Values) (r *appserver.EmptyResponse, statusCode int, err error) {
	action := "RestartWorkload"
	r = &appserver.EmptyResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) RollbackWorkload(body string, query url.Values) (r *appserver.EmptyResponse, statusCode int, err error) {
	action := "RollbackWorkload"
	r = &appserver.EmptyResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) ListWorkloadPods(body string, query url.Values) (r *appserver.ListWorkloadPodsResponse, statusCode int, err error) {
	action := "ListWorkloadPods"
	r = &appserver.ListWorkloadPodsResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) ListWorkloadServices(body string, query url.Values) (r *appserver.ListServicesResponse, statusCode int, err error) {
	action := "ListWorkloadServices"
	r = &appserver.ListServicesResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) ListWorkloadHistories(body string, query url.Values) (r *appserver.ListWorkloadHistoriesResponse, statusCode int, err error) {
	action := "ListWorkloadHistories"
	r = &appserver.ListWorkloadHistoriesResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) ListObjectEvents(body string, query url.Values) (r *appserver.ListWorkloadEventsResponse, statusCode int, err error) {
	action := "ListObjectEvents"
	r = &appserver.ListWorkloadEventsResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) PatchWorkloadParallelism(body string, query url.Values) (r *appserver.EmptyResponse, statusCode int, err error) {
	action := "PatchWorkloadParallelism"
	r = &appserver.EmptyResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) GetContainerLogs(body string, query url.Values) (r *appserver.GetContainerLogsResponse, statusCode int, err error) {
	action := "GetContainerLogs"
	r = &appserver.GetContainerLogsResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) ListWorkloadHorizontalPodAutoscalers(body string, query url.Values) (r *appserver.ListWorkloadHorizontalPodAutoscalersResponse, statusCode int, err error) {
	action := "ListWorkloadHorizontalPodAutoscalers"
	r = &appserver.ListWorkloadHorizontalPodAutoscalersResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) GetDeployment(body string, query url.Values) (r *appserver.GetDeploymentResponse, statusCode int, err error) {
	action := "GetDeployment"
	r = &appserver.GetDeploymentResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) ListDeployments(body string, query url.Values) (r *appserver.ListDeploymentsResponse, statusCode int, err error) {
	action := "ListDeployments"
	r = &appserver.ListDeploymentsResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) CreateDeployment(body string, query url.Values) (r *appserver.GetDeploymentResponse, statusCode int, err error) {
	action := "CreateDeployment"
	r = &appserver.GetDeploymentResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) UpdateDeployment(body string, query url.Values) (r *appserver.GetDeploymentResponse, statusCode int, err error) {
	action := "UpdateDeployment"
	r = &appserver.GetDeploymentResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) DeleteDeployment(body string, query url.Values) (r *appserver.EmptyResponse, statusCode int, err error) {
	action := "DeleteDeployment"
	r = &appserver.EmptyResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) GetStatefulSet(body string, query url.Values) (r *appserver.GetStatefulSetResponse, statusCode int, err error) {
	action := "GetStatefulSet"
	r = &appserver.GetStatefulSetResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) ListStatefulSets(body string, query url.Values) (r *appserver.ListStatefulSetsResponse, statusCode int, err error) {
	action := "ListStatefulSets"
	r = &appserver.ListStatefulSetsResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) CreateStatefulSet(body string, query url.Values) (r *appserver.GetStatefulSetResponse, statusCode int, err error) {
	action := "CreateStatefulSet"
	r = &appserver.GetStatefulSetResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) UpdateStatefulSet(body string, query url.Values) (r *appserver.GetStatefulSetResponse, statusCode int, err error) {
	action := "UpdateStatefulSet"
	r = &appserver.GetStatefulSetResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) DeleteStatefulSet(body string, query url.Values) (r *appserver.EmptyResponse, statusCode int, err error) {
	action := "DeleteStatefulSet"
	r = &appserver.EmptyResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) GetJob(body string, query url.Values) (r *appserver.GetJobResponse, statusCode int, err error) {
	action := "GetJob"
	r = &appserver.GetJobResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) ListJobs(body string, query url.Values) (r *appserver.ListJobsResponse, statusCode int, err error) {
	action := "ListJobs"
	r = &appserver.ListJobsResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) CreateJob(body string, query url.Values) (r *appserver.GetJobResponse, statusCode int, err error) {
	action := "CreateJob"
	r = &appserver.GetJobResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) UpdateJob(body string, query url.Values) (r *appserver.GetJobResponse, statusCode int, err error) {
	action := "UpdateJob"
	r = &appserver.GetJobResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) DeleteJob(body string, query url.Values) (r *appserver.EmptyResponse, statusCode int, err error) {
	action := "DeleteJob"
	r = &appserver.EmptyResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) GetSecret(body string, query url.Values) (r *appserver.GetSecretResponse, statusCode int, err error) {
	action := "GetSecret"
	r = &appserver.GetSecretResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) ListSecrets(body string, query url.Values) (r *appserver.ListSecretsResponse, statusCode int, err error) {
	action := "ListSecrets"
	r = &appserver.ListSecretsResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) CreateSecret(body string, query url.Values) (r *appserver.GetSecretResponse, statusCode int, err error) {
	action := "CreateSecret"
	r = &appserver.GetSecretResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) UpdateSecret(body string, query url.Values) (r *appserver.GetSecretResponse, statusCode int, err error) {
	action := "UpdateSecret"
	r = &appserver.GetSecretResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) DeleteSecret(body string, query url.Values) (r *appserver.EmptyResponse, statusCode int, err error) {
	action := "DeleteSecret"
	r = &appserver.EmptyResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) GetDaemonSet(body string, query url.Values) (r *appserver.GetDaemonSetResponse, statusCode int, err error) {
	action := "GetDaemonSet"
	r = &appserver.GetDaemonSetResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) ListDaemonSets(body string, query url.Values) (r *appserver.ListDaemonSetsResponse, statusCode int, err error) {
	action := "ListDaemonSets"
	r = &appserver.ListDaemonSetsResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) CreateDaemonSet(body string, query url.Values) (r *appserver.GetDaemonSetResponse, statusCode int, err error) {
	action := "CreateDaemonSet"
	r = &appserver.GetDaemonSetResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) UpdateDaemonSet(body string, query url.Values) (r *appserver.GetDaemonSetResponse, statusCode int, err error) {
	action := "UpdateDaemonSet"
	r = &appserver.GetDaemonSetResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) DeleteDaemonSet(body string, query url.Values) (r *appserver.EmptyResponse, statusCode int, err error) {
	action := "DeleteDaemonSet"
	r = &appserver.EmptyResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) GetCronJob(body string, query url.Values) (r *appserver.GetCronJobResponse, statusCode int, err error) {
	action := "GetCronJob"
	r = &appserver.GetCronJobResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) ListCronJobs(body string, query url.Values) (r *appserver.ListCronJobsResponse, statusCode int, err error) {
	action := "ListCronJobs"
	r = &appserver.ListCronJobsResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) ListJobsFromCronJob(body string, query url.Values) (r *appserver.ListJobsFromCronJobResponse, statusCode int, err error) {
	action := "ListJobsFromCronJob"
	r = &appserver.ListJobsFromCronJobResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) CreateCronJob(body string, query url.Values) (r *appserver.GetCronJobResponse, statusCode int, err error) {
	action := "CreateCronJob"
	r = &appserver.GetCronJobResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) UpdateCronJob(body string, query url.Values) (r *appserver.GetCronJobResponse, statusCode int, err error) {
	action := "UpdateCronJob"
	r = &appserver.GetCronJobResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) DeleteCronJob(body string, query url.Values) (r *appserver.EmptyResponse, statusCode int, err error) {
	action := "DeleteCronJob"
	r = &appserver.EmptyResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) PatchCronJobSuspend(body string, query url.Values) (r *appserver.EmptyResponse, statusCode int, err error) {
	action := "PatchCronJobSuspend"
	r = &appserver.EmptyResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) GetResourceYaml(body string, query url.Values) (r *appserver.GetResourceYamlResponse, statusCode int, err error) {
	action := "GetResourceYaml"
	r = &appserver.GetResourceYamlResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) CreateResourceByYaml(body string, query url.Values) (r *appserver.EmptyResponse, statusCode int, err error) {
	action := "CreateResourceByYaml"
	r = &appserver.EmptyResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) UpdateResourceByYaml(body string, query url.Values) (r *appserver.EmptyResponse, statusCode int, err error) {
	action := "UpdateResourceByYaml"
	r = &appserver.EmptyResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) GetHorizontalPodAutoscaler(body string, query url.Values) (r *appserver.GetHorizontalPodAutoscalerResponse, statusCode int, err error) {
	action := "GetHorizontalPodAutoscaler"
	r = &appserver.GetHorizontalPodAutoscalerResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) ListHorizontalPodAutoscalers(body string, query url.Values) (r *appserver.ListHorizontalPodAutoscalersResponse, statusCode int, err error) {
	action := "ListHorizontalPodAutoscalers"
	r = &appserver.ListHorizontalPodAutoscalersResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) CreateHorizontalPodAutoscaler(body string, query url.Values) (r *appserver.GetHorizontalPodAutoscalerResponse, statusCode int, err error) {
	action := "CreateHorizontalPodAutoscaler"
	r = &appserver.GetHorizontalPodAutoscalerResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) UpdateHorizontalPodAutoscaler(body string, query url.Values) (r *appserver.GetHorizontalPodAutoscalerResponse, statusCode int, err error) {
	action := "UpdateHorizontalPodAutoscaler"
	r = &appserver.GetHorizontalPodAutoscalerResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) DeleteHorizontalPodAutoscaler(body string, query url.Values) (r *appserver.EmptyResponse, statusCode int, err error) {
	action := "DeleteHorizontalPodAutoscaler"
	r = &appserver.EmptyResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) GetPod(body string, query url.Values) (r *appserver.GetPodResponse, statusCode int, err error) {
	action := "GetPod"
	r = &appserver.GetPodResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) ListPods(body string, query url.Values) (r *appserver.ListPodsResponse, statusCode int, err error) {
	action := "ListPods"
	r = &appserver.ListPodsResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) DeletePod(body string, query url.Values) (r *appserver.EmptyResponse, statusCode int, err error) {
	action := "DeletePod"
	r = &appserver.EmptyResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) GetAPIGroup(body string, query url.Values) (r *appserver.GetAPIGroupResponse, statusCode int, err error) {
	action := "GetAPIGroup"
	r = &appserver.GetAPIGroupResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) ListAPIGroups(body string, query url.Values) (r *appserver.ListAPIGroupsResponse, statusCode int, err error) {
	action := "ListAPIGroups"
	r = &appserver.ListAPIGroupsResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) ListCRDs(body string, query url.Values) (r *appserver.ListCRDsResponse, statusCode int, err error) {
	action := "ListCRDs"
	r = &appserver.ListCRDsResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) DeleteCRD(body string, query url.Values) (r *appserver.EmptyResponse, statusCode int, err error) {
	action := "DeleteCRD"
	r = &appserver.EmptyResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) GetResource(body string, query url.Values) (r *appserver.ResourceItem, statusCode int, err error) {
	action := "GetResource"
	r = &appserver.ResourceItem{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) ListResources(body string, query url.Values) (r *appserver.ListResourcesResponse, statusCode int, err error) {
	action := "ListResources"
	r = &appserver.ListResourcesResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) DeleteResource(body string, query url.Values) (r *appserver.EmptyResponse, statusCode int, err error) {
	action := "DeleteResource"
	r = &appserver.EmptyResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) ListCharts(body string, query url.Values) (r *appserver.ListChartsResponse, statusCode int, err error) {
	action := "ListCharts"
	r = &appserver.ListChartsResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) ListChartCategories(body string, query url.Values) (r *appserver.ListChartCategoriesResponse, statusCode int, err error) {
	action := "ListChartCategories"
	r = &appserver.ListChartCategoriesResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) GetChart(body string, query url.Values) (r *appserver.GetChartResponse, statusCode int, err error) {
	action := "GetChart"
	r = &appserver.GetChartResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) ListReleases(body string, query url.Values) (r *appserver.ListReleasesResponse, statusCode int, err error) {
	action := "ListReleases"
	r = &appserver.ListReleasesResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) GetRelease(body string, query url.Values) (r *appserver.GetReleaseResponse, statusCode int, err error) {
	action := "GetRelease"
	r = &appserver.GetReleaseResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) CreateRelease(body string, query url.Values) (r *appserver.CreateReleaseResponse, statusCode int, err error) {
	action := "CreateRelease"
	r = &appserver.CreateReleaseResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) UpdateRelease(body string, query url.Values) (r *appserver.CreateReleaseResponse, statusCode int, err error) {
	action := "UpdateRelease"
	r = &appserver.CreateReleaseResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) RollbackRelease(body string, query url.Values) (r *appserver.CreateReleaseResponse, statusCode int, err error) {
	action := "RollbackRelease"
	r = &appserver.CreateReleaseResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) DeleteRelease(body string, query url.Values) (r *appserver.EmptyResponse, statusCode int, err error) {
	action := "DeleteRelease"
	r = &appserver.EmptyResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) ListCronHorizontalPodAutoscalers(body string, query url.Values) (r *appserver.ListCronHorizontalPodAutoscalersResponse, statusCode int, err error) {
	action := "ListCronHorizontalPodAutoscalers"
	r = &appserver.ListCronHorizontalPodAutoscalersResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) CreateCronHorizontalPodAutoscaler(body string, query url.Values) (r *appserver.GetCronHorizontalPodAutoscalerResponse, statusCode int, err error) {
	action := "CreateCronHorizontalPodAutoscaler"
	r = &appserver.GetCronHorizontalPodAutoscalerResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) UpdateCronHorizontalPodAutoscaler(body string, query url.Values) (r *appserver.GetCronHorizontalPodAutoscalerResponse, statusCode int, err error) {
	action := "UpdateCronHorizontalPodAutoscaler"
	r = &appserver.GetCronHorizontalPodAutoscalerResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) DeleteCronHorizontalPodAutoscaler(body string, query url.Values) (r *appserver.EmptyResponse, statusCode int, err error) {
	action := "DeleteCronHorizontalPodAutoscaler"
	r = &appserver.EmptyResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) ListLogCollectRules(body string, query url.Values) (r *appserver.ListLogCollectRulesResp, statusCode int, err error) {
	action := "ListLogCollectRules"
	r = &appserver.ListLogCollectRulesResp{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) CreateLogCollectRule(body string, query url.Values) (r *appserver.CreateLogCollectRuleResp, statusCode int, err error) {
	action := "CreateLogCollectRule"
	r = &appserver.CreateLogCollectRuleResp{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) UpdateLogCollectRule(body string, query url.Values) (r *appserver.EmptyResponse, statusCode int, err error) {
	action := "UpdateLogCollectRule"
	r = &appserver.EmptyResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) DeleteLogCollectRule(body string, query url.Values) (r *appserver.EmptyResponse, statusCode int, err error) {
	action := "DeleteLogCollectRule"
	r = &appserver.EmptyResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) GetLogCollectRule(body string, query url.Values) (r *appserver.LogCollectRule, statusCode int, err error) {
	action := "GetLogCollectRule"
	r = &appserver.LogCollectRule{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}

func (p *AppapiserviceRawApi) ForwardKubernetesApi(body string, query url.Values) (r *appserver.ForwardKubernetesApiResponse, statusCode int, err error) {
	action := "ForwardKubernetesApi"
	r = &appserver.ForwardKubernetesApiResponse{}
	statusCode, err = p.Client.CommonHandler(action, query, body, r)
	return r, statusCode, err
}
