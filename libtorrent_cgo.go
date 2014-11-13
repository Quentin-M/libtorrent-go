package libtorrent

// #cgo pkg-config: libtorrent-rasterbar openssl
// #cgo darwin LDFLAGS: -lm -lstdc++
// #cgo linux LDFLAGS: -lm -lstdc++ -ldl -lrt
// #cgo windows CXXFLAGS: -DIPV6_TCLASS=39 -O0 -g
// #cgo windows LDFLAGS: -static-libgcc -static-libstdc++
import "C"
