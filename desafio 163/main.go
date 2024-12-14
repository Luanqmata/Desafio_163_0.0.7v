package main

import (
	"carteira_163/encoding"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strings"
	"sync"
)

const (
	caracters_btc = "0123456789abcdef"
	carteira_163  = "1Hoyt6UBzwL5vvUSTLMQC2mwvvE5PpeSC" // CARTEIRA 163
	wif           = "403d3x4xcxfx6x9xfx3xaxcx5x0x4xbxbx7x2x6x8x7x8xax4x0x8x3x3x3x7x3x"
)

var (
	memBuffer  = make([]byte, 4*1024*1024*1024)
	mutex      sync.Mutex
	encontrado bool
	wg         sync.WaitGroup
)

func random_random() string {
	randomIndex := rand.Intn(len(caracters_btc))
	randomChar := string(caracters_btc[randomIndex])
	return randomChar
}

func gerador_wif_163() string {
	var wifGerado string
	wifSplit := strings.Split(wif, "x")

	for i, part := range wifSplit {
		wifGerado += part
		if i < len(wifSplit)-1 {
			wifGerado += random_random()
		}
	}

	return wifGerado
}
func worker(id int) {
	defer wg.Done()

	for {
		if encontrado {
			return
		}

		wif_gerada := gerador_wif_163()
		pubKeyHash := encoding.CreatePublicHash160(wif_gerada)
		carteira := encoding.EncodeAddress(pubKeyHash)

		mutex.Lock()
		if encontrado {
			mutex.Unlock()
			return
		}
		if carteira == carteira_163 {
			encontrado = true
			mutex.Unlock()

			output := fmt.Sprintf("\n\t\t|--------------%s----------------|\n", carteira) +
				"\t\t|----------------------ATENÇÃO-PRIVATE-KEY-----------------------|\n" +
				fmt.Sprintf("\t\t|%s|\n", wif_gerada)

			fmt.Print(output)

			// Escreve no arquivo
			file, err := os.OpenFile("carteira_encontradas.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatalf("Erro ao abrir arquivo: %v", err)
			}
			defer file.Close()

			_, err = file.WriteString(output)
			if err != nil {
				log.Printf("Erro ao escrever no arquivo: %v", err)
			}
			return
		} //else {   	// para testes
		//	fmt.Print("\n", wif_gerada)
		//	fmt.Print("\n", carteira)
		//}
		mutex.Unlock()
	}
}

func main() {
	fmt.Print("Bem vindo ao desafio 163 -- By China/gpt 0.2v\n\n")

	fmt.Print("1 ~~ 4\n2 ~~ 8\n3 ~~ 12\n4 ~~ 16 \n5 ~~ 20 \nDigite sua escolha: ")
	var escolha_modo int
	fmt.Scanln(&escolha_modo)

	var numThreads int
	switch escolha_modo {
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

	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go worker(i)
	}
	wg.Wait()
}
