package main

import (
	"LemmyPiefedApi/config"
	"LemmyPiefedApi/dto/response/piefed"
	"LemmyPiefedApi/helper"
	appHttp "LemmyPiefedApi/http"
	"LemmyPiefedApi/router"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	gracefulShutdown := make(chan os.Signal, 1)
	signal.Notify(gracefulShutdown, syscall.SIGTERM, syscall.SIGINT)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if config.CorsRegex != nil && config.CorsRegex.MatchString(request.Header.Get("Origin")) {
			writer.Header().Add("Access-Control-Allow-Origin", request.Header.Get("Origin"))
			writer.Header().Add("Access-Control-Allow-Headers", "Content-Type, User-Agent, Authorization")
		}

		if request.Method == "OPTIONS" {
			appHttp.WriteHttpResponse(appHttp.NoContent(), writer)
			return
		}

		body, err := io.ReadAll(request.Body)
		if err != nil {
			appHttp.WriteHttpResponse(appHttp.InternalProxyError(), writer)
			log.Println(err)
			return
		}
		defer request.Body.Close()

		appRequest := appHttp.NewRequest(
			body,
			request.Method,
			helper.SingularizeMapValue(request.Header),
		)

		query := request.URL.Query()
		if query != nil {
			appRequest.QueryParams = helper.SingularizeMapValue(query)
		}

		var result *appHttp.Response
		for _, route := range config.AppRouter.Routes {
			var httpMethod router.HttpMethod
			httpMethod, err = router.HttpMethodFromString(request.Method)
			if err != nil {
				break
			}

			matches, params, errRoute := router.RouteMatches(route, httpMethod, request.URL.Path)
			if errRoute != nil {
				err = errRoute
				break
			}
			if !matches {
				continue
			}

			appRequest.RouteParams = params
			result, err = route.ControllerMethod(appRequest)
			break
		}

		var responseError *piefed.ErrorResponse
		if errors.As(err, &responseError) {
			appHttp.WriteHttpResponse(&appHttp.Response{
				StatusCode: responseError.StatusCode,
				Body:       helper.ConvertPiefedErrorToLemmyError(responseError),
			}, writer)
			return
		}

		if err != nil {
			appHttp.WriteHttpResponse(appHttp.InternalProxyError(), writer)
			log.Println(err)
			return
		}

		if result == nil {
			appHttp.WriteHttpResponse(appHttp.NotFoundProxyError(), writer)
			return
		}

		appHttp.WriteHttpResponse(result, writer)
	})

	go func() {
		var err error
		serveAddr := fmt.Sprintf(":%d", config.AppHttpPort)
		if config.CertificateFile != "" && config.CertificateKey != "" {
			err = http.ListenAndServeTLS(serveAddr, config.CertificateFile, config.CertificateKey, nil)
		} else {
			err = http.ListenAndServe(serveAddr, nil)
		}
		if err != nil {
			panic(err)
		}
	}()

	log.Println("Server started...")
	<-gracefulShutdown
	log.Println("Server shutting down...")
}
