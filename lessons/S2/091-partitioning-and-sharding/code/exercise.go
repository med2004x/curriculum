package main

import "hash/fnv"

func shardFor(key string, shards int) int {
	// TODO: handle invalid shard counts and define the stable hashing contract.
	h := fnv.New32a()
	_, _ = h.Write([]byte(key))
	return int(h.Sum32() % uint32(shards))
}

func main() {}
