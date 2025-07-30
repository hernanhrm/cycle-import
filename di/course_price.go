package di

import (
	"cycle-imports/application/courseprice"
	"cycle-imports/domain/ports"
	"fmt"
)

func InitCoursePriceUseCase() ports.CoursePriceUseCase {
	fmt.Println("🔄 DI: Inicializando CoursePriceUseCase...")
	
	// ¡AQUÍ ESTÁ EL PROBLEMA! Para crear CoursePrice necesito Prices,
	// pero para crear Prices necesito CoursePriceUseCase
	pricesService := InitPriceUseCase() // Esta llamada creará un ciclo infinito!
	
	fmt.Println("🏗️  DI: Creando CoursePrice con dependencia Prices...")
	return courseprice.NewCoursePrice(pricesService)
}
