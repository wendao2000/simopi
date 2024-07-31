package app

var cache = make(map[string]map[string]Simopi)

func GetCacheByCred(c Cred) (Simopi, bool) {
	m, ok := cache[c.User][c.Endpoint+":"+c.Method]
	return m, ok
}

func SetCache(m Simopi) {
	if cache[m.User] == nil {
		cache[m.User] = make(map[string]Simopi)
	}

	cache[m.User][m.Endpoint+":"+m.Method] = m
}

func DeleteCacheByCred(c Cred) {
	delete(cache[c.User], c.Endpoint+":"+c.Method)
}

func DeleteCacheByUser(user string) {
	delete(cache, user)
}

func ClearCache() {
	cache = make(map[string]map[string]Simopi)
}
