# S2 COURSE EXERCISE - OS2
# Study the prompts below, then complete the exercises in this file.

# # FSEGS LIG — Systèmes d'exploitation 2
#
# **Semester:** S2
# **Attached engineering lesson:** 058 — Processes, Threads, and Scheduling Revisited
#
# ## Scope
#
# Process/thread model; CPU scheduling; synchronization; critical sections; semaphores/mutexes; deadlocks; virtual memory; pages/frames/page tables; allocation; filesystems; I/O; IPC; protection; performance trade-offs.
#
# ## Teaching notes
#
# Distinguish the process abstraction from the thread abstraction before discussing scheduling. A process provides an isolated execution environment; threads are execution paths that share important process resources. A context switch changes the execution context and has a cost.
#
# For scheduling questions, write the policy, arrival times, burst times, and time quantum before drawing the timeline. Waiting time and turnaround time are different quantities. In synchronization questions, identify the shared state, critical section, required guarantee, and mechanism.
#
# For deadlock, check the actual conditions rather than using the word as a synonym for “slow”. For virtual memory, write the page size and page-table mapping before translating an address. For IPC, distinguish communication semantics and protection boundaries instead of listing mechanisms from memory.
#
# ## Worked practice
#
# ### Practice A — Round Robin
#
# Take four jobs with supplied arrival times and CPU bursts. Draw the timeline for a stated quantum, compute completion time, turnaround time, and waiting time, then repeat after changing the quantum.
#
# ### Practice B — Race and mutex
#
# Given two threads incrementing a shared counter, identify the non-atomic read/modify/write sequence. Explain the possible incorrect result, then describe exactly what the mutex changes.
#
# ### Practice C — Address translation
#
# Given a page size and page table, split a virtual address into page number and offset, translate it to a physical frame, and reconstruct the physical address.
#
# ## Exercises
#
# 1. Draw process/thread resources and identify what is shared.
# 2. Build FCFS, SJF, and Round Robin schedules and compute waiting/turnaround times.
# 3. Reproduce a race on a shared counter and explain why ordering matters.
# 4. Protect the same critical section with a mutex and explain the guarantee.
# 5. Analyse a resource-allocation graph for deadlock.
# 6. Distinguish deadlock from starvation and blocking.
# 7. Translate sample virtual addresses using a supplied page table.
# 8. Calculate page counts and memory overhead from page size.
# 9. Compare file allocation strategies conceptually.
# 10. Explain how IPC mechanisms differ in isolation and communication semantics.
#
# ## Timed assessment
#
# Complete one scheduling table, one synchronization/deadlock problem, one virtual-address translation, and one short filesystem/IPC explanation in 75 minutes. Show every intermediate calculation.
#
# ## Exam traps
#
# Never confuse waiting time with turnaround time. Do not call every blocked task deadlocked. State the scheduling policy and time quantum. State the page size and mappings before translating addresses. In synchronization answers, identify the shared resource and critical section before naming a primitive.
#
# ## Mastery gate
#
# Solve one scheduling table, one race/deadlock problem, one address-translation problem, and one filesystem/IPC explanation from unseen data. Your explanations must justify the result, not merely state it.
