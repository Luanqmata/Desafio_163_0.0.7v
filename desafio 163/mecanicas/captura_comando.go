package mecanicas

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// CapturaComando captura sinais do sistema para encerrar o programa
func CapturaSinal() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-signalChan
		fmt.Println("\nO usuário escolheu sair...")
		os.Exit(0)
	}()
}
