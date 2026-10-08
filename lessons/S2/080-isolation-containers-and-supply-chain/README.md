# Lesson 080 — Isolation, Containers, and Supply Chain

Read this explanation, then solve the problems in [WORK.md](WORK.md). The work card names the files to edit and how to check your work.

## 1. Problem and model

**What problem this solves.** Understand what container isolation actually provides (and its real limits) at the OS level — namespaces and capabilities — and apply least-privilege thinking (system-engineering Lesson 10) concretely to container configuration.

**The contract.** Containers normally share the host kernel while isolating selected process, mount, network, user, and other namespaces and limiting resources through cgroups. A VM presents a guest machine with its own kernel boundary. Neither model removes the need to patch, configure least privilege, and protect secrets.

**Worked example.** Start with the smallest concrete case described by the source material: Run a container (Docker or equivalent) and, from inside it, inspect `/proc/1/status` and process listing — confirm you see a restricted process tree (your container's own processes only, with your container's main process as PID 1), directly observing PID namespace isolation in action, distinct from what `ps aux` on the actual host machine would show.

## 2. Mechanism and vocabulary

**Mechanism.** Namespaces shape a process's view of selected resources; cgroups account for and limit resource use. Linux capabilities split some root privileges into separate permissions. A process running as UID 0 inside a container is not automatically equivalent to host root, but misconfiguration or kernel vulnerabilities can weaken the boundary.

**Second point.** Supply-chain controls include pinning base images by digest where appropriate, minimizing/updating dependencies, generating an SBOM, verifying signatures or provenance when available, and reviewing vulnerability reports. These controls improve traceability but do not prove software is vulnerability-free.

**Worked example.** Compare a default isolated network namespace with host networking; inspect visible interfaces and explain the effect on isolation. Then run as a non-root user, drop unnecessary capabilities, use a read-only root filesystem where practical, and document writable mounts.

## 3. Worked case from the reference material

Worked example: Run a container (Docker or equivalent) and, from inside it, inspect `/proc/1/status` and process listing — confirm you see a restricted process tree (your container's own processes only, with your container's main process as PID 1), directly observing PID namespace isolation in action, distinct from what `ps aux` on the actual host machine would show.

## 4. Verification, failure, and evidence

**Verification.** Record process/network/mount visibility, effective user/capabilities, read-only/writable paths, image digest, dependency/image scan results, and the exact attempted operation that was denied. Distinguish configured settings from observed behavior.

**Failure mode.** Attempt a controlled operation that should be blocked by the selected capability, user, filesystem, or namespace restriction. If it succeeds, identify the missing boundary and adjust the configuration; do not probe the host or unrelated systems.

**Completion criterion.** You can explain namespace/cgroup/capability boundaries, demonstrate least privilege, and describe image/dependency provenance and residual supply-chain risk.

## Next step

Open [WORK.md](WORK.md), complete its problems in order, and record your commands and results in [EVIDENCE.md](EVIDENCE.md).
