package main

import (
	"fmt"
	"strings"
	"transporte/src/library/sqlquery"
)

func main() {
	l := "Alexa"
	a := l[0:1]
	fmt.Println(a)

	user := generateUser("Alexa Daniela", "Cartolin", "Tovar")
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
	for i := 0; i <= 10; i++ {
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
		case 2:
			l := ""
			if existNombre2 {
				l = nombre2[0:2]
			} else {
				l = nombre1[0:1]
			}
			user = l + apl1
		case 3:
			l := ""
			if existNombre2 {
				l = nombre1[0:2] + apl1[0:2] + apl2[0:2]
			} else {
				l = nombre1[0:1]
			}
			user = l
		case 4:
			l := ""
			if existNombre2 {
				l = nombre1[0:3] + apl1[0:2] + apl2[0:3]
			} else {
				l = nombre1[0:1]
			}
			user = l
		case 5:
			l := ""
			if existNombre2 {
				l = nombre1[0:4] + apl1[0:2] + apl2[0:2]
			} else {
				l = nombre1[0:1]
			}
			user = l
		case 6:
			l := ""
			if existNombre2 {
				l = nombre1 + apl1[0:1] + apl2[0:2]
			} else {
				l = nombre1[0:1]
			}
			user = l
		case 7:
			l := ""
			if existNombre2 {
				l = nombre1[0:2] + apl1 + apl2[0:1]
			} else {
				l = nombre1[0:1]
			}
			user = l
		case 8:
			l := ""
			if existNombre2 {
				l = nombre1[0:3] + apl1[0:2] + apl2
			} else {
				l = nombre1[0:1]
			}
			user = l
		case 9:
			l := ""
			if existNombre2 {
				l = nombre1 + apl1[0:2] + apl2[0:2]
			} else {
				l = nombre1[0:1]
			}
			user = l
		case 10:
			l := ""
			if existNombre2 {
				l = nombre1 + apl1[0:3] + apl2[0:4]
			} else {
				l = nombre1[0:1]
			}
			user = l
		}
		//consulta\
		usuario := sqlquery.NewQuerys("Seguridad").Select("users").Where("users","=",user).Exec().Text("users")
		
		if usuario == user {
			break
		}
	}
	return user

}
