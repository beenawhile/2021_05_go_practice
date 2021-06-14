package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"testing"
	"todos/model"

	"github.com/stretchr/testify/assert"
)

func TestTodos(t *testing.T) {

	os.Remove("./test.db")

	assert := assert.New(t)

	ah := MakeHandler()
	defer ah.Close()

	ts := httptest.NewServer(ah)
	defer ts.Close()

	// 1. add test
	resp, err := http.PostForm(ts.URL+"/todos", url.Values{"name": {"Test Todo1"}})
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	var todo *model.Todo
	err = json.NewDecoder(resp.Body).Decode(&todo)
	assert.NoError(err)
	assert.Equal("Test Todo1", todo.Name)
	assert.False(todo.Completed)

	id1 := todo.ID

	resp, err = http.PostForm(ts.URL+"/todos", url.Values{"name": {"Test Todo2"}})
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	err = json.NewDecoder(resp.Body).Decode(&todo)
	assert.NoError(err)
	assert.Equal("Test Todo2", todo.Name)
	assert.False(todo.Completed)

	id2 := todo.ID

	// 2. get list
	resp, err = http.Get(ts.URL + "/todos")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	todos := []*model.Todo{}

	err = json.NewDecoder(resp.Body).Decode(&todos)
	assert.NoError(err)
	assert.Equal(2, len(todos))

	for _, t := range todos {
		if t.ID == id1 {
			assert.Equal("Test Todo1", t.Name)
		} else if t.ID == id2 {
			assert.Equal("Test Todo2", t.Name)
		} else {
			assert.Error(fmt.Errorf("testID should be id1 or id2"))
		}
	}

	// 3. complete test
	resp, err = http.Get(ts.URL + "/complete-todo/" + strconv.Itoa(id1) + "?complete=true")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	resp, err = http.Get(ts.URL + "/todos")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	err = json.NewDecoder(resp.Body).Decode(&todos)
	assert.NoError(err)
	assert.Equal(2, len(todos))

	for _, t := range todos {
		if t.ID == id1 {
			assert.True(t.Completed)
		}
	}

	// 4. delete test
	req, err := http.NewRequest("DELETE", ts.URL+"/todos/"+strconv.Itoa(id1), nil)
	assert.NoError(err)

	resp, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	resp, err = http.Get(ts.URL + "/todos")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	err = json.NewDecoder(resp.Body).Decode(&todos)
	assert.NoError(err)
	assert.Equal(1, len(todos))

	for _, t := range todos {
		assert.Equal(id2, t.ID)
	}

}
