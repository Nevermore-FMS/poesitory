# Poesitory CLI

The Poesitory CLI is a tool to push and pull plugins to and from Poesitory

## Installation

### Prebuilt Binaries

Poesitory currently provides pre-built binaries for the following:

- macOS (Darwin) for 386 and amd64 architectures
- Linux for 386 and amd64 architectures
- Windows for 386 and amd64 architectures

Download the appropriate version for your platform from the [Releases page](https://github.com/Nevermore-FMS/poesitory/releases)

Once downloaded, the binary can be run from anywhere. You don’t need to install it into a global location.

If you would like Poesitory to be accessible from anywhere, you should put it in a directory and add that directory to your PATH
- On Linux, this is most likely `/usr/local/bin`
- On Windows, you should create a directory like `C:\Program Files\poesitory\bin`. Then, In PowerShell or your preferred CLI, add the poesitory.exe executable to your PATH by navigating to `C:\Program Files\poesitory\bin` (or the location of your poesitory.exe file) and use the command `set PATH=%PATH%;C:\Program Files\poesitory\bin`. You may need to run this as administrator.

### Build and install with Go

Poesitory may be compiled from source wherever the Go toolchain can run; e.g., on other operating systems such as DragonFly BSD, OpenBSD, Plan 9, Solaris, and others. See https://golang.org/doc/install/source for the full set of supported combinations of target operating systems and compilation architectures.

To build and install, run `go install github.com/Nevermore-FMS/poesitory/cli/poesitory@latest`

The Poesitory CLI tool will now be available in your PATH.

## Documentation

See the [Docs Folder](docs)