package main

import (
	"bufio"
	"fmt"
	"os"
	
)

func main() {
	signal := "Pokedex >"
	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Printf("%v", signal)
		if !scanner.Scan() {
			break
		}
		userInput := cleanInput(scanner.Text())
		if len(userInput) > 0 {
			if cmd, exist := commandList[userInput[0]]; exist {
				if err := cmd.callback(areaPointer); err != nil {
					fmt.Printf("Error executing command: %v\n", err)
				}
			} else {
				fmt.Printf("Unknown command: %s\n", userInput[0])
			}
		} else {
			fmt.Println("请输入非空命令")
		}
	}
}
