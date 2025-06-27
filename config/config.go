package config

import (
	"LemmyPiefedApi/router"
	"LemmyPiefedApi/service"
	piefedService "LemmyPiefedApi/service/piefed"
	"os"
	"regexp"
	"strconv"
)

var AppHttpPort = 8080
var AppRouter *router.Router
var CertificateFile string
var CertificateKey string
var CorsRegex *regexp.Regexp

var simulateLemmy bool
var piefed *piefedService.Piefed
var activityPub *service.ActivityPub

func init() {
	if port, exists := os.LookupEnv("PORT"); exists {
		parsed, err := strconv.Atoi(port)
		if err != nil {
			panic(err)
		}

		AppHttpPort = parsed
	}
	if simulate, exists := os.LookupEnv("SIMULATE"); exists {
		var err error
		simulateLemmy, err = strconv.ParseBool(simulate)
		if err != nil {
			panic(err)
		}
	}

	piefedInstance, exists := os.LookupEnv("PIEFED_INSTANCE")
	if !exists {
		panic("PIEFED_INSTANCE environment variable not set")
	}
	piefed = piefedService.NewPiefed(piefedInstance)
	activityPub = service.NewActivityPub()

	AppRouter = router.NewRouter()

	if certFile, exists := os.LookupEnv("CERTIFICATE_FILE"); exists {
		CertificateFile = certFile
	}
	if certKey, exists := os.LookupEnv("CERTIFICATE_KEY"); exists {
		CertificateKey = certKey
	}
	if regexStr, exists := os.LookupEnv("CORS_REGEX"); exists {
		CorsRegex = regexp.MustCompile(regexStr)
	}
}
