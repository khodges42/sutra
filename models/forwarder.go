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

type forwarder struct {
	uri string
	port string
	rootHash string //this wont be a string later
	client string
}


