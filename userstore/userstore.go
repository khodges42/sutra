package userstore

import "sutra/models"

/*
{
   "user":{
      "userNick":"eve",
      "userHash":"foo",
      "trustedForwarders":["bar"],
      "publicKey":"baz"
   }
}
 */

func NewUserStore() *models.Users {
	eve := &models.User{
		UserNick: "eve",
		UserHash: "000",
		TrustedForwarders: "",
		PublicKey: "hunter2",
	}
	Users := &models.Users{Users: []models.User{*eve}}
	return Users
}