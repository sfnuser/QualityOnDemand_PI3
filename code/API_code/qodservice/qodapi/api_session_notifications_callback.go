// Copyright 2023 Spry Fox Networks
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
 * QoD for enhanced communication
 *
 * Service Enabling Network Function API for QoS control
 *
 * API version: 0.8.0
 * Contact: project-email@sample.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package qodapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PostNotification - Session notifications callback
func PostNotification(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{})
}