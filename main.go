package main

import "fmt"

func main() {
	var name string
	name = "revathy"

	fmt.Println(name)
}

//chapter 1:formatting verbs/words
// fmt.Printf.....fmt.SPrintf()---->STRING FORMATTING
// %v------->FORMTTING DEFAULT REPRESENTATION(v:value)
//%s---->INTERPOLATE A STRING(s:string)
// %f----->interpolate A DECIMAL(f:decimal or float say)

//CONDITIONALS--> NO PARENTHESES AROUND THE CASES
 if age>30 {
	fmt.Println("the person is over 30")
} else if age = 30{
	fmt.Println("the person is 30")
} else {
	fmt.Println("the person comes under 30")
}
// if INITIAL_STATEMENT; CONDITION{}
if length:= getLength(email); length<1{
	fmt.Println("Email is invalid")
}
//CHAPTER 2:FUNCTIONS IN GO
// in go var type comes after var name for readability
func add(i int , j int) int{
	return i+j
}
// HOW TO IGNORE RETURN VALUES--->
func addvalue() (i int, j int){
	return 1,3
}
x,_ :=addvalue()// so 3 is ommitted

//keep in mind in GO , go compiler will throw err if u have unsused var decl in ur code so ignore anything not in use.
//naked returns----returning the named values returns
//explict returns--->
func getPoints() (x,y int){
	return x,y// EXPLICTLY RETURNS X Y
}

//CHAPTER 3:STRUCTS

//CHAPTER 4:INTERFACES
//CHAPTER 5:ERRORS
//CHAPTER 6:LOOPS
//CHAPTOER 7:ARRAYS N SLICES
//CHAPTER 8: MAPS
//CHAPTER 9:ADVANCED FUNCTIONS
//CHAPTER 10:POINTERS
//CHAPTER 11:LOCAL DEVELOPMENTS
//CHAPTER 12:CHANNELS
//CHAPTER 13:MUTEXES
//CHAPTER 14:GENERICS