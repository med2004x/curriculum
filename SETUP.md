# Setup the tools needed to start

For Lessons 001 through 024, you need Python 3 and a C compiler with the standard headers. Later modules add Go, PostgreSQL, Rust, and other tools; install them when the path reaches those modules.

## Windows (PowerShell)

This checkout currently has Python, but its available `clang.exe` has no C standard headers and there is no GCC. It cannot build the first C lesson yet. Install a complete compiler toolchain:

1. Install [MSYS2](https://www.msys2.org/) using its official installer.
2. Open the **MSYS2 UCRT64** terminal, update it with `pacman -Syu`, reopen the terminal if it closes, then install the compiler and Make:

   ```sh
   pacman -S --needed mingw-w64-ucrt-x86_64-gcc mingw-w64-ucrt-x86_64-make
   ```

3. From the repository root in PowerShell, add GCC and `make` to your user PATH:

   ```powershell
   powershell.exe -NoProfile -ExecutionPolicy Bypass -File .\tools\setup_windows_make_path.ps1
   gcc --version
   make --version
   python --version
   ```

4. Open a new PowerShell terminal so it loads the updated PATH. Then return to [Start Today](START-TODAY.md).

If compilation says `stdio.h` is missing, the compiler is installed without its C development headers or the wrong compiler is being selected. `clang --version` alone does not prove that a usable C toolchain is installed.

## Linux / WSL

Install GCC, GNU Make, and Python 3 using your distribution's package manager (for example, `sudo apt install build-essential make python3`). Confirm with `gcc --version`, `make --version`, and `python3 --version`, then follow the shell commands in the lesson.

## Later tools

The first lesson is the only required setup now. Each later module should state its own tool before the learner reaches it. Do not install PostgreSQL, QEMU, or database project dependencies just to begin Lesson 001.
