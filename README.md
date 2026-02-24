# NKNU-Core
> 🤖 Provides core functionality for [NKNU-Connect](https://github.com/GDG-on-Campus-NKNU/NKNU-Connect) and other NKNU-related school projects.

NKNU-Core is a Go-based library designed to provide reusable core functionality, compiled into a shared object (.so) file for use in Android and other related projects. This eliminates the need to repeatedly implement the same features across multiple projects.

[![State-of-the-art Shitcode](https://img.shields.io/static/v1?label=State-of-the-art&message=Shitcode&color=7B5804)](https://github.com/trekhleb/state-of-the-art-shitcode)
[![Go](https://img.shields.io/badge/Go-%2300ADD8.svg?&logo=go&logoColor=white)](#)
[![Android](https://img.shields.io/badge/Android-3DDC84?logo=android&logoColor=white)](#)

## Prerequisites
Before you begin, ensure the following tools are installed:
- Go (version 1.24.0)
- Android NDK (version r26)
- GNU Make (version 4.4.1)

> ⚠️ Higher version may work but have not been tested

### Environment Setup
1. Go:
    - Verify installation: `go version`
2. Android NDK:
    - Set the `ANDROID_NDK_HOME` environment variable:
      ```bash
      export ANDROID_NDK_HOME=/path/to/android-ndk
      ```
3. GNU Make:
    - Verify installation: `make --version`

## Project Structure
```
NKNU-Core/
├── sso/                # NKNU Single Sign-On (SSO) related functionality
├── utils/              # Utility functions and helper tools
├── Makefile            # Build automation script for compiling the library
└── README.md           # Project documentation
```

## Building the Project
The project uses a Makefile to automate the compilation process, generating a shared object (`.so`) file for integration with Android or other platforms.

To compile the library:
```bash
go work init ./api ./schoolbusschedule ./sso ./utils
make compile
```

This command builds the core functionality into a `.so` file, ready for use in dependent projects like [NKNU-Connect](https://github.com/GDG-on-Campus-NKNU/NKNU-Connect).

## Usage
The compiled `.so` file can be integrated into Android projects or other platforms that support native libraries. Refer to the documentation of the consuming project (e.g., [NKNU-Connect](https://github.com/GDG-on-Campus-NKNU/NKNU-Connect)) for specific integration instructions.

## Development Guidelines
To maintain modularity and reusability, each independent feature should be implemented in a separate Go package. For example, if developing a traffic information query feature, create a dedicated `traffic` package. All functions intended for export to other projects must be placed in a file named `api.go` within the respective package. Additionally, in the `main` package's `main.go`, use side effect imports (e.g., `import _ "path/to/package"`) to ensure that the package is included during compilation.

## Testing
To run tests, set the following environment variables for SSO authentication:

- `account`: Your SSO account username
- `password`: Your SSO account password

Example:
```bash
export account=your_sso_username
export password=your_sso_password
```

Then, execute the tests using:
```bash
make test
```