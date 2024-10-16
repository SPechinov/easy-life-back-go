package http_routes_auth

import (
	"easy-life-back-go/internal/server/deliveries/http_routes/constants"
	"easy-life-back-go/internal/server/deliveries/http_routes/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Controller struct {
}

func newController() *Controller {
	httpAuthController := &Controller{}
	return httpAuthController
}

func (c *Controller) login(w http.ResponseWriter, r *http.Request) {
	var loginForm LoginForm

	if err := utils.BodyParser(r, &loginForm, true); err != nil {
		errData := err.Error()

		w.WriteHeader(http.StatusBadRequest)

		log.Println(errData.ServerError.Path, errData.ServerError.Message)

		if err := json.NewEncoder(w).Encode(errData.ClientError); err != nil {
			http.Error(w, http_errors_codes.SomethingHappen, http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (c *Controller) registration(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("registration")
}

func (c *Controller) registrationSuccess(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("registrationSuccess")
}

func (c *Controller) forgotPassword(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("forgotPassword")
}

func (c *Controller) forgotPasswordSuccess(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("forgotPasswordSuccess")
}
