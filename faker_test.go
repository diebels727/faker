package faker

import (
        "testing"
        "fmt"
       )

func TestFirstNamesLengthZero(t *testing.T) {
  f := Faker{}
  if len(f.FirstNames) != 0 {
    t.Error("Expected fake_data[\"first_names\"] to be zero, got ",len(f.FirstNames))
  }
}

func TestFirstNamesNonZero(t *testing.T) {
  f := New()
  if len(f.FirstNames) == 0 {
    t.Error("Expected fake_data[\"first_names\"] to be zero, got ",len(f.FirstNames))
  }
}

func TestRandomFirstNameNonNil(t *testing.T) {
  f := New()
  name := f.FirstName()
  fmt.Println("name")
  if len(name) == 0 {
    t.Error("Expected name to not have length zero.")
  }
}

func TestRandomShouldNotMatch(t *testing.T) {
  f := New()
  first := f.FirstName()
  second := f.FirstName()
  if first == second {
    t.Error("First name",first,"should not equal second name",second)
  }
}