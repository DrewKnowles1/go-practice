package main 

import (
	"fmt"
)
//build struct below. Just a 'worker', with some attributes
type (
	worker struct{
		firstname string 
		lastname string
		employeeNumber int
		title string
		salary int
	}
)

//Getters for func properties
func (w1 *worker)getFirstname() string {
	return w1.firstname
}

func (w1 *worker)getLastname() string {
	return w1.lastname
}

func (w1 *worker)getSalary() int {
	return w1.salary
}

//Setters for func properties
func (w1 *worker)setSalary(salary int) {
	 w1.salary = salary
}

func returnAllWorkers()(workers []worker){
	// function that creates a slices of type worker, adds 2 workes to it, and returns it
	 w1 := worker { firstname: "Andrew", lastname: "Knowles", employeeNumber: 123, title: "Engineer", salary: 999999}
	 w2 := worker { firstname: "Jay", lastname: "Marshall", employeeNumber: 174, title: "Engineer", salary: 1}
	 workers = append(workers, w1, w2)
	 return workers
}

func workerSalaryAlignment(workerPreAdjust worker)(workerPostAdjust worker){
	//If your name is me, you get a raise
	if workerPreAdjust.getFirstname() == "Andrew" && workerPreAdjust.getLastname() == "Knowles"{ 
		//set value of old salary to struct's salary value
		oldSalary := workerPreAdjust.getSalary()
		//set new slary 1000 x that, because im going a good job :) 
		newSalary := oldSalary * 1000
		//set salary of struct to the new caluclated value
		workerPreAdjust.setSalary(newSalary)

		//formatting int to string
		oldSalaryText := fmt.Sprintf("\n\nHi Andrew! \nRaise for u!\n \nold salary: " + "%d", oldSalary)
		newSalaryText := fmt.Sprintf("New salary: " + "%d", workerPreAdjust.getSalary() )

		//print message informing me of my fantastic new salary
		fmt.Printf(oldSalaryText + "\n" + newSalaryText + "\n\n")

	//If your name isnt me, you do not get a raise
	} else {
		fmt.Println("hey, " + workerPreAdjust.getFirstname() + " " + workerPreAdjust.getLastname() + " - You aren't me, no more money for you!")
	}
	
	workerPostAdjust = workerPreAdjust
	return workerPostAdjust
}