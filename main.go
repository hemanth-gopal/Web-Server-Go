package main

import (
	"fmt"
	"log"
	"http/net"
)

func formHandler(w http.ResponseWriter,r *http.Request){
	if err:= r.ParseForm(); err!=nil{
		fmt.FPrintf(w,"ParseForm() err: %v",err)
		return
	}
	fmt.Fprintf(w,"POST request successfull")
	name := r.FormValue("name")
	address := r.FormValue("Address")
	fmt.Fprintf(w,"Name = %s\n",name)
	fmt.Fprintf(w,"Address = %s\n",address)
}

func helloHandler(w http.ResponseWriter,r *http.Request){
	if r.URL.Path != "/hello"{
		http.Error(w,"404 not found",http.StatusNotFound)
		return
	}
	if r.Method !="GET"{
		http.Error(w,"Method is not supported",http.StatusNotFound)
		return
	}
	fmt.FPrintf(w,"Hello Handler Function")
}
func main(){
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/",fileServer)
	http.HandleFunc("/form",formHandler)
	http.HandleFunc("/hello",helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err:= http.ListerAndServer(":8080",nil); err!=nil{
		log.Fatal(err)
	}
}