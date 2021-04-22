package httpapi

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserDB struct {
	mock.Mock
}

func (m MockUserDB) FindByID(ID string) (string, error) {
	args := m.Called()
	return args.Get(0).(string), args.Error(1)
}

func TestGetUser(t *testing.T) {
	t.Run("can get user", func(t *testing.T) {
		userDBMock := new(MockUserDB)
		userDBMock.On("FindByID", mock.Anything).Return("Saya", nil)

		req, err := http.NewRequest(http.MethodGet, "/users", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := GetUser(userDBMock)

		handler.ServeHTTP(rr, req)

		assert.Equal(t, http.StatusOK, rr.Code, "handler returning wrong status code : got %v want %v", rr.Code, http.StatusOK)

		expected := `{"Name" : "Saya"}`
		assert.JSONEq(t, expected, rr.Body.String(), "handler returned unexpected body : got %v want %v", rr.Body.String(), expected)
	})
}
