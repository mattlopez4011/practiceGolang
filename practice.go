package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, Student!")

	yourAge := 2

	switch yourAge {
	case 16 : fmt.Println("Go Drive")
	case 18 : fmt.Println("Go Vote")
	default : fmt.Println("Go Have Fun")
	}

}



//Variable Declaration
//var foo int
//var foo int = 21
//foo := 21

// For Loop
/* i := 1
for i <= 10 {
fmt.println(i)
i++ // or // i = i + 1
}

for j := 0; j < 5; j++ {
fmt.Println(j)
}

 */

//Conditional
/*if 7%2 == 0 {
	fmt.Println("7 is even")
} else{
	fmt.Println("7 is odd")
}
*/

/*if num := 9; num < 0 {
fmt.Println(num, "is negative")
} else if num < 10 {
fmt.Println(num, "has 1 digit")
} else {
fmt.Println(num, "has multiple digits")
}
*/

//Switch Statement
/*
yourAge := 5

switch yourAge {
	case 0 <= 3  : fmt.Println("Hi Baby!")
	case 16 : fmt.Println("Go Drive")
	case 18 : fmt.Println("Go Vote")
	default : fmt.Println("Go Have Fun")
}
*/
/*	switch hour := time.Now().Hour(); { // missing expression means "true"
	case hour < 12:
		fmt.Println("Good morning!")
		fmt.Println("Time is " + strconv.Itoa(hour))
	case hour < 17:
		fmt.Println("Good afternoon!")
	case hour <= 10:
		fmt.Println("Oh Yeahhhhhh!!")
	default:
		fmt.Println("Good evening!")
	}

*/
/*
Array
var favNums2[5] float64

favNums2[0] = 163
favNums2[1] = 164
favNums2[2] = 165
favNums2[3] = 166
favNums2[4] = 167

fmt.Println(favNums2[3])
favNums3 := [5]float64 {1,2,3,4,5}
 */


