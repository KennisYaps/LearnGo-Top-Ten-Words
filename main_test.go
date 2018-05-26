package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHomePageHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Error(err)
	}
	res := httptest.NewRecorder()
	HomePageHandler(res, req)
	if res.Code != http.StatusOK {
		t.Errorf("Response code was %v; want 200", res.Code)
	}
}

func TestResultPageHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/results/", nil)
	if err != nil {
		t.Error(err)
	}
	res := httptest.NewRecorder()
	ResultPageHandler(res, req)
	if res.Code != http.StatusOK {
		t.Errorf("Response code was %v; want 200", res.Code)
	}
}
