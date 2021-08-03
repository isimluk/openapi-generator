/*
 * OpenAPI Petstore
 *
 * This is a sample server Petstore server. For this sample, you can use the api key `special-key` to test the authorization filters.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstoreserver

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// PetApiController binds http requests to an api service and writes the service results to the http response
type PetApiController struct {
	service PetApiServicer
	errorHandler ErrorHandler
}

// PetApiOption for how the controller is set up.
type PetApiOption func(*PetApiController)

// WithPetApiErrorHandler inject ErrorHandler into controller
func WithPetApiErrorHandler(h ErrorHandler) PetApiOption {
	return func(c *PetApiController) {
		c.errorHandler = h
	}
}

// NewPetApiController creates a default api controller
func NewPetApiController(s PetApiServicer, opts ...PetApiOption) Router {
	controller := &PetApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all of the api route for the PetApiController
func (c *PetApiController) Routes() Routes {
	return Routes{ 
		{
			"AddPet",
			strings.ToUpper("Post"),
			"/v2/pet",
			c.AddPet,
		},
		{
			"DeletePet",
			strings.ToUpper("Delete"),
			"/v2/pet/{petId}",
			c.DeletePet,
		},
		{
			"FindPetsByStatus",
			strings.ToUpper("Get"),
			"/v2/pet/findByStatus",
			c.FindPetsByStatus,
		},
		{
			"FindPetsByTags",
			strings.ToUpper("Get"),
			"/v2/pet/findByTags",
			c.FindPetsByTags,
		},
		{
			"GetPetById",
			strings.ToUpper("Get"),
			"/v2/pet/{petId}",
			c.GetPetById,
		},
		{
			"UpdatePet",
			strings.ToUpper("Put"),
			"/v2/pet",
			c.UpdatePet,
		},
		{
			"UpdatePetWithForm",
			strings.ToUpper("Post"),
			"/v2/pet/{petId}",
			c.UpdatePetWithForm,
		},
		{
			"UploadFile",
			strings.ToUpper("Post"),
			"/v2/pet/{petId}/uploadImage",
			c.UploadFile,
		},
	}
}

// AddPet - Add a new pet to the store
func (c *PetApiController) AddPet(w http.ResponseWriter, r *http.Request) {
	pet := &Pet{}
	if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.AddPet(r.Context(), *pet)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// DeletePet - Deletes a pet
func (c *PetApiController) DeletePet(w http.ResponseWriter, r *http.Request) {
	petId, err := parseInt64Parameter(chi.URLParam(r, "petId"), true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	apiKey := r.Header.Get("api_key")
	result, err := c.service.DeletePet(r.Context(), petId, apiKey)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// FindPetsByStatus - Finds Pets by status
func (c *PetApiController) FindPetsByStatus(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	status := strings.Split(query.Get("status"), ",")
	result, err := c.service.FindPetsByStatus(r.Context(), status)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// FindPetsByTags - Finds Pets by tags
func (c *PetApiController) FindPetsByTags(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	tags := strings.Split(query.Get("tags"), ",")
	result, err := c.service.FindPetsByTags(r.Context(), tags)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// GetPetById - Find pet by ID
func (c *PetApiController) GetPetById(w http.ResponseWriter, r *http.Request) {
	petId, err := parseInt64Parameter(chi.URLParam(r, "petId"), true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	result, err := c.service.GetPetById(r.Context(), petId)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// UpdatePet - Update an existing pet
func (c *PetApiController) UpdatePet(w http.ResponseWriter, r *http.Request) {
	pet := &Pet{}
	if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	result, err := c.service.UpdatePet(r.Context(), *pet)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// UpdatePetWithForm - Updates a pet in the store with form data
func (c *PetApiController) UpdatePetWithForm(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	petId, err := parseInt64Parameter(chi.URLParam(r, "petId"), true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

				name := r.FormValue("name")
				status := r.FormValue("status")
	result, err := c.service.UpdatePetWithForm(r.Context(), petId, name, status)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}

// UploadFile - uploads an image
func (c *PetApiController) UploadFile(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	petId, err := parseInt64Parameter(chi.URLParam(r, "petId"), true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

				additionalMetadata := r.FormValue("additionalMetadata")
	
	file, err := ReadFormFileToTempFile(r, "file")
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
			result, err := c.service.UploadFile(r.Context(), petId, additionalMetadata, file)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, result.Headers, w)

}
