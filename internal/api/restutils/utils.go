package restutils

import (
	"errors"
	"fmt"
	"github.com/go-chi/render"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ent "github.com/go-playground/validator/v10/translations/en"
	"net/http"
)

func ToSuccessPayloadResponse(w http.ResponseWriter, r *http.Request, resp interface{}) {
	render.Status(r, http.StatusOK)
	render.Respond(w, r, resp)
}

func ToErrorResponse(w http.ResponseWriter, r *http.Request, err error, httpStatusCode int) {
	render.Status(r, httpStatusCode)
	errorResponse := BookStoreErrorResponse{
		Msg: err.Error(),
	}
	render.Respond(w, r, errorResponse)
}

func UnmarshalJSONRequest(req *http.Request, payload interface{}) error {
	if payload == nil {
		return errors.New("input payload is empty")
	}

	if err := render.Decode(req, payload); err != nil {
		return fmt.Errorf("error decoding request body: %w", err)
	}
	return nil
}

// ValidateJSONRequest unmarshalls the JSON request body and performs the validation
func ValidateJSONRequest(payload interface{}) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = ent.RegisterDefaultTranslations(validate, trans)
	if err := validate.Struct(payload); err != nil {
		errStr := ""
		for _, inputErr := range TranslateError(err, trans) {
			errStr += fmt.Sprintf("%s;", inputErr.Error())
		}
		return fmt.Errorf("payload format has issues: %s", errStr)
	}
	return nil
}

// TranslateError translate validation error to human-readable language
func TranslateError(err error, trans ut.Translator) (errs []error) {
	if err == nil {
		return nil
	}
	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := fmt.Errorf(e.Translate(trans))
		errs = append(errs, translatedErr)
	}
	return errs
}

// ValidateUser todo implement user validation
func ValidateUser(r *http.Request) (string, error) {
	userId := r.URL.Query().Get("userId")
	return userId, nil
}
