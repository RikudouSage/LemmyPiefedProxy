package router

func RouteMatches(route *Route, httpMethod HttpMethod, path string) (bool, map[string]string, error) {
	if route.HttpMethod != httpMethod && route.HttpMethod != HttpMethodAll {
		return false, nil, nil
	}

	pathRegex, err := RegexifyRoute(route.Path)
	if err != nil {
		return false, nil, err
	}

	params := make(map[string]string)

	if pathRegex.MatchString(path) {
		matchesRealValues := pathRegex.FindStringSubmatch(path)
		matchesRoute := pathRegex.FindStringSubmatch(route.Path)

		if len(matchesRealValues) > 1 {
			for i, value := range matchesRealValues[1:] {
				paramName := matchesRoute[i+1][1 : len(matchesRoute[i+1])-1]
				params[paramName] = value
			}
		}

		return true, params, nil
	}

	return false, nil, nil
}
