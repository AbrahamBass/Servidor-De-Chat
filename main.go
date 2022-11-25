package main

import "fmt"

type IProduct interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}

type Computer struct {
	name  string
	stock int
}

func (c *Computer) setStock(stock int) {
	c.stock = stock
}

func (c *Computer) setName(name string) {
	c.name = name
}

func (c *Computer) getStock() int {
	return c.stock
}

func (c *Computer) getName() string {
	return c.name
}

type Lapto struct {
	Computer
}

type Desktop struct {
	Computer
}

func newLapto() IProduct {
	return &Lapto{
		Computer: Computer{
			name:  "Lapto Computer",
			stock: 22,
		},
	}
}

func newDesktop() IProduct {
	return &Desktop{
		Computer: Computer{
			name:  "Desktop Computer",
			stock: 30,
		},
	}
}

func GetComputerFactory(computerType string) (IProduct, error) {
	if computerType == "laptop" {
		return newLapto(), nil
	}

	if computerType == "desktop" {
		return newDesktop(), nil
	}

	return nil, fmt.Errorf("Invalid Computer")
}


func printNameAndStock (p IProduct) {
	fmt.Printf("Product Name: %v And Stock %v\n", p.getName(), p.getStock())
}


func main() {
	laptop, err := GetComputerFactory("laptop")
	desktop, err := GetComputerFactory("desktop")

	if err != nil {
		fmt.Println(err)
	}

	printNameAndStock(laptop)
	printNameAndStock(desktop)

}
