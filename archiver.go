package backup

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

// Archiver represents type capable of archiving and
// restoring files.
type Archiver interface {
	Archive(src, dest string) error
}

type zipper struct{}

func (z *zipper) Archive(src, dest string) error {
	if err := os.MkdirAll(filepath.Dir(dest), 0777); err != nil {
		return err
	}
	out, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer out.Close()
	w := zip.NewWriter(out)
	defer w.Close()
	return filepath.Walk(src, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil // skip
		}

		in, err := os.Open(path)
		if err != nil {
			return err
		}
		defer in.Close()

		f, err := w.Create(path)
		if err != nil {
			return err
		}

		_, err = io.Copy(f, in)
		if err != nil {
			return err
		}
		return nil
	})
}

// ZIP is an Archiver that zips and unzips files.
var ZIP Archiver = (*zipper)(nil)
