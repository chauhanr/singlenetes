package cache

// IndexFunc knows how to compute the set of Indexed values for an object
type IndexFunc func(object interface{}) ([]string, error)

// Indexers map the name to the IndexFunc
type Indexers map[string]IndexFunc

// Index maps the indexed values to a set of keys in the store that match the value.
type Index map[string]string

// Indices maps a name to an Index
type Indices map[string]Index

// Indexer extends the store with more indices to make searching it faster.
type Indexer interface {
	Store
	Index(indexName string, object interface{}) ([]interface{}, error)
	GetIndexers() Indexers
	AddIndexers(newIndexer Indexers) error
	// IndexKeys(indexName, indexedValue string) ([]string, error)
	// ByIndex(indexName, indexedValue string) ([]interface{}, eror)
}
