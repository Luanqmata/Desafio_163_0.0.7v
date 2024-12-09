package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
	"carteira_163/encoding"
)

const (
	caracters_btc = "0123456789abcdef"
	endereco_163  = "1Hoyt6UBzwL5vvUSTLMQC2mwvvE5PpeSC" // CARTEIRA 163
)

var (
	_wifindices_carteira = []string{"4", "3", "3", "4", "c", "f", "6", "9", "f", "3", "a", "c", "5", "0", "4", "b", "b", "7", "2", "6", "8", "7", "8", "a", "4", "0", "8", "3", "3", "3", "7", "3"}
	memBuffer            = make([]byte, 4*1024*1024*1024)
	ultima_Wif_gerada    string
	mutex                sync.Mutex
)

func random_random() string {
	randomIndex := rand.Intn(len(caracters_btc))
	randomChar := string(caracters_btc[randomIndex])
	return randomChar
}

func gerador_chave_163() {
	var chaveBuilder string
	for i := 0; i < 32 && i < len(_wifindices_carteira); i++ {
		chaveBuilder += _wifindices_carteira[i]
		randomChar := random_random()
		chaveBuilder += randomChar
	}

	mutex.Lock()
	ultima_Wif_gerada = chaveBuilder
	mutex.Unlock()
}

func worker() {
	for {
		mutex.Lock()
		chave := ultima_Wif_gerada
		mutex.Unlock()

		hash_chave_pub := encoding.CreatePublicHash160(chave)
		endereco := encoding.EncodeAddress(hash_chave_pub)

		if endereco == endereco_163 {
			fmt.Println("Endereço correto encontrado:", endereco)
			break
		} else {
			fmt.Println("Carteira:",endereco,"\n Wif: ",chave)

			mutex.Lock()
			ultima_Wif_gerada = ""
			mutex.Unlock()
			gerador_chave_163()
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	runtime.GOMAXPROCS(8)

	go worker()

	gerador_chave_163()

	select {}
}
