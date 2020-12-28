package resource

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sohaha/zlsgo/zfile"
)

func (r *Resource) Compress() (err error) {
	if r.tmpPath == "" {
		return errors.New("please execute the download method first")
	}
	tmp := zfile.RealPathMkdir(r.tmpPath+"/tmp", true)
	switch r.Ext {
	case "gz":
		err = zfile.GzDeCompress(r.TmpFile, tmp)
	case "zip":
		err = zfile.ZipDeCompress(r.TmpFile, tmp)
	default:
		err = fmt.Errorf("compressed package format is not supported: %s", r.Ext)
	}
	if err != nil {
		return err
	}
	r.tmpPath = fixPath(tmp)
	return err
}

func fixPath(path string) string {
	n := 0
	dir := path
	_ = filepath.Walk(path, func(p string, i os.FileInfo, err error) error {
		if path == p {
			return nil
		}
		n++
		if i.IsDir() {
			dir = p
			return filepath.SkipDir
		}
		return nil
	})
	if n == 1 {
		path = dir
	}
	return zfile.RealPath(path, true)
}
