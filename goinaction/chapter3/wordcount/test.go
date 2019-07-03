package main


import(
	"fmt"
)

type notifier interface {
	notify()
}
type changer interface {
	change(string)
}

type user struct{
	name string
	email string
}

type admin struct{
	name user
	id int
}

func (u *user)notify(){
	fmt.Printf("Send User email to %s<%s>\n",
		u.name,u.email)
}

func (u* admin)notify(){
	fmt.Printf("Send Admin email to %s<%s>\n",
		u.name.name,u.name.email)
}

func (u *user)change(email string){
	u.email = email
}

func (u *admin)change(email string){
	u.name.email = email
}

func sendNotification(n notifier){
	n.notify()
}

func sendChange(n changer,email string){
	n.change(email)
}

func main(){

	bill:=user{"Bill","Bill@163.com"}
	lisa:=admin{user{"Lisa","lisa@qq.com"},1}
	sendNotification(&bill)
	sendNotification(&lisa)

	sendChange(&bill,"newbill")
	sendNotification(&bill)

	sendChange(&lisa,"newlisa")
	sendNotification(&lisa)

	//bill.change("Bill@new.com")
	//bill.notify()
	//lisa.change("Lisa@new.com")
	//lisa.notify()
}
