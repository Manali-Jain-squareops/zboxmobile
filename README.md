# zboxmobile

This is a wrapper over zboxcore package of [gosdk](https://github.com/0chain/gosdk), This can be used to perform zbox(storage) related operations on the chain.
Since this is written in Go, It's pretty easy and fast to use it directly with the mobile applications.

### Mobile Builds (iOS and Android) ###
- gosdk can be build to use on Mobile platforms iOS and Android using gomobile.
- Xcode Command Line Tools is required to build SDK for iOS.
- Android studio with NDK is required to build SDK for Android
- Run below command for the first time to setup gomobile environment

        make setup-gomobile

- Use below commands in the root folder of the repo to build Mobile SDK

        For iOS and Android:
                make build-mobilesdk IOS=1 ANDROID=1
        For iOS only:
                make build-mobilesdk IOS=1
        For Android only:
                make build-mobilesdk ANDROID=1
        For Mac(xcode 12+) only:
                make build-mobilesdk MAC=1

### Notes
- For iOS: If you are already using the SDK and  getting the older version after updating, then you need to manually remove the SDK from the location where is was already placed (Most probably it will be Project Folder > SDK > zboxmobile.framework).
- For Mac: Since XCode 12 you can't import ios library/framework into mac project (xcode 11 still allowing). Before compiling to Mac, be sure to complie bls-go-binary with xcode 12 script. Follow up with external guide: /tools/xcode12-build.md

### FAQ ###

- [How to install GO on any platform](https://golang.org/doc/install)
- [How to install different version of GO](https://golang.org/doc/install#extra_versions)
- [How to use go mod](https://blog.golang.org/using-go-modules)
- [What is gomobile](https://godoc.org/golang.org/x/mobile/cmd/gomobile)
- [About XCode](https://developer.apple.com/xcode/)
- [Android Studio](https://developer.android.com/studio)
- [Android NDK](https://developer.android.com/ndk/)