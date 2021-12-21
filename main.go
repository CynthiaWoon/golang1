package main //every go file will be part of a package which is a collection of files and codes

//usually call the entry file as main.go
//here is the demo code from golang website

/*
if the package is called main, this tells the go compiler that the code should be compiled into an executable program at the end
so golang will create a standalone executable files for us if we choose to build the application and running the file will start the program
the package will be specified at the start of the file
if you are making some kind of shared library or utility code that could be used in other programs, then you will call the package something else
depending on what you are making, something that describes that library or code
in this case, main package is called since we are making a program to run on the computer
*/

/*
the standard library contains packages for all kinds of different functionality
fmt package if for formatting string and printing the messages to console
*/
import (
	"fmt" //importing packages from the go standard library
	"math"
	"strings"
)

//package scope variable
var score = 99.5

//function declaration
//main() is the entry point of the application
/*
when running the program, go will look through the files in the program and look
for this main function and fire this first one automatically
if the function name is something else beside "main", then go will not automatically fire this
when running the program

And there must be one and only one main function in the application
we do not have one in every file that inside the package, just have one main function in the entry file
*/
func main() {
	/*
		the methods all started with Capital when you use the method from the packages that you imported
		such as Println here, this is because inside the packages, you have to use a capital for whatever
		you are exporting
	*/
	fmt.Println("Hello, Golang") //fmt method

	//variables
	// remember every variable that is declared must be used, otherwise, it will cause an error
	//String
	var nameOne string = "mario"
	var nameTwo = "luigi"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree) //nameThree by default is empty since no value assigned

	//you can change the variable values
	nameOne = "peach"
	nameTwo = "bowser"
	nameThree = "anything"

	fmt.Println(nameOne, nameTwo, nameThree)

	//only use this when initializing the variable for the first time
	nameFour := "yoshi"
	fmt.Println(nameFour)

	//numbers - integer, float
	//you can specify the bit size for integer or float variations
	//the larger number we use for the variable, the higher number of bits you need
	//there will be a specfic range of numbers that we can use when we talk about the bit size
	//integer
	var ageOne int = 20
	var ageTwo = 30
	var ageThree int //by default is 0 for integer
	ageFour := 40
	fmt.Println(ageOne, ageTwo, ageThree, ageFour)

	//bits & memory
	var numOne int8 = 25    //8 bits
	var numTwo uint = 25    //unsigned int, which means you cannot have negative number
	var numThree uint8 = 25 //you can also specify bit size for unsigned number, but now you can go beyond the range that specified in signned number because you are not including minuses

	fmt.Println(numOne, numTwo, numThree)

	//float
	//unlike integers, you have to specify the bit size which dictates the range of numbers you can use
	var scoreOne float32 = -1.5
	var scoreTwo float64 = 223334455.98 //for the most part, float64 is going to be used because it has a slightly higher precision
	scoreThree := 1.5                   //in fact, if you use the operator :=, float64 will be default type

	fmt.Println(scoreOne, scoreTwo, scoreThree)

	//Sprintf (can save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", 35, "Nobody") //return formatted string and stored into the variable
	fmt.Println("the saved string is", str)

	//array declaration
	var ages [3]int = [3]int{20, 25, 30} //three items inside the array
	var ages2 = [4]int{20, 25, 30, 40}   //Go can also infer the type for the array

	fmt.Println(ages, ages2)
	//can also print out the length of the array
	fmt.Println(len(ages), len(ages2))

	//shorthand
	names := [5]string{"yoshi", "mario", "peach", "bowser", "nobody"}
	fmt.Println(names, len(names))

	//change the value
	names[1] = "luigi" //change mario to luigi
	fmt.Println(names, len(names))

	//slices (still use the array type under the hood)
	var scores = []int{100, 50, 60} //this time, you do not place number inside the square brackets which it will create the slice, not array

	//change the value
	scores[2] = 25 //change 60 to 25
	//append the item to the slice (it is not doable in array type)
	/*
		note: append() is not change the scores variable, it is going to return a new slice
		so if you want to update it with the new slice, you have to say "scores = append(...)"
		so you are overwritting this scores variable with a new slice
	*/
	scores = append(scores, 85) //append 85 to the scores slice
	fmt.Println(scores, len(scores))

	//slice range (is a way to get a range of elements from an existing array or slice and store them in new slice)
	rangeOne := names[1:3]  //go from position one to three, position three is excluded
	rangeTwo := names[2:]   //go from position two and to the end of the slice
	rangeThree := names[:3] //go from the start and get up to the position three but position three is excluded
	fmt.Println(rangeOne, rangeTwo, rangeThree)

	//you can append these ranges as well because they are stored as slices
	rangeOne = append(rangeOne, "koopa")
	fmt.Println(rangeOne)

	//loop
	x := 0

	//while loop
	for x < 5 { //while x is less than 5
		fmt.Println("value of x is:", x)
		x++ //add one to x
	}

	//traditional for loop
	for i := 0; i < 5; i++ {
		fmt.Println("value of i is:", i)
	}

	//loop through a slice
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	//for end loop
	/*
		cycle through the slice, and each time go around the slice,
		get the index(position) element and the value at that index
	*/
	for index, value := range names {
		fmt.Printf("the value at index %v is %v\n", index, value)
	}
	/*
		note: if you just wanna print the value but not index, you can replace the index with _
		e.g., for _, value
		and vice versa
		e.g., for index, _ (if you just want to use the index)
	*/

	//call function inside main
	sayGreeting("mario")
	sayGreeting("luigi")
	sayBye("mario")

	namesTwo := []string{"adik", "kakak", "abang"}
	//the function will be invoked inside the cycleNames function
	cycleNames(namesTwo, sayGreeting) //we dont invoke the function here, we pass it as reference of the function
	cycleNames(namesTwo, sayBye)

	a1 := circleArea(10.5) //the function is called and return a value and store inside a1
	a2 := circleArea(15)
	fmt.Printf("the circle 1 is %0.2f and the circle 2 is %0.2f\n", a1, a2)

	//since returning two values, you cannot only use one variable
	//so you need two variables
	firstname1, secondname1 := getInitial("tifa lockhart")
	fmt.Println(firstname1, secondname1)

	firstname2, secondname2 := getInitial("cloud strife")
	fmt.Println(firstname2, secondname2)

	firstname3, secondname3 := getInitial("cloud")
	fmt.Println(firstname3, secondname3)

	//call function from the other file
	// sayHello("mario")

	// showScore()

	//map
	menu := map[string]float64{ //specify the type of the keys in [], specify the values outside brackets
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55, //put the comma also after the last one, otherwise it will give the error
	}
	fmt.Println(menu)
	fmt.Println(menu["pie"]) //get the value of one of the items, make sure the key you enter the type is the same as you declared to get the value out

	//you can loop through the map
	for key, value := range menu { //instead of index and value, this time is key and value
		fmt.Println(key, "-", value)
	}

	//ints as key type
	phonebook := map[int]string{
		123456789: "mario",
		111222333: "luigi",
		456734566: "peach",
	}
	fmt.Println(phonebook)
	fmt.Println(phonebook[123456789]) //get mario

	//update an item inside the map
	phonebook[111222333] = "bowser" //update the luigi to bowser
	fmt.Println(phonebook)

	phonebook[456734566] = "yoshi"
	fmt.Println(phonebook)

}

//function with return value
func circleArea(r float64) float64 {
	return math.Pi * r * r
}

//function
func sayGreeting(n string) {
	fmt.Printf("Good morning %v\n", n)
}

func sayBye(n string) {
	fmt.Printf("Goodbye %v\n", n)
}

//function with multiple arguement
//note: we can also pass function in as arguement to another function
/*
this func(string) must take in string as that arguement which matches
the function that has string arguement in this case, sayGreeting or
sayBye can be passed in
*/
func cycleNames(n []string, f func(string)) {
	for _, value := range n {
		f(value) //call the function for each element inside the slice
	}
}

//function that return multiple values
func getInitial(n string) (string, string) { //return multiple values using extra parenthesis to write down the type of each returned value
	s := strings.ToUpper(n)
	names := strings.Split(s, " ") //split the name wherever there is a space [tifa, lockhart]

	var initials []string
	for _, value := range names {
		initials = append(initials, value[:1]) //get the initial for each element (value)
	}

	if len(initials) > 1 {
		return initials[0], initials[1] //return multiple values
	}

	return initials[0], "_" //if the length of the slice initials is smaller than one
}
