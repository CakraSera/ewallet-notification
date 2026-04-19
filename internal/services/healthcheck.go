package services

import "ewallet-notification/internal/interfaces"

type Healthcheck struct {
	HealthycheckRespository interfaces.IHealthycheckRepo
}

func (s *Healthcheck) HealthcheckServices() (string, error) {
	return "service is healthy", nil
}
