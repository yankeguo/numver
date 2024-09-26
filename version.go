package numver

import (
	"strconv"
	"strings"
)

// Version is a sequence of integers.
type Version []int

// Parse parses a version string, extract all numeric parts found.
func Parse(s string) Version {
	v := Version{}

	var chunk string
	for _, c := range s {
		if c <= '9' && c >= '0' {
			chunk += string(c)
		} else {
			if chunk != "" {
				n, _ := strconv.Atoi(chunk)
				v = append(v, n)
				chunk = ""
			}
		}
	}

	if chunk != "" {
		n, _ := strconv.Atoi(chunk)
		v = append(v, n)
	}

	return v
}

// IsEmpty returns true if the version is empty.
func (v Version) IsEmpty() bool {
	return len(v) == 0
}

// String returns the version as a string, with dot separated.
func (v Version) String() string {
	var parts []string
	for _, n := range v {
		parts = append(parts, strconv.Itoa(n))
	}
	return strings.Join(parts, ".")
}

// Match returns true if the version matches the constraint.
func (v Version) Match(constraint Version) bool {
	if len(constraint) == 0 {
		return true
	}
	if len(constraint) > len(v) {
		return false
	}
	for i, n := range constraint {
		if v[i] != n {
			return false
		}
	}
	return true
}

// Compare returns -1 if v < other, 0 if v == other, 1 if v > other.
func (v Version) Compare(other Version) int {
	for i := 0; i < len(v) && i < len(other); i++ {
		if v[i] < other[i] {
			return -1
		}
		if v[i] > other[i] {
			return 1
		}
	}
	switch {
	case len(v) < len(other):
		return -1
	case len(v) > len(other):
		return 1
	default:
		return 0
	}
}
