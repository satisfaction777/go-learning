package main

import (
	"fmt"
	"sync"
)

type Order struct{
	ID int 
	Item string 
	Amount float64
}


func processOrder(o Order, wg *sync.WaitGroup, ch chan float64){ // Симуляция обработки заказа
	defer wg.Done()
	fmt.Printf("Заказ %d: %s - %.2f\n", o.ID, o.Item, o.Amount)
	ch <- o.Amount // Отправляем сумму в канал
}

func main(){
	orders := []Order{
	{1, "Ноутбук", 129990.00},
	{2, "Мышка", 4990.00},
	{3, "Монитор", 75990.00},
    {4, "Клавиатура", 13990.00},
	{5, "Наушники", 10990.00},
}

ch:= make(chan float64, len(orders)) // буфер = колво заказов



	var wg sync.WaitGroup

	for _, order := range orders { // Процессинг каждого заказа в отдельной горутине
		wg.Add(1)
		go processOrder(order, &wg, ch)
	}
	wg.Wait()
	close(ch) // закрываем канал после обработки всех заказов

	total:= 0.0
	for amount := range ch{
		total+= amount
	} 
	fmt.Println("Все заказы обработаны!")
	fmt.Printf("Итого: %.2f\n", total)
}



