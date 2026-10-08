# S1 COURSE EXERCISE - OS1
# Study the prompts below, then complete the exercises in this file.

# # FSEGS LIG — Systèmes d'exploitation 1
#
# **Semester:** S1  
# **Attached engineering lesson:** 009 — Files, Paths, Arguments, and Permissions
#
# ## What you must be able to do
#
# Explain the role of an operating system and reason about processes, process states, system calls, files, file descriptors, permissions, environment variables, processes/threads, signals, pipes, basic scheduling, virtual memory, and synchronization/deadlock concepts.
#
# ## Course sequence
#
# 1. OS role, kernel/user mode, system-call boundary.
# 2. Processes, process image, PID, parent/child, lifecycle and states.
# 3. Files and directories, metadata, file descriptors, paths, permissions.
# 4. Shell/process composition, pipes, redirection, environment.
# 5. Signals, process control, IPC basics.
# 6. Threads, shared state, race condition and critical section.
# 7. CPU scheduling and simple performance measures.
# 8. Virtual memory and deadlock foundations.
#
# ## Teaching notes
#
# A process is a running program instance with its own execution context and address space. A thread is an execution path inside a process; threads share the process's address space and many resources.
#
# System calls are the controlled boundary through which user programs request privileged kernel services. Typical examples include opening files, creating processes, reading/writing file descriptors, and changing process behavior.
#
# A Unix file descriptor is a small process-local integer referring to an open file or stream. `0`, `1`, and `2` are conventionally stdin, stdout, and stderr.
#
# Permissions are normally represented for owner, group, and others, with read/write/execute bits. Convert octal permissions to symbolic bits before deciding whether a particular access is allowed.
#
# A race condition occurs when the result depends on an uncontrolled ordering of concurrent operations. A deadlock is a stronger state in which processes/threads wait indefinitely for resources held by one another under the stated assumptions.
#
# ## Worked examples
#
# ### Example A — Permission decoding
#
# Decode `-rwxr-x---` into owner/group/other permissions and then into octal. Predict whether a process belonging only to the file's group can modify it.
#
# ### Example B — Process tree
#
# Start a shell command that creates two children. Draw the parent/child relation and identify which process waits for which child.
#
# ### Example C — Pipe direction
#
# For `producer | consumer`, state which process writes the pipe, which reads it, and what EOF means when the writer closes the descriptor.
#
# ## Exercises
#
# 1. Draw the lifecycle states of a process and describe each transition.
# 2. Explain the difference between a program, process, and thread.
# 3. Decode five permission strings and predict read/write/execute access for different users.
# 4. Use `ps`, `/proc`, and `ls -l` to collect evidence about one running process and one file.
# 5. Explain stdin/stdout/stderr and redirection using concrete shell commands.
# 6. Build a process tree and identify parent/child relationships.
# 7. Explain what happens when a parent exits before a child and what happens when a child exits before its parent.
# 8. Given four jobs with arrival and CPU-burst times, construct FCFS and Round Robin timelines and compute waiting time.
# 9. Identify whether each concurrent scenario is a race, deadlock, starvation, or none.
# 10. Draw the virtual-address → page → frame translation at a conceptual level.
#
# ## Systems lab
#
# Use WSL/Linux. Capture the exact commands and outputs in the lesson's `EVIDENCE.md`. Do not run permissions experiments as root.
#
# Tasks: process observation; file descriptor inspection; permission matrix; pipeline construction; one signal experiment; one simple race demonstration; one scheduling table.
#
# ## Exam preparation
#
# Expect definition questions, state diagrams, permission decoding, process-tree interpretation, shell command questions, scheduling tables, race/deadlock classification, and short explanations of kernel/user boundaries.
#
# ## Mastery gate
#
# Without notes, explain process versus thread, syscall boundary, file descriptor, permissions, race condition, and deadlock. Then solve one scheduling table and one permission problem from fresh data and verify one claim in Linux.
