package main

type DistancePost struct {
	Source string `json:"source"`
	Target string `json:"target"`
}

type DistanceResponse struct {
	Success   bool
	Message   string
	Distances []Distance `json:",omitempty"`
}

type Distance struct {
	Source   string
	Target   string
	Distance int
}
