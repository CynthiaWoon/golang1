//you create a new file to define custom type
//you do not use classes in go but instead you are using struct to define a object
package main

import "fmt"

/*
to make this custom struct type, write "type" to make a new type
and "struct" to say it is a structure
inside the curly brances, you define different properties of this structure
*/
type bill struct { //this is the blueprint for any bill object
	name  string
	items map[string]float64
	tip   float64
}

//make new bills
func newBill(name string) bill {
	//this is how you create bill object
	//equal to the type and curly braces which is used to
	//define the initial values for the properties
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0, //remember to also put the comma after the last one
	}
	return b
}

//receiver function to aassociate the function to the object
//format new bill
//Go also copy the bill object to pass it into the function, just like how they did
//when passing the variables to the function
func (b bill) format() string { //now this function belongs to bill object, then you cannot call this function without delcare a bill object first
	fs := "Bill breakdown: \n"
	var total float64 = 0

	//list items
	for key, value := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", key+":", value) //-25v means add the spaces to the rest of the characters and minus means add to the right
		total += value
	}

	//tip
	fs += fmt.Sprintf("%-25v ...$%v \n", "tip:", b.tip)

	//total
	fs += fmt.Sprintf("%-25v ...$%v \n", "total:", total)

	return fs
}

//to update the original bill object, it is the same as how the original variable is updated
//create a pointer and pass the pointer into the function to update the original value since the
//pointer points to the memory location where the bill object is stored
//however, for structs, inside the function, you do not need to explicitly dereference it because go will automatically
//dereference it, all you need to do is only pass the pointer as the receiver

/*
another reason to pass the pointer is because we will no be making copies everytime when we call
the function if you want to call the function quite frequently
if you do not have the pointer, then you would be creating a new copy everytime of the bill, if you just pass
in the pointer, you are only copy the pointer
besides, if your bill object is larger and more complex, then to make a copy of that is more work than just
copy the pointer
*/

func (b *bill) newFormat() string { //now this function belongs to bill object, then you cannot call this function without delcare a bill object first
	fs := "Bill breakdown: \n"
	var total float64 = 0

	//list items
	for key, value := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", key+":", value) //-25v means add the spaces to the rest of the characters and minus means add to the right
		total += value
	}

	//tip
	fs += fmt.Sprintf("%-25v ...$%v \n", "tip:", b.tip)

	//total
	fs += fmt.Sprintf("%-25v ...$%v \n", "total:", total+b.tip)

	return fs
}

func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}
