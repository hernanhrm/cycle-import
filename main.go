package main

import (
	"cycle-imports/di"
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("🚀 Iniciando programa con DEPENDENCIA CÍCLICA...")
	fmt.Printf("CPUs disponibles: %d\n", runtime.NumCPU())
	fmt.Println()

	fmt.Println("⚠️  ADVERTENCIA CRÍTICA:")
	fmt.Println("💀 Este programa tiene una dependencia cíclica en los constructores!")
	fmt.Println("🔥 Causará Stack Overflow y consumo infinito de memoria!")
	fmt.Println("🛑 Para detenerlo usa Ctrl+C inmediatamente!")
	fmt.Println()

	fmt.Println("🔄 Explicación del problema:")
	fmt.Println("   📦 InitCoursePriceUseCase() necesita InitPriceUseCase()")
	fmt.Println("   📦 InitPriceUseCase() necesita InitCoursePriceUseCase()")
	fmt.Println("   🔁 = Recursión infinita en la inicialización!")
	fmt.Println()

	// Countdown para que el usuario pueda cancelar
	for i := 5; i > 0; i-- {
		fmt.Printf("⏰ Iniciando en %d segundos... (Ctrl+C para cancelar)\n", i)
		time.Sleep(1 * time.Second)
	}

	fmt.Println()
	fmt.Println("💥 INICIANDO DEPENDENCIA CÍCLICA...")
	fmt.Println("🔄 Intentando crear CoursePriceUseCase...")

	// ¡AQUÍ OCURRE EL DESASTRE!
	// Esta línea causará recursión infinita:
	// main -> InitCoursePriceUseCase -> InitPriceUseCase -> InitCoursePriceUseCase -> InitPriceUseCase -> ...
	coursePriceService := di.InitCoursePriceUseCase()

	// Este código nunca se ejecutará debido al stack overflow anterior
	fmt.Println("✅ Si ves este mensaje, algo salió mal - el programa debería haber crasheado!")
	coursePriceService.ConsumeResources()

	// Código que nunca se alcanzará
	for {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		fmt.Printf("💾 Memoria: %d MB | Goroutines: %d\n",
			m.Alloc/1024/1024, runtime.NumGoroutine())
		time.Sleep(1 * time.Second)
	}
}

