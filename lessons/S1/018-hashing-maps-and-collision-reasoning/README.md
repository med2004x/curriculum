# Lesson 018 — Hashing, Maps, and Collision Reasoning

## Threads
**Primary:** C/Algorithms

**Supporting:** hash functions; collisions; maps; probing

**Math companion:** expected cost and probability

**Time:** 3–4 hours. **Prerequisites:** Lesson 017.

## What you will be able to do

1. Explain what a hash table does and why it is useful for dictionary-like lookup.
2. Explain the difference between a hash function, a collision, and a load factor.
3. Describe the difference between separate chaining and open addressing.
4. Explain why a table with good distribution has near-constant average lookup time.

## Retrieve first (no notes)

1. What is the difference between LIFO and FIFO?
2. Why is recursion tied to a call stack?
3. What is a recurrence?
4. What does it mean for a stack to overflow?
5. What does it mean for a queue to underflow?
6. Why are proofs about algorithmic cost useful in systems design?

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

## Predict, then observe

1. Predict what happens when many keys hash to the same slot.
2. Predict whether a full table with poor collision resolution still provides fast lookup.
3. Predict why a resize can reduce average lookup cost.
4. Predict what happens if the table is not rehashed after growth.
5. Predict how insertion differs from lookup in an open-addressed table.
6. Predict how a load factor greater than 1 behaves in a hash table.

## Build from blank

1. Build a small integer-key hash table using an array of buckets.
2. Implement `insert`, `lookup`, and `delete`.
3. Add a collision scenario where at least two keys map to the same slot.
4. Explain what happens in both chaining and open addressing.

## Failure drill

1. Create a set of adversarial keys that all map to the same slot.
2. Explain why certain operations degrade to O(n).
3. Add a resize and compare the cost before and after resizing.
4. Record the invariant: every live key is stored at a valid location and is reachable by the correct hash path.

## Transfer

Change exactly one condition: use a different hash function, a different table size, or a different collision strategy. Explain which operation becomes faster or slower and why the result follows from the load factor and collision pattern.

## Retrieval questions

1. What is a hash table?
2. What is a collision?
3. What is the load factor?
4. What is separate chaining?
5. What is open addressing?
6. Why does resizing help?

## Evidence to record

Create `lesson018/EVIDENCE.md` with exactly this structure:

```text
# Lesson 018 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Thinking that a hash table is always O(1) in the worst case.
- Ignoring the load factor.
- Using a weak hash function without checking collisions.
- Forgetting to resize or rehash after growth.

## Precision checks

1. State the difference between a hash function and a collision in one sentence.
2. Explain why a large load factor worsens lookup cost.
3. State one difference between chaining and open addressing.
4. Explain how a resize changes the expected cost.
5. State one example of a poorly distributed key set.

## University revision bridge

The academic layer is the probabilistic model: a good hash function distributes keys broadly, and average-case behavior is close to constant. The engineering layer is to observe actual collision behavior under real keys and decide when resizing or alternative probes are necessary.

## Assigned material

- `07-data-structures-and-algorithms/lesson-04-hash-tables/README.md`
- `08-mathematics/step-05-probability/lesson-02-random-variables/README.md`

## Mastery gate

You pass when you can explain the hash table model, describe collision handling, and explain why the average-case lookup cost is constant but the worst-case cost can be linear. You should be able to show a small insertion and lookup sequence and explain the bucket or probe path.
