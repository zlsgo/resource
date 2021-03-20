package resource

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zlsgo/zstring"
)

func (r *Resource) MoveFile() error {
	d := zfile.RealPathMkdir(r.Dir, true)
	return filepath.Walk(r.tmpPath, func(path string, info os.FileInfo, err error) error {
		path = zfile.RealPath(path, info.IsDir())
		newPath := strings.Replace(path, r.tmpPath, d, 1)
		shortPath := strings.Replace(path, r.tmpPath, "/", 1)
		if len(r.ignore) > 0 {
			for _, v := range r.ignore {
				if pathMatche(shortPath, v) {
					return filepath.SkipDir
				}
			}
		}
		for k, v := range r.moveRule {
			if shortPath == k {
				newPath = d + strings.TrimLeft(v, "/")
				delete(r.moveRule, k)
				break
			}
		}
		if r.keep != "" && zfile.FileExist(newPath) {
			_ = os.Rename(newPath, newPath+"."+r.keep)
		}
		if info.IsDir() {
			_ = os.MkdirAll(newPath, info.Mode())
			return nil
		}
		return zfile.CopyFile(path, newPath)
	})
}

func pathMatche(path, rule string) bool {
	return zstring.RegexMatch(rule, path)
}
