package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaClaudio := contas.ContaCorrente{
		Titular:       clientes.Titular{Nome: "Claudio"},
		NumeroAgencia: 0,
		NumeroConta:   0,
	}

	contaClaudio.Depositar(100)
	PagarBoleto(&contaClaudio, 50)
	fmt.Println("Conta: ", contaClaudio.Titular.Nome, "| Saldo:", contaClaudio.Obtersaldo())

	contaJose := contas.ContaPoupanca{
		Titular:       clientes.Titular{Nome: "Jose"},
		NumeroAgencia: 0,
		NumeroConta:   0,
	}

	contaJose.Depositar(100)
	PagarBoleto(&contaJose, 100)
	fmt.Println("Conta: ", contaJose.Titular.Nome, "| Saldo:", contaJose.Obtersaldo())

}
