package main

type KeyStore map[string]string

func applyOnce(store KeyStore, key, result string) (string, bool) {
	// TODO: make repeated delivery return the same stored result.
	if old, ok := store[key]; ok {
		return old, false
	}
	store[key] = result
	return result, true
}

func main() {}
