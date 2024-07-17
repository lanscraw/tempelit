package main

import (
    "errors"
    "fmt"
)

type Service struct {
    Name string
}

// Mock function to simulate service retrieval or creation.
func GetInstance(serviceName, context string, factory func() (any, error)) (any, error) {
    // Here you would typically check if the service instance already exists.
    // If not, use the factory function to create it.
    if serviceName == "" || context == "" {
        return nil, errors.New("invalid service name or context")
    }

    instance, err := factory()
    if err != nil {
        return nil, err
    }

    // Returning the created instance
    return instance, nil
}

func main() {
    tmallOrdersvc, err := GetInstance("ordersvc", "tmall", func() (any, error) {
        // Creating a new instance of the service
        return &Service{Name: "Tmall Order Service"}, nil
    })

    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Type assertion to convert from `any` to `*Service`
    ordersvc := tmallOrdersvc.(*Service)
    fmt.Printf("Service Name: %s\n", ordersvc.Name)
}
