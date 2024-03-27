package main

import "github.com/samber/do"

func main() {
	injector := do.New()

	// provides CarService
	do.Provide(injector, NewCarService)

	// provides EngineService
	do.Provide(injector, NewEngineService)

	car := do.MustInvoke[*CarService](injector)
	car.Start()
	// prints "car starting"

	do.HealthCheck[EngineService](injector)
	// returns "engine broken"

	// injector.ShutdownOnSIGTERM()    // will block until receiving sigterm signal
	injector.Shutdown()
	// prints "car stopped"
}
