package main

import (
	"carteira_163/encoding"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"sync"
	"time"
)

const (
	caracters_btc = "0123456789abcdef"
	carteira_163  = "1Hoyt6UBzwL5vvUSTLMQC2mwvvE5PpeSC" // CARTEIRA 163
)

var (
	_wifindices_carteira = []string{"4", "3", "3", "4", "c", "f", "6", "9", "f", "3", "a", "c", "5", "0", "4", "b", "b", "7", "2", "6", "8", "7", "8", "a", "4", "0", "8", "3", "3", "3", "7", "3"}
	memBuffer            = make([]byte, 4*1024*1024*1024)
	ultima_Wif_gerada    string
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
	var chaveBuilder string
	for i := 0; i < 32 && i < len(_wifindices_carteira); i++ {
		chaveBuilder += _wifindices_carteira[i]
		randomChar := random_random()
		chaveBuilder += randomChar
	}
	return chaveBuilder
}

func worker(id int) {
	defer wg.Done()

	for {
		mutex.Lock()
		if encontrado {
			mutex.Unlock()
			return
		}
		mutex.Unlock()

		wif := gerador_wif_163()
		pubKeyHash := encoding.CreatePublicHash160(wif)
		carteira := encoding.EncodeAddress(pubKeyHash)

		mutex.Lock()
		ultima_Wif_gerada = wif
		if carteira == carteira_163 {
			output := fmt.Sprint("\n carteira: ", carteira, "\n wif: ", wif)
			fmt.Print(output)
			file, err := os.OpenFile("carteira_encontradas.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatalf("Erro ao abrir arquivo: %v", err)
			}
			defer file.Close()

			_, err = file.WriteString(output)
			if err != nil {
				log.Printf("Erro ao escrever no arquivo: %v", err)
			}
			encontrado = true
			mutex.Unlock() // Libera o mutex
			return         // Encerra a execução do worker
		}
		mutex.Unlock()

		if encontrado {
			return
		}
	}
}

func main() {
	//-----------------------------------VERIFICA N° THREADS---------------------------------------
	numCPUs := runtime.NumCPU()
	cpuInfo, _ := cpu.Info()
	cpuModelName := "Desconhecido"
	if len(cpuInfo) > 0 {
		cpuModelName = cpuInfo[0].ModelName
	}
	fmt.Printf("\n\t	Obs: O Seu Computador tem %d threads. (Processador: %s)\n", numCPUs, cpuModelName)
	//-----------------------------------VERIFICA N° THREADS---------------------------------------
	//------------------------------------MENSAGEM DE MODOS----------------------------------------

	style.Modos()
	var escolha_modo int
	fmt.Scanln(&escolha_modo)

	var numThreads int
	switch escolha_modo {
	case 1:
		numThreads = 3
	case 2:
		numThreads = 5
	case 3:
		numThreads = 11
	case 4:
		numThreads = 17
	case 5:
		numThreads = 24
	case 6:
		numThreads = runtime.NumCPU()
	default:
		fmt.Println("	Escolha inválida. Usando o SECURE MODE...  (20%) - CPU 58°C - 117K Chaves P/seg.")
		numThreads = 4
	}
	style.Mensagem_iniciando(numThreads)

	runtime.GOMAXPROCS(numThreads)
	startTime = time.Now()

	// Inicia goroutines
	for i := 0; i < numThreads; i++ {
		wg.Add(1)
		go worker(i)
	}
	//--------------------------------------EXIBIÇÃO EM TEMPO REAL--------------------------------------

	style.MonitorarChaves(&mu, &contador, &encontrado, startTime, &ultimaChaveGerada)

	//-------------------------------------------MENSAGEM FINAL----------------------------------
	rand.Seed(time.Now().UnixNano())
	runtime.GOMAXPROCS(1)

	// Inicia as goroutines
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go worker(i)
	}

	wg.Wait()
}
