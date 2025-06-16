package main

import (
	"os"
	"bufio"
	"fmt"
)

func main(){
		signal := "Pokedex >"
		scanner := bufio.NewScanner(os.Stdin)

		for {
			
			fmt.Printf("%v", signal)
			if !scanner.Scan(){
				break;
			}
			userInput := cleanInput(scanner.Text())
			if len(userInput) > 0{
				fmt.Println("Your command was:",userInput[0])
				}else{
				fmt.Println("请输入非空命令")
			}
		}
}