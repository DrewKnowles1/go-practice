package main 

//build struct below. Just a 'worker', with some attributes
type (
	worker struct{
		firstname string 
		lastname string
		employeeNumber int
		title string
	}
)

func returnAllWorkers()(workers []worker){
	// function that creates a slices of type worker, adds 2 workes to it, and returns it
	 w1 := worker { firstname: "Andrew", lastname: "knowles", employeeNumber: 123, title: "Engineer"}
	 w2 := worker { firstname: "Jay", lastname: "Marshall", employeeNumber: 174, title: "Engineer"}
	 workers = append(workers, w1, w2)
	 return workers
}