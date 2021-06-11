package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"testing"
	"todos_practice/model"

	"github.com/stretchr/testify/assert"
)

func TestTodos(t *testing.T) {

	err := os.Remove("./test.db")
	if err == nil {
		log.Print("deleted")
	} else {
		log.Print(err)
	}

	assert := assert.New(t)

	ah := MakeHandler()
	defer ah.Close()

	ts := httptest.NewServer(ah)
	defer ts.Close()

	// 1.add test
	resp, err := http.PostForm(ts.URL+"/todos", url.Values{"name": {"Test Todo1"}})
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	var todo model.Todo
	err = json.NewDecoder(resp.Body).Decode(&todo)
	assert.NoError(err)
	assert.Equal("Test Todo1", todo.Name)

	id1 := todo.ID

	resp, err = http.PostForm(ts.URL+"/todos", url.Values{"name": {"Test Todo2"}})
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)

	err = json.NewDecoder(resp.Body).Decode(&todo)
	assert.NoError(err)
	assert.Equal("Test Todo2", todo.Name)

	id2 := todo.ID

	// 2. get todos list test
	resp, err = http.Get(ts.URL + "/todos")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	todos := []*model.Todo{}
	err = json.NewDecoder(resp.Body).Decode(&todos)
	assert.NoError(err)
	assert.Equal(2, len(todos))

	for _, v := range todos {
		if v.ID == id1 {
			assert.Equal("Test Todo1", v.Name)
		} else if v.ID == id2 {
			assert.Equal("Test Todo2", v.Name)
		} else {
			assert.Error(fmt.Errorf("testId should be id1 or id2"))
		}
	}

	// 3. complete
	resp, err = http.Get(ts.URL + "/complete-todo/" + strconv.Itoa(id1) + "?complete=true")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	// complete result test
	resp, err = http.Get(ts.URL + "/todos")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	todos = []*model.Todo{}
	err = json.NewDecoder(resp.Body).Decode(&todos)
	assert.NoError(err)
	assert.Equal(2, len(todos))

	for _, v := range todos {
		if v.ID == id1 {
			assert.True(v.Completed)
		} else {
			assert.Error(fmt.Errorf("testId1 should be completed!"))
		}
	}

	// delete
	req, _ := http.NewRequest("DELETE", ts.URL+"/todos/"+strconv.Itoa(id1), nil)
	resp, err = http.DefaultClient.Do(req)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	resp, err = http.Get(ts.URL + "/todos")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)

	todos = []*model.Todo{}
	err = json.NewDecoder(resp.Body).Decode(&todos)
	assert.NoError(err)
	assert.Equal(1, len(todos))

	for _, v := range todos {
		assert.Equal(id2, v.ID)
	}

}
