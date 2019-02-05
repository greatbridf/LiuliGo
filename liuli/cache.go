package liuli

import (
    "bufio"
    "io"
    "os"
    "strings"
    "io/ioutil"
)

type CacheItem struct {
    id string
    hash string
}

type Cache struct {
    index *os.File
    valid bool
    list []CacheItem
}

func ReadCacheItem(line string) CacheItem {
    tmp := strings.Split(line, " ")
    item := CacheItem{}
    item.id = tmp[0]
    item.hash = tmp[1]
    return item
}

func (item CacheItem) ToString() string {
    tmp := item.id + " " + item.hash + "\n"
    return tmp
}

func (c *Cache) Init(indexPath string) {
    tmp_file, err := os.OpenFile(indexPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    if err != nil {
        PrintError("Cannot open cache file!")
    }
    c.index = tmp_file
    br := bufio.NewReader(c.index)
    for {
        line, _, flag := br.ReadLine()
        str_line := string(line[:])
        if flag == io.EOF {
            break
        }
        item := ReadCacheItem(str_line)
        c.list = append(c.list, item)
    }
}

func (c *Cache) Add(id string, hash string, data string) {
    item := CacheItem{id, hash}
    c.list = append(c.list, item)
    c.index.WriteString(item.ToString())
    err := ioutil.WriteFile("caches/" + hash, []byte(data), 0666)
    if err != nil {
        PrintError("Cannot write cache!")
    }
}

func (c *Cache) Get(id string) string {
    for i := 0; i < len(c.list); i++ {
        if c.list[i].id == id {
            fileName := c.list[i].hash
            data, err := ioutil.ReadFile("caches/" + fileName)
            if err != nil {
                PrintError("Cannot read from cache!")
            }
            return string(data[:])
        }
    }
    return ""
}

func (c *Cache) Find(id string) bool {
    for i := 0; i < len(c.list); i++ {
        if c.list[i].id == id {
            return true
        }
    }
    return false
}

func (c *Cache) Close() {
    c.index.Close()
}

