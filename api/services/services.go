package services

import "github.com/mrinjamul/gitignore-service/api/controllers"

type Services interface {
	TemplatesService() controllers.Templates
	GitignoreService() controllers.Gitignore
	HealthCheckService() controllers.HealthCheck
}

type services struct {
	templates   controllers.Templates
	gitignore   controllers.Gitignore
	healthCheck controllers.HealthCheck
}

func (svc *services) TemplatesService() controllers.Templates {
	return svc.templates
}

func (svc *services) GitignoreService() controllers.Gitignore {
	return svc.gitignore
}

func (svc *services) HealthCheckService() controllers.HealthCheck {
	return svc.healthCheck
}

// NewServices initializes services
func NewServices() Services {
	return &services{
		templates:   controllers.NewTemplates(),
		gitignore:   controllers.NewGitignore(),
		healthCheck: controllers.NewHealthCheck(),
	}
}
