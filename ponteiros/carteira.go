package ponteiros

import "errors"

var ErroSaldoInsuficiente = errors.New("Não é possível retirar: saldo insuficiente")

type Carteira struct {
	saldo Bitcoin
}

func (carteira *Carteira) Depositar(quantidade Bitcoin) {
	carteira.saldo += quantidade
}

func (carteira *Carteira) Saldo() Bitcoin {
	return carteira.saldo
}

func (carteira *Carteira) Retirar(quantidade Bitcoin) error {
	if quantidade > carteira.saldo {
		return ErroSaldoInsuficiente
	}

	carteira.saldo -= quantidade
	return nil
}
