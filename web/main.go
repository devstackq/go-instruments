package main

import (
	"web/handlers"
)



func main(){

	t := handlers.NewTriangle()
	t.X=11
	t.Y=12
	t.Z=11
	// p.R=12
	var p  handlers.Point
	p.R = 14

	handlers.DoThing(&p)
	handlers.DoThing(t)

	//func(s *store)GetUserById(id int){}
	//store := NewStore(db)
	//serv := service.ServiceStruct.NewService(store)
	//serv.GetUser(10)
}