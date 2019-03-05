package Model

import (
	"fmt"
)

type CreatedRobots struct {
}

var(
	listOfRobots [1000] string
	Count int
)
func (Cr *CreatedRobots)SetRobots(name string){
	if Count < len(listOfRobots) {
		listOfRobots[Count] = name
		Count ++
	}else {
		for i:=0;i<Count;i++{
			if listOfRobots[i]=="" {
				listOfRobots[i]=name
			}
		}
	}
}
func (Cr *CreatedRobots)ListarRoboces(){
	if Count!= 0{
		for i:=0;i<Count;i++{
			if listOfRobots[i]!=""{
				fmt.Printf("Los roboces que hay en el arreglo %s\n",listOfRobots[i])
			}
		 }
	}else {
		fmt.Printf("No existen Roboces \n")
	}
}
func (Cr *CreatedRobots)FindRobot(name string)bool{

	for i:=0;i<Count ; i++ {
		if name == listOfRobots[i]{
			return true
		}
	}
	return false
	}
func (Cr *CreatedRobots)DeleteRobot(name string){
	if name =="all"{
		for i:=0;i<len(listOfRobots) ; i++ {
			listOfRobots[i]=""
		}
	}else{
		for i:=0;i<len(listOfRobots) ; i++ {
			if name == listOfRobots[i]{
				listOfRobots[i]=""
				break
			}
		}
	}
}