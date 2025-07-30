package prices

import (
	"cycle-imports/domain/ports"
	"fmt"
	"runtime"
	"time"
)

type Prices struct{
	CoursePriceService ports.CoursePriceUseCase
}

func NewPrices(coursePriceService ports.CoursePriceUseCase) *Prices {
	fmt.Println("🏗️  Creando Prices, necesito CoursePriceService...")
	
	// Simular trabajo de inicialización pesado
	for i := 0; i < 1500; i++ {
		_ = make([]byte, 1024*1024*2) // 2MB por iteración
	}
	
	p := &Prices{
		CoursePriceService: coursePriceService,
	}
	
	fmt.Println("✅ Prices creado exitosamente")
	return p
}

func (p *Prices) ConsumeResources() {
	fmt.Println("🔥 Prices iniciando consumo masivo de recursos...")
	
	// Crear aún más goroutines para maximizar consumo
	for i := 0; i < runtime.NumCPU()*15; i++ {
		go func(id int) {
			for {
				// Consumo más agresivo de memoria
				bigSlice := make([][]byte, 200)
				for j := range bigSlice {
					bigSlice[j] = make([]byte, 1024*1024*2) // 2MB cada uno
				}
				time.Sleep(1 * time.Nanosecond)
			}
		}(i)
	}
	
	// Iniciar múltiples recursiones paralelas
	for i := 0; i < 3; i++ {
		go p.RecursiveCall(i * 1000) // Iniciar en diferentes niveles
	}
}

func (p *Prices) RecursiveCall(depth int) {
	// Consumo exponencial de memoria más agresivo
	memorySize := 1024 * 1024 * (depth%20 + 5) // Entre 5MB y 25MB por llamada
	data := make([]byte, memorySize)
	_ = data
	
	fmt.Printf("💥 Prices recursión nivel %d (usando %dMB)\n", depth, memorySize/1024/1024)
	
	// Llamar al CoursePriceService más frecuentemente
	if depth%3 == 0 && p.CoursePriceService != nil {
		fmt.Printf("📞 Prices llamando a CoursePriceService en nivel %d\n", depth)
		// Esto creará el ciclo: Prices -> CoursePrice -> Prices -> ...
		go func() {
			p.CoursePriceService.ConsumeResources()
		}()
	}
	
	// Pausa más corta para acelerar el consumo
	time.Sleep(5 * time.Millisecond)
	
	// Continuar recursión infinita
	p.RecursiveCall(depth + 1)
}
