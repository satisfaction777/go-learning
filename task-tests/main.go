package main


import (
	"fmt"
)

func calcBalance (income, food, transport float64) float64 {
	return income - food - transport
}



func canAfford(balance, price float64) (bool, float64) {
	if balance > price{
		return true, balance - price
	} else {
		return false, price - balance
	}
}


func validateAge(age int) error {
    if age < 0 {
        return fmt.Errorf("возраст не может быть отрицательным")
    }
    if age > 150 {
        return fmt.Errorf("возраст не может быть больше 150")
    }
    return nil
}

func main() {
	
}



