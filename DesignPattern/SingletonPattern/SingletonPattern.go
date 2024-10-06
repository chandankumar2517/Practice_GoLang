package designpattern

import "sync"

var (
	Instance *Singleton

	once sync.Once
)

type Singleton struct{}

func GetInstance() *Singleton {
	once.Do(func() {
		Instance = &Singleton{}
	})
	return Instance
}
