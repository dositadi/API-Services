package posts

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

var (
	notFound = errors.New("Not found!")
)

type PostStore struct {
	list map[string]Post
	Path string
}

func NewPostStore(path string) *PostStore {
	list := make(map[string]Post)
	return &PostStore{
		list: list,
		Path: path,
	}
}

func (p *PostStore) ReadFromDB() (map[string]Post, error) {
	file, err := os.OpenFile(p.Path, os.O_RDONLY, 0644)
	if err != nil {
		return map[string]Post{}, fmt.Errorf("Error Reading (Opening) DB: %w", err)
	}

	defer file.Close()

	fileStat, err2 := file.Stat()
	if err2 != nil {
		return map[string]Post{}, fmt.Errorf("Error Reading (Stat) DB: %w", err2)
	}

	postBuffer := make([]byte, fileStat.Size())

	_, err3 := file.Read(postBuffer)
	if err3 != nil {
		return map[string]Post{}, fmt.Errorf("Error Reading (Read) DB: %w", err3)
	}

	recievePosts := make(map[string]Post)

	err4 := json.Unmarshal(postBuffer, &recievePosts)
	if err4 != nil {
		return map[string]Post{}, fmt.Errorf("Error Reading (Json) DB: %w", err4)
	}
	return recievePosts, nil
}

func (p *PostStore) WriteToDB(posts map[string]Post) error {
	file, err := os.OpenFile(p.Path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("Error Writing (Opening) to DB, %w, %w", notFound, err)
	}

	postJson, err2 := json.MarshalIndent(posts, "", " ")
	if err2 != nil {
		return fmt.Errorf("Error Writing (json) to DB %w", err2)
	}

	_, err3 := file.Write(postJson)
	if err3 != nil {
		return fmt.Errorf("Error Writing (Write) to DB %w", err3)
	}
	return nil
}

/*
type PostsMethods interface {
	Get(id string) (m.Post, error)
	Put(id string, post m.Post) error
	List() (map[string]m.Post, error)
	Delete(id string) error
	Update(id string, post m.Post) error
} */

func (p *PostStore) Get(id string) (Post, error) {
	DB, err := p.ReadFromDB()
	if err != nil {
		return Post{}, err
	}

	if post, ok := DB[id]; ok {
		return post, nil
	}
	return Post{}, notFound
}

func (p *PostStore) Put(id string, post Post) error {
	DB, err := p.ReadFromDB()
	if err != nil {
		DB[id] = post

		err1 := p.WriteToDB(DB)
		if err1 != nil {
			return err
		}
		return nil
	}

	if _, ok := DB[id]; ok {
		return errors.New("Error: Post already exists in DB, Do you want to update instead!")
	}

	if p.CheckPostKey(post.Id) {
		return errors.New("Error: Post already exists in DB, Do you want to update instead!")
	}

	DB[id] = post

	err1 := p.WriteToDB(DB)
	if err1 != nil {
		return err
	}
	return nil
}

func (p *PostStore) List() (map[string]Post, error) {
	DB, err := p.ReadFromDB()
	if err != nil {
		return map[string]Post{}, err
	}

	return DB, nil
}

func (p *PostStore) Delete(id string) error {
	DB, err := p.ReadFromDB()
	if err != nil {
		return nil
	}

	dB_Id := ""

	for storeID := range DB {
		if strings.Compare(id, DB[storeID].Id) == 0 {
			dB_Id = storeID
			break
		}
	}

	if dB_Id != "" {
		delete(DB, dB_Id)

		err1 := p.WriteToDB(DB)
		if err1 != nil {
			return err
		}
		return nil
	}
	return notFound
}

func (p *PostStore) Update(id string, post Post) error {
	DB, err := p.ReadFromDB()
	if err != nil {
		return nil
	}

	dB_Id := ""

	for storeID := range DB {
		if strings.Compare(id, DB[storeID].Id) == 0 {
			dB_Id = storeID
			break
		}
	}

	if dB_Id != "" {
		DB[dB_Id] = post

		err1 := p.WriteToDB(DB)
		if err1 != nil {
			return err
		}
		return nil
	}
	return notFound
}

func (p *PostStore) CheckPostKey(id string) bool {
	DB, err := p.ReadFromDB()
	if err != nil {
		return false
	}

	for storeID := range DB {
		if strings.Compare(id, DB[storeID].Id) == 0 {
			return true
		}
	}
	return false
}
