package interfaces

type IHealthycheckServices interface {
	HealthcheckServices() (string, error)
}

type IHealthycheckRepo interface{}
