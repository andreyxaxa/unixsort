package app

import "github.com/andreyxaxa/unixsort/pkg/unixsort"

// Run init and run all components
func Run() error {
	params := unixsort.NewParams()
	if err := params.Start(); err != nil {
		return err
	}

	return nil
}
