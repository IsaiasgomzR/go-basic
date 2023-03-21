	package struc

	type car struct{
		brand string
		year int
	}

	func main (){
		mycar:= car{
			brand: "kia",
			year:2022
		}
		fmt.Println(mycar)
	}