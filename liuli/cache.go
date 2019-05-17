package liuli

import (
	"database/sql"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Cache struct {
	index *sql.DB
	list  map[string]string
	rev   map[string]string
}

func (c *Cache) Init() error {
	c.list = make(map[string]string)
	c.rev = make(map[string]string)
	_, err := os.Stat("caches")
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir("caches", 0775)
			if err != nil {
				return HE(err)
			}
		} else {
			return HE(err)
		}
	}
	index, err := sql.Open("sqlite3", "index.db")
	c.index = index
	if err != nil {
		return HE(err)
	}
	// Test if the database exists
	rows, err := c.index.Query("SELECT hash,name FROM `index`")
	if err != nil {
		create, err := c.index.Prepare("CREATE TABLE `index`(id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT, hash TEXT)")
		if err != nil {
			return HE(err)
		}
		_, err = create.Exec()
		if err != nil {
			return HE(err)
		}
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		// Read to memory
		var hash string
		var name string
		err := rows.Scan(&hash, &name)
		if err != nil {
			return HE(err)
		}
		c.list[name] = hash
		c.rev[hash] = name
	}
	return nil
}

func (c *Cache) Add(id string, data []byte) error {
	hash := Hash(data)
	c.list[id] = hash
	c.rev[hash] = id
	err := ioutil.WriteFile("caches/"+hash, data, 0666)
	if err != nil {
		return HEM(err, "Cannot write cache file")
	}
	stmt, err := c.index.Prepare("INSERT INTO `index`(name, hash) VALUES(?, ?)")
	if err != nil {
		return HE(err)
	}
	_, err = stmt.Exec(id, hash)
	if err != nil {
		return HE(err)
	}
	Log.D("Add " + id + " to cache!")
	return nil
}

func (c Cache) Get(id string) ([]byte, error) {
	data, err := ioutil.ReadFile("caches/" + c.GetHash(id))
	if err != nil {
		return nil, HE(errors.WithStack(err))
	}
	PrintDebug("Get " + id + " from cache")
	return data, nil
}

func (c Cache) Find(id string) bool {
	_, exists := c.list[id]
	return exists
}

func (c *Cache) Remove(id string) error {
	if !c.Find(id) {
		return errors.New("Resource does not exist")
	}
	delete(c.rev, c.GetHash(id))
	delete(c.list, id)
	stmt, err := c.index.Prepare("DELETE FROM `index` WHERE name=?")
	if err != nil {
		return HE(err)
	}
	_, err = stmt.Exec(id)
	if err != nil {
		return HE(err)
	}
	return nil
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
