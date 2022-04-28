package main

type NodeResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ForksCount  int    `json:"forksCount"`
}

type NodesResponse struct {
	Nodes []NodeResponse `json:"nodes"`
}

type DataResponse struct {
	Projects NodesResponse `json:"projects"`
}

type Response struct {
	Data DataResponse `json:"data"`
}
