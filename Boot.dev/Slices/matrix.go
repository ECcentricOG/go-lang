package main

import "fmt"

func createMatrix(rows, cols int) [][]int {
    matrix := make([][]int, rows)
    for i:=0; i<rows; i++ {
        row := make([]int, cols)

        for j:=0; j<cols; j++ {
            row[j] = i * j
        }
       matrix[i] = row 
    }
    return matrix
}

func main() {
    fmt.Println(createMatrix(3, 3))
    fmt.Println(createMatrix(4, 4))
}
