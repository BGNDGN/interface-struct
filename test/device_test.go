package test

import (
	"testing"
)

type MockPhone struct {
	Markasi string
	Modeli  string
}

func (m MockPhone) Gelis() string {
	return "BeklenenDeger"
}

func (m MockPhone) İcerik() string {
	return "BeklenenDeger2"
}

func TestMockPhoneGelis(t *testing.T) {
	mock := MockPhone{
		Markasi: "Iphone",
		Modeli:  "13 Pro",
	}

	expected := "BeklenenDeger"
	result := mock.Gelis()

	if result != expected {
		t.Errorf("Alınması Gereken Değer:%s\nAlınan Değer:%s", expected, result)
	}
}

func TestMockPhoneİcerik(t *testing.T) {
	mock := MockPhone{
		Markasi: "Iphone",
		Modeli:  "13 Pro",
	}

	expected := "BeklenenDeger2"
	result := mock.İcerik()

	if result != expected {
		t.Errorf("Alınması Gereken Değer: %s\nAlınan Değer:%s", expected, result)
	}
}
