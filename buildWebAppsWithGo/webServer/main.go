package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func helloWorld(w http.ResponseWriter,r *http.Request){

	r.ParseForm()//解析参数
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v := range r.Form{
		fmt.Println("key",k)
		fmt.Println("val:",strings.Join(v,""))
	}
	fmt.Fprintf(w,"Hello alex!")
}

//实现登录
func login(w http.ResponseWriter,r*http.Request){
	fmt.Println("method: ",r.Method)
	if r.Method =="GET"{
		t,_:=template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w,nil))
	}else{
		r.ParseForm()
		fmt.Println("username: ",r.Form["username"])
		fmt.Println("password: ",r.Form["password"])
	}
}


func main(){
	http.HandleFunc("/",helloWorld) // 注册了请求/的路由规则,
	//当请求uri为"/",路由就会转到函数helloWorld
	http.HandleFunc("/login",login)
	err:=http.ListenAndServe(":9090",nil)
	if err!=nil{
		log.Fatal("ListenAndServe: ",err)
	}
}
