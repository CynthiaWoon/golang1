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
package main //every go file will be part of a package which is a collection of files and codes

/*
the standard library contains packages for all kinds of different functionality
fmt package if for formatting string and printing the messages to console
*/
import "fmt" //importing packages from the go standard library

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

}
