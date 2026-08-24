package main

import (
	"fmt"
	"math"
)

// type Point struct {
// 	X int
// 	Y int
// }

// func (p Point) Distance(other Point) float64 {

// 	dx := float64(other.X - p.X)
// 	dy := float64(other.Y - p.Y)

// 	sqr := math.Sqrt(dx*dx+dy*dy)

// 	fmt.Println(sqr)

// 	return sqr
// }
// // ---------------------------------------------------------
// type Circle struct {
// 	radius float64
// }

// func (c *Circle) Area() float64 {
// 	return math.Pi * (c.radius * c.radius)
// }

// func (c Circle) Circumference() float64 {
// 	return 2 * math.Pi * c.radius
// }

// func NewCircle(radius float64) (*Circle,error) {

// 	if radius < 0 {
// 		return nil,errors.New("Radius can't be negative")
// 	}

// 	return &Circle{radius: radius},nil

// }
// // ----------------------------------------------------------

// type Vehicle struct {
// 	Make string
// 	Model string
// }

// func (v Vehicle) Describe() string {
// 	return "this car is " + v.Make + " and model is " + v.Model
// }

// type Car struct {
// 	Vehicle
// 	Year int
// }

// type Motorcycle struct {
// 	Vehicle
// 	Seat int
// }

// // --------------------------------------------------------

// type BankAccount struct {
// 	balance float64
// }

// func (ba *BankAccount) Deposit (amount float64) error {
// 	if amount < 0 {
// 		return errors.New("Can't deposit negative amount")
// 	}

// 	ba.balance += amount

// 	fmt.Println("Amount is deposited")

// 	return nil
// }

// func (ba *BankAccount) Withdraw(amount float64) error {
// 	if amount > ba.balance {
// 		return errors.New("Insufficient balance")
// 	}
// 	if amount < 0 {
// 		return errors.New("Can't withdraw negative amount")
// 	}

// 	ba.balance -= amount

// 	fmt.Println("amount is withdrawed")

// 	return nil
// }

// func (ba *BankAccount) Balance() float64 {
// 	return ba.balance
// }

// --------------------------------------------------------

// type Person struct {
// 	Name 		string 		`json:"name"`
// 	Email 		string 		`json:"email"`
// 	Password 	string 		`json:"password"`
// }

// --------------------------------------------------------

type Rectangle struct {
	Height float64
	Width float64
}

type Circle struct {
	Raduis float64
}

type Triangle struct {
	Side1 float64
	Side2 float64 
	Side3 float64
	Height float64
}

func (r Rectangle) Area() float64 { 
	return r.Height* r.Width
}

func (r Rectangle) Perimeter() float64 { 
	return 2 * (r.Height + r.Width)
}

func (c Circle) Area() float64 { 
	return math.Pi * (c.Raduis * c.Raduis)
}

func (c Circle) Perimeter() float64 { 
	return 2 * math.Pi * c.Raduis
}

func (t Triangle) Area() float64 { 
	return  ( t.Side1 * t.Height ) / 2 
}

func (t Triangle) Perimeter() float64 { 
	return t.Side1 + t.Side2 + t.Side3
}

func PrintRectangleAreas(rects []Rectangle) { 
	 for _,rect := range rects {
		fmt.Println(rect.Area())
	 }
}

func PrintCricleAres(circles []Circle) {
	for _, circle := range circles {
		fmt.Println(circle.Area())
	}
}

func PrintTriangleAreas(triangles []Triangle) { 
	for _, triangle := range triangles { 
		fmt.Println(triangle.Area())
	}	
}

// -------------------------------------------

type Matrix struct {
	
}

func main(){






	// p := Person{Name: "Aashish",Email: "aashish@gmail.com",Password: "123123"}
	// j,err := json.Marshal(p)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(string(j))


	// var p2 Person

	// err  = json.Unmarshal(j,&p2)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(p2)

	// fmt.Printf("type of p is %T \n",p)
	// fmt.Printf("type of j is %T \n",j)

	// acc := BankAccount{balance: 1000}

	// acc.Deposit(500)

	// acc.Balance()

	// acc.Withdraw(1000)

	// acc.Balance()
	
	// myPoint := Point{X: 5,Y: 9}

	// otherPoint := Point{X: 10,Y: 20}

	// distance := myPoint.Distance(otherPoint)

	// fmt.Println(distance)

	// circle := Circle{radius: 5}	

	// fmt.Println(circle.Area())
	
	// fmt.Println(circle.Circumference())
	

	// car1 := Car{Vehicle: Vehicle{Make: "Toyota",Model: "4500Y"},Year: 2005 }

	// fmt.Println(car1)

	// fmt.Println(car1.Make)

	// fmt.Println(car1.Model)

	// fmt.Println(car1.Describe())

}