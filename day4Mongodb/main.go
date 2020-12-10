package main

import (
	mgo "gopkg.in/mgo.v2"
)

const (
	CONN_PORT = "8080"
	CONN_HOST = "localhost"
)

var session *mgo.Session

func init() {
	session, connectionError := mgo.Dial(mongoDBURL)
}
func main() {

}
