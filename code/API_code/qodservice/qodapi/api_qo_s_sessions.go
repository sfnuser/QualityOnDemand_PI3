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
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sfnuser/camara/qodmodels/api"
	"github.com/sfnuser/qodservice/logger"
	"github.com/sfnuser/qodservice/producer"
	"github.com/sfnuser/qodservice/util"
	"go.uber.org/zap"
)

// CreateSession - Creates a new session
func CreateSession(c *gin.Context) {
	requestBody, err := c.GetRawData()
	if err != nil {
		logger.Api.Sugar().Errorf("failed to get request body: %v", err)
		data := util.NewQoDErrorInfo("INTERNAL", "Session could not be created")
		c.Data(http.StatusInternalServerError, CONTENT_TYPE_DATA, data)
		return
	}
	var sessionReq api.CreateSession
	err = json.Unmarshal(requestBody, &sessionReq)
	if err != nil {
		logger.Api.Sugar().Errorf("failed to unmarshal request: %v", err)
		data := util.NewQoDErrorInfo("INVALID_INPUT", "Schema validation failed")
		c.Data(http.StatusBadRequest, CONTENT_TYPE_DATA, data)
		return
	}
	err = util.ValidateSessionReq(&sessionReq)
	if err != nil {
		data := util.NewQoDErrorInfo("INVALID_INPUT", err.Error())
		c.Data(http.StatusBadRequest, CONTENT_TYPE_DATA, data)
		return
	}

	// Log the JSON request
	reqBodyStr, err := json.MarshalIndent(&sessionReq, "", "  ")
	if err == nil {
		// Print only in non-error cases
		logger.Api.Sugar().Debugf("CreateSession: Req: JSON(createSession): %s", reqBodyStr)
	}

	// Handle the Create Session request
	rsp := producer.HandleCreateSessionRequest(&util.CreateSessionReq{SessionReq: &sessionReq})
	var contentType string
	var rspBody []byte
	var statusCode int
	if rsp.ErrorInfo != nil {
		contentType = CONTENT_TYPE_DATA
		statusCode = util.ConvertErrorToHttpStatusCode(rsp.ErrorInfo.Code)
		rspBody, err = json.Marshal(rsp.ErrorInfo)
		if err != nil {
			logger.Api.Sugar().Errorf("failed to encode error info. err %v", err)
		}
		logger.Api.Sugar().Errorf("CreateSession: failed. errorInfo %v", rsp.ErrorInfo)
	} else {
		// Success case
		contentType = CONTENT_TYPE_DATA
		statusCode = http.StatusCreated
		rspBody, err = json.Marshal(rsp.SessionInfo)
		if err != nil {
			logger.Api.Sugar().Errorf("failed to encode error info. err %v", err)
			statusCode = http.StatusInternalServerError
			contentType = CONTENT_TYPE_DATA
		}

		// Log the JSON response
		respBodyStr, err := json.MarshalIndent(rsp.SessionInfo, "", "  ")
		if err == nil {
			// Print only in non-error cases
			logger.Api.Sugar().Debugf("CreateSession: Resp: JSON(sessionInfo): %s", respBodyStr)
		}
	}
	c.Data(statusCode, contentType, rspBody)
}

// DeleteSession - Free resources related to QoS session
func DeleteSession(c *gin.Context) {
	sessionId := c.Params.ByName("sessionId")
	logger.Api.Info("Delete Session", zap.String("sessionId", sessionId))

	// Handle the Create Session request
	rsp := producer.HandleDeleteSessionRequest(&util.DeleteSessionReq{SessionId: sessionId})
	if rsp.ErrorInfo != nil {
		contentType := CONTENT_TYPE_DATA
		statusCode := util.ConvertErrorToHttpStatusCode(rsp.ErrorInfo.Code)
		rspBody, err := json.Marshal(rsp.ErrorInfo)
		if err != nil {
			logger.Api.Sugar().Errorf("failed to encode error info. err %v, statusCode %v", err, statusCode)
		}
		logger.Api.Sugar().Errorf("DeleteSession: failed. errorInfo %v", rsp.ErrorInfo)
		c.Data(statusCode, contentType, rspBody)
		return
	}
	c.JSON(http.StatusNoContent, gin.H{})
	return
}

// GetSession - Get session information
func GetSession(c *gin.Context) {
	logger.Api.Sugar().Warnf("@todo: not implemented yet")
	c.JSON(http.StatusOK, gin.H{})
}
