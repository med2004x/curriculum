# Lesson 044 — IP Addressing, Subnets, and Routing

## Threads
**Primary:** Networking/Math

**Supporting:** CIDR; routing; address spaces

**Math companion:** binary and logarithms

**Time:** 3–4 hours. **Prerequisites:** Lesson 043.

**Time:** 3–4 hours. **Prerequisites:** Lesson 043.

## What you will be able to do

1. Calculate an IPv4 network address, broadcast address, and usable host range from a CIDR prefix.
2. Decide whether a destination is on-link or must be sent to a gateway.
3. Apply longest-prefix match to a routing table.
4. Distinguish Layer 2 neighbor resolution from Layer 3 routing.

## Retrieve first (no notes)

1. What does UDP not guarantee about delivery?
2. What is a DNS TTL?
3. What is the difference between `NXDOMAIN` and a timeout?
4. Why can DNS use more than one transport protocol?

## 1. Problem and model

**What problem this solves.** Understand how packets get from one IP network to another.

**The contract.** CIDR `/p` fixes the first `p` network bits; the remaining bits identify addresses within the prefix. A host first determines whether the destination is on-link according to its configured prefix. For off-link destinations it selects a next hop from the routing table, generally using the most specific matching prefix. Ethernet/neighbor discovery resolves the next-hop link-layer address; it does not choose the IP route.

**Worked example.** For `192.0.2.64/26`, the mask is `255.255.255.192`, the block size in the last octet is 64, and the prefix covers `.64` through `.127`. Decide whether `192.0.2.100` is on-link (yes) and whether `192.0.2.130` is on-link (no; use the configured route/gateway).

Self-check: `/26` leaves 6 host bits, so the address block contains 64 total IPv4 addresses; do not confuse total addresses with generally usable host addresses, especially for special prefix lengths.

## 2. Mechanism and vocabulary

**Mechanism.** IPv4 has 32 bits. The prefix mask has `p` leading one-bits; network address is obtained by bitwise AND with the mask. A routing table can contain overlapping prefixes; choose the matching entry with the longest prefix, subject to route metrics and operating-system policy for equal prefixes.

**Second point.** A host sends an on-link packet in a link-layer frame addressed to the destination's resolved MAC address. For an off-link destination it sends the frame to the next-hop gateway's MAC address while the IP packet retains the remote destination IP.

IPv6 uses 128-bit addresses and Neighbor Discovery rather than IPv4 ARP. Do not apply IPv4 broadcast-address rules to IPv6.

**Worked example.** Given routes `10.0.0.0/8 via G1`, `10.4.0.0/16 via G2`, and `default via G0`, destination `10.4.8.9` matches all three but uses the `/16` route because it is the longest matching prefix.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Given a small network diagram, calculate subnets and determine which gateway handles several destinations.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** For each supplied destination, write the binary prefix comparison, selected route, next-hop IP, and which link-layer address must be resolved. Compare the hand calculation with the host's routing table without assuming the local machine has the example routes.

**Failure mode.** Select a less-specific route despite a more-specific match, or resolve the remote destination's MAC instead of the gateway's. Explain which network-layer or link-layer decision was confused.

**Completion criterion.** You can calculate a prefix boundary, choose the longest matching route, and explain how the selected next hop is delivered over the local link.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Calculate the block size and range for `192.0.2.64/26`.
2. Predict whether `192.0.2.100` is on-link for that prefix.
3. Predict the next hop for an off-link destination.
4. Select between `/8`, `/16`, and default routes for `10.4.8.9`.
5. Predict which MAC address is resolved for an off-link packet on Ethernet.
6. State how IPv6 neighbor discovery differs from IPv4 ARP.
7. Predict what changes when a more-specific route is added.
8. Separate route selection from neighbor resolution in your explanation.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Build the smallest useful version: Given a small network diagram, calculate subnets and determine which gateway handles several destinations.

2. Extend it with one explicit boundary case: Inspect a real routing table on Linux.

3. Rebuild the same behavior with one changed constraint: Write a second case with one changed constraint and verify it.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Break one assumption and record the first observable failure.

2. Change one boundary condition and classify the resulting failure.

3. Remove one verification step and record what evidence is lost.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change one route or prefix: add a more-specific destination prefix, or change the local network from `/26` to `/27`. Recompute on-link membership and the longest-prefix route without changing the packet destination.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. What does an IPv4 CIDR prefix specify?
2. How do you calculate a network address from an address and prefix?
3. What is longest-prefix match?
4. When is a destination on-link?
5. Which MAC address is used for an off-link packet's Ethernet frame?
6. What does ARP resolve, and what does routing decide?

## Evidence to record

Create `lesson044/EVIDENCE.md` with exactly this structure:

```text
# Lesson 044 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Confusing the IP destination with the Ethernet next-hop address.
- Ignoring a more-specific matching route.
- Treating ARP as a routing protocol.
- Counting all addresses in a subnet as ordinary host addresses.
- Applying IPv4 broadcast assumptions to IPv6.

## Precision checks

1. Calculate network and host portions from a CIDR prefix.
2. Decide on-link versus gateway delivery.
3. Apply longest-prefix match to overlapping routes.
4. Identify the IP next hop and corresponding link-layer resolution.
5. Distinguish IPv4 ARP from IPv6 Neighbor Discovery.

Routing selects a Layer 3 next hop; neighbor resolution supplies the link-layer destination for the next transmission.
## University revision bridge

For the LIG track, connect CIDR prefixes to binary masks and graph-like route selection, then distinguish the IP routing decision from local link delivery.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 044 — IP Addressing, Subnets, and Routing without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You can explain the route decision for a packet without invoking “the router just knows”. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
