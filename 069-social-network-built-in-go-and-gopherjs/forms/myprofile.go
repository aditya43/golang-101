package forms

import (
	"go.isomorphicgo.org/go/isokit"
)

type MyProfileForm struct {
	isokit.BasicForm
}

func NewMyProfileForm(formParams *isokit.FormParams) *MyProfileForm {
	prefillFields := []string{"firstName", "lastName", "email", "messageBody"}
	fields := make(map[string]string)
	errors := make(map[string]string)
	m := &MyProfileForm{}
	m.SetPrefillFields(prefillFields)
	m.SetFields(fields)
	m.SetErrors(errors)
	m.SetFormParams(formParams)
	return m
}

func (m *MyProfileForm) Validate() bool {
	m.RegenerateErrors()
	m.PopulateFields()

	// Check if the about text box was filled out
	if isokit.FormValue(m.FormParams(), "aboutTextArea") == "" {
		m.SetError("aboutTextArea", "The about text box is required.")
	}

	// Check if the location field was filled out
	if isokit.FormValue(m.FormParams(), "locationInput") == "" {
		m.SetError("locationInput", "The location field is required.")
	}

	// Check if the interests field was filled out
	if isokit.FormValue(m.FormParams(), "interestsInput") == "" {
		m.SetError("interestsInput", "The interests field must be filled.")
	}

	if len(m.Errors()) > 0 {
		return false

	} else {
		return true
	}
}
