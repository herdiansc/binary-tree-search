package main

import (
    "fmt"
    "math/rand"
    "time"
)

func ask(original, guest int) string {
    if original < guest {
        return "lower"
    } else if original > guest {
        return "higher"
    } else {
        return "equal"
    }
}

func main() {
    fmt.Println("Go: Lower-Biger")
    rand.Seed(time.Now().UnixNano())

    var original int = rand.Intn(1000)
    
    fmt.Printf("Original number is %d\n", original)
    
    var response string
    var guest, low, i int
    var high int = 1000
    for response != "equal" {
        guest = ((high-low)/2)+low
        response = ask(original, guest)
        if response == "lower" {
            high = guest
            fmt.Printf("Iteration: %d: guest:%d response: %s\n", (i+1), guest, response)
            guest = low
        } else if response == "higher" {
            low = guest
            fmt.Printf("Iteration: %d: guest:%d response: %s\n", (i+1), guest, response)
            guest = high
        } else {
            fmt.Printf("Iteration: %d: guest:%d response: %s\n", (i+1), guest, response)
        }
        
        i++
    }
}