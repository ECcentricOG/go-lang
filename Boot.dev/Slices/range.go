package main

import "fmt"

func indexOfFirstBadWord(msg []string, badWords []string) int {
    bI := len(msg)
    isPresnet := false
    for _, bw := range badWords {
        for i, m := range msg {
            if m == bw && bI > i{
                bI = i 
                isPresnet = true
            }
        }
    }
    if isPresnet {
        return bI
    }
    return -1
}

func main() {
    bWords := [4]string{"crap", "shoot", "frick", "dang"}
    fmt.Println(indexOfFirstBadWord([]string{"hey", "there", "john"}, bWords[:]))
    fmt.Println(indexOfFirstBadWord([]string{"ugh", "oh", "my", "frick"}, bWords[:]))
    fmt.Println(indexOfFirstBadWord([]string{"what", "the", "shoot", "I", "hate", "that", "crap"}, bWords[:]))
}
