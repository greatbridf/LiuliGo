package liuli

import (
	"bufio"
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

func (c *Cache) Init(indexPath string) {
	c.list = make(map[string]string)
	c.rev = make(map[string]string)
	tmp_file, err := os.OpenFile(indexPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		PrintError("Cannot open cache file!")
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
}

func (c *Cache) Add(id string, data []byte) {
	hash := Hash(data)
	c.list[id] = hash
	c.rev[hash] = id
	c.index.WriteString(id + " " + hash + "\n")
	err := ioutil.WriteFile("caches/"+hash, data, 0666)
	if err != nil {
		PrintError("Cannot write cache!")
	}
	PrintDebug("Add " + id + " to cache!")
}

func (c Cache) Get(id string) ([]byte, bool) {
	data, err := ioutil.ReadFile("caches/" + c.GetHash(id))
	if err != nil {
		PrintError("Cannot read from cache!")
		return nil, false
	}
	PrintDebug("Get " + id + " from cache")
	return data, true
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
