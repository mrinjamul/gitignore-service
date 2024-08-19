package services

import "github.com/mrinjamul/gitignore-service/api/controllers"

type Services interface {
	TemplatesService() controllers.Templates
	APIService() controllers.API
	HealthCheckService() controllers.HealthCheck
}

type services struct {
	templates   controllers.Templates
	api         controllers.API
	healthCheck controllers.HealthCheck
}

func (svc *services) TemplatesService() controllers.Templates {
	return svc.templates
}

func (svc *services) APIService() controllers.API {
	return svc.api
}

func (svc *services) HealthCheckService() controllers.HealthCheck {
	return svc.healthCheck
}

// NewServices initializes services
func NewServices() Services {
	return &services{
		templates:   controllers.NewTemplates(),
		api:         controllers.NewAPI(),
		healthCheck: controllers.NewHealthCheck(),
	}
}
