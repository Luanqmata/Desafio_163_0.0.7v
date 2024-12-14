package workers

import (
	"carteira_163/encoding"
	"fmt"
	"log"
	"os"
	"sync"
)

const carteiraAlvo = "1Hoyt6UBzwL5vvUSTLMQC2mwvvE5PpeSC"

var (
	mutex      sync.Mutex
	Encontrado bool
)

func Worker(id int, wif string, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		if Encontrado {
			return
		}

		wifGerado := encoding.GeradorWif(wif)
		pubKeyHash := encoding.CreatePublicHash160(wifGerado)
		carteira := encoding.EncodeAddress(pubKeyHash)

		mutex.Lock()
		if Encontrado {
			mutex.Unlock()
			return
		}

		if carteira == carteiraAlvo {
			Encontrado = true
			mutex.Unlock()

			output := fmt.Sprintf("\n\t\t|--------------%s----------------|\n", carteira) +
				"\t\t|----------------------ATENÇÃO-PRIVATE-KEY-----------------------|\n" +
				fmt.Sprintf("\t\t|%s|\n", wifGerado)

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
