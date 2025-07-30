package courseprice

import (
	"cycle-imports/domain/ports"
	"fmt"
	"runtime"
	"time"
)

type CoursePrice struct{
	PriceService ports.Prices
}

func NewCoursePrice(priceService ports.Prices) *CoursePrice {
	fmt.Println("🏗️  Creando CoursePrice, necesito PriceService...")
	
	// Simular trabajo de inicialización que consume recursos
	for i := 0; i < 1000; i++ {
		_ = make([]byte, 1024*1024) // 1MB por iteración
	}
	
	cp := &CoursePrice{
		PriceService: priceService,
	}
	
	fmt.Println("✅ CoursePrice creado exitosamente")
	return cp
}

func (cp *CoursePrice) ConsumeResources() {
	fmt.Println("🔥 CoursePrice iniciando consumo de recursos...")
	
	// Crear muchas goroutines para consumir CPU
	for i := 0; i < runtime.NumCPU()*10; i++ {
		go func(id int) {
			for {
				// Bucle infinito que consume CPU y memoria
				data := make([][]byte, 100)
				for j := range data {
					data[j] = make([]byte, 1024*1024) // 1MB cada uno
				}
				time.Sleep(1 * time.Nanosecond)
			}
		}(i)
	}
	
	// Iniciar recursión infinita
	go cp.RecursiveCall(0)
}

func (cp *CoursePrice) RecursiveCall(depth int) {
	// Consume memoria exponencialmente con cada llamada
	memorySize := 1024 * 1024 * (depth%10 + 1) // Entre 1MB y 10MB
	data := make([]byte, memorySize)
	_ = data
	
	fmt.Printf("🔄 CoursePrice recursión nivel %d (usando %dMB)\n", depth, memorySize/1024/1024)
	
	// Cada cierto nivel, intentar usar el servicio de precios
	if depth%5 == 0 && cp.PriceService != nil {
		fmt.Printf("📞 CoursePrice llamando a PriceService en nivel %d\n", depth)
		// Esto podría causar más recursión si PriceService también llama a este servicio
		go func() {
			cp.PriceService.ConsumeResources()
		}()
	}
	
	// Pequeña pausa para no saturar completamente
	time.Sleep(10 * time.Millisecond)
	
	// Continuar recursión infinita
	cp.RecursiveCall(depth + 1)
}
