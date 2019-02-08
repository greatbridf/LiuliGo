package liuli

import (
	"bufio"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Cache struct {
	index *os.File
	valid bool
	list  map[string]string
	rev   map[string]string
}

func (c *Cache) Init() error {
	c.list = make(map[string]string)
	c.rev = make(map[string]string)
	tmp_file, err := os.OpenFile("caches/index", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return errors.WithStack(err)
	}
	c.index = tmp_file
	br := bufio.NewReader(c.index)
	for {
		line, _, flag := br.ReadLine()
		if flag == io.EOF {
			break
		}
		id_and_hash := strings.Split(string(line[:]), " ")
		c.list[id_and_hash[0]] = id_and_hash[1]
		c.rev[id_and_hash[1]] = id_and_hash[0]
	}
	return nil
}

func (c *Cache) Add(id string, data []byte) error {
	hash := Hash(data)
	c.list[id] = hash
	c.rev[hash] = id
	c.index.WriteString(id + " " + hash + "\n")
	err := ioutil.WriteFile("caches/"+hash, data, 0666)
	if err != nil {
		return errors.Wrap(err, "Cannot write cache file")
	}
	Log.D("Add " + id + " to cache!")
	return nil
}

func (c Cache) Get(id string) ([]byte, error) {
	data, err := ioutil.ReadFile("caches/" + c.GetHash(id))
	if err != nil {
		Log.D(err.Error())
		return nil, errors.WithStack(err)
	}
	PrintDebug("Get " + id + " from cache")
	return data, nil
}

func (c Cache) Find(id string) bool {
	_, exists := c.list[id]
	return exists
}

func (c Cache) GetHash(id string) string {
	hash := c.list[id]
	return hash
}

func (c Cache) GetIDByHash(hash string) string {
	id := c.rev[hash]
	return id
}

func (c Cache) HasHash(hash string) bool {
	_, exists := c.rev[hash]
	return exists
}

func (c *Cache) Close() {
	c.index.Close()
}
