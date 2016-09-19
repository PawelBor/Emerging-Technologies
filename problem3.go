package main

import (
	"bufio"
    "fmt"
    "os"
	)

func reverse(value string) string {
 // Convert the string to rune.
 //Runes: characters
    data := []rune(value)
    result := []rune{}
 // Add runes in reverse order.
    for i := len(data) - 1; i >= 0; i-- {
		result = append(result, data[i])
    }

 // Returns the string.
    return string(result)
}

func main() {
//reader
	reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter Desired Word: ")
    value1, _ := reader.ReadString('\n')
    fmt.Println("Your word is: " + value1)

    reversed1 := reverse(value1)
    
    fmt.Println("Your Word in Reverse: " + reversed1)
}
