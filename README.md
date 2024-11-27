# Terminal Input Scanner Utility for Go

This Go program provides utility functions to read user input from the terminal with prompts, supporting strings, integers, and floating-point numbers. It simplifies the process of reading and parsing user input in command-line applications.

## Features

- **GetString**: Read a string input from the user.
- **GetInt**: Read and parse an integer input from the user.
- **GetFloat**: Read and parse a floating-point number input from the user.
- Customizable prompt messages for each input.

## Usage

### Installation

To use these utility functions in your Go project, you can copy the functions into your code or package them into a module.

### Example Code

Below is an example of how to use these functions in a Go program:

```go
package main

import (
    "fmt"
)

func main() {
    name, err := GetString("Enter your name: ")
    if err != nil {
        fmt.Println("Error reading name:", err)
        return
    }

    age, err := GetInt("Enter your age: ")
    if err != nil {
        fmt.Println("Error reading age:", err)
        return
    }

    salary, err := GetFloat("Enter your desired salary: ")
    if err != nil {
        fmt.Println("Error reading salary:", err)
        return
    }

    fmt.Printf("\nName: %s\nAge: %d\nDesired Salary: %.2f\n", name, age, salary)
}
```
