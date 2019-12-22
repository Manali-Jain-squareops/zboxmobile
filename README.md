# zboxmobile
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

### FAQ ###

- [How to install GO on any platform](https://golang.org/doc/install)
- [How to install different version of GO](https://golang.org/doc/install#extra_versions)
- [How to use go mod](https://blog.golang.org/using-go-modules)
- [What is gomobile](https://godoc.org/golang.org/x/mobile/cmd/gomobile)
- [About XCode](https://developer.apple.com/xcode/)
- [Android Studio](https://developer.android.com/studio)
- [Android NDK](https://developer.android.com/ndk/)