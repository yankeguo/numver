package numver

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	require.Equal(t, Version{}, Parse("a.b.c"))
	require.Equal(t, Version{}, Parse(""))
	require.Equal(t, Version{1, 2, 3}, Parse("1.2.3"))
	require.Equal(t, Version{234, 35, 22}, Parse(" as 234 35_22  "))
	require.Equal(t, Version{1, 0, 3}, Parse("1.0.3"))
	require.Equal(t, Version{1, 2, 3}, Parse("1.0002.3"))
	require.Equal(t, Version{0, 2, 3}, Parse("0.0002.3"))
}

func TestVersionIsEmpty(t *testing.T) {
	require.True(t, Version{}.IsEmpty())
	require.False(t, Version{0}.IsEmpty())
}

func TestVersionMatch(t *testing.T) {
	require.True(t, Version{}.Match(Version{}))
	require.False(t, Version{}.Match(Version{1}))
	require.True(t, Version{1, 2, 3}.Match(Version{1, 2, 3}))
	require.True(t, Version{1, 2, 3}.Match(Version{1, 2}))
	require.False(t, Version{1, 2, 3}.Match(Version{1, 2, 2}))
	require.False(t, Version{1, 2, 3}.Match(Version{1, 2, 4}))
	require.False(t, Version{1, 2, 3}.Match(Version{1, 2, 3, 1}))
}

func TestVersionString(t *testing.T) {
	require.Equal(t, "1.2.3", Version{1, 2, 3}.String())
	require.Equal(t, "1.0.3", Version{1, 0, 3}.String())
	require.Equal(t, "234.4.22", Version{234, 4, 22}.String())
}

func TestVersionCompare(t *testing.T) {
	require.Equal(t, 0, Version{1, 2, 3}.Compare(Version{1, 2, 3}))
	require.Equal(t, -1, Version{1, 2, 3}.Compare(Version{1, 2, 4}))
	require.Equal(t, 1, Version{1, 2, 3}.Compare(Version{1, 2, 2}))
	require.Equal(t, -1, Version{1, 2, 3}.Compare(Version{1, 3}))
	require.Equal(t, 1, Version{1, 2, 3, 3}.Compare(Version{1, 2, 3}))
	require.Equal(t, -1, Version{1, 2}.Compare(Version{1, 2, 3}))
}
