package main

import (
	"carteira_163/mecanicas"
	"carteira_163/workers"
	"fmt"
	"runtime"
	"sync"
)

const wif = "403d3x4xcxfx6x9xfx3xaxcx5x0x4xbxbx7x2x6x8x7x8xax4x0x8x3x3x3x7x3x"

func main() {
	fmt.Print("Bem vindo ao desafio 163 -- By China/gpt 0.2v\n\n")

	fmt.Print("1 ~~ 4\n2 ~~ 8\n3 ~~ 12\n4 ~~ 16 \n5 ~~ 20 \nDigite sua escolha: ")
	var escolhaModo int
	fmt.Scanln(&escolhaModo)

	var numThreads int
	switch escolhaModo {
	case 1:
		numThreads = 4
	case 2:
		numThreads = 8
	case 3:
		numThreads = 12
	case 4:
		numThreads = 16
	case 5:
		numThreads = runtime.NumCPU()
	default:
		numThreads = 4
	}

	runtime.GOMAXPROCS(numThreads)

	mecanicas.CapturaSinal()

	var wg sync.WaitGroup
	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go workers.Worker(i, wif, &wg)
	}

	wg.Wait()
}
