package main

import (
	"fmt"
	"log"
	"strconv"
	mypackage "./mypackage/mypackage.go"
)


func normalFunction(message string)  {
	fmt.Println(message)
}

func manyArgs(a int, b int, c string){
	fmt.Println(a, b, c)
}

func returnValue(a int) int{
	return a*2
}

func doubeReturnValue(a int)(c , d int){
	return a , a * 2
}

func isPalindromo(text string){
	var textReverse string
	for i := len(text) -1; i>= 0; i--{
		textReverse += string(text[i])
	}
	if textReverse == text{
		fmt.Println("this is palindromo")
	}else{
		fmt.Println("this is not palindromo")
	}
}

type car struct {
	brand string
	year int
}

func main() {
	basic()

	mycar := car{brand: "kia",year:202}
	fmt.Println(mycar)


	var otherCar car
	otherCar.brand = "MG"

	fmt.Println(otherCar)



}

func basic (){

	fmt.Println("hello world")

	const pi float64 = 3.14
	const pi2 = 3.1415
		fmt.Println("this the value of the PI :",pi, "this is other value:", pi2 )
	base := 12
	var altura int = 16
	var area int
	fmt.Println(base, altura, area)

	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)

	const baseCuadrado int = 10
	areaCuadrado := baseCuadrado* baseCuadrado
	fmt.Println(areaCuadrado)


	x := 10
	y := 50

	result := x + y
	fmt.Println(result)

	result = x - y
	fmt.Println(result)

	result = x * y
	fmt.Println(result)

	result = x / y
	fmt.Println(result)

	result = x % y
	fmt.Println(result) 

	 helloMessage := "hello"
	wordMessage := "word"

	fmt.Println(helloMessage, wordMessage) 



	platzi := "platzi"
	cursos := 500

	fmt.Printf("%s tiene mas de %d curso\n",platzi,cursos)

	message := fmt.Sprintf("%v tiene mas de %v cursos \n",platzi, cursos)
	fmt.Println(message)


	fmt.Printf("helloMessage:\n", helloMessage)
	fmt.Printf("wordMessage:\n", wordMessage)

	normalFunction("helloMessage")
	manyArgs(1,2, "helloMessage") 

	value := returnValue(5)
	fmt.Println(value)

	value1, value2 := doubeReturnValue(5)
	fmt.Println(value1, value2) 


	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	counter:= 0
	for counter < 10 {
		fmt.Println(counter)
        counter++
	}

	// counterForever := 0
	// for  {
	// 	fmt.Println(counterForever)
	// 	counterForever++
	// }

	valueOne := 2
	valueTwo := 3
    
	if valueOne == 2 {
		fmt.Println("valueone is equal a 2")
	}else{
		fmt.Println("valueone is not equal a 2")
	}

	if valueOne == 2 && valueTwo == 3{
	fmt.Println("es verdad")
	}

	if valueOne == 0 || valueTwo == 3 {
		fmt.Println("valueOne is 0 or valueTwo is 3")
	}

	value, errr := strconv.Atoi("45")
	if errr!= nil {
        log.Fatal(errr)
	}
	fmt.Println("vlaue:", value)

	modulo := 4% 2
	switch modulo {
		case 0:
            fmt.Println("modulo is par")
        default:
            fmt.Println("modulo isn't par")
    
	}

	valueForSwitch := 200

	switch  {
		case valueForSwitch > 100 :
            fmt.Println("valueForSwitch is more than 100")
		case valueForSwitch < 0: 
		    fmt.Println("valueForSwitch is less than 0")
        default:
            fmt.Println("valueForSwitch isn't number")
	}


	defer	fmt.Println("hello")
	fmt.Println("word")

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 3{
			fmt.Println("this is",i)
			continue
		}
		if i == 8 {
			fmt.Println("break",i)
            break
		}
	}

	// array 
	var array  [4]int
	array[0] = 1
	array[1] = 2
	fmt.Println(array) 
	fmt.Println(len(array), cap(array ))


	//slices 

	slice := []int{1,2,3,4,5,6}	
	fmt.Println(slice, len(slice), cap(slice))

	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// append
	slice = append(slice, 7)
	fmt.Println(slice)

	newSlice := []int{8,9}

	slice = append(slice, newSlice...)
	fmt.Println(slice)


	sliceText:= []string {"hola", "que", "hace"}
	for i:= range sliceText{
		fmt.Println(sliceText[i])
	}

	isPalindromo("aman")
	

	m := make(map[string]int)

	m["jose"] = 13
	m["isaias"] = 20
	fmt.Println(m)

	for i, v := range m {
		fmt.Println(i, v)
	}

	noTFount := m["josep"]
	fmt.Println(noTFount)
}

