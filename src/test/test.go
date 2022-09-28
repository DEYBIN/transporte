package main

import (
	"fmt"
	"strings"
)

func main() {
	l := "Brayan Lei"
	a := l[0:1]
	fmt.Println(a)

	user := generateUser("brayan lei", "basurto", "huaman")
	fmt.Println(user)

}

func generateUser(name string, apl1 string, apl2 string) string {
	existNombre2 := false
	array_nombre := strings.Split(name, " ")
	nombre1 := array_nombre[0]
	nombre2 := ""

	if len(array_nombre) == 2 {
		existNombre2 = true
		nombre2 = array_nombre[1]
	}
	var user string
	for i := 0; i < 10; i++ {
		switch i {
		case 0:
			l := nombre1[0:1]
			user = l + apl1
		case 1:
			l := ""
			if existNombre2 {
				l = nombre2[0:1]

			} else {
				l = nombre1[0:1]

			}
			user = l + apl1
		case 3:
			l := ""
			if existNombre2 {
				l = nombre2[0:2]
			} else {
				l = nombre1[0:1]
			}
			user = l + apl1
		case 4:
			l := ""
			if existNombre2 {
				l = nombre1[0:3] + apl1[0:2] + apl2[0:3]
			} else {
				l = nombre1[0:1]
			}
			user = l
		
		}
		//consulta\
		if "lhuaman" == user {

			break
		}
	}
	return user

}
