package handler

import (
	"encoding/json"
	"github.com/bryanfree66/titanic-go-gin/app/model"
	"github.com/bryanfree66/titanic-go-gin/app/model/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPassenger(t *testing.T) {
	// Setup
	gin.SetMode(gin.TestMode)

	// Success Test
	t.Run("Success", func(t *testing.T) {
		uuidString := "613930750e0a7f5d95a5dadb"
		uuid, _ := primitive.ObjectIDFromHex(uuidString)

		mockPassengerResp := &model.Passenger{
			UUID:      uuid,
			FirstName: "Bob",
			LastName:  "Smith",
		}

		mockPassengerService := new(mocks.MockPassengerService)
		mockPassengerService.On("Get", mock.AnythingOfType("*gin.Context"), uuid).Return(mockPassengerResp, nil)

		rr := httptest.NewRecorder()
		router := gin.Default()

		NewHandler(&Config{
			R:                router,
			PassengerService: mockPassengerService,
		})

		request, err := http.NewRequest(http.MethodGet, "/"+uuidString, nil)
		assert.NoError(t, err)

		router.ServeHTTP(rr, request)

		respBody, err := json.Marshal(gin.H{
			"passenger": mockPassengerResp,
		})
		assert.NoError(t, err)

		assert.Equal(t, 200, rr.Code)
		assert.Equal(t, respBody, rr.Body.Bytes())
		mockPassengerService.AssertExpectations(t) // assert that PassengerService.Get was called
	})

	// Not Found Test
	//t.Run("NotFound", func(t *testing.T) {
	//	uuidString := "613930750e0a7f5d95a5dadb"
	//	uuid, _ := primitive.ObjectIDFromHex(uuidString)
	//	mockPassengerService := new(mocks.MockPassengerService)
	//	mockPassengerService.On("Get", mock.Anything, uuid).Return(nil, fmt.Errorf("Some error down call chain"))
	//
	//	// a response recorder for getting written http response
	//	rr := httptest.NewRecorder()
	//
	//	router := gin.Default()
	//	router.Use(func(c *gin.Context) {
	//		c.Set("passenger", &model.Passenger{
	//			UUID: uuid,
	//		},
	//		)
	//	})
	//
	//	NewHandler(&Config{
	//		R:                router,
	//		PassengerService: mockPassengerService,
	//	})
	//
	//	request, err := http.NewRequest(http.MethodGet, "/:uuid", nil)
	//	assert.NoError(t, err)
	//
	//	router.ServeHTTP(rr, request)
	//
	//	respErr := errors.NewNotFound("passenger", uuid.String())
	//
	//	respBody, err := json.Marshal(gin.H{
	//		"error": respErr,
	//	})
	//	assert.NoError(t, err)
	//
	//	assert.Equal(t, respErr.Status(), rr.Code)
	//	assert.Equal(t, respBody, rr.Body.Bytes())
	//	mockPassengerService.AssertExpectations(t) // assert that PassengerService.Get was called
	//})
}
