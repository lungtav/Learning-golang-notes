package main

import "fmt"


type rect struct{
		height int
		width int
	}

func (r rect) area() int{
		return r.height * r.width
	}

func main(){


	//struct is used to represent structured data. it is used to represent data that objects and dictionary would be used for in other languages


	type car struct{
		brand string
		model string
		doors int
		mileage int
	}

	myCar:= car{
		brand: "Mercedez",
		model: "GLB",
		doors: 4,
		mileage: 1000,
	}
	fmt.Println(myCar)


	//nested structs
	type wheel struct{
		radius int
		material string
	}
	
	type car2 struct {
		brand string
		model string
		doors int
		mileage int
		frontWheel wheel
		backWheel wheel
	}

	mySecondCar:= car2{
		brand: "Mercedez",
		model: "GLB",
		doors: 4,
		mileage: 1000,
		frontWheel: wheel{
			radius: 4,
			material: "cement",
		},
		backWheel: wheel{
			radius: 3,
			material: "plastic",
		},

	}

	fmt.Println(mySecondCar)
	
	
	//STRUCTS METHODS
	
	rectangle1:= rect{
		height: 5,
		width: 5,
	}

	fmt.Println(rectangle1.height, rectangle1.width, rectangle1.area())

}