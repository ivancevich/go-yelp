package yelp

import (
	"testing"
	"time"
)

func TestIsValid(t *testing.T) {
	t.Run("Location and Coordinates are set", func(t *testing.T) {
		so := SearchOptions{
			Location: StringPtr("Kanto"),
			Coordinates: &Coordinates{
				Longitude: 31.54,
				Latitude:  3.22,
			},
		}
		if so.IsValid() {
			t.Fatal("This should not be valid.")
		}
	})

	t.Run("OpenNow and OpenAt are set", func(t *testing.T) {
		so := SearchOptions{
			Location: StringPtr("Hoenn"),
			OpenNow:  BoolPtr(true),
			OpenAt:   Int64Ptr(int64(time.Now().Second())),
		}
		if so.IsValid() {
			t.Fatal("This should not be valid.")
		}
	})

	t.Run("Only Location is set", func(t *testing.T) {
		so := SearchOptions{
			Location: StringPtr("Johto"),
		}
		if !so.IsValid() {
			t.Fatal("This should be valid.")
		}
	})

	t.Run("Only Coordinates is set", func(t *testing.T) {
		so := SearchOptions{
			Coordinates: &Coordinates{
				Longitude: 20.17,
				Latitude:  3.14159,
			},
		}
		if !so.IsValid() {
			t.Fatal("This should be valid.")
		}
	})

	t.Run("OpenNow is set", func(t *testing.T) {
		so := SearchOptions{
			Location: StringPtr("Unova"),
			OpenNow:  BoolPtr(true),
		}
		if !so.IsValid() {
			t.Fatal("This should be valid.")
		}
	})

	t.Run("Only OpenAt is set", func(t *testing.T) {
		so := SearchOptions{
			Location: StringPtr("Alola"),
			OpenAt:   Int64Ptr(int64(time.Now().Second())),
		}
		if !so.IsValid() {
			t.Fatal("This should be valid.")
		}
	})
}

func TestURLValues(t *testing.T) {
	t.Run("All url.Values are set correctly", func(t *testing.T) {
		so := SearchOptions{
			Coordinates: &Coordinates{
				Longitude: 132.231,
				Latitude:  123.57,
			},
			Term:       StringPtr("restaurants"),
			Radius:     Int64Ptr(1987),
			Categories: StringPtr("chinese"),
			Locale:     StringPtr("en_US"),
			Limit:      Int64Ptr(23),
			Offset:     Int64Ptr(13),
			SortBy:     StringPtr("rating"),
			Price:      StringPtr("1,4"),
			OpenNow:    BoolPtr(false),
			OpenAt:     Int64Ptr(int64(time.Now().Second())),
			Attributes: StringPtr("hot_and_new"),
		}
		vals := so.URLValues()
		for k, _ := range vals {
			v := vals.Get(k)
			switch k {
			case "longitude":
				if v != "132.231" {
					t.Fatalf("Longitude: Expected %s to equal 132.231.", v)
				}
			case "latitude":
				if v != "123.57" {
					t.Fatalf("Latitude: Expected %s to equal 123.57.", v)
				}
			case "location":
				if v != "" {
					t.Fatalf("Location: Expected \"%s\" to equal \"\".", v)
				}
			case "term":
				if v != "restaurants" {
					t.Fatalf("Term: Expected \"%s\" to equal \"restaurants\".", v)
				}
			}
		}
	})
}
