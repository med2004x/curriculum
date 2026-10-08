# Lesson 100 — Graduation: Design, Build, Measure, Explain

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Build a complete, production-shaped network service integrating correct networking, persistence, observability, and failure handling — the first of four terminal capstones drawing on nearly every track in this curriculum, proving the pieces compose into something a real team could actually operate.

**The contract.** The capstone adds no isolated technology; it tests integration. A small, complete vertical path with explicit limits is stronger evidence than a broad feature list that has not been run end to end.

**Worked example.** Start with the smallest concrete case described by the source material: Build a real, non-trivial network service (your choice of domain — a genuine extension of TARDOC, Mahall, or Lead Sourcer is encouraged, since real motivation produces better engineering than an arbitrary toy) meeting this minimum integrated scope:

## 2. Mechanism and vocabulary

**Mechanism.** Contract-first design fixes the user-visible interface before implementation. Persistence invariants, idempotency, bounded work, timeouts, telemetry, migrations, and recovery must agree at the boundaries between components. A load test states offered rate, concurrency, data shape, duration, and the point where latency or errors show saturation.

**Second point.** Production-shaped means that the service has repeatable setup, health checks, bounded behavior under failure, traceable releases, and documented recovery. Its scope is intentionally modest; evidence and operability matter more than feature count.

**Worked example.** **Contract-first design** (API engineering Lessons 1-3): a written requirements document and OpenAPI spec before implementation.

## 3. Worked case from the reference material

Worked example: Build a real, non-trivial network service (your choice of domain — a genuine extension of TARDOC, Mahall, or Lead Sourcer is encouraged, since real motivation produces better engineering than an arbitrary toy) meeting this minimum integrated scope:

## 4. Verification, failure, and evidence

**Verification.** **Load test it** (production Lesson 6): a real, open-loop load test finding your actual capacity and bottleneck, compared against a theoretical capacity estimate (system-engineering Lesson 2).

**Failure mode.** Break a cross-component assumption that isolated tests miss: for example, let a retry repeat a committed write without an idempotency key, or make a circuit-breaker rejection invisible to its metric. Diagnose from user-path evidence, repair it, and add an integration test that fails on the broken behavior.

**Completion criterion.** Another engineer can build and run the service from documented steps, reproduce its measured workload and synthetic incident, and identify remaining risks without relying on unstated assumptions.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
