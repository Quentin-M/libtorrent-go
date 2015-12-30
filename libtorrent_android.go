// +build android

package libtorrent

// #cgo pkg-config: libtorrent-rasterbar openssl
// #cgo android CXXFLAGS: -I/opt/android-toolchain-arm/include/libtorrent -I/opt/android-toolchain-arm/include -Wno-deprecated-declarations
// #cgo android LDFLAGS: -lm -ldl -L/opt/android-toolchain-arm/lib

import "C"
