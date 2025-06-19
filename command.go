package main

import (
	"fmt"
	"os"
	"github.com/cry0404/pokdexcli/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.Config) error
}



// 注意这里是包级别的声明，所以需要 var 关键字
var apiClient *pokeapi.Client
var commandList map[string]cliCommand
var areaPointer *pokeapi.Config //用来指向当前的区域是什么，通过当前的区域来构建 get 请求
func init() {
	apiClient = pokeapi.NewClient()
	areaPointer = &pokeapi.Config{}

	commandList = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "print the helpful message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "list next twenty areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "list last twenty areas",
			callback:    commandMapb,
		},
	}
}
//get https://pokeapi.co/api/v2/location-area/{id or name}/
func commandExit(p *pokeapi.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(p *pokeapi.Config) error {
	//这里应该通过遍历 map 表来得到最新的信息
	fmt.Println(" Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	for _, cmd := range commandList {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(p *pokeapi.Config) error {
	resp, err := apiClient.GetLocationAreas(p.NextURL)
	if err != nil{
		return fmt.Errorf("failed to get location areas: %w", err)
	}
	for _, area := range resp.Results {
		fmt.Printf("%s\n", area.Name)
	}

	if resp.Next != nil {
		p.NextURL = *resp.Next
	}else{
		p.NextURL = ""
	}

	if resp.Previous != nil {
		p.PreviousURL = *resp.Previous
	}else{
		p.PreviousURL = ""
	}

	return nil
}

func commandMapb(p *pokeapi.Config) error {
	if p.PreviousURL == "" {
		fmt.Println("您处于第一页")
		return nil
	}

	resp, err := apiClient.GetLocationAreas(p.PreviousURL)
	if err != nil {
		return fmt.Errorf("failed to get previous location area %w", err)
	}

	fmt.Println("Previous Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf("%s\n", area.Name)
	}

	// 更新配置状态
	if resp.Next != nil {
		p.NextURL = *resp.Next
	} else {
		p.NextURL = ""
	}
	if resp.Previous != nil {
		p.PreviousURL = *resp.Previous
	} else {
		p.PreviousURL = ""
	}

	return nil
}
