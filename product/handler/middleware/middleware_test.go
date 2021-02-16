package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/require"
)

const (
	baseURL = "/v1/product"
)

func TestUserID(t *testing.T) {
	t.Parallel()
	t.Run("ID does not exist", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", baseURL, faker.UUIDHyphenated()), nil)
		is.Nil(err)
		rec := httptest.NewRecorder()

		fn := func(w http.ResponseWriter, r *http.Request) {
			obtainedID, _ := r.Context().Value("userID").(string)
			is.Empty(obtainedID)
		}

		s := UserID(http.HandlerFunc(fn))
		s.ServeHTTP(rec, req)

		is.Equal(http.StatusOK, rec.Code)
	})
	t.Run("ID is not valid", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		id := faker.Sentence()
		fn := func(w http.ResponseWriter, r *http.Request) {
			obtainedID, _ := r.Context().Value(UserIDKey).(string)
			is.Equal(id, obtainedID)
		}
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", baseURL, faker.UUIDHyphenated()), nil)
		is.Nil(err)
		req.Header.Set("X-USER-ID", id)
		rec := httptest.NewRecorder()

		s := UserID(http.HandlerFunc(fn))
		s.ServeHTTP(rec, req)

		is.Equal(http.StatusBadRequest, rec.Code)
		is.Contains(rec.Body.String(), errInvalidUserID.Error())
	})
	t.Run("ID is valid", func(t *testing.T) {
		t.Parallel()
		is := require.New(t)

		id := faker.UUIDHyphenated()
		fn := func(w http.ResponseWriter, r *http.Request) {
			obtainedID, _ := r.Context().Value(UserIDKey).(string)
			is.Equal(id, obtainedID)
		}
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/%s", baseURL, faker.UUIDHyphenated()), nil)
		is.Nil(err)
		req.Header.Set("X-USER-ID", id)
		rec := httptest.NewRecorder()

		s := UserID(http.HandlerFunc(fn))
		s.ServeHTTP(rec, req)
	})
}
