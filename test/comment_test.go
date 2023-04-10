// +build e2e


package test

import(
     _"fmt"
	"testing"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)


const(
	BASE_URL= "http://localhost:8080"
)


func TestGetCommnets(t *testing.T){
	client := resty.New()
	resp, err := client.R().Get(BASE_URL + "/api/comment")
	if err != nil{
		t.Fail()
	}
	assert.Equal(t, 200, resp.StatusCode())
}

func TestPostCommnets(t *testing.T){
	client := resty.New()
	resp, err := client.R().
     SetBody(`{"slug":"/", "author":"12345","body":"hello world"}`).
	Post(BASE_URL + "/api/comment")
	if err != nil{
		t.Fail()
	}
	assert.NoError(t, err);
	assert.Equal(t, 200, resp.StatusCode())
}