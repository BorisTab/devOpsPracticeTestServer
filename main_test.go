package main

import (
	"net/http/httptest"
	"strings"

	"testing"

	"github.com/dailymotion/allure-go"
)

func TestUppercase(t *testing.T) {
	allure.Test(t, allure.Action(func() {
		strService := stringService{}
		str := "hello"
		expected := "HELLO"

		actual, err := strService.Uppercase(str)
		if err != nil {
			t.Error(err)
		}

		if actual != expected {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
	}))
}

func TestCount(t *testing.T) {
	allure.Test(t, allure.Action(func() {
		strService := stringService{}
		str := "hello"
		expected := 5

		actual := strService.Count(str)

		if actual != expected {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	}))
}

func TestMakeUppercaseEndpoint(t *testing.T) {
	allure.Test(t, allure.Action(func() {
		strService := stringService{}
		endpoint := makeUppercaseEndpoint(strService)
		str := "hello"
		expected := "HELLO"

		actual, err := endpoint(nil, uppercaseRequest{str})
		if err != nil {
			t.Error(err)
		}

		if actual.(uppercaseResponse).V != expected {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
	}))
}

func TestMakeCountEndpoint(t *testing.T) {
	allure.Test(t, allure.Action(func() {
		strService := stringService{}
		endpoint := makeCountEndpoint(strService)
		str := "hello"
		expected := 5

		actual, err := endpoint(nil, countRequest{str})
		if err != nil {
			t.Error(err)
		}

		if actual.(countResponse).V != expected {
			t.Errorf("Expected %d, got %d", expected, actual)
		}
	}))
}

func TestDecodeUppercaseRequest(t *testing.T) {
	allure.Test(t, allure.Action(func() {
		str := "hello"
		expected := uppercaseRequest{str}

		req := httptest.NewRequest("POST", "/uppercase", strings.NewReader(`{"s":"`+str+`"}`))

		actual, err := decodeUppercaseRequest(nil, req)
		if err != nil {
			t.Error(err)
		}

		if actual.(uppercaseRequest).S != expected.S {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
	}))
}

func TestDecodeCountRequest(t *testing.T) {
	allure.Test(t, allure.Action(func() {
		str := "hello"
		expected := countRequest{str}

		req := httptest.NewRequest("POST", "/count", strings.NewReader(`{"s":"`+str+`"}`))

		actual, err := decodeCountRequest(nil, req)
		if err != nil {
			t.Error(err)
		}

		if actual.(countRequest).S != expected.S {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
	}))
}

func TestEncodeResponse(t *testing.T) {
	allure.Test(t, allure.Action(func() {
		str := "hello"
		expected := "{\"v\":\"" + str + "\"}\n"

		res := httptest.NewRecorder()
		err := encodeResponse(nil, res, uppercaseResponse{str, ""})
		if err != nil {
			t.Error(err)
		}

		if res.Body.String() != expected {
			t.Errorf("Expected %s, got %s", expected, res.Body.String())
		}
	}))
}
