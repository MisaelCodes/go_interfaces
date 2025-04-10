package advanced_level

import "fmt"

type Cache interface {
	Get(key string) (string, bool)
	Set(key, value string)
}

type InMemoryCache struct {
	data map[string]string
}

func NewInMemoryCache(data map[string]string) *InMemoryCache {
	return &InMemoryCache{data}
}

func (im *InMemoryCache) Get(key string) (string, bool) {
	res, ok := im.data[key]
	return res, ok
}

func (im *InMemoryCache) Set(key, value string) {
	im.data[key] = value
	fmt.Println(im.data)
}

type NoCache struct{}

func (nc *NoCache) Get(key string) (string, bool) {
	return "", false
}

func (nc *NoCache) Set(key, value string) {}

type UserInfo struct {
	C Cache
}

func (ui *UserInfo) Get(id string) string {
	info, ok := ui.C.Get(id)
	if !ok {
		info = "current user info unavailable"
		ui.C.Set(id, info)
	}
	return info
}
