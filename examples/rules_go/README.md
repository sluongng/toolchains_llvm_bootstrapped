# rules_go with cgo Example

This example demonstrates how to use the `toolchains_llvm_bootstrapped` toolchain with `rules_go` to build a Go binary that has cgo dependencies.

## Overview

This example includes:
- A C library (`greet.c` and `greet.h`) with functions that can be called from Go
- A Go binary (`main.go`) that uses cgo to call the C functions
- Proper Bazel configuration to link the Go code with the C library

## Structure

```
.
├── BUILD.bazel       # Bazel build definitions
├── MODULE.bazel      # Bazel module dependencies
├── README.md         # This file
├── greet.h           # C header file
├── greet.c           # C implementation
├── main.go           # Go main file with cgo
├── .bazelrc          # Bazel configuration
└── .bazelversion     # Bazel version specification
```

## How it Works

1. **C Library**: The `cc_library` target compiles `greet.c` into a library that exports two functions:
   - `greet(const char* name)`: Returns a greeting message
   - `add_numbers(int a, int b)`: Adds two integers

2. **Go Binary**: The `go_binary` target embeds a `go_library` that:
   - Has `cgo = True` to enable cgo
   - Uses `cdeps = [":greet"]` to link against the C library
   - Imports the C functions using cgo's `import "C"` syntax

3. **LLVM Toolchain**: Both the C compilation and the cgo compilation use the hermetic LLVM toolchain provided by `toolchains_llvm_bootstrapped`, ensuring consistent cross-platform builds.

## Building

To build the example:

```bash
cd examples/rules_go
bazel build :cgo_example
```

## Running

To run the example:

```bash
cd examples/rules_go
bazel run :cgo_example
```

Expected output:
```
Hello from C, Go Developer!
C says: 15 + 27 = 42

Successfully called C functions from Go using cgo!
This demonstrates rules_go working with the LLVM bootstrapped toolchain.
```

## Cross-Compilation

You can cross-compile this example to different platforms using the `--platforms` flag:

```bash
# Cross-compile to Linux ARM64
bazel build :cgo_example --platforms=//platforms:linux_aarch64

# Cross-compile to Linux AMD64
bazel build :cgo_example --platforms=//platforms:linux_amd64

# Cross-compile to macOS ARM64 (from macOS)
bazel build :cgo_example --platforms=//platforms:darwin_arm64
```

## Key Configuration

The `.bazelrc` file includes important settings for cgo:

- `--incompatible_enable_cc_toolchain_resolution`: Enables proper C++ toolchain resolution for cgo
- `--@io_bazel_rules_go//go/config:pure=false`: Disables pure Go mode, enabling cgo

## Notes

- This example requires the LLVM bootstrapped toolchain to be properly registered (done in `MODULE.bazel`)
- The cgo code uses the same LLVM toolchain for C compilation, ensuring compatibility
- The C library is compiled with the hermetic LLVM toolchain, not the system compiler
