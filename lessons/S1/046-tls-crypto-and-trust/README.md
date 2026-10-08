# Lesson 046 — TLS, Cryptography, and Trust

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Understand what TLS actually provides — encryption, authentication, integrity — and specifically what a certificate proves (and doesn't prove), by inspecting a real handshake rather than treating "HTTPS" as an opaque padlock icon.

**The contract.** **Three separate guarantees, commonly conflated into "it's secure."** TLS provides: **confidentiality** (an eavesdropper on the network path can't read the content — the connection is encrypted), **integrity** (the content can't be silently modified in transit without detection — any tampering breaks a cryptographic check), and **authentication** (the client.

**Worked example.** Start with the smallest concrete case described by the source material: Generate a local, self-signed TLS certificate (e.g.

## 2. Mechanism and vocabulary

**Mechanism.** These are genuinely separable properties; understanding them separately matters because a system can have some without others (e.g. an encrypted-but-unauthenticated connection is still vulnerable to a man-in-the-middle who simply presents their own valid encryption keys, since nothing was checking *whose* keys they were).

**Second point.** **What a certificate actually proves.** A certificate is a signed statement, from a Certificate Authority (CA) that the client already trusts (via a pre-installed list of trusted root CAs), asserting "this public key belongs to this domain name." When your browser connects to `example.com` and receives a certificate, verification means:

**Worked example.** `openssl req -x509 -newkey rsa:2048 -keyout key.pem -out cert.pem -days 1 -nodes`) and use it to run a minimal HTTPS server (Go's `net/http` with `ListenAndServeTLS`, pointing at your generated cert/key, is straightforward for this).

## 3. Worked case from the reference material

Worked example: Generate a local, self-signed TLS certificate (e.g.

## 4. Verification, failure, and evidence

**Verification.** For step 3, report the actual certificate chain length (how many certificates were presented) and the negotiated cipher suite for a real production server, and state in your own words what each certificate in the chain is vouching for relative to the next one up the chain.

**Failure mode.** Take your step 1 self-signed server and, using a packet capture tool (Lesson 7 preview) or a network proxy tool, attempt to read the actual HTTP request/response content passing over your `curl -k` connection at the raw TCP byte level (not through the TLS-aware tool itself, but genuinely at the wire).

**Completion criterion.** You've run a real TLS handshake against your own self-signed server and observed both the encrypted-but-unverified case and the properly-rejected-for-lack-of-trust case, you've inspected a real production certificate chain and can explain what each link in it vouches for, and you can state clearly, using your own failure-drill observation of genuinely unread.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
