
### XCode 12 build process for Mac Application ###

## Go 1.13.5
## GOPATH & GOROOT point to same directory

export GOROOT=/usr/local/Cellar/go/
export GOPATH=/usr/local/Cellar/go/

## Step by step guide

- You have to remove/rename go.mod file before building for Mac
- `rm -r vendor`
- `go mod vendor`
- `mkdir work && cd work`
- `git clone https://github.com/herumi/mcl && \ git clone https://github.com/herumi/bls && \ git clone --recurse-submodules https://github.com/herumi/bls-go-binary`
- `cd bls-go-binary` in bls-go-binary add to `bls-go-binary/Makefile` script from `tools/herumi-build`
- run build with `make ios-mac`. As result you will see compiled binary in folder `ios-mac/libbls384_256.a`
- copy file `ios-mac/libbls384_256.a` into your go folder `/usr/local/Cellar/go/1.16/libexec/src/github.com/herumi/bls-go-binary/bls/lib/ios/libbls384_256.a`
- go back to root folder
- `rm -r vendor/github.com/herumi/bls-go-binary`
- `cp -r work/bls-go-binary/ vendor/github.com/herumi/bls-go-binary/`
- `mv go.mod go.mod.tmp`
- After all steps done build zboxmobile with `make build-mobilesdk MAC=1`
- `mv go.mod.tmp go.mod`


### XCode 12 building libbls384_256.a for Mac application ###

- Follow up steps below to compile ios-mac library
- Copy-past libbls384_256.a to Mac Application project

### Troubleshooting

```
ype-checking package "github.com/0chain/zboxmobile/zbox" failed (/Users/artem.bogomaz/documents/git2/0chain/zboxmobile/zbox/allocation.go:9:2: could not import github.com/0chain/gosdk/zboxcore/sdk (type-checking package "github.com/0chain/gosdk/zboxcore/sdk" failed (/Users/artem.bogomaz/Documents/GIT2/0chain/go/gosdk/zboxcore/sdk/allocation.go:18:2: could not import github.com/0chain/gosdk/zboxcore/client (type-checking package "github.com/0chain/gosdk/zboxcore/client" failed (/Users/artem.bogomaz/Documents/GIT2/0chain/go/gosdk/zboxcore/client/entity.go:6:2: could not import github.com/0chain/gosdk/core/zcncrypto (type-checking package "github.com/0chain/gosdk/core/zcncrypto" failed (/Users/artem.bogomaz/Documents/GIT2/0chain/go/gosdk/core/zcncrypto/bls0chain.go:12:2: could not import github.com/tyler-smith/go-bip39 (type-checking package "github.com/tyler-smith/go-bip39" failed (/Users/artem.bogomaz/Documents/GIT2/0chain/go/pkg/mod/github.com/tyler-smith/go-bip39@v1.0.0/bip39.go:12:2: could not import github.com/tyler-smith/go-bip39/wordlists (cannot find package "github.com/tyler-smith/go-bip39/wordlists" in any of:
	/usr/local/go/src/github.com/tyler-smith/go-bip39/wordlists (from $GOROOT)
	/Users/USER/Documents/GIT2/0chain/go/src/github.com/tyler-smith/go-bip39/wordlists (from $GOPATH)))))))))))
```

To resolve it run:

```go get golang.org/x/mobile/cmd/gomobile```
```gomobile init```

## Troubleshooting 2 - IOS compile

```building for macOS, but linking in object file built for iOS Simulator, for architecture x86_64```

Open compiled herumi directory in GOPATH folder and copy paste from darwin to ios

## Troubleshooting 3 - MAC compile

```ld: in /var/folders/l7/sc5qmw916hv9tm87tyc9q1bh0000gp/T/gomobile-work-767411242/pkg/mod/github.com/herumi/bls-go-binary@v0.0.0-20191119080710-898950e1a520/bls/lib/ios/libbls384_256.a(bls_c384_256.o), building for macOS, but linking in object file built for iOS Simulator, for architecture x86_64
clang: error: linker command failed with exit code 1 (use -v to see invocation)```

go mod vendor
