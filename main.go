package main

import "fmt"

//To anyone reading this, The point of this project is to test what i remember from the course ive been doing
//As well as function as a quick reference guide for syntax should i need it in future 
//I apologiese for the horrid code ;) 

func main(){

	fmt.Println("Hello World \n")
	fmt.Println("My function to add strings: " + stringAdd("I'm", "Redundant")) 
	printSliceContents(returnTeamMembers())
}

func stringAdd(string1 string, string2 string)(concatString string){
	//Im aware this function is pointless :)
	concatString = string1 + " " + string2
	return concatString
} 

func returnTeamMembers() (teamMembers []string){
	teamMembers = append(teamMembers, "Andrew Knowles", "Izzie duthie")
	return teamMembers
}

func printSliceContents(slice []string){
	fmt.Println("My function to print the contents of a slice: ")
	for _, index := range slice{
		fmt.Println(index)
	}
}