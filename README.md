# Desafio 163 - Carteira Bitcoin

## Informações Gerais

- **Canal**: [Investidor Internacional](https://www.youtube.com/@investidorint)
- **Descrição**: O desafio consiste em encontrar a chave privada Bitcoin que corresponde ao padrão fornecido.
- **Recompensa**: R$ 900 para quem encontrar a chave privada correta.

## O Desafio

### Padrão da Chave Privada:

```
403b3d4xcxfx6x9xfx3xaxcx5x0x4xbxbx7x2x6x8x7x8xax4x0x8x3x3x3x7x3x
```

**Nota**: Onde há "x", substitua por um dos caracteres hexadecimais:
- `0-9`
- `A-F`

### Exemplo de Solução Válida:

Se o padrão fosse:
```
40xx
```
Uma chave privada válida seria:
```
40A3
```

## Ambiente Recomendado

- **Processadores Testados**:
  - Xeon E5-2666 v3
  - Xeon E5-2680 v4
  - i9-14900K

- **Requisitos Mínimos**:
  - Memória RAM: 24 GB (ou mais, para otimização do processo).
  - Threads: 28 ou mais.
  
## Etapas do Desafio

1. **Preparação do Ambiente**:
    - Configure seu sistema com o máximo de RAM possível para melhorar a geração de chaves privadas por segundo.
    - Utilize threads paralelas para maior eficiência no cálculo.

2. **Estratégia para Gerar Chaves**:
    - Gerar combinações aleatórias de caracteres hexadecimais para preencher os "x".
    - Verificar se a chave gerada é válida no contexto do desafio.

3. **Ferramentas e Linguagem**:
    - Recomendação: **Go** pela sua eficiência em tarefas paralelas.
    - Utilize buffers de memória para otimizar a geração.

## Benchmarking (Referências)

### Teste em i9-14900K

- **Velocidade de Geração**: 50 milhões de chaves/s.

![Benchmark i9-14900K](https://github.com/user-attachments/assets/979f526a-7e6b-4b0d-a749-06bb5f168296)

### Teste em Xeon 2680 v4

- **Velocidade de Geração**: ~35 milhões de chaves/s (estimado).

## Código de Referência

Para iniciar, você pode usar o código abaixo em Go:

```go
package main

import (
    "crypto/rand"
    "encoding/hex"
    "fmt"
    "math/big"
    "strings"
    "sync"
)

func gerarChavePrivada(padrao string) string {
    caracteres := "0123456789ABCDEF"
    resultado := make([]byte, len(padrao))

    for i, c := range padrao {
        if c == 'x' {
            idx, _ := rand.Int(rand.Reader, big.NewInt(int64(len(caracteres))))
            resultado[i] = caracteres[idx.Int64()]
        } else {
            resultado[i] = byte(c)
        }
    }

    return string(resultado)
}

func main() {
    padrao := "403b3d4xcxfx6x9xfx3xaxcx5x0x4xbxbx7x2x6x8x7x8xax4x0x8x3x3x3x7x3x"

    var wg sync.WaitGroup
    threads := 28 // Configure o número de threads desejado

    for i := 0; i < threads; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for {
                chave := gerarChavePrivada(padrao)
                fmt.Println("Chave Gerada:", chave)
                // Adicione lógica para verificar se a chave é a correta
            }
        }()
    }

    wg.Wait()
}
```

## Como Participar

1. Execute o código acima em sua máquina configurada.
2. Gere chaves dentro do padrão especificado.
3. Submeta a chave encontrada ao canal [Investidor Internacional](https://www.youtube.com/@investidorint).

---

Boa sorte no desafio! 🚀
