# Lesson 044 — IP Addressing, Subnets, and Routing

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Understand how packets get from one IP network to another.

**The contract.** CIDR `/p` fixes the first `p` network bits; the remaining bits identify addresses within the prefix. A host first determines whether the destination is on-link according to its configured prefix. For off-link destinations it selects a next hop from the routing table, generally using the most specific matching prefix. Ethernet/neighbor discovery resolves the next-hop link-layer address; it does not choose the IP route.

**Worked example.** For `192.0.2.64/26`, the mask is `255.255.255.192`, the block size in the last octet is 64, and the prefix covers `.64` through `.127`. Decide whether `192.0.2.100` is on-link (yes) and whether `192.0.2.130` is on-link (no; use the configured route/gateway).

## 2. Mechanism and vocabulary

**Mechanism.** IPv4 has 32 bits. The prefix mask has `p` leading one-bits; network address is obtained by bitwise AND with the mask. A routing table can contain overlapping prefixes; choose the matching entry with the longest prefix, subject to route metrics and operating-system policy for equal prefixes.

**Second point.** A host sends an on-link packet in a link-layer frame addressed to the destination's resolved MAC address. For an off-link destination it sends the frame to the next-hop gateway's MAC address while the IP packet retains the remote destination IP.

IPv6 uses 128-bit addresses and Neighbor Discovery rather than IPv4 ARP. Do not apply IPv4 broadcast-address rules to IPv6.

**Worked example.** Given routes `10.0.0.0/8 via G1`, `10.4.0.0/16 via G2`, and `default via G0`, destination `10.4.8.9` matches all three but uses the `/16` route because it is the longest matching prefix.

## 3. Worked case from the reference material

Worked example: Given a small network diagram, calculate subnets and determine which gateway handles several destinations.

## 4. Verification, failure, and evidence

**Verification.** For each supplied destination, write the binary prefix comparison, selected route, next-hop IP, and which link-layer address must be resolved. Compare the hand calculation with the host's routing table without assuming the local machine has the example routes.

**Failure mode.** Select a less-specific route despite a more-specific match, or resolve the remote destination's MAC instead of the gateway's. Explain which network-layer or link-layer decision was confused.

**Completion criterion.** You can calculate a prefix boundary, choose the longest matching route, and explain how the selected next hop is delivered over the local link.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
