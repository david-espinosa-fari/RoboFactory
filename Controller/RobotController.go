package RobotController

import (
	"github.com/RobotFactory/Model"
	"github.com/RobotFactory/Views"
	"regexp"
)

var Len = 1000 //la cantidad de roboces a hacer
type RobotController struct {

}
func (Rc *RobotController)MakeAllRobots(){
	for i:=0;i<Len;i++{
		Rc.MakeRobot()
	}
}
func (Rc *RobotController)MakeRobot(){
	Mr := Model.MakeRobotModel{}
	var NameRobot = Mr.MakeRobot() //genera un nombre para el robot
	Cr := Model.CreatedRobots{}
	if Cr.FindRobot(NameRobot) == true {
		Vr := Views.MessageFromRobots{}
		Vr.RobotExits(NameRobot)
	}else {
		Cr.SetRobots(NameRobot)			//else Agrega robot
		Vr := Views.MessageFromRobots{}
		Vr.CreatedNewRobot(NameRobot)
	}
}
func (Rc *RobotController)DeleteRobot(name string)bool{
	Vr := Views.MessageFromRobots{}
	Cr := Model.CreatedRobots{} //donde estan los roboces ya creados
	if Cr.FindRobot(name)==true || name == "all"{
		Cr.DeleteRobot(name)
		if name == "all"{
			Vr.DeletedAllRobotMesage()
		}else{
			Vr.DeletedRobotMesage(name)
		}
	return true
	}else {
		Vr.RobotNotExitsMesage (name)
		return false
	}
}
func (Rc *RobotController)ListaRobots()  {
	Cr := Model.CreatedRobots{}
	Cr.ListarRoboces()
}
func (Rc *RobotController)UpdateRobots(nameOld string){
	Vr :=Views.MessageFromRobots{}
	var check = regexp.MustCompile(`^[A-Z]{2}\d{3}$`)
	if !check.MatchString(nameOld) {
		Vr.NameRobotNotValidMesage(nameOld)
	}else{
		if Rc.DeleteRobot(nameOld){
			Rc.MakeRobot()
			Vr.UpdateSuccessMesagge()
		}else {
			Vr.RobotFindOrNotDeleted ()
		}

	}
}





