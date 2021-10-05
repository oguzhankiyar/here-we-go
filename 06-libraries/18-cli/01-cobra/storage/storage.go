package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
)

var Storage = NewStorage()

type storage struct {
	items map[int]string
}

func NewStorage() *storage {
	s := storage{
		items: make(map[int]string),
	}
	s.load()
	return &s
}

func (s *storage) GetAll() map[int]string {
	return s.items
}

func (s *storage) Get(id int) (string, error) {
	if value, ok := s.items[id]; !ok {
		return "", errors.New("not found")
	} else {
		return value, nil
	}
}

func (s *storage) Set(id int, name string) error {
	if _, ok := s.items[id]; ok {
		return errors.New("already exist")
	} else {
		s.items[id] = name
		s.save()
		return nil
	}
}

func (s *storage) Del(id int) error {
	if _, ok := s.items[id]; !ok {
		return errors.New("not found")
	} else {
		delete(s.items, id)
		s.save()
		return nil
	}
}

func (s *storage) load() {
	bytes, err := ioutil.ReadFile("storage/storage.json")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	err = json.Unmarshal(bytes, &s.items)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
}

func (s *storage) save() {
	bytes, err := json.Marshal(s.items)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	err = ioutil.WriteFile("storage.json", bytes, fs.ModePerm)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
}