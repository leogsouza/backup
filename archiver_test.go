package backup_test

import (
	"os"
	"testing"

	"github.com/leogsouza/backup"
	"github.com/stretchr/testify/require"
)

func setup(t *testing.T) {
	os.MkdirAll("test/output", 0777)
}

func teardown(t *testing.T) {
	os.RemoveAll("test/output")
}

func TestZipArchive(t *testing.T) {
	setup(t)
	defer teardown(t)

	err := backup.ZIP.Archive("test/hash1", "test/output/1.zip")
	require.NoError(t, err)
}

func TestZipErrors(t *testing.T) {
	setup(t)
	defer teardown(t)

	errorTests := []struct {
		name string
		src  string
		dest string
	}{
		{name: "Destination path invalid", src: "/test/hash1", dest: "//.../test/#$$output/.#../1.zip"},
		{name: "Source path invalid", src: "`%?.ska", dest: "test/output/1.zip"},
		{name: "Filename invalid", src: "/test/hash1", dest: "~/test//output/1/fasfsdfasdfasdf/.."},
	}

	for _, tt := range errorTests {
		t.Run(tt.name, func(t *testing.T) {
			err := backup.ZIP.Archive(tt.src, tt.dest)
			require.Error(t, err)
		})
	}
}
