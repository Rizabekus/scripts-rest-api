package validators

import (
	"regexp"

	"github.com/go-playground/validator"
)

func IsValidBashScript(fl validator.FieldLevel) bool {
	script := fl.Field().String()

	bashScriptPattern := `^#!/bin/bash\n([\s\S]*)$`

	re := regexp.MustCompile(bashScriptPattern)

	return re.MatchString(script)
}
