# Lesson 080 — Isolation, Containers, and Supply Chain

## Threads
**Primary:** Security/Linux/Production

**Supporting:** process isolation; container boundaries; dependencies

**Math companion:** risk and reliability

**Time:** 3–4 hours. **Prerequisites:** Lesson 079.

**Time:** 3–4 hours. **Prerequisites:** Lesson 079.

## What you will be able to do

1. Distinguish container namespaces/cgroups from a virtual machine's separate guest kernel.
2. Inspect process, network, mount, and capability boundaries in a container.
3. Apply least-privilege runtime settings and verify the denied operation.
4. Describe supply-chain controls for base images, dependencies, provenance, and vulnerability findings.

## Retrieve first (no notes)

1. Why do SQL parameters separate values from query syntax?
2. Why is output escaping context-specific?
3. What does CSRF exploit in a browser?
4. Why must SSRF defenses consider redirects and resolved addresses?

## 1. Problem and model

**What problem this solves.** Understand what container isolation actually provides (and its real limits) at the OS level — namespaces and capabilities — and apply least-privilege thinking (system-engineering Lesson 10) concretely to container configuration.

**The contract.** Containers normally share the host kernel while isolating selected process, mount, network, user, and other namespaces and limiting resources through cgroups. A VM presents a guest machine with its own kernel boundary. Neither model removes the need to patch, configure least privilege, and protect secrets.

**Worked example.** Start with the smallest concrete case described by the source material: Run a container (Docker or equivalent) and, from inside it, inspect `/proc/1/status` and process listing — confirm you see a restricted process tree (your container's own processes only, with your container's main process as PID 1), directly observing PID namespace isolation in action, distinct from what `ps aux` on the actual host machine would show.

Self-check: write the input, the state that changes, and the observable result before running anything.

## 2. Mechanism and vocabulary

**Mechanism.** Namespaces shape a process's view of selected resources; cgroups account for and limit resource use. Linux capabilities split some root privileges into separate permissions. A process running as UID 0 inside a container is not automatically equivalent to host root, but misconfiguration or kernel vulnerabilities can weaken the boundary.

**Second point.** Supply-chain controls include pinning base images by digest where appropriate, minimizing/updating dependencies, generating an SBOM, verifying signatures or provenance when available, and reviewing vulnerability reports. These controls improve traceability but do not prove software is vulnerability-free.

Keep the terminology exact. If the source uses a named boundary, command, data structure, protocol, or formula, use that name rather than replacing it with a vague synonym.

**Worked example.** Compare a default isolated network namespace with host networking; inspect visible interfaces and explain the effect on isolation. Then run as a non-root user, drop unnecessary capabilities, use a read-only root filesystem where practical, and document writable mounts.

Self-check: explain why the example works using the mechanism above, not by describing the final output alone.

## 3. Worked case from the reference material

Worked example: Run a container (Docker or equivalent) and, from inside it, inspect `/proc/1/status` and process listing — confirm you see a restricted process tree (your container's own processes only, with your container's main process as PID 1), directly observing PID namespace isolation in action, distinct from what `ps aux` on the actual host machine would show.

Then write the same case in your own words as a three-step trace: starting state, operation, resulting state.

Self-check: which step is directly observable from the terminal and which step is only an inference?

## 4. Verification, failure, and evidence

**Verification.** Record process/network/mount visibility, effective user/capabilities, read-only/writable paths, image digest, dependency/image scan results, and the exact attempted operation that was denied. Distinguish configured settings from observed behavior.

**Failure mode.** Attempt a controlled operation that should be blocked by the selected capability, user, filesystem, or namespace restriction. If it succeeds, identify the missing boundary and adjust the configuration; do not probe the host or unrelated systems.

**Completion criterion.** You can explain namespace/cgroup/capability boundaries, demonstrate least privilege, and describe image/dependency provenance and residual supply-chain risk.

Worked example: run one success case, then one deliberate failure. Keep the commands and evidence together so another person can reproduce the observation.

Self-check: what observation would prove that the implementation is correct rather than merely lucky on one input?

## Predict, then observe

