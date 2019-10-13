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

	"github.com/developer-kikikaikai/githubapi/server/data"
	"github.com/developer-kikikaikai/githubapi/server/usecases"
	"github.com/gin-gonic/gin"
	"github.com/savaki/swag"
	"github.com/savaki/swag/endpoint"
	"github.com/savaki/swag/swagger"
)

func handle(c *gin.Context) {
	code := c.DefaultQuery("code", "")
	if token, err := usecases.GenerateToken(code); err != nil {
		c.String(http.StatusForbidden, err.Error())
	} else {
		c.JSON(http.StatusOK, Token{Token: token})
	}
	c.Abort()
}

type Token struct {
	Token string `json:"access_token"`
}

func main() {
	get := endpoint.New("get", "/auth", "Send ID from GitHub",
		endpoint.Handler(handle),
		endpoint.Query("code", "string", "authorization code", true),
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

	http.ListenAndServeTLS(":"+data.GetPort(), data.GetCert(), data.GetKey(), router)
}
