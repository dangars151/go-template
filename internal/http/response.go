package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	net_http "net/http"
	"strings"
)

type ResponseData struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrType int

const (
	InternalServerErr ErrType = 1
	ValidationErr     ErrType = 2
	PostgresErr       ErrType = 3
)

const (
	ForeignKeyViolation = "23503"
	UniqueViolation     = "23505"
	CheckViolation      = "23514"
)

func Success(c *gin.Context, data interface{}, message ...string) {
	c.JSON(net_http.StatusOK, ResponseData{
		Status:  net_http.StatusOK,
		Message: strings.Join(message, ", "),
		Data:    data,
	})
}

// handle error
func HandleError(c *gin.Context, err error, errType ErrType, data ...interface{}) bool {
	if err != nil {
		switch errType {
		case InternalServerErr:
			errInternal(c, err, data)
		case ValidationErr:
			errValidation(c, err)
		case PostgresErr:
			switch rootError := errors.Cause(err).(type) {
			case pg.Error:
				if rootError.IntegrityViolation() {
					// Refer to https://www.postgresql.org/docs/10/protocol-error-fields.html for all error fields and detail
					errCode := rootError.Field(byte('C'))
					errTable := rootError.Field(byte('t'))
					key := rootError.Field(byte('D'))
					switch errCode {
					case UniqueViolation:
						errBadRequest(c, fmt.Sprintf("%s %s", errTable, key))
					case ForeignKeyViolation:
						errBadRequest(c, fmt.Sprintf("%s %s", errTable, strings.Split(key, " in table")[0]))
					case CheckViolation:
						errBadRequest(c, fmt.Sprintf("%s is already checked violation", errTable))
					default:
						errBadRequest(c, "Bad request")
						return true
					}
				} else {
					errInternal(c, err, data)
				}
			default:
				errInternal(c, err, data)
			}
		}
		return true
	}
	return false
}

func errInternal(c *gin.Context, err error, data ...interface{}) {
	logInternalServerError(err, data)
	c.JSON(net_http.StatusInternalServerError, ResponseData{
		Status:  net_http.StatusInternalServerError,
		Message: "An error appeared! Please try again",
	})
}

func errValidation(c *gin.Context, err error) {
	if newErr, ok := err.(validator.ValidationErrors); ok {
		errMap := make(map[string]interface{})
		for _, err := range newErr {
			fields := strings.Split(err.Namespace(), ".")
			lastField := formatField(fields[len(fields)-1])
			message := ""
			tagsMap := map[string]string{
				"eq":  "equal",
				"gt":  "greater than",
				"gte": "greater than or equal",
				"lt":  "less than",
				"lte": "less than or equal",
				"ne":  "not equat",
			}
			switch err.Tag() {
			case "required":
				message = lastField + " is required"
			case "eq", "gt", "gte", "lt", "lte", "ne":
				message = fmt.Sprintf("%s must be %s %s", lastField, tagsMap[err.Tag()], err.Param())
			default:
				message = lastField + " must be a valid " + err.Tag()
			}
			if len(fields) == 2 {
				errMap[lastField] = message
			} else {
				data := map[string]interface{}{
					lastField: message,
				}
				for i := len(fields) - 2; i > 1; i-- {
					data = map[string]interface{}{
						formatField(fields[i]): data,
					}
				}
				errMap[formatField(fields[1])] = data
			}
		}
		c.JSON(net_http.StatusBadRequest, ResponseData{
			Status:  net_http.StatusBadRequest,
			Message: errMap,
		})
		return
	}
	if strings.Contains(err.Error(), "invalid UUID") {
		c.JSON(net_http.StatusBadRequest, ResponseData{
			Status:  net_http.StatusBadRequest,
			Message: "UUID invalid",
		})
		return
	}
	if strings.Contains(err.Error(), "parsing time") {
		c.JSON(net_http.StatusBadRequest, ResponseData{
			Status:  net_http.StatusBadRequest,
			Message: "Time invalid",
		})
		return
	}
	if strings.Contains(err.Error(), "cannot unmarshal string into") {
		c.JSON(net_http.StatusBadRequest, ResponseData{
			Status:  net_http.StatusBadRequest,
			Message: "invalid json field format",
		})
		return
	}
	c.JSON(net_http.StatusBadRequest, ResponseData{
		Status:  net_http.StatusBadRequest,
		Message: err.Error(),
	})
}

func errBadRequest(c *gin.Context, message ...string) {
	c.JSON(net_http.StatusBadRequest, ResponseData{
		Status:  net_http.StatusBadRequest,
		Message: strings.Join(message, ", "),
	})
}

// helpers
func logInternalServerError(err error, data interface{}) {
	logrus.WithFields(logrus.Fields{
		"data": data,
	}).Errorf("%+v", err)
}

func formatField(field string) string {
	formattedField := strings.ReplaceAll(field, "ID", "Id")
	formattedField = strings.ToLower(formattedField[:1]) + formattedField[1:]
	return formattedField
}
