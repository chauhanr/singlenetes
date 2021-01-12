package cache

import "github.com/chauhanr/singlenetes/apimachinery/watch"

type Lister interface {
}

type Watcher interface {
	Watch() (watch.Interface, error)
}
