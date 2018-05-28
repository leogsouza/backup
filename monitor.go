package backup

import (
	"fmt"
	"path/filepath"
	"time"
)

// Monitor monitors files and folder changes
type Monitor struct {
	Paths       map[string]string
	Achiver     Archiver
	Destination string
}

// Now reports how many files are changed
func (m *Monitor) Now() (int, error) {
	var counter int
	for path, lastHash := range m.Paths {
		newHash, err := DirHash(path)
		if err != nil {
			return counter, err
		}
		if newHash != lastHash {
			err := m.act(path)
			if err != nil {
				return counter, err
			}
			m.Paths[path] = newHash // update the hash
			counter++
		}
	}
	return counter, nil
}

func (m *Monitor) act(path string) error {
	dirname := filepath.Base(path)
	filename := fmt.Sprintf("%d.zip", time.Now().UnixNano())
	return m.Achiver.Archive(path, filepath.Join(m.Destination,
		dirname, filename))
}
