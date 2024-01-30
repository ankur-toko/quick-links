package repository

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ankur-toko/quick-links/core/models"
)

const separator = ":$%@$"
const default_db_location = "./data/filedb"

type FileDB struct {
	dbFile      *os.File
	AbsFilePath string
	dbpath      string
	CacheDB     QuickLinkRepo
	ch          chan models.QuickLink
}

func (db *FileDB) initDBPath() (string, error) {
	p, e := filepath.Abs(db.dbpath + "/data.db")
	if e != nil {
		return "", e
	}
	db.AbsFilePath = p
	return db.AbsFilePath, e

}

func (db *FileDB) getDataFilePath() string {
	return db.AbsFilePath
}

func GetFileDB(cacheDB QuickLinkRepo) (*FileDB, error) {
	dbpath := default_db_location

	if cacheDB == nil {
		return nil, errors.New("cache db cannot be null for file db repo")
	}

	err := os.MkdirAll(dbpath, 0777)
	if err != nil && !os.IsExist(err) {
		fmt.Print(err)
		return nil, err
	}

	db := FileDB{}
	db.dbpath = dbpath
	db.CacheDB = cacheDB
	db.ch = make(chan models.QuickLink, 100)

	db.initDBPath()

	_, e := db.createFileIfNotExist()
	if e != nil {
		return nil, e
	}

	e = db.initializeDb()
	if e != nil {
		return nil, e
	}
	file, err := os.OpenFile(db.getDataFilePath(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)

	if err != nil {
		return nil, e
	}
	db.dbFile = file

	go db.processSaveRequests()

	return &db, nil
}

func (db *FileDB) createFileIfNotExist() (*os.File, error) {
	var file *os.File
	_, err := os.Stat(db.getDataFilePath())
	if os.IsNotExist(err) {
		// Create the db file here
		f, e := os.Create(db.getDataFilePath())
		if e != nil {
			fmt.Print("err", e)
			return nil, e
		}
		defer f.Close()
		file = f
	} else {
		fmt.Println("DB File Already Exists!")
	}

	return file, nil
}

func (db *FileDB) initializeDb() error {
	// Reads the file from the start and repopulates the memory
	file, err := os.OpenFile(db.getDataFilePath(), os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		datum := strings.Split(scanner.Text(), separator)
		log.Print(datum)
		if len(datum) == 2 {
			db.CacheDB.Save(models.QuickLink{Key: datum[0], URL: datum[1]})
		} else {
			// unable to process this line of the data file
			fmt.Print("error reading file db line: ", datum)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func (db *FileDB) Save(r models.QuickLink) error {
	db.CacheDB.Save(r)

	// quick link is quickly pushed to Cache DB and a channel which is then consumed eventually.
	// Sometimes this may lead to lost data but should be rare
	db.ch <- r
	return nil
}

func (db *FileDB) processSaveRequests() {
	for data := range db.ch {
		db.persistData(data)
	}
}

func (db *FileDB) persistData(r models.QuickLink) error {
	_, e := db.dbFile.Write([]byte(serialize(r)))
	if e != nil {
		fmt.Println("unable to store in persistant db", e)
		return e
	}
	return nil
}

func serialize(r models.QuickLink) string {
	return r.Key + separator + r.URL + "\n"
}

func (db *FileDB) Get(key string) *models.QuickLink {
	return db.CacheDB.Get(key)
}
