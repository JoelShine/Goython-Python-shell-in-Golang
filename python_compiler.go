package main

import (
    "fmt"
    "os"
    "strings"
    "bufio"
    "os/signal"
    "syscall"
)

// main function
func main() {

	fmt.Println("Goython - Python Compiler, written in Golang")
	fmt.Println("Type \"help\", \"copyright\" or \"license\" for more information. Type \"exit()\" to properly exit Goython.")

	for true {
		// KeyboardInterrupt exception
		c := make(chan os.Signal)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		go func() {
			<-c
			fmt.Println("Thanks for trying out Goython !")
			os.Exit(0)
		}()

		// Main code

		fmt.Print("\n>>> ")
		scanner := bufio.NewScanner(os.Stdin)

		if scanner.Scan() {
			input := scanner.Text()

			if strings.Contains(input, "print(\"") && strings.Contains(input, "\")") {

				print_remove := strings.Trim(input, "print")
				open_parantheses_remove := strings.Trim(print_remove, "(")
				close_parantheses_remove := strings.Trim(open_parantheses_remove, ")")

				text := strings.TrimLeft(strings.TrimRight(close_parantheses_remove,"\""),"\"")
				fmt.Println(text)

			} else if input == "exit()" {
				fmt.Println("Thanks for trying out Goython !")
				os.Exit(0)

			} else if input == "exit" || input == "close" || input == "quit"{
				fmt.Println("Press Ctrl+c or type 'exit()' to exit Goython")

			} else if input == "license" {
				fmt.Println(`
MIT License

Copyright (c) 2021 Joel Shine

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.`)

			} else if input == "copyright" {
				fmt.Println("Copyright (c) 2021 Joel Shine")

			} else if input == "help" {

				fmt.Println(`
Welcome to Goython help

How to start working with Python in Golang 

Just type your usual commands in python. Currently Goython supports only print statements.

To print a statement - print("Your statement here")
For help - help
To exit the shell - exit()

New features will be added. Stay Tuned to Goython.

Official website of Goython - https://github.com/JoelShine/Goython-Python-shell-in-Golang`)

			} else {
				fmt.Println("NameError: name '"+input+"' is not defined")
			}
		}
	}
}
//else: fmt.Println("NameError: name '"+input+"' is not defined")
