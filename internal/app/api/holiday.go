package api

import (
	"github.com/gotasma/internal/app/holiday"
)

func newHolidayService(policy holiday.PolicyServices) (*holiday.Service, error) {
	s, err := dialDefaultMongoDB()
	if err != nil {
		return nil, err
	}

	repo := holiday.NewMongoRepository(s)
	return holiday.New(repo, policy), nil
}

func newHolidayHandler(srv *holiday.Service) *holiday.Handler {
	return holiday.NewHanlder(srv)
}
