package main

type Animal interface {
	Dog
	Cat
}

type Dog interface {
	bark()
}

type Cat interface {
	screem()
}

func main() {

}
