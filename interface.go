package main

import (
	"fmt"
)

type empty interface {
}

type USB interface {
	Name() string
	Connector
}
type Connector interface {
	Connect()
}

type PhoneConnector struct {
	// implement USB interface
	name string
}

func (pc PhoneConnector) Name() string {
	return pc.name
}

func (pc PhoneConnector) Connect() {
	fmt.Println("connect: ", pc.name)
}

type TVConnector struct { // designed to implement only Connect(), but not Name()
	name string
}

func (tv TVConnector) Connect() {
	fmt.Println("connect: ", tv.name)
}

func Disconnect(usb interface{}) {
	// if pc, ok := usb.(PhoneConnector); ok {
	// 	fmt.Println("Disconnected: ", pc.name)
	// 	return
	// }
	// fmt.Println("unknown device")
	switch v := usb.(type) {
	case PhoneConnector:
		fmt.Println("Disconnected: ", v.name)
	default:
		fmt.Println("unknown device")
	}

}
func main() {
	// var a USB //interface type
	a := PhoneConnector{"PhoneConnectorA"} // PhoneConnector implements both Name and Connect func
	a.Connect()
	Disconnect(a)
	fmt.Println(a.Name())

	pc := PhoneConnector{"PhoneConnectorB"}
	var b Connector
	b = Connector(pc) // type convert
	b.Connect()
	Disconnect(b)
	pc.name = "changed name" // pc copied to b, but a value copy
	b.Connect()
	// b.Name() // not allowed, since b was converted into Connector, even tho it was PhoneConnector type

	// tv := TVConnector{"TVConnectorC"}
	// var c USB
	// c = USB(tv) // not allowed, since TVConnector did not impelemtn Name

	var d interface{}
	fmt.Println(d == nil)

	var p *int = nil
	d = p
	fmt.Println(d == nil) // only when both type and stored object were nil, then the interface equal to nil
}
