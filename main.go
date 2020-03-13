package main

import (
	"bufio"
	decoder "falabella-tlv/tlv-decoder"
	"fmt"
	"os"
	"strings"
)

func main() {

	decoder.Decode([]byte{})

	var option string
	inputType := bufio.NewScanner(os.Stdin)

	fmt.Println("- - - - -")
	showInstructions()
	fmt.Println("- - - - -")

	for {

		inputType.Scan()
		option = inputType.Text()
		option = strings.ToLower(option)

		switch option {
		case "decode":
			fmt.Printf("Enter the TLV fields\n")
			inputType.Scan()
			data := inputType.Text()
			res, err := decoder.Decode([]byte(data))
			if err != nil {
				fmt.Printf("An error occurred while processing the fields: %v\n", err.Error())
				fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - - - -")
				continue
			}
			printMap(res)

		case "help":
			showInstructions()

		case "exit":
			return

		default:
			fmt.Printf("Unknown command %v\n", option)
		}

		fmt.Println("- - - - - - - - - - - - - - - - - - - - - - - - - - - - - -")
	}
}

func printMap(m map[string]string) {
	fmt.Printf("Resulting Map:\n")
	for k, v := range m {
		fmt.Printf("%v: %v\n", k, v)
	}
}

func showInstructions() {
	fmt.Println("decode		-> decode entered TLV fields into a map")
	fmt.Println("help		-> shows this list of commands")
	fmt.Println("exit 		-> to exit")
}
