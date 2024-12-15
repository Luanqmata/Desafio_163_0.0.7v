package style

import (
	"fmt"
)

func Logo_163() {
	fmt.Println(`
			8888888b.                              .d888 d8b               d888    .d8888b.   .d8888b.
			888  "Y88b                            d88P"  Y8P              d8888   d88P  Y88b d88P  Y88b
			888    888                            888                       888   888             .d88P
			888    888  .d88b.  .d8888b   8888b.  888888 888  .d88b.        888   888d888b.      8888"
			888    888 d8P  Y8b 88K          "88b 888    888 d88""88b       888   888P "Y88b      "Y8b.
			888    888 88888888 "Y8888b. .d888888 888    888 888  888       888   888    888 888    888
			888  .d88P Y8b.          X88 888  888 888    888 Y88..88P       888   Y88b  d88P Y88b  d88P
			8888888P"   "Y8888   88888P' "Y888888 888    888  "Y88P"      8888888  "Y8888P"   "Y8888P"     Moldado para : Xeon E5 2666 v3
	`)
}

func Bem_vindo() {
	fmt.Print("\n\n\t\t\t\t\t\t\t\t~~ Bem vindo ~~\n\n")
	Logo_163()
	fmt.Print("\n\t\t\t\t\t\t\tDesafio do Investidor Internacional (Dii)\n\n")
	fmt.Print("\n\t\t\t\t\t\t\t\tBy: China /gpt 0.3v\n\n")
}

func Opcoes_uso_proc() {
	fmt.Println("\t\t\t\t\t\t========== Menu de Uso do Processador ==========\n")

	fmt.Println("\t\t\t\t\t\t\t[1]  25%   ▓▓▓░░░░░░░░░░░░░░")
	fmt.Println("\t\t\t\t\t\t\t[2]  50%   ▓▓▓▓▓▓▓▓░░░░░░░░░")
	fmt.Println("\t\t\t\t\t\t\t[3]  75%   ▓▓▓▓▓▓▓▓▓▓▓▓░░░░░")
	fmt.Println("\t\t\t\t\t\t\t[4] 100%   ▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓▓")
	fmt.Println("\n\t\t\t\t\t\t==================================================")
	fmt.Print("\t\t\t\t\t\t\tDigite sua escolha: ")
}
