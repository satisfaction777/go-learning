package main

import(
	"fmt"
)

type Transfer struct{
	FromAccount string
	ToAccount string
	Amount float64
}

type TransferError struct {
	Field string
	Message string
}

func (e *TransferError) Error() string {
	return fmt.Sprintf("ошибка перевода [поле '%s']: %s", e.Field, e.Message)
}


func validateTransfer(t Transfer) error {
	if t.Amount <= 0 {
		return &TransferError{Field: "Amount", Message: "сумма должна быть больше нуля"}
	}
	if t.Amount > 100000 {
		return &TransferError{Field: "Amount", Message: "превышен лимит перевода 100000"}
	}
	if t.FromAccount == "" {
		return &TransferError{Field: "FromAccount", Message: "счёт отправителя не указан"}
	}
	if t.ToAccount == "" {
		return &TransferError{Field: "ToAccount", Message: "счёт получателя не указан"}
	}
	if t.FromAccount == t.ToAccount {
		return &TransferError{Field: "ToAccount", Message: "нельзя переводить на тот же счёт"}
	}
	return nil
}

func processTransfer(t Transfer) error {
	err := validateTransfer(t)
	if err != nil {
		return fmt.Errorf("processTransfer: %w", err)
	}
	fmt.Printf("Перевод выполнен: %s -> %s, сумма: %.2f рублей\n", t.FromAccount, t.ToAccount, t.Amount)
	return nil
}




func main() {
	transfers := []Transfer{
    {"acc001", "acc002", 9999},      // ок
    {"acc001", "acc002", -100},     // ошибка суммы
    {"acc001", "acc002", 200000},   // превышен лимит
    {"", "acc002", 500},            // нет отправителя
    {"acc001", "acc001", 500},      // одинаковые счета
}

for _, t := range transfers{
	err := processTransfer(t)
	if err != nil {
		fmt.Println(err)
	}
}
}