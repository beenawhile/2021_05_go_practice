package myapp

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	assert := assert.New(t)

	// http 서버를 모의한 서버 생성
	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	// 첫번째 테스트 : 페이지 잘 접속하는지
	res, err := http.Get(ts.URL)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	// 두번째 테스트 : 결과가 예상된 것이 맞는지

	data, _ := ioutil.ReadAll(res.Body)
	assert.Equal("hello, world", string(data))

}

func TestUsers(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/users")
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "Get UserInfo")

}
func TestUserInfo(t *testing.T) {

	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/users/89")
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "No User Id:89")

}

func TestCreateUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	res, err := http.Post(ts.URL+"/users", "application/json", strings.NewReader(`{"first_name":"tucker", "last_name":"kim","email":"tucker@naver.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)

	user := new(User)
	err = json.NewDecoder(res.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	id := user.ID
	// 다시 get 한번해봄 => 실제 정보 가져오기
	res, err = http.Get(ts.URL + "/users/" + strconv.Itoa(id))
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	user2 := new(User)
	err = json.NewDecoder(res.Body).Decode(user2)

	assert.NoError(err)
	assert.Equal(user.ID, user2.ID)
	assert.Equal(user.FirstName, user2.FirstName)

}

func TestDeleteUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	// delete는 기본적인 웹브라우저가 제공하는 메소드가 아니라 기본적으로 제공 안함
	req, _ := http.NewRequest("DELETE", ts.URL+"/users/1", nil)
	res, err := http.DefaultClient.Do(req)

	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "No User Id:1")

	res, err = http.Post(ts.URL+"/users", "application/json", strings.NewReader(`{"first_name":"tucker", "last_name":"kim","email":"tucker@naver.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)

	user := new(User)
	err = json.NewDecoder(res.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	req, _ = http.NewRequest("DELETE", ts.URL+"/users/1", nil)
	res, err = http.DefaultClient.Do(req)

	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)
	data, _ = ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "Deleted User Id:1")

}

func TestUpdateUser(t *testing.T) {
	assert := assert.New(t)

	ts := httptest.NewServer(NewHandler())
	defer ts.Close()

	req, _ := http.NewRequest("PUT", ts.URL+"/users", strings.NewReader(`{"id":1,"first_name":"updated","last_name":"updated","email":"updated@naver.com"}`))
	res, err := http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	// update and create 방식을 .쓸지 update or error 방식을 쓸지 고려해서 작성해야함
	data, _ := ioutil.ReadAll(res.Body)
	assert.Contains(string(data), "No User Id:1")

	// create 해서 확인해보기
	res, err = http.Post(ts.URL+"/users", "application/json", strings.NewReader(`{"first_name":"tucker", "last_name":"kim","email":"tucker@naver.com"}`))
	assert.NoError(err)
	assert.Equal(http.StatusCreated, res.StatusCode)

	user := new(User)
	err = json.NewDecoder(res.Body).Decode(user)
	assert.NoError(err)
	assert.NotEqual(0, user.ID)

	// 문자열 동적으로 만들기 위해 따로 뺌
	updateStr := fmt.Sprintf(`{"id":%d,"first_name":"updated"}`, user.ID)

	req, _ = http.NewRequest("PUT", ts.URL+"/users", strings.NewReader(updateStr))
	res, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, res.StatusCode)

	updateUser := new(User)
	err = json.NewDecoder(res.Body).Decode(updateUser)
	assert.NoError(err)
	assert.Equal(updateUser.ID, user.ID)
	assert.Equal("updated", updateUser.FirstName)
	assert.Equal("updated", updateUser.LastName)
	assert.Equal("updated@naver.com", updateUser.Email)

}
