package resource

import (
	"errors"
	"path"
	"strings"

	"github.com/sohaha/zlsgo/zfile"
	"github.com/sohaha/zlsgo/zhttp"
	"github.com/sohaha/zlsgo/zstring"
	"github.com/sohaha/zlsgo/zutil"
)

// Download Download
func (r *Resource) Download(progress func(current, total int64)) error {
	var data []interface{}
	if progress != nil {
		data = append(data, zhttp.DownloadProgress(progress))
	}
	res, err := zhttp.Get(r.Remote, data...)
	if err != nil {
		return err
	}
	name := path.Base(r.Remote)
	ext := path.Ext(name)
	tmp := zfile.TmpPath()
	r.TmpFile = tmp + "/" + name
	err = res.ToFile(r.TmpFile)
	if err != nil {
		return err
	}
	fileMd5, _ := zstring.Md5File(r.TmpFile)
	if r.md5 != "" && r.md5 != fileMd5 {
		return errors.New("file validation failed")
	}
	r.tmpPath = tmp
	r.Ext = zutil.IfVal(ext == "", "zip", strings.TrimLeft(ext, ".")).(string)
	return nil
}
