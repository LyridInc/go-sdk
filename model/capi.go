package model

type (
	OpenStackInfrastructureClient struct {
		CloudsYaml           string `json:"cloudsYaml"`
		NetworkEndpoint      string `json:"networkEndpoint"`
		LoadBalancerEndpoint string `json:"loadBalancerEndpoint"`
		DnsNameServers       string `json:"dnsNameServers"`
		ExternalNetworkId    string `json:"externalNetworkId"`
		SshKeyName           string `json:"sshKeyName"`
		FailureDomain        string `json:"failureDomain"`
		IgnoreVolumeAZ       bool   `json:"ignoreVolumeAz"`
	}

	OnboardVegaRequest struct {
		UUID                   string `json:"uuid"`
		VegaTag                string `json:"vegaTag"`
		XVegaTag               string `json:"xVegaTag"`
		VegaId                 string `json:"vegaId"`
		JoinToken              string `json:"joinToken"`
		VegaHostName           string `json:"vegaHostname"`
		Prometheus             string `json:"prometheus"`
		VegaUseSecure          string `json:"vegaUseSecure"`
		Ingress                string `json:"ingress"`
		LyridNamespace         string `json:"lyridNamespace"`
		VegaHelmChartLocation  string `json:"vegaHelmChartLocation"`
		XVegaHelmChartLocation string `json:"xVegaHelmChartLocation"`
		UserToken              string `json:"userToken"`
		VegaName               string `json:"vegaName"`
		MongoDbUrl             string `json:"mongoDbUrl"`
		RedisEndpoint          string `json:"redisEndpoint"`
		RedisPassword          string `json:"redisPassword"`
		RedisDbNo              string `json:"redisDbNo"`
		MessagingEndpoint      string `json:"messagingEndpoint"`
		MessagingPassword      string `json:"messagingPassword"`
		MessagingDbNo          string `json:"messagingDbNo"`
		RegistryEndpoint       string `json:"registryEndpoint"`
		RegistryPort           string `json:"registryPort"`
		RegistryUsername       string `json:"registryUsername"`
		RegistryPassword       string `json:"registryPassword"`
		RegistrySecure         string `json:"registrySecure"`
		ClusterName            string `json:"clusterName"`
		ConfigAutoscalerUrl    string `json:"configAutoscalerUrl"`
		ConfigDomainUrl        string `json:"configDomainUrl"`
		LyridSaUrl             string `json:"lyridSaUrl"`
		LyridSaConfigUrl       string `json:"lyridSaConfigUrl"`
		IngressValuesUrl       string `json:"ingressValuesUrl"`
		LyraToken              string `json:"lyraToken"`

		VegaPort             uint   `json:"vegaPort"`
		VegaEngineId         string `json:"engineId"`
		VegaVendorShortName  string `json:"vegaVendorShortName"`
		VegaRegionId         string `json:"vegaRegionId"`
		UserClusterSettingId string `json:"userClusterSettingId"`
		AccountId            string `json:"accountId"`
	}

	ProvisionClusterRequest struct {
		Id                        string                 `json:"id"`
		InfrastructureProvider    string                 `json:"infrastructureProvider"`
		InfrastructureClient      map[string]interface{} `json:"infrastructureClient"` // cloudsYaml, flavor, etc
		KubernetesVersion         string                 `json:"kubernetesVersion"`
		ClusterName               string                 `json:"clusterName"`
		ControlPlaneMachineFlavor string                 `json:"controlPlaneMachineFlavor"`
		ControlPlaneMachineCount  float64                `json:"controlPlaneMachineCount"`
		WorkerMachineFlavor       string                 `json:"workerMachineFlavor"`
		WorkerMachineCount        float64                `json:"workerMachineCount"`
		ImageName                 string                 `json:"imageName"`
		NodeCidr                  string                 `json:"nodeCidr"`
	}
)
