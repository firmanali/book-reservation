package forms

import (
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	form := New(url.Values{})
	isValid := form.Valid()
	if !isValid {
		t.Error("Got invalid, should be valid")
	}
}

func TestForm_Required(t *testing.T) {
	form := New(url.Values{})

	form.Required("a","b","c")
	if form.Valid() {
		t.Error("Form shows valid when required fields missing")
	}

	postedData := url.Values{}
	postedData.Add("a","a")
	postedData.Add("b","b")
	postedData.Add("c","c")

	form = New(postedData)
	form.Required("a", "b", "c")
	if !form.Valid() {
		t.Error("Form shows invalid when required fields present")
	}
}

func TestForm_Has(t *testing.T) {
	form := New(url.Values{})

	if form.Has("a") {
		t.Error("Form shows valid when field missing")
	}

	isErr := form.Errors.Get("a")
	if isErr == "" {
		t.Error("Should have an error but didn't get one")
	}

	postedData := url.Values{}
	postedData.Add("a","a")
	form = New(postedData)

	if !form.Has("a") {
		t.Error("Form shows invalid when field present")
	}

	isErr = form.Errors.Get("a")
	if isErr != "" {
		t.Error("Should not have an error but got one")
	}
}

func TestForm_MinLength(t *testing.T) {
	data := url.Values{}
	data.Add("a", "123")
	form := New(data)
	if form.MinLength("a", 4) {
		t.Error("Form shows valid when value less than min length")
	}

	data.Add("b","1234")
	form = New(data)
	if !form.MinLength("b", 3) {
		t.Error("Form shows invalid when value more than min length")
	}
}

func TestForm_IsEmail(t *testing.T) {
	data := url.Values{}
	data.Add("a","abc")
	form := New(data)
	form.IsEmail("a")
	if form.Valid() {
		t.Error("Form shows valid when value is not email")
	}

	data.Add("b", "asd@sd.com")
	form = New(data)
	form.IsEmail("b")
	if !form.Valid() {
		t.Error("Form shows invalid when value is email")
	}
}