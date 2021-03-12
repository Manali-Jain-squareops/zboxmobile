
### XCode 12 build process for Mac Application ###

- You have to remove/rename go.mod file before building for Mac
- `rm -r vendor`
- `go mod vendor`
- `mkdir work && cd work`
- `git clone https://github.com/herumi/mcl && \ git clone https://github.com/herumi/bls && \ git clone https://github.com/herumi/bls-go-binary`
- `cd bls-go-binary` in bls-go-binary add to `bls-go-binary/Makefile` script from `tools/herumi-build`
- run build with `make ios-mac`. As result you will see compiled binary in folder `ios-mac/libbls384_256.a`
- go back to root folder 
- `rm -r vendor/github.com/herumi/bls-go-binary`
- `cp -r work/bls-go-binary/ vendor/github.com/herumi/bls-go-binary/`
- `mv go.mod go.mod.tmp`
- After all steps done build zboxmobile with `make build-mobilesdk MAC=1`
- `mv go.mod.tmp go.mod`


### XCode 12 building libbls384_256.a for Mac application ###

- Follow up steps below to compile ios-mac library
- Copy-past libbls384_256.a to Mac Application project