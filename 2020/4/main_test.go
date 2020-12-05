package main

import (
	"testing"
)

func TestByrValid(t *testing.T) {
	valid := isValid("2002", "byr")
	if !valid {
		t.Errorf("Wrong, expected true")
	}
}

func TestByrInvalid(t *testing.T) {
	valid := isValid("2003", "byr")
	if valid {
		t.Errorf("Wrong, expected false")
	}
}

func TestHgtValidIn(t *testing.T) {
	valid := isValid("60in", "hgt")
	if !valid {
		t.Errorf("Wrong, expected true")
	}
}

func TestHgtValidCm(t *testing.T) {
	valid := isValid("190cm", "hgt")
	if !valid {
		t.Errorf("Wrong, expected true")
	}
}

func TestHgtNInvalidCm(t *testing.T) {
	valid := isValid("190in", "hgt")
	if valid {
		t.Errorf("Wrong, expected false")
	}
}

func TestHgtNInvalidMissingUnit(t *testing.T) {
	valid := isValid("190", "hgt")
	if valid {
		t.Errorf("Wrong, expected false")
	}
}

func TestHclValid(t *testing.T) {
	valid := isValid("#123abc", "hcl")
	if !valid {
		t.Errorf("Wrong, expected true")
	}
}

func TestHclInvalidLetter(t *testing.T) {
	valid := isValid("#123abz", "hcl")
	if valid {
		t.Errorf("Wrong, expected false")
	}
}

func TestHclInvalidMissingHashtag(t *testing.T) {
	valid := isValid("123abc", "hcl")
	if valid {
		t.Errorf("Wrong, expected false")
	}
}

func TestEclValid(t *testing.T) {
	valid := isValid("brn", "ecl")
	if !valid {
		t.Errorf("Wrong, expected true")
	}
}

func TestEclInvalid(t *testing.T) {
	valid := isValid("wat", "ecl")
	if valid {
		t.Errorf("Wrong, expected false")
	}
}

func TestPidValid(t *testing.T) {
	valid := isValid("000000001", "pid")
	if !valid {
		t.Errorf("Wrong, expected true")
	}
}

func TestPidInvalid(t *testing.T) {
	valid := isValid("0123456789", "pid")
	if valid {
		t.Errorf("Wrong, expected false")
	}
}
