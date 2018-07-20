package server

import (
	"fmt"
	"net/http"
	"github.com/pkg/errors"
	"github.com/eug48/fhir/models"
	"github.com/eug48/fhir/models2"
	"github.com/eug48/fhir/search"
)

func ErrorToOpOutcome(err interface{}) (statusCode int, outcome *models.OperationOutcome) {
	switch x := err.(type) {
	case *search.Error:
		return x.HTTPStatus, x.OperationOutcome
	case error:
		cause := errors.Cause(x)
		_, isSchemaError := cause.(models2.FhirSchemaError)
		_, isVersionConflict := cause.(ErrConflict)
		if isSchemaError {
			outcome := models.NewOperationOutcome("fatal", "structure", cause.Error())
			return http.StatusBadRequest, outcome
		} else if isVersionConflict {
			outcome := models.NewOperationOutcome("error", "conflict", cause.Error())
			return http.StatusConflict, outcome // TODO (FHIR R4): changed to 412
		} else {
			// fmt.Printf("handlePanics: recovered %+v\n", r)
			// runtime_debug.PrintStack()

			outcome := models.NewOperationOutcome("fatal", "exception", x.Error())
			return http.StatusInternalServerError, outcome
		}
	default:
		// fmt.Printf("handlePanics: recovered %+v\n", r)
		// runtime_debug.PrintStack()

		str := fmt.Sprintf("%#v", err)
		outcome := models.NewOperationOutcome("fatal", "exception", str)
		return http.StatusInternalServerError, outcome
	}
}
