package main

import (
	"fmt"
)

type Strategy interface {
	Execute(a, b int) int
}
type ConcreteStrategyAdd struct{}
func (csa ConcreteStrategyAdd) Execute(a, b int) int {
	return a + b
}
type ConcreteStrategySubtract struct{}

func (css ConcreteStrategySubtract) Execute(a, b int) int {
	return a - b
}

type ConcreteStrategyMultiply struct{}

func (csm ConcreteStrategyMultiply) Execute(a, b int) int {
	return a * b
}
type Context struct {
	strategy Strategy
}
func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}
func (c *Context) ExecuteStrategy(a, b int) int {
	return c.strategy.Execute(a, b)
}
type ExampleApplication struct{}

func (ea ExampleApplication) Main() {
	context := Context{}

	var firstNumber int
	fmt.Print("Enter the first number: ")
	fmt.Scan(&firstNumber)
	var secondNumber int
	fmt.Print("Enter the second number: ")
	fmt.Scan(&secondNumber)

	var action string
	fmt.Print("Enter the desired action (addition/subtraction/multiplication): ")
	fmt.Scan(&action)

	var strategy Strategy

	switch action {
	case "addition":
		strategy = ConcreteStrategyAdd{}
	case "subtraction":
		strategy = ConcreteStrategySubtract{}
	case "multiplication":
		strategy = ConcreteStrategyMultiply{}
	default:
		fmt.Println("Invalid action.")
		return
	}

	context.SetStrategy(strategy)
	result := context.ExecuteStrategy(firstNumber, secondNumber)
	fmt.Printf("Result: %d\n", result)
}

func main() {
	app := ExampleApplication{}
	app.Main()
}
