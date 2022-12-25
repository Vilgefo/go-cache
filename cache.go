package go_cache

type Cache map[string]interface{}

func New() Cache {
	return Cache(make(map[string]interface{}))
}

func (c Cache) Set(name string, value interface{}) {
	c[name] = value
}

func (c Cache) Get(name string) (interface{}, bool) {
	value, exists := c[name]
	return value, exists
}

func (c Cache) Delete(name string) {
	delete(c, name)
}
