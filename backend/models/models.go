package models

type Response struct {
	Entity        string      `json:"entity"`
	IptableOutput interface{} `json:"iptableOutput"`
}

type KubernetesDefaultResponse struct {
	CurrentPodName string      `json:"currentPodName"`
	PodList        []string    `json:"podNames"`
	IptableOutput  interface{} `json:"iptableOutput"`
}
