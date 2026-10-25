package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// Lab 1

func lab1(){
	fmt.Println("Hellow, World!")

	var a int = 5
	var b int = 6
	c := a + b
	fmt.Println(c)


	c+=a
	fmt.Println(c)

	c-=b
	fmt.Println(c)

	c*=a
	fmt.Println(c)
	
	c/=b
	fmt.Println(c)

	c%=a
	fmt.Println(c)
}

// Lab 2

func lab2(){
	var a int = 5
	var b int = 5

	fmt.Println(a>0 && b>0) 

	fmt.Println(a>b || b>a)

	fmt.Println(a!=b)
}

// Lab 3

func myAND(a, b int){fmt.Printf("Before a = %b\n", a); fmt.Printf("Before b = %b\n", b); a&=b; fmt.Printf("After a&=b = %b\n", a)}

func myXOR(a, b int){fmt.Printf("Before a = %b\n", a); fmt.Printf("Before b = %b\n", b); a^=b; fmt.Printf("After a^=b = %b\n", a)}

func myOR(a, b int){fmt.Printf("Before a = %b\n", a); fmt.Printf("Before b = %b\n", b); a|=b; fmt.Printf("After a|=b = %b\n", a)}

func myRight(a, b int){fmt.Printf("Before a = %b\n", a); fmt.Printf("Before b = %b\n", b); a>>=b; fmt.Printf("After a>>=b = %b\n", a)}

func myLeft(a, b int){fmt.Printf("Before a = %b\n", a); fmt.Printf("Before b = %b\n", b); a<<=b; fmt.Printf("After a<<=b = %b\n", a)}

func lab3(){
	var a int = 4
	var b int = 12

	myAND(a,b)
	fmt.Println()

	myXOR(a,b)
	fmt.Println()

	myOR(a,b)
	fmt.Println()

	myRight(a,b)
	fmt.Println()

	myLeft(a,b)
}

// Lab 4

func add(a, b int) int{
	return a+b
}

func sub(a, b int)int{
	return a-b
}

func mul(a, b int) int {
	return a*b
}

func div(a, b int) float32{
	return float32(a)/float32(b) 
}

func mod(a, b int) int{
	return a%b
}

func lab4(){
	var a int
	var b int
	var choice int
	for (choice != 6){
		fmt.Println("===== Mini Calculator =====\n1. Add\n2. Sub\n3. Mul\n4. Div\n5. Mod\n6. Exit")
		fmt.Println()

		fmt.Print("Choose : ")
		fmt.Scanln(&choice)

		if(choice==6 || choice > 6){
			fmt.Print("Exiting....")
			break
		}

		fmt.Print("Enter a : ")
		fmt.Scanln(&a)

		fmt.Print("Enter b : ")
		fmt.Scanln(&b)
		fmt.Println()

		if (choice==1){
			result := add(a,b)
			fmt.Println("Result :", result)
			fmt.Println()
		}else if(choice==2){
			result := sub(a,b)
			fmt.Println("Result :", result)
			fmt.Println()
		}else if(choice==3){
			result := mul(a,b)
			fmt.Println("Result :", result)
			fmt.Println()
		}else if(choice==4){
			result := div(a,b)
			fmt.Printf("Result : %.2f\n", result)
			fmt.Println()
		}else{
			result := mod(a,b)
			fmt.Println("Result :", result)
			fmt.Println()
		}
	}
}

// Lab 5
func lab5(){
	var input string
	fmt.Print("Enter : ")
	fmt.Scanln(&input)

	data := []byte(input)
	fmt.Printf("Input in binary : %b\n", data)

	hex := hex.EncodeToString(data)
	fmt.Println("Input in Hexadecimal :", hex)

	base64String := base64.StdEncoding.EncodeToString(data)
	fmt.Println("Input in Base64 :", base64String)

}

// Lab 6

func xorEncrypt(text string, key byte) string{
	data := []byte(text)
	
	for i := 0; i<len(data); i++{
		data[i] = data[i] ^ key
	}
	return  string(data)
}

func lab6(){
	var text string
	fmt.Print("Enter a message : ")
	fmt.Scanln(&text)

	var key byte
	fmt.Print("Enter a key : ")
	fmt.Scanln(&key)
	var output = xorEncrypt(text, key)
	fmt.Printf("Result : %v", output)
}

func main(){
	// lab1()
	// lab2()
	// lab3()
	// lab4()
	// lab5()
	lab6()
}