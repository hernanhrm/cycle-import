package di

import (
	"cycle-imports/application/prices"
	"cycle-imports/domain/ports"
	"fmt"
)

func InitPriceUseCase() ports.Prices {
	fmt.Println("🔄 DI: Inicializando PriceUseCase...")
	
	// ¡AQUÍ ESTÁ EL OTRO LADO DEL PROBLEMA! Para crear Prices necesito CoursePriceUseCase,
	// pero para crear CoursePriceUseCase necesito Prices
	coursePriceService := InitCoursePriceUseCase() // Esta llamada completará el ciclo infinito!
	
	fmt.Println("🏗️  DI: Creando Prices con dependencia CoursePriceUseCase...")
	return prices.NewPrices(coursePriceService)
}
