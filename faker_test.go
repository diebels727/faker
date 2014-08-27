package faker

import "testing"

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