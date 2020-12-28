package resource

import (
	"github.com/sohaha/zlsgo/zfile"
)

type Resource struct {
	Remote   string
	Dir      string
	TmpFile  string
	Ext      string
	tmpPath  string
	md5      string
	keep     string
	ignore   []string
	moveRule map[string]string
}

func New(remote string) *Resource {
	return &Resource{
		Remote: remote,
	}
}

func (r *Resource) SilentRun(progress func(current, total int64)) error {
	err := r.Download(progress)
	if err != nil {
		return err
	}

	err = r.Compress()
	if err != nil {
		return err
	}

	return r.MoveFile()
}

func (r *Resource) SetMd5(m string) {
	r.md5 = m
}

func (r *Resource) SetDeCompressPath(path string) {
	r.Dir = zfile.RealPath(path, true)
}

func (r *Resource) SetKeepOldFile(keep string) {
	r.keep = keep
}

func (r *Resource) SetFilterRule(ignore []string) {
	r.ignore = ignore
}

func (r *Resource) SetMoveRule(moveRule map[string]string) {
	r.moveRule = moveRule
}