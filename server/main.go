// Copyright 2019 developer-kikikaikai
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/savaki/swag"
	"github.com/savaki/swag/endpoint"
	"github.com/savaki/swag/swagger"
	"github.com/developer-kikikaikai/githubapi/server/data"
)

func handle(c *gin.Context) {
	token := Token{
		Code: c.Param("code"),
	}
	c.JSON(http.StatusOK, token)
	c.Abort()
}

type Token struct {
	Code string `json:"code"`
}

func main() {
	get := endpoint.New("get", "/auth/{code}", "Send ID from GitHub",
		endpoint.Handler(handle),
		endpoint.Path("code", "string", "ID from GitHub", true),
		endpoint.Response(http.StatusOK, Token{}, "successful operation"),
	)

	api := swag.New(
		swag.Endpoints(get),
	)

	router := gin.New()
	api.Walk(func(path string, endpoint *swagger.Endpoint) {
		h := endpoint.Handler.(func(c *gin.Context))
		path = swag.ColonPath(path)
		router.Handle(endpoint.Method, path, h)
	})

	http.ListenAndServeTLS(":"+ data.GetPort(), data.GetCert(), data.GetKey(), router)
}
