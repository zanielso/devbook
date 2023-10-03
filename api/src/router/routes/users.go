package routes

import (
	"api/src/controllers"

	"net/http"
)

var UsersRoutes = []Route{
	{
		Uri:                    "/users",
		Method:                 http.MethodPost,
		Function:               controllers.Create,
		AuthenticationRequired: false,
	},
	{
		Uri:                    "/users",
		Method:                 http.MethodGet,
		Function:               controllers.FindAll,
		AuthenticationRequired: false,
	},
	{
		Uri:                    "/users/{userId}",
		Method:                 http.MethodGet,
		Function:               controllers.GetUser,
		AuthenticationRequired: false,
	},
	{
		Uri:                    "/users/{userId}",
		Method:                 http.MethodPut,
		Function:               controllers.Update,
		AuthenticationRequired: false,
	},
	{
		Uri:                    "/users/{userId}",
		Method:                 http.MethodDelete,
		Function:               controllers.Delete,
		AuthenticationRequired: false,
	},
}
