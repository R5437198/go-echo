package di

import "go-echo/internal/presentation/handler/health_check"

func HealthCheckHandler() *health_check.Handler {
	return health_check.New()
}
