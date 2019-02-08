package liuli

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type Cache struct {
	index     *os.File
	valid     bool
	list      map[string]string
	hash_list []string
}

func (c *Cache) Init(indexPath string) {
	c.list = make(map[string]string)
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
		c.hash_list = append(c.hash_list, id_and_hash[1])
	}
}

func (c *Cache) Add(id string, hash string, data []byte) {
	c.list[id] = hash
	c.hash_list = append(c.hash_list, hash)
	c.index.WriteString(id + " " + hash + "\n")
	err := ioutil.WriteFile("caches/"+hash, data, 0666)
	if err != nil {
		PrintError("Cannot write cache!")
	}
	PrintDebug("Add " + id + " to cache!")
}

func (c Cache) Get(id string) ([]byte, bool) {
	hash, exists := c.list[id]
	if !exists {
		return nil, false
	}
	data, err := ioutil.ReadFile("caches/" + hash)
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
	hash, _ := c.list[id]
	return hash
}

func (c *Cache) Close() {
	c.index.Close()
}
