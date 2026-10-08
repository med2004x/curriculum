# Lesson 018 — Hashing, Maps, and Collision Reasoning

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

A hash table stores key-value pairs in an array, using a hash function to map a key to an index. The table supports:

- insert
- lookup
- delete

The main challenge is that multiple different keys can map to the same slot. This is a collision.

A table with `m` slots and `n` entries has a load factor:

```text
alpha = n / m
```

If the table is too full, collisions become frequent and lookup cost increases. A good hash table grows by rehashing when the load factor crosses a threshold, usually around 0.5 to 0.75 depending on the design.

## 2. Mechanism and vocabulary

There are two common collision strategies:

1. Separate chaining
   - each bucket stores a linked list (or another container)
   - collision means “more than one key shares one slot”

2. Open addressing
   - each key is stored directly in the table
   - if a slot is occupied, the algorithm probes a sequence of alternative positions

The main probe patterns are:

- linear probing: check next slot in a fixed step
- quadratic probing: skip by a quadratic sequence
- double hashing: use a second hash function to determine step size

The key idea is that expected lookup cost is near constant when the table is not too full. In the worst case, however, all keys can hash to the same slot, and then the cost becomes O(n).

## 3. Worked case: mapping keys to slots

Suppose we use a table of 8 slots and hash function `h(k) = k mod 8`.

Keys: 10, 18, 26, 35, 43

- `10 mod 8 = 2`
- `18 mod 8 = 2`
- `26 mod 8 = 2`
- `35 mod 8 = 3`
- `43 mod 8 = 3`

This creates collisions at slots 2 and 3. In a chaining table, each slot contains a list of key-value pairs. In an open-address table, the algorithm keeps probing until it finds an empty slot.

A good hash table is not magical. It is just a space-time trade-off: a larger table reduces collisions but uses more memory.

## 4. Verification, failure, and evidence

You can verify a hash table by recording:

- the hash function used
- the keys inserted
- the bucket or probe sequence
- the final state of the table

Failure modes include:

- poor hash function causing many collisions
- overloaded table causing long probe chains
- incorrect resize logic causing lost entries
- use-after-free or stale pointers in chained nodes

A data structure is correct only if every key can be found again after insertion, even when collisions occur.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
