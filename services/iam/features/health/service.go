package health

import "github.com/raufhm/golang-project-convention/modules/logger"

type Service struct {
	// Add dependencies like database, cache, etc. if needed
	logger *logger.Logger
}

func NewService(logger *logger.Logger) *Service {
	return &Service{logger: logger}
}

func (s *Service) Status() map[string]string {
	return map[string]string{
		"status":  "OK",
		"version": "1.0.0",
	}
}
