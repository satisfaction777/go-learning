package main

import "fmt"

type Payment interface {
	Pay(amount float64) string // платежная операция
	Name() string       // название платежного метода
}

type Card struct {
	Owner  string
	Number string // последние 4 цифры карты
}

type Cash struct {
	Currency string // валюта платежа
}

type Crypto struct {
	Wallet string // адрес кошелька
}


// Реализация интерфейса Payment для Card

func (c Card) Name() string {
	return fmt.Sprintf("Карта **** %s", c.Number)
}

func (c Card) Pay(amount float64) string {
	return fmt.Sprintf("Оплачено картой **** %s: %.2f", c.Number, amount)
}

// Реализация интерфейса Payment для Cash

func (c Cash) Name() string {
	return fmt.Sprintf("Наличные %s", c.Currency)
}

func (c Cash) Pay(amount float64) string {
	return fmt.Sprintf("Оплачено наличными ( %s ): %.2f", c.Currency, amount)
}

// Реализация интерфейса Payment for Crypto

func (c Crypto) Name() string {
	return fmt.Sprintf("Крипто кошелек %s", c.Wallet)
}

func (c Crypto) Pay(amount float64) string {
	return fmt.Sprintf("Оплачено криптой с кошелька %s: %.2f", c.Wallet, amount)
}

func checkout(p Payment, amount float64) {
	fmt.Printf("Способ оплаты:%s\n", p.Name())
	fmt.Println(p.Pay(amount))
}

func main(){
	methods := []Payment{
		Card{Owner: "Павел Наумов", Number: "1234"},
		Cash{Currency: "RUB"},
		Crypto{Wallet: "ab12...ef421"},
	}	

	for _, method := range methods {
	checkout(method, 666.66)
	fmt.Println()
	}
}




