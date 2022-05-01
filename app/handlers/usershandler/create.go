package usershandler

import (
	"net/http"

	"github.com/daison12006013/lucid/app/models/users"
	"github.com/daison12006013/lucid/app/validations"
	"github.com/daison12006013/lucid/pkg/engines"
	"github.com/daison12006013/lucid/pkg/errors"
	"github.com/daison12006013/lucid/pkg/facade/session"
	"github.com/gorilla/csrf"
)

func Create(T engines.EngineContract) *errors.AppError {
	engine := T.(engines.MuxEngine)
	w := engine.ResponseWriter
	r := engine.HttpRequest
	ses := session.File(w, r)
	req := engine.Request
	res := engine.Response

	data := map[string]interface{}{
		"title":          "Create Form",
		"record":         &users.Model{},
		"isCreate":       true,
		csrf.TemplateTag: csrf.TemplateField(r),

		//> to retrieve the flashes from Store() or somewhere
		//> that redirects back to Create()
		"success": ses.GetFlash("success"),
		"error":   ses.GetFlash("error"),
		"fails":   ses.GetFlashMap("fails"),
		"inputs":  ses.GetFlashMap("inputs"),
	}

	if req.IsJson() && req.WantsJson() {
		return res.Json(data, http.StatusOK)
	}

	return res.View(
		[]string{"base", "users/show"},
		data,
	)
}

func Store(T engines.EngineContract) *errors.AppError {
	engine := T.(engines.MuxEngine)
	w := engine.ResponseWriter
	r := engine.HttpRequest
	ses := session.File(w, r)
	req := engine.Request
	res := engine.Response
	url := engine.Url

	//> validate the inputs
	validator := req.Validator(validations.UserValidateCreate())
	if validator != nil {
		if req.IsJson() && req.WantsJson() {
			return res.Json(map[string]interface{}{
				"fails": validator.ValidationError,
			}, http.StatusUnauthorized)
		}

		ses.SetFlashMap("fails", validator.ValidationError)
		ses.SetFlashMap("inputs", req.All())
		url.RedirectPrevious()
		return nil
	}

	//> prepare message and status
	message := "Successfully Created!"
	status := http.StatusOK

	//> create user based on all request inputs
	data, appErr := users.Create(req.All())
	if appErr != nil {
		return appErr
	}

	//> for api based
	if req.IsJson() && req.WantsJson() {
		return res.Json(map[string]interface{}{
			"success": message,
			"data":    data,
		}, status)
	}

	//> for form based, just redirect
	ses.SetFlash("success", message)
	url.RedirectPrevious()
	return nil
}
