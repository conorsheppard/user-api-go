package impl

import (
	"encoding/json"
	"fmt"
	mockdb "github.com/conorsheppard/user-api-go/internal/db/mock"
	db "github.com/conorsheppard/user-api-go/internal/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestUserServiceImpl_GetAll(t *testing.T) {
	tests := []struct {
		name             string
		usersMockDBInput []db.User
		requestParams    db.GetAllUsersParams
		expectedResult   []db.User // todo: update to a usersResponse struct
	}{
		{
			name:             "Get all with wildcard",
			usersMockDBInput: extractJSON("../../../../internal/assets/test/json/user-service/input/get-all-db-input-1.json"),
			requestParams: db.GetAllUsersParams{
				Country: "%",
				Limit:   int32(10),
				Offset:  int32(0),
			},
			expectedResult: extractJSON("../../../../internal/assets/test/json/user-service/expected/get-all-db-expected-1.json"), // todo: update to a usersResponse struct
		},
		{
			name:             "Return only IE",
			usersMockDBInput: extractJSON("../../../../internal/assets/test/json/user-service/input/get-all-db-input-2.json"),
			requestParams: db.GetAllUsersParams{
				Country: "IE",
				Limit:   int32(10),
				Offset:  int32(0),
			},
			expectedResult: extractJSON("../../../../internal/assets/test/json/user-service/expected/get-all-db-expected-2.json"), // todo: update to a usersResponse struct
		},
	}

	for i, test := range tests {
		fmt.Printf("test %d: %s\n", i, test.name)
		controller := gomock.NewController(t)
		controller.Finish() // check if all methods were called

		store := mockdb.NewMockStore(controller)

		recorder := httptest.NewRecorder()
		context, _ := gin.CreateTestContext(recorder)
		requestUrl := "/users?country=" + test.requestParams.Country
		request, err := http.NewRequest(http.MethodGet, requestUrl, nil)
		if err != nil {
			log.Fatalf("error creating request: %s\n", err.Error())
		}
		context.Request = request

		// generate stubs
		store.EXPECT().
			GetAllUsers(context, test.requestParams).
			Times(1).
			Return(test.expectedResult, nil)

		userService := NewUserService(store)
		userService.GetAll(context)
		if context.Errors.Errors() != nil {
			t.Errorf("FAIL: Error when getting all users: %s", context.Errors.Errors())
		}

		var resultUser []db.User
		err = json.Unmarshal(recorder.Body.Bytes(), &resultUser)
		if err != nil {
			log.Fatal(err)
		}

		if !reflect.DeepEqual(test.expectedResult, resultUser) {
			t.Errorf("FAIL: expected %v, got %v\n", test.expectedResult, resultUser)
		} else {
			fmt.Printf("\tPASS\n")
		}
	}
}

func getCurrentTime() time.Time {
	currentTime, err := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		log.Fatalf("unable to create current time struct: %s\n", err.Error())
	}
	return currentTime
}

func extractJSON(filename string) []db.User {
	inputBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("fatal: failed to read test file " + filename)
	}
	s := strings.ReplaceAll(string(inputBytes), "$currentTime", getCurrentTime().String())
	var resultUser []db.User
	err = json.Unmarshal([]byte(s), &resultUser)
	if err != nil {
		log.Fatal(err)
	}
	return resultUser
}
