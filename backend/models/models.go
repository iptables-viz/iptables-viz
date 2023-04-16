package models

type Response struct {
	Entity        string      `json:"entity"`
	IptableOutput interface{} `json:"iptablesOutput"`
}

type KubernetesDefaultResponse struct {
	CurrentPodName string      `json:"podName"`
	PodList        []string    `json:"podNames"`
	IptableOutput  interface{} `json:"iptablesOutput"`
}

type LinuxIptableOutput struct {
	TableName     string      `json:"tableName"`
	IptableOutput interface{} `json:"iptablesOutput"`
}
