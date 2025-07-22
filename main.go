package main

import (
	"fmt"
)

type Pessoa struct {
	Nome	string
	Idade 	int
}

type Conta struct {
	Dono 	Pessoa
	Saldo 	float64
}

type ContaPoupanca struct {
	Dono		Pessoa
	Saldo   	float64
	JurosAno	float64
}

type Operacoes interface {
	Depositar(valor float64)
	Sacar(valor float64) bool
	VerSaldo() float64
}

func (c *Conta) Depositar(valor float64) {
	c.Saldo += valor
}

func (c *Conta) Sacar(valor float64) bool {
	if c.Saldo >= valor {
		c.Saldo -= valor
		return true
	} else {
		fmt.Println("Saldo insuficiente")
		return false
	}
}

func (c Conta) VerSaldo() float64 {
	return c.Saldo
}

func (c *ContaPoupanca) Depositar(valor float64) {
	c.Saldo += valor
}

func (c *ContaPoupanca) Sacar(valor float64) bool {
	if c.Saldo >= valor{
		c.Saldo += valor
		return true
	}else {
		fmt.Println("Saldo Insuficiente")
		return false
	}
}

func (c ContaPoupanca) VerSaldo() float64 {
	return c.Saldo
}