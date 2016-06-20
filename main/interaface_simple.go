package main

import "fmt"

type USB interface {
	Name() string
	Connecter
}

type Connecter interface {
	Connection()
}

type PhoneConnecter struct {
	name string
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connection() {
	fmt.Println("Connection:", pc.name)
}

func (pc *PhoneConnecter) dial(num string) {
	pc.name = num
	fmt.Println("dial:", num)
}

func main() {
	v := PhoneConnecter{name:"iphone"}
	var a Connecter
	a = Connecter(v)
	a.Connection()
	//Disconnected(v)

	v.name = "pc"
	v.Connection()
}

func Disconnected(usb interface{}) {
	// open pattern
	//if pc, ok := usb.(PhoneConnecter); ok {
	//	fmt.Println("Disconnected:", pc.name)
	//	return
	//}
	//fmt.Println("Unknow device.")

	// type switch
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnected:", v.name)
	default:
		fmt.Println("Unknow device.")
	}
}