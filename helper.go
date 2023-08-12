package main 

import (
	"strings"
	
)

func validate(Firstname string, lastname string, email string, usertickets uint) (bool, bool) {

	isvalid := len(Firstname) >= 2 && len(lastname) >= 2 && strings.Contains(email, "@")
	isvalidd := strings.Contains(email, ".") && usertickets > 0

	return isvalid, isvalidd

}
