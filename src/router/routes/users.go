package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route {
	{
		URI: "/users",
		Method: http.MethodPost,
		Function: controllers.CreateUser,
		RequiresAuthentication: false,
	},
	{
		URI: "/users",
		Method: http.MethodGet,
		Function: controllers.ReadUsers,
		RequiresAuthentication: true,
	},
	{
		URI: "/users/{userId}",
		Method: http.MethodGet,
		Function: controllers.ReadUser,
		RequiresAuthentication: true,
	},
	{
		URI: "/users/{userId}",
		Method: http.MethodPut,
		Function: controllers.UpdateUser,
		RequiresAuthentication: true,
	},
	{
		URI: "/users/{userId}",
		Method: http.MethodDelete,
		Function: controllers.DeleteUser,
		RequiresAuthentication: true,
	},
	{
		URI: "/users/{userToFollowId}/follow",
		Method: http.MethodPost,
		Function: controllers.FollowUser,
		RequiresAuthentication: true,
	},
}