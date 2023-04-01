package main

type Response struct {
	Entity        string      `json:"entity"`
	IptableOutput interface{} `json:"iptableOutput"`
}

type KubernetesDefaultResponse struct {
    PodName       string      `json:"podName"`
    IptableOutput interface{} `json:"iptableOutput"`
}

type KubeProxy struct {
	PodNames   []string       `json:"podNames"`
}