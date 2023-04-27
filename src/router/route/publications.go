package route

import (
	"api/src/controllers"
	"net/http"
)

var routePublications = []Route{
	{
		URI:                   "/publications",
		Method:                http.MethodPost,
		Function:              controllers.CreatePublication,
		RequireAuthentication: true,
	},
	{
		URI:                   "/publications",
		Method:                http.MethodGet,
		Function:              controllers.GetPublications,
		RequireAuthentication: true,
	},
	{
		URI:                   "/publications/{publicationId}",
		Method:                http.MethodGet,
		Function:              controllers.GetPublication,
		RequireAuthentication: true,
	},
	{
		URI:                   "/publications/{publicationId}",
		Method:                http.MethodPut,
		Function:              controllers.UpdatePublication,
		RequireAuthentication: true,
	},
	{
		URI:                   "/publications/{publicationId}",
		Method:                http.MethodDelete,
		Function:              controllers.DeletePublication,
		RequireAuthentication: true,
	},
	{
		URI:                   "/users/{userId}/publications",
		Method:                http.MethodGet,
		Function:              controllers.GetPublicationsByUser,
		RequireAuthentication: true,
	},
	{
		URI:                   "/publications/{publicationId}/like",
		Method:                http.MethodPost,
		Function:              controllers.LikePublication,
		RequireAuthentication: true,
	},
}
