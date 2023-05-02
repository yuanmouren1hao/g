package git

import (
	"sync"

	"github.com/Equationzhao/g/cached"
)

var (
	ignored         *cached.Map[GitRepoPath, *FileGits]
	IgnoredInitOnce sync.Once
)

func GetCache() *cached.Map[GitRepoPath, *FileGits] {
	IgnoredInitOnce.Do(func() {
		ignored = cached.NewCacheMap[GitRepoPath, *FileGits](20)
	})
	return ignored
}

func FreeCache() {
	ignored.Free()
}

func DefaultInit(repoPath GitRepoPath) func() *FileGits {
	return func() *FileGits {
		res := make(FileGits, 0)
		out, err := GetShortGitStatus(repoPath)
		if err == nil {
			res = ParseShort(out)
		}
		return &res
	}
}
