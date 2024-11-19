package tests

import "testing"

// import (
//
//	"encoding/json"
//	"fmt"
//	"github.com/labstack/echo/v4"
//	"github.com/stretchr/testify/assert"
//	"go-clean/config"
//	"go-clean/internal/api/rest"
//	"go-clean/internal/api/rest/constants"
//	"go-clean/internal/api/rest/routes/auth"
//	"go-clean/internal/api/rest/routes/auth/mocks"
//	"go-clean/internal/api/rest/utils/rest_error"
//	globalConstants "go-clean/internal/constants"
//	"go-clean/pkg/helpers"
//	"go.uber.org/mock/gomock"
//	"net/http"
//	"net/http/httptest"
//	"strings"
//	"testing"
//	"time"
//
// )
func TestHandler_updateJWT(t *testing.T) {

}

//func TestHandler_updateJWT(t *testing.T) {
//	type mockBehavior func(s *mock_auth.MockuseCases, jwt, expectedSessionID, expectedAccessJWT, expectedRefreshJWT string)
//	testTable := []struct {
//		name               string
//		mockBehavior       mockBehavior
//		expectedStatusCode int
//		sessionID          string
//		refreshJWT         string
//		expectedResponse   string
//		expectedSessionID  string
//		expectedAccessJWT  string
//		expectedRefreshJWT string
//		expectedError      *rest.ResponseBad
//	}{
//		{
//			name: "OK",
//			mockBehavior: func(s *mock_auth.MockuseCases, jwt, expectedSessionID, expectedAccessJWT, expectedRefreshJWT string) {
//				s.EXPECT().UpdateJWT(gomock.Any(), "", "", jwt).Return(expectedSessionID, expectedAccessJWT, expectedRefreshJWT, nil)
//			},
//			expectedStatusCode: http.StatusNoContent,
//			sessionID:          "47689d53-4fb5-4863-87f9-4cd895463e62",
//			refreshJWT: func() string {
//				jwt, jwtErr := helpers.CreateJWT("test_jwt_secret_kay", time.Minute, map[string]string{
//					globalConstants.UserIDInJWTKey: "1231231231",
//				})
//				if jwtErr != nil {
//					return ""
//				}
//
//				return jwt
//			}(),
//			expectedSessionID:  "session_id",
//			expectedAccessJWT:  "new_access_jwt",
//			expectedRefreshJWT: "new_refresh_jwt",
//		},
//		{
//			name: "Unauthorized",
//			mockBehavior: func(s *mock_auth.MockuseCases, jwt, expectedSessionID, expectedAccessJWT, expectedRefreshJWT string) {
//			},
//			expectedStatusCode: http.StatusUnauthorized,
//			expectedError:      rest.NewResponseBad(rest_error.ErrNotAuthorized.Code),
//		},
//	}
//
//	for _, testCase := range testTable {
//		t.Run(testCase.name, func(t *testing.T) {
//			c := gomock.NewController(t)
//			defer c.Finish()
//
//			// Create mock
//			mockUseCases := mock_auth.NewMockuseCases(c)
//			testCase.mockBehavior(mockUseCases, testCase.refreshJWT, testCase.expectedSessionID, testCase.expectedAccessJWT, testCase.expectedRefreshJWT)
//
//			// Test server
//			server, group := StartTestServer()
//			cfg := new(config.Config)
//			cfg.HTTPAuth.JWTSecretKey = "test_jwt_secret_kay"
//			authRoutes := auth.New(cfg, mockUseCases)
//			authRoutes.Register(group)
//
//			// Test request
//			w := httptest.NewRecorder()
//			req := httptest.NewRequest(http.MethodPost, "/auth/update-jwt", strings.NewReader(""))
//			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
//
//			// Funcs
//			assertRefreshJWTInResponse := func(expect bool) {
//				cookies := w.Result().Cookies()
//				cookiesFound := false
//				for _, cookie := range cookies {
//					if cookie.Name == constants.CookieRefreshJWT && cookie.Value == testCase.expectedRefreshJWT {
//						cookiesFound = true
//						break
//					}
//				}
//
//				assert.Equal(t, expect, cookiesFound)
//			}
//
//			if testCase.refreshJWT != "" {
//				req.AddCookie(&http.Cookie{
//					Name:     constants.CookieRefreshJWT,
//					Value:    testCase.refreshJWT,
//					Path:     "/",
//					MaxAge:   int(globalConstants.RestAuthRefreshWTDuration / time.Second),
//					Secure:   true,
//					HttpOnly: true,
//				})
//				req.AddCookie(&http.Cookie{
//					Name:     constants.CookieSessionID,
//					Value:    testCase.sessionID,
//					Path:     "/",
//					MaxAge:   int(globalConstants.RestAuthRefreshWTDuration / time.Second),
//					Secure:   true,
//					HttpOnly: true,
//				})
//
//				fmt.Println(11)
//				// Perform request
//				server.ServeHTTP(w, req)
//
//				// Assert
//				assert.Equal(t, testCase.expectedStatusCode, w.Code)
//				assert.Equal(t, testCase.expectedAccessJWT, w.Result().Header.Get(constants.HeaderResponseAccessJWT))
//				assertRefreshJWTInResponse(true)
//			} else {
//				// Perform request
//				server.ServeHTTP(w, req)
//
//				// Assert
//				assert.Equal(t, testCase.expectedStatusCode, w.Code)
//				assert.Equal(t, "", w.Result().Header.Get(constants.HeaderResponseAccessJWT))
//				assertRefreshJWTInResponse(false)
//
//				var responseBad *rest.ResponseBad
//				err := json.Unmarshal(w.Body.Bytes(), &responseBad)
//				assert.NoError(t, err)
//				assert.Equal(t, testCase.expectedError.Error(), responseBad.Error())
//			}
//		})
//	}
//}
