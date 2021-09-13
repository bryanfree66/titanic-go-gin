package mocks

import (
	"context"
	"github.com/bryanfree66/titanic-go-gin/app/model"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MockPassengerService is a mock type for model.PassengerService
type MockPassengerService struct {
	mock.Mock
}

// Get is mock of UserService Get
func (m *MockPassengerService) Get(ctx context.Context, uuid primitive.ObjectID) (*model.Passenger, error) {
	// args that will be passed to "Return" in the tests
	ret := m.Called(ctx, uuid)

	// first value passed to "Return"
	var r0 *model.Passenger
	if ret.Get(0) != nil {
		// we can just return this if we know we won't be passing function to "Return"
		r0 = ret.Get(0).(*model.Passenger)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}
