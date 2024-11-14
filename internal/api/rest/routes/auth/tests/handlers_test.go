package tests

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go-clean/config"
	"go-clean/internal/api/rest/routes/auth"
	"go-clean/internal/api/rest/routes/auth/mocks"
	"go-clean/internal/entities"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler_registration(t *testing.T) {
	type mockBehavior func(s *mock_auth.MockuseCases, entity entities.UserAdd)

	testTable := []struct {
		name               string
		inputBody          string
		entity             entities.UserAdd
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{
		{
			name:      "OK Email",
			inputBody: `{"email": "qwe@qwe.com"}`,
			entity: entities.UserAdd{
				AuthWay: entities.UserAuthWay{Email: "qwe@qwe.com"},
			},
			mockBehavior: func(s *mock_auth.MockuseCases, entity entities.UserAdd) {
				s.EXPECT().Registration(gomock.Any(), gomock.Eq(entity)).Return(nil)
			},
			expectedStatusCode: http.StatusNoContent,
		},
		{
			name:      "OK Phone",
			inputBody: `{"phone": "+79999999999"}`,
			entity: entities.UserAdd{
				AuthWay: entities.UserAuthWay{Phone: "+79999999999"},
			},
			mockBehavior: func(s *mock_auth.MockuseCases, entity entities.UserAdd) {
				s.EXPECT().Registration(gomock.Any(), gomock.Eq(entity)).Return(nil)
			},
			expectedStatusCode: http.StatusNoContent,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockUseCases := mock_auth.NewMockuseCases(c)
			testCase.mockBehavior(mockUseCases, testCase.entity)

			// Test server
			server, group := StartTestServer()
			authRoutes := auth.New(new(config.Config), mockUseCases)
			authRoutes.Register(group)

			// Test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/auth/registration", strings.NewReader(testCase.inputBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			// Perform request
			server.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
		})
	}
}

func TestHandler_registrationConfirm(t *testing.T) {
	type mockBehavior func(s *mock_auth.MockuseCases, dto entities.UserAddConfirm)

	testTable := []struct {
		name               string
		inputBody          string
		entity             entities.UserAddConfirm
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{
		{
			name:      "OK Email",
			inputBody: `{"email": "qwe@qwe.com", "firstName": "TestFirstName", "password": "TestPassword", "code": "123456"}`,
			entity: entities.UserAddConfirm{
				AuthWay:   entities.UserAuthWay{Email: "qwe@qwe.com"},
				FirstName: "TestFirstName",
				Password:  "TestPassword",
				Code:      "123456",
			},
			mockBehavior: func(s *mock_auth.MockuseCases, dto entities.UserAddConfirm) {
				s.EXPECT().RegistrationConfirm(gomock.Any(), gomock.Eq(dto)).Return(nil)
			},
			expectedStatusCode: http.StatusCreated,
		},
		{
			name:      "OK Phone",
			inputBody: `{"phone": "+79999999999", "firstName": "TestFirstName", "password": "TestPassword", "code": "123456"}`,
			entity: entities.UserAddConfirm{
				AuthWay:   entities.UserAuthWay{Phone: "+79999999999"},
				FirstName: "TestFirstName",
				Password:  "TestPassword",
				Code:      "123456",
			},
			mockBehavior: func(s *mock_auth.MockuseCases, dto entities.UserAddConfirm) {
				s.EXPECT().RegistrationConfirm(gomock.Any(), gomock.Eq(dto)).Return(nil)
			},
			expectedStatusCode: http.StatusCreated,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockUseCases := mock_auth.NewMockuseCases(c)
			testCase.mockBehavior(mockUseCases, testCase.entity)

			// Test server
			server, group := StartTestServer()
			authRoutes := auth.New(new(config.Config), mockUseCases)
			authRoutes.Register(group)

			// Test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/auth/registration-confirm", strings.NewReader(testCase.inputBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			// Perform request
			server.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
		})
	}
}

func TestHandler_forgotPassword(t *testing.T) {
	type mockBehavior func(s *mock_auth.MockuseCases, dto entities.UserForgotPassword)

	testTable := []struct {
		name               string
		inputBody          string
		entity             entities.UserForgotPassword
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{
		{
			name:      "OK Email",
			inputBody: `{"email": "qwe@qwe.com"}`,
			entity: entities.UserForgotPassword{
				AuthWay: entities.UserAuthWay{Email: "qwe@qwe.com"},
			},
			mockBehavior: func(s *mock_auth.MockuseCases, dto entities.UserForgotPassword) {
				s.EXPECT().ForgotPassword(gomock.Any(), gomock.Eq(dto)).Return(nil)
			},
			expectedStatusCode: http.StatusNoContent,
		},
		{
			name:      "OK Phone",
			inputBody: `{"phone": "+79999999999"}`,
			entity: entities.UserForgotPassword{
				AuthWay: entities.UserAuthWay{Phone: "+79999999999"},
			},
			mockBehavior: func(s *mock_auth.MockuseCases, dto entities.UserForgotPassword) {
				s.EXPECT().ForgotPassword(gomock.Any(), gomock.Eq(dto)).Return(nil)
			},
			expectedStatusCode: http.StatusNoContent,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockUseCases := mock_auth.NewMockuseCases(c)
			testCase.mockBehavior(mockUseCases, testCase.entity)

			// Test server
			server, group := StartTestServer()
			authRoutes := auth.New(new(config.Config), mockUseCases)
			authRoutes.Register(group)

			// Test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/auth/forgot-password", strings.NewReader(testCase.inputBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			// Perform request
			server.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
		})
	}
}

func TestHandler_forgotPasswordConfirm(t *testing.T) {
	type mockBehavior func(s *mock_auth.MockuseCases, dto entities.UserForgotPasswordConfirm)

	testTable := []struct {
		name               string
		inputBody          string
		entity             entities.UserForgotPasswordConfirm
		mockBehavior       mockBehavior
		expectedStatusCode int
	}{
		{
			name:      "OK Email",
			inputBody: `{"email": "qwe@qwe.com", "password": "TestPassword", "code": "123456"}`,
			entity: entities.UserForgotPasswordConfirm{
				AuthWay:  entities.UserAuthWay{Email: "qwe@qwe.com"},
				Password: "TestPassword",
				Code:     "123456",
			},
			mockBehavior: func(s *mock_auth.MockuseCases, dto entities.UserForgotPasswordConfirm) {
				s.EXPECT().ForgotPasswordConfirm(gomock.Any(), gomock.Eq(dto)).Return(nil)
			},
			expectedStatusCode: http.StatusNoContent,
		},
		{
			name:      "OK Phone",
			inputBody: `{"phone": "+79999999999", "password": "TestPassword", "code": "123456"}`,
			entity: entities.UserForgotPasswordConfirm{
				AuthWay:  entities.UserAuthWay{Phone: "+79999999999"},
				Password: "TestPassword",
				Code:     "123456",
			},
			mockBehavior: func(s *mock_auth.MockuseCases, dto entities.UserForgotPasswordConfirm) {
				s.EXPECT().ForgotPasswordConfirm(gomock.Any(), gomock.Eq(dto)).Return(nil)
			},
			expectedStatusCode: http.StatusNoContent,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			mockUseCases := mock_auth.NewMockuseCases(c)
			testCase.mockBehavior(mockUseCases, testCase.entity)

			// Test server
			server, group := StartTestServer()
			authRoutes := auth.New(new(config.Config), mockUseCases)
			authRoutes.Register(group)

			// Test request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/auth/forgot-password-confirm", strings.NewReader(testCase.inputBody))
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			// Perform request
			server.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
		})
	}
}
