
### XCode 12 build process for Mac Application ###

- You have to remove/rename go.mod file before building for Mac
- `go vendor`
- `mkdir work && cd work`
- `git clone https://github.com/herumi/mcl && \ git clone https://github.com/herumi/bls && \ git clone https://github.com/herumi/bls-go-binary`
- `cd bls-go-binary` in bls-go-binary add to Makefile script from herumi-build
- run build with `make ios-mac`. As result you will see compiled binary in folder ios-mac/libbls384_256.a
- cp ios-mac/libbls384_256.a /vendor/herumi/bls-go-library/bls/lib/ios
- After all steps done build zboxmobile with `make build-mobilesdk MAC=1`


### XCode 12 building libbls384_256.a for Mac application ###

- Follow up steps below to compile ios-mac library
- Copy-past libbls384_256.a to Mac Application project