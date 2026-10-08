# Lesson 077 — Threat Modeling and Trust Boundaries

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Identify what must be protected, who can influence it, and where trust changes before choosing a security mechanism.

**The contract.** Cover assets, actors, trust boundaries, attack surfaces, abuse cases, least privilege, and a practical STRIDE-style review.

**Worked example.** Start with the smallest concrete case described by the source material: Threat-model a small task API with browser clients, an API service, PostgreSQL, object storage, and an administrator.

## 2. Mechanism and vocabulary

**Mechanism.** Draw the data-flow diagram with external actors, processes, data stores, and trust-boundary crossings. For each flow, identify what data crosses, who controls it, and what identity or integrity claim is relied upon.

**Second point.** STRIDE prompts are Spoofing, Tampering, Repudiation, Information disclosure, Denial of service, and Elevation of privilege. Apply them to concrete components/flows, prioritize by impact and feasibility, then identify controls and remaining risk.

**Worked example.** For `browser → API → PostgreSQL/object storage`, consider a forged identity token, a cross-tenant record ID, an oversized upload, a leaked signed URL, and a forged webhook. Map each threat to the relevant trust boundary and control (token validation, ownership check, request limits, short-lived scoped URL, HMAC verification and deduplication).

## 3. Worked case from the reference material

Worked example: Threat-model a small task API with browser clients, an API service, PostgreSQL, object storage, and an administrator.

## 4. Verification, failure, and evidence

**Verification.** Produce a diagram and a threat register containing asset, actor, entry point/flow, threat, impact, mitigation, owner, and residual risk. Ensure each mitigation addresses the stated threat, not merely a neighboring concern.

**Failure mode.** Remove the resource-ownership check while leaving authentication and role checks in place. Show which trust assumption permits cross-tenant access and identify the missing authorization predicate.

**Completion criterion.** Every high-priority threat has a concrete boundary, mitigation, and residual-risk owner; controls are justified by the system's assets and flows.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
