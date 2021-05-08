package main

import "os"

var (
	mongoURL          = "mongodb://%s:%s@%s:27017/?authSource=admin"
	mongoIP           = os.Getenv("MONGO_IP")
	mongoUser         = os.Getenv("ME_CONFIG_MONGODB_ADMINUSERNAME")
	mongoPass         = os.Getenv("ME_CONFIG_MONGODB_ADMINPASSWORD")
	ginPortYarbDB     = os.Getenv("GIN_PORT_YARB_DB")
	yarbBasicAuthUser = os.Getenv("YARB_BASIC_AUTH_USER")
	yarbBasicAuthPass = os.Getenv("YARB_BASIC_AUTH_PASS")
)
