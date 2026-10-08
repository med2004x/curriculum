# Lesson 078 — Cryptography and Secure Protocols

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Use standard cryptographic primitives (hashing, HMAC, signatures) correctly in real code, and understand precisely why "rolling your own crypto" is dangerous — not as a platitude, but by directly experiencing a specific, common way naive cryptographic code fails.

**The contract.** A cryptographic hash is deterministic and fixed-length; security properties include preimage, second-preimage, and collision resistance. A hash is not encryption and does not in general hide low-entropy input. HMAC combines a hash with a secret key for message authentication; a digital signature uses a private key to sign and a public key to verify.

**Worked example.** Compute SHA-256 for two inputs differing by one byte using the standard library. Observe that the digests differ; do not treat one pair as proof of collision resistance or claim every output bit must change.

## 2. Mechanism and vocabulary

**Mechanism.** Use a password-hashing scheme designed to be salted and computationally expensive, such as Argon2id, scrypt, bcrypt, or PBKDF2 with an appropriate current work factor. Store the encoded parameters/salt/hash format returned by the library; do not invent a custom password hash.

**Second point.** For HMAC verification, compare MAC bytes with the cryptographic library's constant-time verification/comparison function. For digital signatures, ensure the public key is trusted through the protocol's certificate/key distribution rules; mathematical verification alone does not identify a trusted publisher.

**Worked example.** Implement HMAC-based message signing and verification (extending API engineering Lesson 11's practical webhook use directly), and confirm that altering either the message or the signature (but not both consistently) causes verification to correctly fail — test both cases explicitly, not just the happy path.

## 3. Worked case from the reference material

Worked example: Compute SHA-256 hashes of several inputs (using your language's standard crypto library, never a hand-rolled implementation) and confirm the hash changes completely for even a single-character input change (the "avalanche effect" — a defining property of a good cryptographic hash), demonstrating this isn't a simple checksum where small input changes produce.

## 4. Verification, failure, and evidence

**Verification.** Present your step 2 HMAC verification results (both the correct-pass and correct-fail cases), your step 3 signature verification results (correct key pair passes, wrong key pair fails), and your step 4 brute-force time estimates for bcrypt versus plain SHA-256, showing the concrete, large difference in attacker cost.

**Failure mode.** Replace the expected HMAC key or public key with an unrelated key and verify that authentication fails. Separately compare password-verification work for fast SHA-256 and a standard password hash; explain why the fast hash is unsuitable for password storage. Do not implement a real endpoint that accepts a hand-rolled password verifier.

**Completion criterion.** You can state what each primitive guarantees and requires, verify valid and tampered messages using supported libraries, and select a password-hashing scheme rather than a fast general-purpose hash.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
