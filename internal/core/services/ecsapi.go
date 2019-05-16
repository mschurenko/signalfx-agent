package services

// ECSAPI ...
type ECSAPI struct {
	EndpointCore `yaml:",inline"`
	IsFargate    bool `yaml:"is_fargate"`
}
