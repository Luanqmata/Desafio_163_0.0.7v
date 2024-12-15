package main

import (
	"carteira_163/encoding"
	"carteira_163/mecanicas"
	"carteira_163/style"
	"carteira_163/workers"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	style.Bem_vindo()

	encoding.InitWif("403b3d4fcxfx6x9xfx3xaxcx5x0x4xbxbx7x2x6x8x7x8xax4x0x8x3x3x3x7x3x")

	style.Opcoes_uso_proc()

	var escolhaModo int
	fmt.Scanln(&escolhaModo)

	var numThreads int
	switch escolhaModo {
	case 1:
		numThreads = 5
	case 2:
		numThreads = 11
	case 3:
		numThreads = 17
	case 4:
		numThreads = runtime.NumCPU()
	default:
		numThreads = 4
	}

	runtime.GOMAXPROCS(numThreads)

	mecanicas.CapturaSinal()

	var wg sync.WaitGroup

	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()

		var ultimoTotal int64

		for range ticker.C {
			geradas := workers.GetChavesGeradas()
			kps := geradas - ultimoTotal
			ultimoTotal = geradas
			fmt.Printf("\r\tChaves geradas: %d | KPS: %d", geradas, kps)

			if workers.Encontrado {
				break
			}
		}
		fmt.Println("\nProcesso concluído!")
	}()

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go workers.Worker(i, "", &wg)
	}

	wg.Wait()
}
