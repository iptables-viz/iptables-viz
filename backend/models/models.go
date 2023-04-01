package models

type Response struct {
	Entity        string         `json:"entity"`
	IptableOutput interface{}    `json:"iptableOutput"`
}

type KubernetesDefaultResponse struct {
    PodName       string         `json:"podName"`
	PodList		  []string       `json:"podNames"`
    IptableOutput interface{}    `json:"iptableOutput"`
}