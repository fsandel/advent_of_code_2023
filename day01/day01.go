package main

import (
    "fmt"
    "os"
    "strings"
    "unicode"
)

func main() {
    dat, err := os.ReadFile("./input")
    if (err != nil) {
        return ;
    }

    splitted := strings.Split(string(dat), "\n");
    for _, word := range splitted {
        var first rune
        var first_set bool = false
        var last rune
        for _, letter := range word {
            if (unicode.IsNumber(letter)) {
                last = letter
                if (first_set == false) {
                    first = letter
                    first_set = false
                }
            }
            fmt.Print(first, last, "\n")
        }
        // fmt.Print( s, "\n")
    }
    // fmt.Print(string(dat))
}