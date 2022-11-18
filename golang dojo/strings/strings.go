package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := "Welcome to the Golang Dojo!"

	fmt.Println(s)
	//BUILTIN STRING FUNCTIONS
	fmt.Println("The length of the string s", len(s))
	fmt.Printf("The first letter of the string is %c", s[0])
	fmt.Println()
	fmt.Println("The first word of the string is", s[0:8])

	fmt.Println("The remainder of the string is", s[8:])
	s += "my friends :)"
	fmt.Println("New string:", s)

	// STRINGS AND BYTE SLICES ARE INTERCHANGEABLE
	abc := "abc"
	b := []byte(abc)
	fmt.Printf("%s %s", abc, b)
	fmt.Println()
	abc2 := string(b)
	fmt.Println("b converted to string gives", abc2)

	// THE STRING LIBRARY
	fmt.Println()
	fmt.Println(strings.ToUpper("Hello to my dojo friends :)!"))
	fmt.Println(strings.ToLower("Hello to my dojo friends :)!"))
	fmt.Println(strings.HasPrefix("Hello to my dojo friends :)!", "hello"))
	fmt.Println(strings.HasPrefix("Hello to my dojo friends :)!", "Hello"))
	fmt.Println(strings.HasSuffix("Hello to my dojo friends :)!", "friends :)!"))
	fmt.Println(strings.Replace("Hello to my dojo friends :)!", "dojo", "programming", 1))

	// TRY THE STRING BUILDER CLASS
	var sb strings.Builder
	fmt.Println()
	fmt.Println("This is a string builder:", sb.String())
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())
	sb.WriteString("Hello")
	fmt.Println("This is a string builder:", sb.String())
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())
	sb.Grow(100)
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())
	fmt.Println("This is a string builder:", sb.String())
	sb.Reset()
	fmt.Println(sb.Cap())
	fmt.Println(sb.Len())
	fmt.Println("This is a string builder:", sb.String())
	fmt.Println()

	sb.Write([]byte{65, 66, 67})
	fmt.Println(sb.String())

	x := 123
	y := strconv.Itoa(x)
	fmt.Println(y)
	fmt.Printf("Type of y is %T", y)
	fmt.Println()

	z, _ := strconv.Atoi(y)
	fmt.Println(z)
	fmt.Printf("Type of z is %T", z)

}
