package config

import (
	"LemmyPiefedApi/controller"
	"LemmyPiefedApi/router"
)

func newRoute(path string, method router.HttpMethod, controller router.ControllerMethod) *router.Route {
	return router.NewRoute("/api/v3"+path, method, controller)
}

func init() {
	userController := controller.NewUserController(piefed)
	siteController := controller.NewSiteController(piefed, activityPub, simulateLemmy)
	postController := controller.NewPostController(piefed)

	AppRouter.AddRoute(newRoute("/user/login", router.HttpMethodPost, userController.Login))
	AppRouter.AddRoute(newRoute("/user/register", router.HttpMethodPost, userController.Register))
	AppRouter.AddRoute(newRoute("/site", router.HttpMethodGet, siteController.Site))
	AppRouter.AddRoute(newRoute("/post/list", router.HttpMethodGet, postController.GetPosts))
}
