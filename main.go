package main

import (
	"context"
)

type Utils struct{}

// Walk the directory, and return a flattened list of all its files
func (m *Utils) Walk(ctx context.Context, dir *Directory) ([]string, error) {
	entries, err := dir.Entries(ctx)
	if err != nil {
		return nil, err
	}
	for _, path := range entries {
		_, err := dir.Entries(ctx, DirectoryEntriesOpts{Path:path})
		if err != nil {
			continue
		}
		subentries, err := m.Walk(ctx, dir.Directory(path))
		if err != nil {
			return nil, err
		}
		for i := range subentries {
			subentries[i] = path + "/" + subentries[i]
		}
		entries = append(entries, subentries...)
	}
	return entries, nil
}
