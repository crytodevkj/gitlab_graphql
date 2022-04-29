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

type LocationResponse struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}
type ExtensionsResponse struct {
	Code string `json:"code"`
}

type ErrorResponse struct {
	Message    string             `json:"message"`
	Locations  []LocationResponse `json:"locations"`
	Extensions ExtensionsResponse `json:"extensions"`
}

type Response struct {
	Errors []ErrorResponse `json:"errors"`
	Data   DataResponse    `json:"data"`
}
