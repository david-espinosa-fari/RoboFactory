package Views

import "fmt"

type MessageFromRobots struct {

}
func (MfR *MessageFromRobots) RobotExits (name string){
	fmt.Printf("El Robot: %s" ,name)
	fmt.Printf(" ya existe \n")
}

func (MfR *MessageFromRobots)CreatedNewRobot(name string){

	fmt.Printf("Se ha creado un nuevo Robot: %s\n" ,name)
}

func (MfR *MessageFromRobots) DeletedRobotMesage (name string){
	fmt.Printf("El Robot: %s" ,name)
	fmt.Printf(" se ha eliminado con exito \n")
}
func (MfR *MessageFromRobots) RobotNotExitsMesage (name string){
	fmt.Printf("El Robot especificado %s" ,name)
	fmt.Printf(" no existe \n")
}
func (MfR *MessageFromRobots) DeletedAllRobotMesage (){
	fmt.Printf("Todos los roboces han sido borrados\n")

}
func (MfR *MessageFromRobots) NameRobotNotValidMesage(name string){
	fmt.Printf("El Nombre de robot proporcionado %s" ,name)
	fmt.Printf(" no es valido. \n")
	fmt.Printf("Debe estar en el formato AA123 \n")
}
func (MfR *MessageFromRobots) UpdateSuccessMesagge(){
	fmt.Printf("El robot se ha agregado con exito. \n")
}
func (MfR *MessageFromRobots) RobotFindOrNotDeleted (){
	fmt.Printf("Lo siento no podemos realizar esta operacion o bien porque el robot nuevo ya existe  \n")
	fmt.Printf("o porque no se encontro un robot para modificar \n")
}
