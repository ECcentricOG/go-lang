package main

import "fmt"

func main() {

	xs := []float64 {98,93,77,82,83}
	fmt.Println(avg(xs))
}

func avg(xs []float64) float64 {

	var total float64

	for _,ele := range xs {

		total += ele
	}
	return total/float64(len(xs))
}
