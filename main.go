package main

import (
    "fmt"
)

type Conta struct {
    Nome  string
    CPF   string
    Saldo float64
}

func (c *Conta) Depositar(valor float64) {
    c.Saldo += valor
}

func (c *Conta) Sacar(valor float64) bool {
    if c.Saldo >= valor {
        c.Saldo -= valor
        return true
    }
    return false
}

func (c Conta) VerSaldo() float64 {
    return c.Saldo
}

func main() {
    contas := make(map[string]*Conta)

    for {
        fmt.Println("\n====== BANCO GOLANG ======")
        fmt.Println("1. Criar conta")
        fmt.Println("2. Depositar")
        fmt.Println("3. Sacar")
        fmt.Println("4. Ver saldo")
        fmt.Println("5. Sair")
        fmt.Print("Escolha uma opção: ")

        var opcao int
        fmt.Scanln(&opcao)

        switch opcao {
        case 1:
            var nome, cpf string
            fmt.Print("Nome: ")
            fmt.Scanln(&nome)
            fmt.Print("CPF: ")
            fmt.Scanln(&cpf)

            if _, existe := contas[cpf]; existe {
                fmt.Println("Já existe uma conta com esse CPF.")
            } else {
                contas[cpf] = &Conta{Nome: nome, CPF: cpf, Saldo: 0}
                fmt.Println("Conta criada com sucesso!")
            }

        case 2:
            var cpf string
            fmt.Print("CPF: ")
            fmt.Scanln(&cpf)

            conta, ok := contas[cpf]
            if !ok {
                fmt.Println("Conta não encontrada.")
                continue
            }

            var valor float64
            fmt.Print("Valor para depositar: R$ ")
            fmt.Scanln(&valor)
            conta.Depositar(valor)
            fmt.Println("Depósito realizado!")

        case 3:
            var cpf string
            fmt.Print("CPF: ")
            fmt.Scanln(&cpf)

            conta, ok := contas[cpf]
            if !ok {
                fmt.Println("Conta não encontrada.")
                continue
            }

            var valor float64
            fmt.Print("Valor para sacar: R$ ")
            fmt.Scanln(&valor)
            if conta.Sacar(valor) {
                fmt.Println("Saque realizado!")
            } else {
                fmt.Println("Saldo insuficiente.")
            }

        case 4:
            var cpf string
            fmt.Print("CPF: ")
            fmt.Scanln(&cpf)

            conta, ok := contas[cpf]
            if !ok {
                fmt.Println("Conta não encontrada.")
            } else {
                fmt.Printf("Titular: %s\n", conta.Nome)
                fmt.Printf("Saldo: R$ %.2f\n", conta.VerSaldo())
            }

        case 5:
            fmt.Println("Saindo...")
            return

        default:
            fmt.Println("Opção inválida. Tente novamente.")
        }
    }
}
