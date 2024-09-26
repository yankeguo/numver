package numver

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearch(t *testing.T) {
	extractor := func(s string) (v string, ok bool) {
		const (
			prefix = "node-v"
			suffix = "-linux-x64.tar.gz"
		)
		ok = strings.HasPrefix(s, prefix) && strings.HasSuffix(s, suffix)
		v = strings.TrimPrefix(strings.TrimSuffix(s, suffix), prefix)
		return
	}

	items := []string{
		"node-v1.3.2-windows-x64.tar.gz",
		"node-v1.3.3-linux-x64.tar.gz",
		"node-v1.2.2-linux-x64.tar.gz",
		"node-v1.2.3-linux-x64.tar.gz",
		"node-v1.2.2-windows-x64.tar.gz",
		"node-v1.2.3-windows-x64.tar.gz",
		"node-v1.3.4-linux-x64.tar.gz",
		"node-v1.3.5-windows-x64.tar.gz",
	}

	item, version, found := Search(
		SearchOptions{
			Items:      items,
			Extractor:  extractor,
			Constraint: "1.3",
		},
	)
	require.Equal(t, "node-v1.3.4-linux-x64.tar.gz", item)
	require.Equal(t, "1.3.4", version.String())
	require.True(t, found)

	item, version, found = Search(
		SearchOptions{
			Items:      items,
			Constraint: "1.3",
		},
	)
	require.Equal(t, "node-v1.3.5-windows-x64.tar.gz", item)
	// this is intentional, because the version only considers numbers
	require.Equal(t, "1.3.5.64", version.String())
	require.True(t, found)

	item, version, found = Search(
		SearchOptions{
			Items:      items,
			Extractor:  extractor,
			Constraint: "1.3",
			Descending: true,
		},
	)
	require.Equal(t, "node-v1.3.3-linux-x64.tar.gz", item)
	require.Equal(t, "1.3.3", version.String())
	require.True(t, found)

	item, version, found = Search(
		SearchOptions{
			Items:      items,
			Extractor:  extractor,
			Constraint: "1.3.8",
			Descending: true,
		},
	)
	require.Equal(t, "", item)
	require.True(t, version.IsEmpty())
	require.False(t, found)
}
