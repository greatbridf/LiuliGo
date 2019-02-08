package liuli

import (
	"github.com/pkg/errors"
)

// GetResource Get resource from cache or internet by hash
func GetResource(hash string) ([]byte, error) {
	cache := Cache{}
	err := cache.Init()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer cache.Close()
	if cache.HasHash(hash) {
		id := cache.GetIDByHash(hash)
		data, err := cache.Get(id)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		return data, nil
	}
	return nil, errors.New("No resource found according to hash")
}
