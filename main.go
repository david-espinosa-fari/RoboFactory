package main

import (
	"fmt"
	"github.com/RobotFactory/Controller"
	"os"

	//"github.com/RobotFactory/Controller"
)
func main(){

	Rc := RobotController.RobotController{}
	Rc.MakeAllRobots()
	for  {
	var option int
	fmt.Printf("Que desea hacer?\n")
	fmt.Printf("Listar todos los roboces `1` \n")
	fmt.Printf("Eliminar un Robot: `2` \n")
	fmt.Printf("Actualizar un Robot: `3` \n")
	fmt.Printf("salir: `4` \n")
		fmt.Scan(&option)
		switch option {
		case 1 :
			Rc.ListaRobots()
		case 2 :
			var RobotToDelete string
			fmt.Printf("Robot a eliminar: \n")
			fmt.Scan(&RobotToDelete)
			Rc.DeleteRobot(RobotToDelete)
			Rc.ListaRobots()
		case 3 :
			var nameOld string
			fmt.Printf("Nombre del robot a actualizar: \n")
			fmt.Scan(&nameOld)
			Rc.UpdateRobots(nameOld)
		case 4:
			os.Exit(0)
		}

	}
}
