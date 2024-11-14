package tests

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go-clean/config"
	"go-clean/internal/api/rest/constants"
	"go-clean/internal/api/rest/routes/auth"
	"go-clean/internal/api/rest/routes/auth/mocks"
	"go-clean/internal/entities"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_login(t *testing.T) {
	type mockBehavior func(s *mock_auth.MockuseCases, entity entities.UserLogin, expectSessionID, expectAccessJWT, expectRefreshJWT string)

	testTable := []struct {
		name               string
		inputBody          string
		entity             entities.UserLogin
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectSessionID    string
		expectAccessJWT    string
		expectRefreshJWT   string
	}{
		{
			name:      "OK Email",
			inputBody: `{"email": "qwe@qwe.com", "password": "1234567890"}`,
			entity: entities.UserLogin{
				AuthWay:  entities.UserAuthWay{Email: "qwe@qwe.com"},
				Password: "1234567890",
			},
			mockBehavior: func(s *mock_auth.MockuseCases, entity entities.UserLogin, expectSessionID, expectAccessJWT, expectRefreshJWT string) {
				s.EXPECT().Login(gomock.Any(), gomock.Eq(entity)).Return(expectSessionID, expectAccessJWT, expectRefreshJWT, nil)
			},
			expectedStatusCode: http.StatusNoContent,
			expectSessionID:    "session_id",
			expectAccessJWT:    "access_jwt",
			expectRefreshJWT:   "refresh_jwt",
		},
		{
			name:      "OK Phone",
			inputBody: `{"phone": "+79999999999", "password": "1234567890"}`,
			entity: entities.UserLogin{
				AuthWay:  entities.UserAuthWay{Phone: "+79999999999"},
				Password: "1234567890",
			},
			mockBehavior: func(s *mock_auth.MockuseCases, entity entities.UserLogin, expectSessionID, expectAccessJWT, expectRefreshJWT string) {
				s.EXPECT().Login(gomock.Any(), gomock.Eq(entity)).Return(expectSessionID, expectAccessJWT, expectRefreshJWT, nil)
			},
			expectedStatusCode: http.StatusNoContent,
			expectSessionID:    "session_id",
			expectAccessJWT:    "access_jwt",
			expectRefreshJWT:   "refresh_jwt",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockUseCases := mock_auth.NewMockuseCases(c)
			testCase.mockBehavior(mockUseCases, testCase.entity, testCase.expectSessionID, testCase.expectAccessJWT, testCase.expectRefreshJWT)

			// Test server
			server, group := StartTestServer()
			authRoutes := auth.New(new(config.Config), mockUseCases)
			authRoutes.Register(group)

			// Test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(testCase.inputBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			// Perform request
			server.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code, "HTTP status code incorrect")
			assert.Equal(t, testCase.expectAccessJWT, w.Result().Header.Get(constants.HeaderResponseAccessJWT), "Has not got AccessJWT in response header")
			cookies := w.Result().Cookies()
			cookiesFound := false
			for _, cookie := range cookies {
				if cookie.Name == constants.CookieRefreshJWT && cookie.Value == testCase.expectRefreshJWT {
					cookiesFound = true
					break
				}
			}
			assert.Truef(t, cookiesFound, "Has not got AccessJWT in response cookie")
		})
	}
}
