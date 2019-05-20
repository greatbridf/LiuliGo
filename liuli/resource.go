package liuli

import (
	"github.com/pkg/errors"
)

type DeleteResult struct {
	Code int
	Msg  string
}

func DeleteResource(id string) (*DeleteResult, error) {
	cache := Cache{}
	err := cache.Init()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer cache.Close()
	if !cache.Find(id) {
		result := DeleteResult{
			404,
			"Not Found",
		}
		Log.D("Queried to delete " + id + " but found nothing")
		return &result, nil
	}
	hash := cache.GetHash(id)
	err = cache.Remove(id)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	Log.D("Deleted " + id + " (was " + hash + ")")
	result := DeleteResult{
		200,
		"Deleted",
	}
	return &result, nil
}
