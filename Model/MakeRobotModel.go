package Model

import (
	"math/rand"
	"strconv"
	"time"
)

type MakeRobotModel struct {
	NameRobot string
}
func (Mr *MakeRobotModel) MakeRobot()string{
		strconv.Itoa(RandNumber())
		RandLetter()
		Mr.NameRobot = RandLetter() + strconv.Itoa(RandNumber())
return Mr.NameRobot
}
func RandNumber() int {
	//rand.Seed(time.Now().UnixNano()) // takes the current time in nanoseconds as the seed
	var Num, Rand int
	for ;Num < 100;  {
		rand.Seed(time.Now().UnixNano() + rand.Int63())
		//rand.Seed(time.Now().UnixNano())
		Rand = rand.Intn(1000)
		Num = Rand
		if Num > 99 && Num < 1000{
			break
		}else {
			if Num + Rand < 1000 && Num + Rand > 100 {
				Num = Num +Rand
				break
			}
		}
	}
//	fmt.Printf("Numero %v\n" ,Num)
	return Num
}
func RandLetter() string {
	var Rand int
	var letter, Randletter string
	var ABC = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//rand.Seed(time.Now().UnixNano())
	for i:=0; i < 2; i++{ /*len(LettersRobot*/
		//rand.Int63() //devuelve un numero entero aleatorio de 64 bits que es el que necesita seed
		rand.Seed(time.Now().UnixNano() + rand.Int63()) //para que existan menos coliciones modifico cada vez la semilla
		Rand = rand.Intn(26)						//si se espicifa fuera del loop menos probabilidades todavia
		Randletter = string(ABC[Rand])
		letter = letter + Randletter
	}
	return letter
}