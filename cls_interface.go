package main

import "fmt"

type Animals interface{
	Sleep()
	GetColor() string
	GetType()	string
}

type Cat struct{
	color string
}

func (this *Cat) Sleep(){
	fmt.Println("Cat is sleeping...")
}

func (this *Cat) GetColor() string{
	return this.color
}

func (this *Cat) GetType() string{
	return "Cat"
}

type Dog struct{
	color string
}

func (this *Dog) Sleep(){
	fmt.Println("Dog is sleeping...")
}

func (this *Dog) GetColor() string{
	return this.color
}

func (this *Dog) GetType() string{
	return "Cat"
}

func ShowAnimals(a Animals){
	a.Sleep()
	fmt.Println("color = ",a.GetColor())
}

func main(){
	// var animals Animals
	// animals = &Cat{"yellow"}
	// animals.Sleep()
	// animals.GetColor()
	// animals = &Dog{"white"}
	// animals.Sleep()
	// animals.GetType()

	cat := Cat{"yellow"}
	dog := Dog{"dark"}

	ShowAnimals(&cat)
	ShowAnimals(&dog)

}