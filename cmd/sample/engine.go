package main

import (
	"fmt"

	"github.com/samber/do"
)

// エンジン
type EngineService interface{}

func NewEngineService(i *do.Injector) (EngineService, error) {
	return &engineServiceImplem{}, nil
}

// repository impl
type engineServiceImplem struct{}

// [Optional] Implements do.Healthcheckable.
func (c *engineServiceImplem) HealthCheck() error {
	return fmt.Errorf("engine broken")
}
