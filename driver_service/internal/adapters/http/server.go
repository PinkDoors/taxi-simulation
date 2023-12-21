package http

import (
	"context"
	generated "driver_service/internal/adapters/http/generate"
	"driver_service/internal/application/trip"
	trip2 "driver_service/internal/mappers/trip"
	"fmt"
)

type TripServer struct {
	tripService trip.Service
}

func NewTripServer(
	tripService trip.Service,
) *TripServer {
	return &TripServer{
		tripService: tripService,
		// userMetrics:    metrics.NewUserMetrics(),
	}
}

// GetTrips обработчик для GetTrips ручки
func (s *TripServer) GetTrips(ctx context.Context, _ generated.GetTripsRequestObject) (generated.GetTripsResponseObject, error) {
	// Логика для обработки GetTrips
	fmt.Println("Handling GetTrips request")
	// Возвращаем фиктивные данные

	foundTrips, err := s.tripService.GetTrips(ctx)
	if err != nil {
		return generated.GetTrips500Response{}, err
	}

	var targetTrips []generated.Trip

	for _, sourceTrip := range foundTrips {
		targetTrip := trip2.ToResponseTrip(sourceTrip)
		targetTrips = append(targetTrips, targetTrip)
	}

	return generated.GetTrips200JSONResponse(targetTrips), nil
}

// GetTripByID обработчик для GetTripByID ручки
func (s *TripServer) GetTripByID(ctx context.Context, request generated.GetTripByIDRequestObject) (generated.GetTripByIDResponseObject, error) {
	// Логика для обработки GetTripByID
	fmt.Println("Handling GetTripByID request")

	foundTrip, err := s.tripService.GetTrip(ctx, request.TripId)
	if err != nil {
		return generated.GetTripByID500Response{}, err
	}

	if foundTrip != nil {
		return generated.GetTripByID404Response{}, err
	}

	targetTrip := trip2.ToResponseTrip(*foundTrip)

	// Возвращаем фиктивные данные
	return generated.GetTripByID200JSONResponse(targetTrip), nil
}

// AcceptTrip обработчик для AcceptTrip ручки
func (s *TripServer) AcceptTrip(ctx context.Context, request generated.AcceptTripRequestObject) (generated.AcceptTripResponseObject, error) {
	// Логика для обработки AcceptTrip
	fmt.Println("Handling AcceptTrip request")

	tripFound, err := s.tripService.AcceptTrip(ctx, request.TripId)
	if err != nil {
		return generated.AcceptTrip500Response{}, err
	}

	if tripFound == false {
		return generated.AcceptTrip404Response{}, nil
	}

	// Возвращаем фиктивные данные
	return generated.AcceptTrip200Response{}, nil
}

// CancelTrip обработчик для CancelTrip ручки
func (s *TripServer) CancelTrip(ctx context.Context, request generated.CancelTripRequestObject) (generated.CancelTripResponseObject, error) {
	// Логика для обработки CancelTrip
	fmt.Println("Handling CancelTrip request")

	tripFound, err := s.tripService.AcceptTrip(ctx, request.TripId)
	if err != nil {
		return generated.CancelTrip500Response{}, err
	}

	if tripFound == false {
		return generated.CancelTrip404Response{}, nil
	}

	// Возвращаем фиктивные данные
	return generated.CancelTrip200Response{}, nil
}

// StartTrip обработчик для StartTrip ручки
func (s *TripServer) StartTrip(ctx context.Context, request generated.StartTripRequestObject) (generated.StartTripResponseObject, error) {
	// Логика для обработки StartTrip
	fmt.Println("Handling StartTrip request")

	tripFound, err := s.tripService.StartTrip(ctx, request.TripId)
	if err != nil {
		return generated.StartTrip500Response{}, err
	}

	if tripFound == false {
		return generated.StartTrip404Response{}, nil
	}

	// Возвращаем фиктивные данные
	return generated.StartTrip200Response{}, nil
}
