package main

import "fmt"

/*
New Struct
	Instructions
You'll need to compose a structure called Point, which will contain the
following fields (in the exact order):
	X - float
	Y - float
	Text - string
You will need to write a function that will return a newly created struct of
this with fields filled from the parameters provided.
*/


// Define a structure called Point
type Point struct{
	X float32
	Y float32
	Text string
}

// Function to create a new Point structure
func MakePoint(x, y float32, text string) Point {
	return Point{X:x, Y:y, Text:text}
}

/*
Point Difference
	Instructions
Using the same structure from the previous, given two Point structures, return the one that is further ahead in the X/Y coordinates. If they are equal, return either.
*/

//Type already defined above

// Function to find the Point with greater coordinates
func PointDiff(p1, p2 Point) Point {
	if p1.X > p2.X {
		return p1
	}else if p1.X < p2.X {
		return p2
	}else {
	if p1.Y > p2.Y {
		return p1
	}else if p1.Y < p2.Y{
		return p2
	}
	}
	return p1
}

/*
Point Text
	Instructions
Again using the same structure from the previous, given a Point structure,
return a new one that includes the same coordinates as the one given, but has
the text formatted in this way: "Text at (X, Y)" (with X and Y replaced with the actual coordinates).
Allowed packages
	fmt
*/

// Function to format the text of a Point structure
func PointText(p Point) Point{
	p.Text = fmt.Sprintf("Text at (%f, %f)", p.X, p.Y)
	return p
}

/*
Circle Result
	Instructions
You'll need a new structure. This will represent a circle and various properties of it.
The Circle structure should include the following fields in the exact order:
	Radius - Float
	Diameter - Float
	Area - Float
	Perimeter - Float
The function, called GetCircle, will take in the radius of the circle. It
should return a new Circle with all the fields filled with the correct values
from the radius given.
For π, you MUST use the value 3.14.
*/

// Define a structure called Circle
type Circle struct{
	Radius float32
	Diameter float32
	Area float32
	Perimeter float32
}

// Function to create a Circle structure with properties calculated from radius
func GetCircle(r float32) Circle {
	return Circle{Radius:r, Diameter: 2*r, Area:3.14*r*r, Perimeter:2*3.14*r}
}

func main() {

	//test cases for New Struct excercise
	samplePoint := MakePoint(5.2, 3.7, "This is a sample point")
	fmt.Printf("Sample Point: X = %.2f, Y = %.2f, Text = %s\n", samplePoint.X,samplePoint.Y, samplePoint.Text)
	//end

	//test cases for PointDiff excercise
	point2 := MakePoint(5.1, 2.8, "Point 2")
	point3 := MakePoint(5.1, 4.3, "Point 3")
	furthestPoint := PointDiff(point2, point3)
	fmt.Printf("Furthest point: %s\n", furthestPoint.Text)
	//end

	//test cases for Point Text excercise
	samPoint := MakePoint(4.7, 1.9, "before format")
	// Print the point before formatting
	fmt.Println("Sample Point before formatting:", samPoint.Text)
	// Call PointText to format the text
	samPoint = PointText(samPoint)
	// Print the point after formatting
	fmt.Println("Sample Point after formatting:", samPoint.Text)
	//end

	//test case for Circle Result excercise
	radius := float32(5)
	newCircle := GetCircle(radius)
	fmt.Printf("Radius: %f Diameter: %f Area: %f Perimeter: %f \n", newCircle.Radius, newCircle.Diameter, newCircle.Area, newCircle.Perimeter)
	//end
}

