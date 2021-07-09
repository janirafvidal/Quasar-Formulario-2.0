package main

import (
	"time"

	"github.com/gopherjs/gopherjs/js"
	"github.com/jaracil/ei"
	"github.com/jaracil/psgo"
	_ "github.com/jaracil/psgo/psjs"
)

func main() {
	subscriber := psgo.NewSubscriber(msgSubscriber)
	subscriber.Subscribe("app")
	js.Global.Set("test123", "test123")

}

func msgSubscriber(msg *psgo.Msg) {
	subscribers := map[string]func(m *psgo.Msg){
		"app.formulario.get_users": getUsers,
		"app.formulario.send":      greeting,
	}
	subscribers[msg.To](msg)
}

func getUsers(msg *psgo.Msg) {
	userList := []map[string]interface{}{
		{"name": "Alex", "age": 21, "email": "alex@mail.com", "phone": 689632214, "address": "Castellón", "isAdmin": true},
		{"name": "Jan", "age": 19, "email": "janira@mail.com", "phone": 654872315, "address": "Castellón", "isAdmin": true},
		{"name": "Jorge", "age": 48, "email": "jorge@mail.com", "phone": 633254789, "address": "Valencia", "isAdmin": false},
		{"name": "María", "age": 36, "email": "maria@mail.com", "phone": 655887441, "address": "Alicante", "isAdmin": true},
		{"name": "Samantha", "age": 18, "email": "samantha@mail.com", "phone": 633225417, "address": "Madrid", "isAdmin": false},
		{"name": "Carlos", "age": 55, "email": "carlos@mail.com", "phone": 699853332, "address": "Castellón", "isAdmin": false},
	}
	time.Sleep(3 * time.Second)
	msg.Answer(userList, nil)
}

func greeting(msg *psgo.Msg) {
	nombre, err := ei.N(msg.Dat).M("nombre").String()
	if err != nil {
		msg.Answer(nil, err)
		return
	}

	apellidos, err := ei.N(msg.Dat).M("apellidos").String()
	if err != nil {
		msg.Answer(nil, err)
		return
	}

	respuesta := "Gracias por registrarte " + nombre + " " + apellidos
	msg.Answer(respuesta, nil)
}
