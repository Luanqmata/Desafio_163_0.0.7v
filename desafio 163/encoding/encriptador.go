// criptografia bitcoinesca
package encoding

import (
	"carteira_163/crypto/base58"
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"strings"

	"github.com/btcsuite/btcd/btcec"
	"golang.org/x/crypto/ripemd160"
)

// Constantes usadas
const caracters_btc = "0123456789abcdef"

// Função para gerar caracteres aleatórios
func Random_random() string {
	randomIndex := rand.Intn(len(caracters_btc))
	randomChar := string(caracters_btc[randomIndex])
	return randomChar
}

// Função que gera WIF 163
func GeradorWif(wif string) string {
	var wifGerado string
	wifSplit := strings.Split(wif, "x")

	for i, part := range wifSplit {
		wifGerado += part
		if i < len(wifSplit)-1 {
			wifGerado += Random_random()
		}
	}

	return wifGerado
}

// Gera o hash160 da chave pública
func CreatePublicHash160(privKeyHex string) []byte {
	privKeyBytes, err := hex.DecodeString(privKeyHex)
	if err != nil {
		panic(err)
	}

	privKey, _ := btcec.PrivKeyFromBytes(btcec.S256(), privKeyBytes)
	compressedPubKey := privKey.PubKey().SerializeCompressed()

	pubKeyHash := Hash160(compressedPubKey)
	return pubKeyHash
}

// Hash SHA256 seguido de RIPEMD160
func Hash160(b []byte) []byte {
	h := sha256.New()
	h.Write(b)
	sha256Hash := h.Sum(nil)

	r := ripemd160.New()
	r.Write(sha256Hash)
	return r.Sum(nil)
}

// Codifica o endereço a partir do hash da chave pública
func EncodeAddress(pubKeyHash []byte) string {
	version := byte(0x00)
	versionedPayload := append([]byte{version}, pubKeyHash...)
	checksum := DoubleSha256(versionedPayload)[:4]
	fullPayload := append(versionedPayload, checksum...)
	return base58.Encode(fullPayload)
}

// Faz o duplo hash SHA256
func DoubleSha256(b []byte) []byte {
	first := sha256.Sum256(b)
	second := sha256.Sum256(first[:])
	return second[:]
}
