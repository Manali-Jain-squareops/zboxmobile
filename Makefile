# Paths
ROOT_DIR:=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
GOSDK_PATH:=$(ROOT_DIR)/../gosdk
ZBOXMOBILE_PATH := $(GOSDK_PATH)/zboxmobile
PATCHES:=$(GOSDK_PATH)/patches
GOSDK_EXP_ROOT := $(GOPATH)/src/github.com/0chain
GOSDK_EXP_PATH := $(GOSDK_EXP_ROOT)/gosdk
ZBOXMOBILE_EXP_PATH := $(GOSDK_EXP_ROOT)/zboxmobile

# Go packages
PKG_0CHAIN	:= github.com/0chain
PKG_GOSDK := $(PKG_0CHAIN)/gosdk
PKG_ZCNCORE := $(PKG_GOSDK)/zcncore
PKG_ZBOXSDK := $(PKG_GOSDK)/zboxcore/sdk
PKG_ZBOXMOBILE := $(PKG_0CHAIN)/zboxmobile/zbox

PKG_EXPORTS := $(PKG_ZBOXMOBILE) $(PKG_ZCNCORE)

# Output dir
OUTDIR := $(ROOT_DIR)/out
IOSMOBILESDKDIR     := $(OUTDIR)/0chainiosmobilesdk
ANDROIDMOBILESDKDIR := $(OUTDIR)/0chainandroidmobilesdk
MACSDKDIR	:= $(OUTDIR)/0chainmacsdk
IOSBINNAME 		:= zbox.framework
ANDROIDBINNAME	:= zbox.aar


BLS_LIB_BASE_PATH=$(GOPATH)/src/github.com/herumi

default: build-mobilesdk

$(IOSMOBILESDKDIR):
	$(shell mkdir -p $(IOSMOBILESDKDIR)/lib)

$(ANDROIDMOBILESDKDIR):
	$(shell mkdir -p $(ANDROIDMOBILESDKDIR)/lib)

$(GOSDK_EXP_PATH):
	@echo "gosdk is not in GOPATH. Creating softlink for $(GOSDK_PATH)..."
ifneq ($(GOPATH), )
	$(shell ln -sf $(GOSDK_PATH) $(GOSDK_EXP_ROOT))
endif

$(ZBOXMOBILE_EXP_PATH):
	@echo "zboxmobile is not in GOPATH. Creating softlink for $(ROOT_DIR)..."
ifneq ($(GOPATH), )
	$(shell ln -sf $(ROOT_DIR) $(ZBOXMOBILE_EXP_PATH))
endif

build-mobilesdk: $(GOSDK_EXP_PATH) $(ZBOXMOBILE_EXP_PATH)  $(if $(IOS),$(IOSMOBILESDKDIR),) $(if $(ANDROID),$(ANDROIDMOBILESDKDIR),)
ifeq ($(filter-out undefined,$(foreach v, IOS ANDROID,$(origin $(v)))),)
	@echo ""
	@echo "Usage:"
	@echo '   For iOS and Android: make build-mobilesdk IOS=1 ANDROID=1'
	@echo '   For iOS only: make build-mobilesdk IOS=1'
	@echo '   For Android only: make build-mobilesdk ANDROID=1'
	@echo '   For Mac(xcode 12) only: make build-mobilesdk MAC=1'
endif
ifneq ($(MAC),)
	@echo "Building MAC framework. Please wait..."
	@@./tools/gomobile bind -ldflags="-s -w -v -X main.version=1.0" -target=ios/amd64 -tags ios -o $(MACSDKDIR)/$(IOSBINNAME) $(PKG_EXPORTS)
	@echo "   $(IOSMOBILESDKDIR)/$(IOSBINNAME). - [OK]"
endif
ifneq ($(IOS),)
	@echo "Building iOS framework. Please wait..."
	@@gomobile bind -ldflags="-s -w" -target=ios -o $(IOSMOBILESDKDIR)/$(IOSBINNAME) $(PKG_EXPORTS)
	@echo "   $(IOSMOBILESDKDIR)/$(IOSBINNAME). - [OK]"
endif
ifneq ($(ANDROID),)
	@echo "Building Android framework. Please wait..."
	@gomobile bind -target=android/arm64,android/amd64 -ldflags=-extldflags=-Wl,-soname,libgojni.so -o $(ANDROIDMOBILESDKDIR)/$(ANDROIDBINNAME) $(PKG_EXPORTS)
	@echo "   $(ANDROIDMOBILESDKDIR)/$(ANDROIDBINNAME). - [OK]"
endif
	@echo ""
