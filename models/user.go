package models


/*
A single message object

{
   "user":{
      "userNick":"eve",
      "userHash":"foo",
      "trustedForwarders":["bar"],
      "publicKey":"baz"
   }
}

*/
type Users struct {
	Users []User
}

type User struct {
	UserNick string
	UserHash  string
	TrustedForwarders string // a person can be their own forwarder
	PublicKey  string
}

