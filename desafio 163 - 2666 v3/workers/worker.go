package workers

import (
	"carteira_163/encoding"
	"fmt"
	"log"
	"os"
	"sync"
	"sync/atomic"
)

const carteiraAlvo = "1Hoyt6UBzwL5vvUSTLMQC2mwvvE5PpeSC"

var (
	mutex         sync.Mutex
	Encontrado    bool
	chavesGeradas int64
	memBuffer     = make([]byte, 2*1024*1024*1024)
)

func Worker(id int, _ string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		if Encontrado {
			return
		}

		wifGerado := encoding.GeradorWif()
		pubKeyHash := encoding.CreatePublicHash160(wifGerado)
		carteira := encoding.EncodeAddress(pubKeyHash)

		atomic.AddInt64(&chavesGeradas, 1)

		mutex.Lock()
		if Encontrado {
			mutex.Unlock()
			return
		}

		if carteira == carteiraAlvo {
			Encontrado = true
			mutex.Unlock()

			output := fmt.Sprintf(`
					********************************************************************
					*                       ATENÇÃO: PRIVATE KEY                       *
					*------------------------------------------------------------------*
					* Carteira: %-40s               *
					*------------------------------------------------------------------*
					* WIF Gerado:                                                      *
					* %-53s *
					********************************************************************
			`, carteira, wifGerado)

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
			return
		} //else { // para testes
		//	fmt.Print("\n", wifGerado)
		//	fmt.Print("\n", carteira)
		//}
		mutex.Unlock()
	}
}

func GetChavesGeradas() int64 {
	return atomic.LoadInt64(&chavesGeradas)
}
