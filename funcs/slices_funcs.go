package funcs

func (c *cache) AddArray(arr []string) {
	for _, element := range arr {
		c.add_to_slice_cache(element, true)
	}
}
