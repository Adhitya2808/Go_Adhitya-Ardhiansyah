package main

import (
	"mvc/configs"
	"mvc/routes"
)

func main (){
	configs.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}