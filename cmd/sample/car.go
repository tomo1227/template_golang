package main

import (
	"github.com/samber/do"
)

func NewCarService(i *do.Injector) (*CarService, error) {
	engine := do.MustInvoke[EngineService](i)
	car := CarService{Engine: engine}
	return &car, nil
}

type CarService struct {
	Engine EngineService
}

func (c *CarService) Start() {
	println("car starting")
}

// [Optional] Implements do.Shutdownable.
func (c *CarService) Shutdown() error {
	println("car stopped")
	return nil
}
