package funcs

func CreatCache() cache {
	return cache{}
}

func (c *cache) add_to_slice_cache(key string, value bool) {
	c.slice_cache[key] = value
}
