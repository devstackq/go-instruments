package grpc

//framework - для взаимодействие сервсивов
// вызовд удаленрй прцоендуры

// grpc - use http2, fast (vs http1)
// - protocol buffer -  binary - small(vs json)
//protobuf - easy, simple
//create .proto file
// use protoc
//bidirectional - vs(client-server)
//limited browser support

//conf & start app
//server package - business logic  app
// makefile - automatization rutine

// protoc --go_out=plugins=grpc:. - генерирвоать код из протофайла

func main(){
	g := server.New()
	g.Start()
	// g.WaitStop()

}

