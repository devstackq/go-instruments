package main

import (
	"log"
	"web/handlers"
)

type Point struct {
	x int
	y int
}

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

	// adminMux := http.NewServeMux()
	// adminMux.HandleFunc("/admin/", adminIndex)

	//func(s *store)GetUserById(id int){}
	//store := NewStore(db)
	//serv := service.ServiceStruct.NewService(store)
	//serv.GetUser(10)
pointer := 2
log.Println(pointer, &pointer)
n := &pointer
	*n = 22


p1 := Point{11, 20}

p2 := &Point{112, 20}

log.Println(p1, p2)

}

func init(){
	log.Println("init first")
}