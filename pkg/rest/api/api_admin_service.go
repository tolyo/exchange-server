/*
 * OPEN OUTCRY API
 *
 * # Introduction This API is documented in **OpenAPI 3.0 format**.  This API the following operations: * Retrieve a list of available instruments * Retrieve a list of executed trades  # Basics * API calls have to be secured with HTTPS. * All data has to be submitted UTF-8 encoded. * The reply is sent JSON encoded.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

import (
	"context"
	"errors"
	"net/http"
)

// AdminAPIService is a service that implements the logic for the AdminAPIServicer
// This service should implement the business logic for every endpoint for the AdminAPI API.
// Include any external packages or services that will be required by this service.
type AdminAPIService struct {
}

// NewAdminAPIService creates a default api service
func NewAdminAPIService() AdminAPIServicer {
	return &AdminAPIService{}
}

// CreateAdminPayment - Create admin payment
func (s *AdminAPIService) CreateAdminPayment(ctx context.Context) (ImplResponse, error) {
	// TODO - update CreateAdminPayment with the required logic for this service method.
	// Add api_admin_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Payment{}) or use other options such as http.Ok ...
	// return Response(200, Payment{}), nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreateAdminPayment method not implemented")
}

// GetAdminPaymentById - Get payment
func (s *AdminAPIService) GetAdminPaymentById(ctx context.Context, paymentId string) (ImplResponse, error) {
	// TODO - update GetAdminPaymentById with the required logic for this service method.
	// Add api_admin_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, Payment{}) or use other options such as http.Ok ...
	// return Response(200, Payment{}), nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetAdminPaymentById method not implemented")
}

// GetAppEntities - Get application entities
func (s *AdminAPIService) GetAppEntities(ctx context.Context) (ImplResponse, error) {
	// TODO - update GetAppEntities with the required logic for this service method.
	// Add api_admin_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, []AppEntity{}) or use other options such as http.Ok ...
	// return Response(200, []AppEntity{}), nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetAppEntities method not implemented")
}

// GetAppEntity - Get application entity
func (s *AdminAPIService) GetAppEntity(ctx context.Context, appEntityId string) (ImplResponse, error) {
	// TODO - update GetAppEntity with the required logic for this service method.
	// Add api_admin_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, AppEntity{}) or use other options such as http.Ok ...
	// return Response(200, AppEntity{}), nil

	// TODO: Uncomment the next line to return response Response(404, {}) or use other options such as http.Ok ...
	// return Response(404, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetAppEntity method not implemented")
}
