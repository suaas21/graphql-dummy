package api

import (
	"github.com/suaas21/graphql-dummy/api/response"
	"github.com/suaas21/graphql-dummy/logger"
	"github.com/suaas21/graphql-dummy/schema"
	"net/http"
)

// BookAuthorController ...
type BookAuthorController struct {
	schema schema.BookAuthor
	lgr    logger.StructLogger
}

// NewBookAuthorController ...
func NewBookAuthorController(schema schema.BookAuthor, lgr logger.StructLogger) *BookAuthorController {
	return &BookAuthorController{
		schema: schema,
		lgr:    lgr,
	}
}

func (baCtrl *BookAuthorController) BookAuthorController(res http.ResponseWriter, req *http.Request) {
	incomingReq := req.URL.Query().Get("query")
	result, err := baCtrl.schema.BookAuthorSchema(incomingReq)
	if err != nil {
		response.ServeJSON(res, nil, http.StatusInternalServerError)
		return
	}
	if result.HasErrors() {
		var errors []string
		for _, e := range result.Errors {
			errors = append(errors, e.Error())
		}
		response.ServeJSON(res, nil, http.StatusBadRequest)
		return
	}

	response.ServeJSON(res, result.Data, http.StatusOK)
	return
}
