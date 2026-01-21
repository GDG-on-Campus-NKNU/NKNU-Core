# Detect OS
ifeq ($(OS),Windows_NT)
  SHELL := cmd.exe
  MKDIR = if not exist bin md bin
  SETVAR = set
  SEP = &&
  CLANG_SUFFIX = .cmd
  HOST_TAG = windows-x86_64
  RM = -rmdir /S /Q
else
  SHELL := /bin/bash
  MKDIR = mkdir -p bin
  SETVAR = export
  SEP = ;
  CLANG_SUFFIX =
  HOST_TAG = linux-x86_64
  RM = rm -rf
endif

# ========== Rules ==========

clean:
	$(RM) bin

create-output-dir:
	$(MKDIR)

compile-sso:
	# DEPRECATED
	gomobile bind -o ./bin/sso.aar -target=android -androidapi 21 -v ./sso

build-android-arm64:
	$(SETVAR) GOOS=android$(SEP) \
	$(SETVAR) GOARCH=arm64$(SEP) \
	$(SETVAR) CGO_ENABLED=1$(SEP) \
	$(SETVAR) CC=$(ANDROID_NDK_HOME)/toolchains/llvm/prebuilt/$(HOST_TAG)/bin/aarch64-linux-android21-clang$(CLANG_SUFFIX)$(SEP) \
	$(SETVAR) CXX=$(ANDROID_NDK_HOME)/toolchains/llvm/prebuilt/$(HOST_TAG)/bin/aarch64-linux-android21-clang++$(CLANG_SUFFIX)$(SEP) \
	go build -o ./bin/android_arm64_nknu_core.so -buildmode=c-shared

build-android-arm:
	$(SETVAR) GOOS=android$(SEP) \
	$(SETVAR) GOARCH=arm$(SEP) \
	$(SETVAR) CGO_ENABLED=1$(SEP) \
	$(SETVAR) CC=$(ANDROID_NDK_HOME)/toolchains/llvm/prebuilt/$(HOST_TAG)/bin/armv7a-linux-androideabi21-clang$(CLANG_SUFFIX)$(SEP) \
	$(SETVAR) CXX=$(ANDROID_NDK_HOME)/toolchains/llvm/prebuilt/$(HOST_TAG)/bin/armv7a-linux-androideabi21-clang++$(CLANG_SUFFIX)$(SEP) \
	go build -o ./bin/android_arm_nknu_core.so -buildmode=c-shared

build-android-x86:
	$(SETVAR) GOOS=android$(SEP) \
	$(SETVAR) GOARCH=386$(SEP) \
	$(SETVAR) CGO_ENABLED=1$(SEP) \
	$(SETVAR) CC=$(ANDROID_NDK_HOME)/toolchains/llvm/prebuilt/$(HOST_TAG)/bin/i686-linux-android21-clang$(CLANG_SUFFIX)$(SEP) \
	$(SETVAR) CXX=$(ANDROID_NDK_HOME)/toolchains/llvm/prebuilt/$(HOST_TAG)/bin/i686-linux-android21-clang++$(CLANG_SUFFIX)$(SEP) \
	go build -o ./bin/android_x86_nknu_core.so -buildmode=c-shared

build-android-x86_64:
	$(SETVAR) GOOS=android$(SEP) \
	$(SETVAR) GOARCH=amd64$(SEP) \
	$(SETVAR) CGO_ENABLED=1$(SEP) \
	$(SETVAR) CC=$(ANDROID_NDK_HOME)/toolchains/llvm/prebuilt/$(HOST_TAG)/bin/x86_64-linux-android21-clang$(CLANG_SUFFIX)$(SEP) \
	$(SETVAR) CXX=$(ANDROID_NDK_HOME)/toolchains/llvm/prebuilt/$(HOST_TAG)/bin/x86_64-linux-android21-clang++$(CLANG_SUFFIX)$(SEP) \
	go build -o ./bin/android_x86_64_nknu_core.so -buildmode=c-shared

build-windows-x86_64:
	$(SETVAR) GOOS=windows$(SEP) \
	$(SETVAR) GOARCH=amd64$(SEP) \
	$(SETVAR) CGO_ENABLED=1$(SEP) \
	$(SETVAR) CC=x86_64-w64-mingw32-gcc$(SEP) \
	go build -o ./bin/windows_x86_64_nknu_core.dll -buildmode=c-shared ./api


# Build all targets
compile: clean create-output-dir build-android-arm64 build-android-arm build-android-x86_64 build-android-x86 build-windows-x86_64

# Test
test-sso:
	go test ./sso

test-school-bus-schedule:
	go test ./schoolbusschedule

test: test-sso test-school-bus-schedule