package main

import (
	"github.com/FerdinaKusumah/fastcrud/internal"
	"github.com/FerdinaKusumah/fastcrud/server"
)

func main(){
	internal.ParseOption()
	server.Run()
}