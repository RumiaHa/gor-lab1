package main

import (
    "fmt"
    "math/rand/v2"
    "sync"
    "time"
)

type Order struct {
    ID       int
    Customer string
    Amount   float64
}

func processOrder(order Order, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Обрабатывается заказ %d для %s...\n", order.ID, order.Customer)
    duration := time.Duration(rand.IntN(3)) * time.Second
    time.Sleep(duration)
    fmt.Printf("Заказ %d на сумму %.2f успешно обработан\n", order.ID, order.Amount)
}

func main() {

    orders := []Order{
        {ID: 101, Customer: "Sofia", Amount: 1999999999.12},
        {ID: 102, Customer: "Marina", Amount: 312.32},
        {ID: 103, Customer: "Misha", Amount: 0.9},
        {ID: 104, Customer: "Sasha", Amount: 0.3},
        {ID: 105, Customer: "Tamik", Amount: 0.8},
    }

    var wg sync.WaitGroup
    for _, order := range orders {

        wg.Add(1)

        go processOrder(order, &wg)
    }

    wg.Wait()

    fmt.Println("Все заказы оформлены")
}
