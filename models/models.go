package models

type TimestampVersion struct {
	Version string `json:"version"`
}

type InRequest struct {
	Source  Source  `json:"source"`
	Version TimestampVersion `json:"version"`
}

type InResponse struct {
	Version  TimestampVersion  `json:"version"`
}

type OutParams struct {
	File string `json:"file"`
}

type OutRequest struct {
	Source Source `json:"source"`
	Params  OutParams `json:"params"`
}

type OutResponse struct {
	Version  TimestampVersion  `json:"version"`
}

type CheckRequest struct {
	Source  Source  `json:"source"`
	Version  TimestampVersion  `json:"version"`
}

type CheckResponse []TimestampVersion

type Source struct {}

