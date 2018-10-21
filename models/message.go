package models


/*
A single message object

{
   "message":{
	  "MessageHash":"foo",
      "rootHash":"foo",
      "toHash":"bar",
      "fromHash":"baz",
      "sentTime":12345,
      "content":"fizz"
   }
}

*/

type Message struct {
	MessageHash string
	RootHash string
	ToHash  string
	FromHash string
	SentTime  int
	Content string
}

