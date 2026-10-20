package usercase_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/lyonnee/hvalid"
	"github.com/lyonnee/hvalid/validators/format"
)

type profile struct {
	Nickname    string
	DisplayName string
	Country     string
	Currency    string
	Color       string
	UserID      string
	ISBN        string
	Updated     string
	Port        string
	Host        string
}

func validateProfile(value profile) error {
	checks := []error{
		hvalid.Validate("nickname", value.Nickname, hvalid.Alpha[string](), hvalid.MinRunes[string](2)),
		hvalid.Validate("displayName", value.DisplayName, hvalid.NotBlank[string]()),
		hvalid.Validate("country", value.Country, format.ISO3166Alpha2[string]()),
		hvalid.Validate("currency", value.Currency, format.ISO4217[string]()),
		hvalid.Validate("color", value.Color, format.HexColor[string]()),
		hvalid.Validate("userID", value.UserID, format.UUID[string]()),
		hvalid.Validate("isbn", value.ISBN, format.ISBN[string]()),
		hvalid.Validate("updated", value.Updated, format.RFC3339[string]()),
		hvalid.Validate("port", value.Port, format.Port[string]()),
		hvalid.Validate("host", value.Host, format.Host[string]()),
	}
	var errs []error
	for _, err := range checks {
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

func sampleProfile() profile {
	return profile{
		Nickname:    "Ann",
		DisplayName: "Ann Lee",
		Country:     "CN",
		Currency:    "CNY",
		Color:       "#336699",
		UserID:      "550e8400-e29b-41d4-a716-446655440000",
		ISBN:        "978-0-306-40615-7",
		Updated:     "2026-10-19T16:42:00Z",
		Port:        "8080",
		Host:        "example.com",
	}
}

func TestProfileWalkthrough(t *testing.T) {
	if err := validateProfile(sampleProfile()); err != nil {
		t.Fatal(err)
	}

	bad := sampleProfile()
	bad.Nickname = "A1"
	bad.DisplayName = ""
	err := validateProfile(bad)
	if err == nil {
		t.Fatal("expected profile errors")
	}

	want := "nickname: must contain only ASCII letters\ndisplayName: must not be blank"
	if err.Error() != want {
		t.Fatalf("got %q, want %q", err.Error(), want)
	}

	group, ok := err.(interface{ Unwrap() []error })
	if !ok {
		t.Fatal("profile errors should keep each field error")
	}
	fields := group.Unwrap()
	if len(fields) != 2 {
		t.Fatalf("got %d field errors", len(fields))
	}
	for _, fieldErr := range fields {
		var ve *hvalid.ValidationError
		if !errors.As(fieldErr, &ve) {
			t.Fatalf("field error %q is not a validation error", fieldErr)
		}
		if strings.Count(fieldErr.Error(), ve.Field+":") != 1 {
			t.Fatalf("field %s was wrapped again: %s", ve.Field, fieldErr)
		}
	}
}
