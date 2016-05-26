package main

/*
 Lesson001-fmt_package
 */

/* Tip-001 : import syntax for multiple imports */
import (
	"fmt"
)

/* Tip : Declaring a constant which is visible only within the package ( because the name start with lower case ) */
const (
	pi = 3.14
)

/* Tip; use of iota for creating const with auto generated values */

const (
	sunday = iota
	monday
	tuesday
	wednesday
	thursday
	friday
	saturday
)

/* Tip : Declaring a variable whose scope is within the package */
var (
	mark = 100
)

func main() {
	/* Tip-002 : variable declaration with type */
	var message string

	/* Tip-003 : variable assignment with value */
	message = "Hello World\nGoLang first program\n"

	/* Tip-004 : Invoking print function in the fmt package */
	fmt.Print(message)

	/* Tip-005 : shortcut to declare and assign value with auto inference */
	name := "My Name is GoLang"

	fmt.Print(name)

	/* Tip-006 : using formatted string to format the output */
	messagefmt := "The Mark is %d"

	/* Tip-007: using the formatted Print method to recognize the format patterns, otherwise it will be treated as string */
	fmt.Print(messagefmt, 100, "\n")

	/* Tip-008: the formatted print (Printf) take a message format and the values as input) */
	messagefmt = "The Mark is %d\n"
	fmt.Printf(messagefmt, 100)

	/* Tip-009: Use the Println method so that an \n is automatically added at the line */
	fmt.Println("First Line is %d" , 1)
	fmt.Printf("Second Line is %d\n" , 2)

	/* print the days const */
	fmt.Println("days: ",sunday, monday, tuesday)
}
