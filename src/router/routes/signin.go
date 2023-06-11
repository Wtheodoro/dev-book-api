package routes

import (
	"api/src/controllers"
	"net/http"
)

var signingRoute = Route {
	URI: "/signin",
	Method: http.MethodPost,
	Function: controllers.Signin,
	RequiresAuthentication: false,
}