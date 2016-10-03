package main

import (
	"os"
	"sort"
)

// AeFileSort - File Sorting type for clarify sake
type AeFileSort int

const (
	sortByNameAsc      AeFileSort = iota
	sortByNameDesc     AeFileSort = iota
	sortByFileSizeAsc  AeFileSort = iota
	sortByFileSizeDesc AeFileSort = iota
	sortByTimeAsc      AeFileSort = iota
	sortByTimeDesc     AeFileSort = iota
	sortByDefault      AeFileSort = sortByNameAsc
)

// byName implements sort.Interface.
type byName []os.FileInfo

func (f byName) Len() int           { return len(f) }
func (f byName) Less(i, j int) bool { return f[i].Name() < f[j].Name() }
func (f byName) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

type bySize []os.FileInfo

func (f bySize) Len() int           { return len(f) }
func (f bySize) Less(i, j int) bool { return f[i].Size() < f[j].Size() }
func (f bySize) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

type byTime []os.FileInfo

func (f byTime) Len() int           { return len(f) }
func (f byTime) Less(i, j int) bool { return f[i].ModTime().Before(f[j].ModTime()) }
func (f byTime) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }

// PsReadDir reads the directory named by dirname and returns
// a list of directory entries sorted by AeFileSort.
func PsReadDir(dirname string, sortType AeFileSort) ([]os.FileInfo, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}

	switch sortType {

	case sortByNameDesc:
		sort.Sort(sort.Reverse(byName(list)))

	case sortByFileSizeAsc:
		sort.Sort(bySize(list))

	case sortByFileSizeDesc:
		sort.Sort(sort.Reverse(bySize(list)))

	case sortByTimeAsc:
		sort.Sort(byTime(list))

	case sortByTimeDesc:
		sort.Sort(sort.Reverse(byTime(list)))

	case sortByNameAsc:
		fallthrough

	case sortByDefault:
		fallthrough

	default:
		sort.Sort(byName(list))
	}

	return list, nil
}

