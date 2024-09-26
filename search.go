package numver

// SearchOptions is a set of options for the Search function.
type SearchOptions struct {
	// Items is a list of items to search.
	Items []string
	// Constraint is a version constraint to match.
	Constraint string
	// Extractor is a function to extract a version from an item.
	Extractor func(src string) (ver string, ok bool)
	// Descending is a flag to search in descending order.
	Descending bool
}

// Search searches for the best matching item in the list of items.
func Search(opts SearchOptions) (itemFound string, versionFound Version, found bool) {
	constraint := Parse(opts.Constraint)

	var (
		items    []string
		versions []Version
	)

	for _, item := range opts.Items {
		rawVersion := item
		if opts.Extractor != nil {
			var ok bool
			if rawVersion, ok = opts.Extractor(item); !ok {
				continue
			}
		}
		if version := Parse(rawVersion); version.Match(constraint) {
			items = append(items, item)
			versions = append(versions, version)
		}
	}

	for i := range items {
		if versionFound == nil ||
			((!opts.Descending) && versions[i].Compare(versionFound) > 0) ||
			(opts.Descending && versions[i].Compare(versionFound) < 0) {
			itemFound = items[i]
			versionFound = versions[i]
			found = true
		}
	}

	return
}
