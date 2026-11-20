# Fundamentos da Computação com Go

Código das aulas da trilha **Fundamentos da Computação com Go**: dez aulas e uma aula bônus, que começam do zero absoluto e terminam num projeto prático, um gerenciador de tarefas no terminal.

Não precisa saber programar. Cada aula tem um vídeo, e o código mostrado nele fica aqui, numa pasta própria.

## As aulas

| Aula | Assunto | Código |
|---|---|---|
| 00 | Boas-vindas à trilha | [`aulas/aula00`](aulas/aula00) |
| 01 | O que acontece quando você roda um código | [`aulas/aula01`](aulas/aula01) |
| 02 | Git e GitHub | em breve |
| 03 | Variáveis e memória | em breve |
| 04 | Decisões com if e switch | em breve |
| 05 | Repetições com for | em breve |
| 06 | Funções e erros | em breve |
| 07 | Arrays e slices | em breve |
| 08 | Maps | em breve |
| 09 | Structs e métodos | em breve |
| 10 | Ponteiros | em breve |
| Bônus | Projeto final: gerenciador de tarefas | em breve |

Toda semana sai uma aula nova, e o código dela entra aqui no mesmo dia.

## Como rodar

1. Instale o Go 1.27.4 ou mais novo: <https://go.dev/dl/>. A instalação passo a passo está na Aula 01.
2. Baixe este repositório. Pelo botão **Code › Download ZIP** desta página, ou com o Git (ensinado na Aula 02):

   ```
   git clone https://github.com/DaksonC/fundamentos-engenharia-go.git
   ```

3. Entre na pasta da aula e rode:

   ```
   cd fundamentos-engenharia-go/aulas/aula01
   go run main.go
   ```

   Saída:

   ```
   Olá, mundo!
   ```

Cada pasta `aulas/aulaNN` é um projeto Go independente, com o seu próprio `go.mod`. Por isso, rode os comandos de dentro da pasta da aula.

## Windows, Mac e Linux

Os vídeos mostram os comandos no PowerShell do Windows. No Mac, no Linux e no WSL, alguns mudam:

| Windows (PowerShell) | Mac e Linux |
|---|---|
| `go build -o ola.exe` | `go build -o ola` |
| `.\ola.exe` | `./ola` |

O `go run main.go` é igual nos três.

## Sobre

Feito por Dakson Cruz · <https://daksonc.github.io/>
