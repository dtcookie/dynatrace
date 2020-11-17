package maintenancewindows

// MonitoredEntityFilterType The type of the Dynatrace entities (for example, hosts or services) you want to pick up by matching.
type MonitoredEntityFilterType string

// MonitoredEntityFilterTypes offers the known enum values
var MonitoredEntityFilterTypes = struct {
	ApmSecurityGateway           MonitoredEntityFilterType
	Application                  MonitoredEntityFilterType
	ApplicationMethod            MonitoredEntityFilterType
	ApplicationMethodGroup       MonitoredEntityFilterType
	AppMonServer                 MonitoredEntityFilterType
	AppMonSystemProfile          MonitoredEntityFilterType
	AutoScalingGroup             MonitoredEntityFilterType
	AuxiliarySyntheticTest       MonitoredEntityFilterType
	AWSApplicationLoadBalancer   MonitoredEntityFilterType
	AWSAvailabilityZone          MonitoredEntityFilterType
	AWSCredentials               MonitoredEntityFilterType
	AWSLambdaFunction            MonitoredEntityFilterType
	AWSNetworkLoadBalancer       MonitoredEntityFilterType
	AzureAPIManagementService    MonitoredEntityFilterType
	AzureApplicationGateway      MonitoredEntityFilterType
	AzureCosmosDB                MonitoredEntityFilterType
	AzureCredentials             MonitoredEntityFilterType
	AzureEventHub                MonitoredEntityFilterType
	AzureEventHubNamespace       MonitoredEntityFilterType
	AzureFunctionApp             MonitoredEntityFilterType
	AzureIotHub                  MonitoredEntityFilterType
	AzureLoadBalancer            MonitoredEntityFilterType
	AzureMgmtGroup               MonitoredEntityFilterType
	AzureRedisCache              MonitoredEntityFilterType
	AzureRegion                  MonitoredEntityFilterType
	AzureServiceBusNamespace     MonitoredEntityFilterType
	AzureServiceBusQueue         MonitoredEntityFilterType
	AzureServiceBusTopic         MonitoredEntityFilterType
	AzureSQLDatabase             MonitoredEntityFilterType
	AzureSQLElasticPool          MonitoredEntityFilterType
	AzureSQLServer               MonitoredEntityFilterType
	AzureStorageAccount          MonitoredEntityFilterType
	AzureSubscription            MonitoredEntityFilterType
	AzureTenant                  MonitoredEntityFilterType
	AzureVM                      MonitoredEntityFilterType
	AzureVMScaleSet              MonitoredEntityFilterType
	AzureWebApp                  MonitoredEntityFilterType
	CfApplication                MonitoredEntityFilterType
	CfFoundation                 MonitoredEntityFilterType
	CinderVolume                 MonitoredEntityFilterType
	CloudApplication             MonitoredEntityFilterType
	CloudApplicationInstance     MonitoredEntityFilterType
	CloudApplicationNamespace    MonitoredEntityFilterType
	ContainerGroup               MonitoredEntityFilterType
	ContainerGroupInstance       MonitoredEntityFilterType
	CustomApplication            MonitoredEntityFilterType
	CustomDevice                 MonitoredEntityFilterType
	CustomDeviceGroup            MonitoredEntityFilterType
	DCRumApplication             MonitoredEntityFilterType
	DCRumService                 MonitoredEntityFilterType
	DCRumServiceInstance         MonitoredEntityFilterType
	DeviceApplicationMethod      MonitoredEntityFilterType
	Disk                         MonitoredEntityFilterType
	DockerContainerGroup         MonitoredEntityFilterType
	DockerContainerGroupInstance MonitoredEntityFilterType
	DynamoDBTable                MonitoredEntityFilterType
	EbsVolume                    MonitoredEntityFilterType
	EC2Instance                  MonitoredEntityFilterType
	ElasticLoadBalancer          MonitoredEntityFilterType
	Environment                  MonitoredEntityFilterType
	ExternalSyntheticTestStep    MonitoredEntityFilterType
	GcpZone                      MonitoredEntityFilterType
	Geolocation                  MonitoredEntityFilterType
	GeolocSite                   MonitoredEntityFilterType
	GoogleComputeEngine          MonitoredEntityFilterType
	Host                         MonitoredEntityFilterType
	HostGroup                    MonitoredEntityFilterType
	HTTPCheck                    MonitoredEntityFilterType
	HTTPCheckStep                MonitoredEntityFilterType
	Hypervisor                   MonitoredEntityFilterType
	KubernetesCluster            MonitoredEntityFilterType
	KubernetesNode               MonitoredEntityFilterType
	MobileApplication            MonitoredEntityFilterType
	NetworkInterface             MonitoredEntityFilterType
	NeutronSubnet                MonitoredEntityFilterType
	OpenStackProject             MonitoredEntityFilterType
	OpenStackRegion              MonitoredEntityFilterType
	OpenStackVM                  MonitoredEntityFilterType
	OS                           MonitoredEntityFilterType
	ProcessGroup                 MonitoredEntityFilterType
	ProcessGroupInstance         MonitoredEntityFilterType
	RelationalDatabaseService    MonitoredEntityFilterType
	Service                      MonitoredEntityFilterType
	ServiceInstance              MonitoredEntityFilterType
	ServiceMethod                MonitoredEntityFilterType
	ServiceMethodGroup           MonitoredEntityFilterType
	SwiftContainer               MonitoredEntityFilterType
	SyntheticLocation            MonitoredEntityFilterType
	SyntheticTest                MonitoredEntityFilterType
	SyntheticTestStep            MonitoredEntityFilterType
	Virtualmachine               MonitoredEntityFilterType
	VmwareDatacenter             MonitoredEntityFilterType
}{
	"APM_SECURITY_GATEWAY",
	"APPLICATION",
	"APPLICATION_METHOD",
	"APPLICATION_METHOD_GROUP",
	"APPMON_SERVER",
	"APPMON_SYSTEM_PROFILE",
	"AUTO_SCALING_GROUP",
	"AUXILIARY_SYNTHETIC_TEST",
	"AWS_APPLICATION_LOAD_BALANCER",
	"AWS_AVAILABILITY_ZONE",
	"AWS_CREDENTIALS",
	"AWS_LAMBDA_FUNCTION",
	"AWS_NETWORK_LOAD_BALANCER",
	"AZURE_API_MANAGEMENT_SERVICE",
	"AZURE_APPLICATION_GATEWAY",
	"AZURE_COSMOS_DB",
	"AZURE_CREDENTIALS",
	"AZURE_EVENT_HUB",
	"AZURE_EVENT_HUB_NAMESPACE",
	"AZURE_FUNCTION_APP",
	"AZURE_IOT_HUB",
	"AZURE_LOAD_BALANCER",
	"AZURE_MGMT_GROUP",
	"AZURE_REDIS_CACHE",
	"AZURE_REGION",
	"AZURE_SERVICE_BUS_NAMESPACE",
	"AZURE_SERVICE_BUS_QUEUE",
	"AZURE_SERVICE_BUS_TOPIC",
	"AZURE_SQL_DATABASE",
	"AZURE_SQL_ELASTIC_POOL",
	"AZURE_SQL_SERVER",
	"AZURE_STORAGE_ACCOUNT",
	"AZURE_SUBSCRIPTION",
	"AZURE_TENANT",
	"AZURE_VM",
	"AZURE_VM_SCALE_SET",
	"AZURE_WEB_APP",
	"CF_APPLICATION",
	"CF_FOUNDATION",
	"CINDER_VOLUME",
	"CLOUD_APPLICATION",
	"CLOUD_APPLICATION_INSTANCE",
	"CLOUD_APPLICATION_NAMESPACE",
	"CONTAINER_GROUP",
	"CONTAINER_GROUP_INSTANCE",
	"CUSTOM_APPLICATION",
	"CUSTOM_DEVICE",
	"CUSTOM_DEVICE_GROUP",
	"DCRUM_APPLICATION",
	"DCRUM_SERVICE",
	"DCRUM_SERVICE_INSTANCE",
	"DEVICE_APPLICATION_METHOD",
	"DISK",
	"DOCKER_CONTAINER_GROUP",
	"DOCKER_CONTAINER_GROUP_INSTANCE",
	"DYNAMO_DB_TABLE",
	"EBS_VOLUME",
	"EC2_INSTANCE",
	"ELASTIC_LOAD_BALANCER",
	"ENVIRONMENT",
	"EXTERNAL_SYNTHETIC_TEST_STEP",
	"GCP_ZONE",
	"GEOLOCATION",
	"GEOLOC_SITE",
	"GOOGLE_COMPUTE_ENGINE",
	"HOST",
	"HOST_GROUP",
	"HTTP_CHECK",
	"HTTP_CHECK_STEP",
	"HYPERVISOR",
	"KUBERNETES_CLUSTER",
	"KUBERNETES_NODE",
	"MOBILE_APPLICATION",
	"NETWORK_INTERFACE",
	"NEUTRON_SUBNET",
	"OPENSTACK_PROJECT",
	"OPENSTACK_REGION",
	"OPENSTACK_VM",
	"OS",
	"PROCESS_GROUP",
	"PROCESS_GROUP_INSTANCE",
	"RELATIONAL_DATABASE_SERVICE",
	"SERVICE",
	"SERVICE_INSTANCE",
	"SERVICE_METHOD",
	"SERVICE_METHOD_GROUP",
	"SWIFT_CONTAINER",
	"SYNTHETIC_LOCATION",
	"SYNTHETIC_TEST",
	"SYNTHETIC_TEST_STEP",
	"VIRTUALMACHINE",
	"VMWARE_DATACENTER",
}