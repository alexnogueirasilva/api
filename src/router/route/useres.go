package route

import (
	"api/src/controllers"
	"net/http"
)

var routesUsers = []Route{
	{
		URI:                   "/users",
		Method:                http.MethodPost,
		Function:              controllers.CreateUser,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users",
		Method:                http.MethodGet,
		Function:              controllers.SearchUsers,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodGet,
		Function:              controllers.SearchUser,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users{userId}",
		Method:                http.MethodPut,
		Function:              controllers.UpdateUser,
		RequireAuthentication: false,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodDelete,
		Function:              controllers.DeleteUser,
		RequireAuthentication: false,
	},
}
