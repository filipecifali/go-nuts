package main

import (
    "math/rand"
    "fmt"
    "time"
)

func CrazeString(s string) string {
    result := []byte("")
    for i := 0; i < len(s); i++ {
        j := rand.Intn(len(s))
        result = append(result,s[j])
    }
    return string(result)
}

func main(){
    word := "Don't use me in production!"
    rand.Seed(time.Now().UnixNano())
    fmt.Println(CrazeString(word))
}
