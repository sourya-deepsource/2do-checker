package main

type Namespace struct {
	Key   string  `json:"key"`
	Value float64 `json:"value"`
}

type Metric struct {
	MetricCode string      `json:"metric_code"`
	Namespaces []Namespace `json:"namespaces"`
}

type MacroResult struct {
	Issues    []Issue     `json:"issues"`
	Metrics   []Metric    `json:"metrics,omitempty"`
	IsPassed  bool        `json:"is_passed"`
	Errors    bool        `json:"errors"`
	ExtraData interface{} `json:"extra_data"`
}

type Coordinate struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}

type Position struct {
	Begin Coordinate `json:"begin"`
	End   Coordinate `json:"end"`
}

type Location struct {
	Path     string   `json:"path"`
	Position Position `json:"position"`
}

type Issue struct {
	Code     string   `json:"issue_code"`
	Title    string   `json:"issue_text"`
	Location Location `json:"location"`
}
