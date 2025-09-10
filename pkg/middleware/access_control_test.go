/*
 * TencentBlueKing is pleased to support the open source community by making
 * 蓝鲸智云 - Go 开发框架 (BlueKing - Go Framework) available.
 * Copyright (C) 2017 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 *	https://opensource.org/licenses/MIT
 *
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * We undertake not to change the open source license (MIT license) applicable
 * to the current version of the project delivered to anyone in the future.
 */

package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/TencentBlueKing/blueapps-go/pkg/common"
	"github.com/TencentBlueKing/blueapps-go/pkg/middleware"
)

func TestAccessControl(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("allows all users when no configure", func(t *testing.T) {
		router := gin.New()
		router.Use(middleware.AccessControl([]string{}))
		router.GET("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "allowed")
		})

		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "allowed", w.Body.String())
	})

	t.Run("allows specified user", func(t *testing.T) {
		router := gin.New()
		allowedUsers := []string{"user1"}
		router.Use(func(c *gin.Context) {
			c.Set(common.UserIDCtxKey, "user1")
		})
		router.Use(middleware.AccessControl(allowedUsers))
		router.GET("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "allowed")
		})

		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "allowed", w.Body.String())
	})

	t.Run("forbids unauthorized user", func(t *testing.T) {
		router := gin.New()
		allowedUsers := []string{"user1"}
		router.Use(func(c *gin.Context) {
			c.Set(common.UserIDCtxKey, "user2")
		})
		router.Use(middleware.AccessControl(allowedUsers))
		router.GET("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "forbidden")
		})

		req, _ := http.NewRequest(http.MethodGet, "/test", nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusForbidden, w.Code)
	})
}