1. Predict process visibility inside and outside a PID namespace.
2. Predict which interfaces differ under host networking.
3. Predict the effect of dropping one Linux capability.
4. Predict what a read-only root filesystem prevents and what mounted volumes still permit.
5. Predict which image input a digest pins.
6. Predict what an SBOM reports and what it does not prove.
7. Predict whether a configuration flag proves an operation is actually denied.
8. State the shared-kernel boundary in a container.

Run the smallest case first. Record prediction and observation side by side.

## Build from blank

1. Inspect process visibility and PID identity inside a container and compare it with the host view.

2. Compare isolated networking with host networking and record visible interfaces/routes.

3. Compare root/default capabilities with a non-root, capability-dropped, read-only configuration; test one operation expected to be denied.

Do not open the reference implementation until your own attempt has been saved.

## Failure drill

1. Take a capability-dropped, non-root container from steps 3-4 and attempt an operation that *should* be prevented by your restricted configuration but that you suspect might still be possible due to some capability or permission you didn't think to drop (e.g.

2. If you find something still works that shouldn't, given your intended security posture, identify specifically what additional restriction (a more limited volume mount, an additional dropped capability, a read-only root filesystem via `--read-only`) would be needed, and apply it.

3. Explain why this kind of "test what you think is restricted, don't just configure and assume" verification step is necessary — configuring capability drops and namespaces correctly is easy to get subtly wrong, and only actually testing the specific operations you intend to prevent (not just checking that the configuration flags are present) gives real confid.

4. Reproduce the failure from a clean terminal session and write the smallest change that returns the system to the expected contract.

## Transfer

Change exactly one constraint in the main case for # 080 — Isolation, Containers, and Supply Chain: change the input shape, data size, interface, or operating condition without changing the underlying objective. Solve it from blank, then verify the new behavior with a terminal command or a reproducible calculation.

Your transfer answer is complete only when the changed constraint is explicit and the evidence distinguishes the new case from the original case.

## Retrieval questions

1. Which kernel is ordinarily shared by containers on one host?
2. What resource views do namespaces isolate?
3. What do cgroups account for or limit?
4. How do capabilities refine root privilege?
5. What does an image digest identify?
6. What do SBOM and provenance records establish, and what remains unproven?

## Evidence to record

Create `lesson080/EVIDENCE.md` with exactly this structure:

```text
# Lesson 080 evidence

Prediction:

Command or calculation:

Observed result:

Failure reproduced:

What changed after the fix:

One sentence I can recall tomorrow:
```

## Common mistakes

- Treating containers as separate guest kernels.
- Assuming UID 0 inside a container has all host-root capabilities.
- Assuming a read-only root filesystem makes mounted volumes read-only.
- Trusting flags without testing the intended denied operation.
- Treating a clean scan or SBOM as proof of safety.

## Precision checks

1. State the central contract in one sentence without using the lesson title as a substitute for the explanation.
2. Re-run the worked case after deleting your own notes. The command, input, and observation must still be reproducible.
3. Change exactly one boundary condition and classify the result as expected behavior, invalid input, implementation failure, or environment failure.
4. Write one sentence explaining what evidence would distinguish two competing explanations for the same observation.
5. Save the terminal command and the result before moving to the next lesson.

This is the point where reading becomes retrieval: the reference material is available, but the first explanation and first implementation must come from memory.
## University revision bridge

When this lesson is used for the LIG track, keep the engineering task and the academic recall separate. First reproduce the formal definition, notation, or standard exercise pattern from the companion pack. Then use this lesson to implement or verify the same concept in the terminal.

The engineering layer does not replace the university syllabus. It is the deeper implementation and reasoning layer.

## Assigned material

The supporting source notes that previously lived in the separate reference tree have been folded into the canonical lesson sequence. Use the lecture first, then the adjacent [`WORK.md`](WORK.md). The unified external/source map is [`../../SOURCES.md`](../../SOURCES.md).

## Mastery gate

You pass when you can explain the central model of # 080 — Isolation, Containers, and Supply Chain without notes, rebuild the core task from blank, reproduce one deliberate failure, and solve the transfer constraint. You've directly observed namespace isolation's effect on process visibility and network interfaces, you've compared and tested real capability differences between default-root, capability-dropped, and non-root container configurations, and you've verified — not just configured and assumed — that your hardening actually prevents the specific operations it's m. **Do not move on because the example looked familiar.** Familiarity is cheap; blank-page recall is what actually learning costs.
