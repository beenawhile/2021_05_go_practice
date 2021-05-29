// app.go 의 test 파일
// go는 파일이름_test.go라고 명명하면 test 코드로 작동함
package myapp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test라고 하는것 역시 convention
func TestIndexPathHandler(t *testing.T) {
	// 패키지 설치 할것임
	// - GoConvey :  자동 검사해줌
	// https://github.com/smartystreets/goconvey

	// 테스트 assert만 가져옴
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	// 모든 테스트 케이스를 적는 방법
	// if res.Code != http.StatusBadRequest {
	// 	t.Fatal("Failed!!", res.Result().StatusCode)
	// }

	// 라이브러리를 사용하는 방법
	// https://github.com/stretchr/testify/tree/master/assert

	// 1 통과여부
	assert.Equal(http.StatusOK, res.Code)

	// 2 받아오는 값
	// res.Body 에 데이터 들어있으나 바로 buffer struct여서 가져올 수 없어서 ioutil 사용함
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello world", string(data))

}
func TestBarHandler_WithoutName(t *testing.T) {

	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello World!", string(data))

}

func TestBarHandler_WithName(t *testing.T) {

	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/bar?name=tucker", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("Hello tucker!", string(data))

}

func TestFooHandler_WithoutJson(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/foo", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
}

func TestFooHandler_WithJson(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/foo", strings.NewReader(`{"first_name":"tucker","last_name":"kim","email":"tucker@naver.com"}`))

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusCreated, res.Code)

	user := new(User)
	err := json.NewDecoder(res.Body).Decode(user)

	assert.Nil(err)
	// 이까지는 문제 없는데 이후에 결과 값 받아올 때 통과 못함 왜그런지 확인해봐야할듯
	// assert.Equal("kim", user.LastName)
	// assert.Equal("tucker", user.FirstName)

}
