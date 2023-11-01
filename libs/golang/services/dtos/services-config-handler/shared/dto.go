package shared

type JobDependencies struct {
	Service string `json:"service"`
	Source  string `json:"source"`
}
