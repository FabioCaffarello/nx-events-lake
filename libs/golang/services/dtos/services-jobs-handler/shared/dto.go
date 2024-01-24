package shared

type UrlDomainDTO struct {
	Url  string `json:"url"`
	Name string `json:"domain"`
}

type ProxyLoaderDTO struct {
	Name     string `json:"name"`
	Priority int    `json:"priority"`
}

type CaptchaSolverDTO struct {
	Name     string `json:"name"`
	Priority int    `json:"priority"`
}
