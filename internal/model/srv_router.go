package model

type SrvRouterCfg struct {
	Version string                      `yaml:"version"`
	Route   map[string]*SrvRouterServer `yaml:"route"`
}

type SrvRouterServer struct {
	RegService string `yaml:"reg_service"`
	Service    string `yaml:"service"`
	Method     string `yaml:"method"`
}
