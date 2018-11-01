package schemes

import "image/color"

// AlphaFire is a gradient color scheme from transparent through red
// to yellow then white.
var AlphaFire []color.Color

func init() {
	AlphaFire = []color.Color{
		color.NRGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0xfa, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0xf5, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0xf0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0xeb, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0xe6, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0xe1, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0xdc, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0xd7, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0xd1, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0xcc, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0xc7, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0xc2, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0xbd, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0xb8, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0xb3, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0xae, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0xa8, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0xa3, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0x9e, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0x99, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0x94, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0x8f, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0x8a, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0x85, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0x7f, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0x7a, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0x75, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0x70, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0x6b, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0x66, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0x61, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0x5c, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0x57, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0x51, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0x4c, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0x47, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0x42, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0x3d, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0x38, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0x33, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0x2e, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0x28, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0x23, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0x1e, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0x19, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0x14, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0xf, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0xa, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0x5, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xff, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xfb, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xf7, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xf3, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xee, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xea, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xe6, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xe2, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xdd, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xd9, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xd5, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xd1, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xcc, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xc8, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xc4, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xbf, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xbb, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xb7, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xb3, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xae, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xaa, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xa6, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xa2, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x9d, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x99, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x95, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x91, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x8c, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x88, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x84, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x7f, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x7b, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x77, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x73, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x6e, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x6a, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x66, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x62, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x5d, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x59, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x55, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x51, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x4c, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x48, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x44, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x3f, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x3b, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x37, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x33, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x2e, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x2a, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x26, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x22, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x1d, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x19, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x15, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x11, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0xc, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x8, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x4, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xff, G: 0x0, B: 0x0, A: 0xff},
		color.NRGBA{R: 0xfe, G: 0x1, B: 0x1, A: 0xff},
		color.NRGBA{R: 0xfd, G: 0x2, B: 0x2, A: 0xff},
		color.NRGBA{R: 0xfc, G: 0x3, B: 0x3, A: 0xfe},
		color.NRGBA{R: 0xfb, G: 0x4, B: 0x4, A: 0xfe},
		color.NRGBA{R: 0xf9, G: 0x5, B: 0x5, A: 0xfe},
		color.NRGBA{R: 0xf9, G: 0x6, B: 0x6, A: 0xfd},
		color.NRGBA{R: 0xf7, G: 0x7, B: 0x7, A: 0xfd},
		color.NRGBA{R: 0xf6, G: 0x8, B: 0x8, A: 0xfd},
		color.NRGBA{R: 0xf5, G: 0x9, B: 0x9, A: 0xfc},
		color.NRGBA{R: 0xf4, G: 0xb, B: 0xb, A: 0xfc},
		color.NRGBA{R: 0xf3, G: 0xc, B: 0xc, A: 0xfc},
		color.NRGBA{R: 0xf2, G: 0xd, B: 0xd, A: 0xfb},
		color.NRGBA{R: 0xf1, G: 0xe, B: 0xe, A: 0xfb},
		color.NRGBA{R: 0xef, G: 0xf, B: 0xf, A: 0xfb},
		color.NRGBA{R: 0xef, G: 0x10, B: 0x10, A: 0xfa},
		color.NRGBA{R: 0xed, G: 0x11, B: 0x11, A: 0xfa},
		color.NRGBA{R: 0xec, G: 0x12, B: 0x12, A: 0xfa},
		color.NRGBA{R: 0xeb, G: 0x13, B: 0x13, A: 0xf9},
		color.NRGBA{R: 0xea, G: 0x15, B: 0x15, A: 0xf9},
		color.NRGBA{R: 0xe9, G: 0x16, B: 0x16, A: 0xf8},
		color.NRGBA{R: 0xe8, G: 0x17, B: 0x17, A: 0xf8},
		color.NRGBA{R: 0xe7, G: 0x18, B: 0x18, A: 0xf8},
		color.NRGBA{R: 0xe6, G: 0x19, B: 0x19, A: 0xf7},
		color.NRGBA{R: 0xe5, G: 0x1a, B: 0x1a, A: 0xf7},
		color.NRGBA{R: 0xe2, G: 0x1b, B: 0x1b, A: 0xf7},
		color.NRGBA{R: 0xe2, G: 0x1d, B: 0x1d, A: 0xf6},
		color.NRGBA{R: 0xe0, G: 0x1e, B: 0x1e, A: 0xf6},
		color.NRGBA{R: 0xdf, G: 0x20, B: 0x20, A: 0xf6},
		color.NRGBA{R: 0xde, G: 0x21, B: 0x21, A: 0xf5},
		color.NRGBA{R: 0xdd, G: 0x22, B: 0x22, A: 0xf5},
		color.NRGBA{R: 0xdc, G: 0x23, B: 0x23, A: 0xf5},
		color.NRGBA{R: 0xdb, G: 0x24, B: 0x24, A: 0xf4},
		color.NRGBA{R: 0xda, G: 0x25, B: 0x25, A: 0xf4},
		color.NRGBA{R: 0xd8, G: 0x26, B: 0x26, A: 0xf4},
		color.NRGBA{R: 0xd7, G: 0x28, B: 0x28, A: 0xf3},
		color.NRGBA{R: 0xd5, G: 0x29, B: 0x29, A: 0xf3},
		color.NRGBA{R: 0xd5, G: 0x2b, B: 0x2b, A: 0xf2},
		color.NRGBA{R: 0xd3, G: 0x2c, B: 0x2c, A: 0xf2},
		color.NRGBA{R: 0xd2, G: 0x2d, B: 0x2d, A: 0xf2},
		color.NRGBA{R: 0xd1, G: 0x2e, B: 0x2e, A: 0xf1},
		color.NRGBA{R: 0xd0, G: 0x2f, B: 0x2f, A: 0xf1},
		color.NRGBA{R: 0xcf, G: 0x30, B: 0x30, A: 0xf1},
		color.NRGBA{R: 0xcd, G: 0x32, B: 0x32, A: 0xf0},
		color.NRGBA{R: 0xcc, G: 0x33, B: 0x33, A: 0xf0},
		color.NRGBA{R: 0xca, G: 0x34, B: 0x34, A: 0xf0},
		color.NRGBA{R: 0xca, G: 0x35, B: 0x35, A: 0xef},
		color.NRGBA{R: 0xc8, G: 0x37, B: 0x37, A: 0xef},
		color.NRGBA{R: 0xc7, G: 0x38, B: 0x38, A: 0xef},
		color.NRGBA{R: 0xc5, G: 0x3a, B: 0x3a, A: 0xee},
		color.NRGBA{R: 0xc4, G: 0x3b, B: 0x3b, A: 0xee},
		color.NRGBA{R: 0xc2, G: 0x3c, B: 0x3c, A: 0xee},
		color.NRGBA{R: 0xc2, G: 0x3d, B: 0x3d, A: 0xed},
		color.NRGBA{R: 0xc1, G: 0x3e, B: 0x3e, A: 0xed},
		color.NRGBA{R: 0xbf, G: 0x3f, B: 0x3f, A: 0xed},
		color.NRGBA{R: 0xbe, G: 0x41, B: 0x41, A: 0xec},
		color.NRGBA{R: 0xbc, G: 0x43, B: 0x43, A: 0xec},
		color.NRGBA{R: 0xbc, G: 0x44, B: 0x44, A: 0xeb},
		color.NRGBA{R: 0xba, G: 0x45, B: 0x45, A: 0xeb},
		color.NRGBA{R: 0xb9, G: 0x46, B: 0x46, A: 0xeb},
		color.NRGBA{R: 0xb7, G: 0x48, B: 0x48, A: 0xea},
		color.NRGBA{R: 0xb6, G: 0x49, B: 0x49, A: 0xea},
		color.NRGBA{R: 0xb5, G: 0x4a, B: 0x4a, A: 0xea},
		color.NRGBA{R: 0xb4, G: 0x4b, B: 0x4b, A: 0xe9},
		color.NRGBA{R: 0xb3, G: 0x4c, B: 0x4c, A: 0xe9},
		color.NRGBA{R: 0xb0, G: 0x4f, B: 0x4f, A: 0xe9},
		color.NRGBA{R: 0xb0, G: 0x50, B: 0x50, A: 0xe8},
		color.NRGBA{R: 0xae, G: 0x51, B: 0x51, A: 0xe8},
		color.NRGBA{R: 0xad, G: 0x52, B: 0x52, A: 0xe8},
		color.NRGBA{R: 0xab, G: 0x54, B: 0x54, A: 0xe7},
		color.NRGBA{R: 0xaa, G: 0x55, B: 0x55, A: 0xe7},
		color.NRGBA{R: 0xa8, G: 0x56, B: 0x56, A: 0xe7},
		color.NRGBA{R: 0xa8, G: 0x57, B: 0x57, A: 0xe6},
		color.NRGBA{R: 0xa6, G: 0x59, B: 0x59, A: 0xe6},
		color.NRGBA{R: 0xa5, G: 0x5b, B: 0x5b, A: 0xe5},
		color.NRGBA{R: 0xa4, G: 0x5c, B: 0x5c, A: 0xe5},
		color.NRGBA{R: 0xa2, G: 0x5d, B: 0x5d, A: 0xe5},
		color.NRGBA{R: 0xa1, G: 0x5f, B: 0x5f, A: 0xe4},
		color.NRGBA{R: 0x9f, G: 0x60, B: 0x60, A: 0xe4},
		color.NRGBA{R: 0x9e, G: 0x61, B: 0x61, A: 0xe4},
		color.NRGBA{R: 0x9c, G: 0x63, B: 0x63, A: 0xe3},
		color.NRGBA{R: 0x9b, G: 0x64, B: 0x64, A: 0xe3},
		color.NRGBA{R: 0x99, G: 0x65, B: 0x65, A: 0xe3},
		color.NRGBA{R: 0x98, G: 0x68, B: 0x68, A: 0xe2},
		color.NRGBA{R: 0x97, G: 0x69, B: 0x69, A: 0xe2},
		color.NRGBA{R: 0x95, G: 0x6a, B: 0x6a, A: 0xe2},
		color.NRGBA{R: 0x95, G: 0x6c, B: 0x6c, A: 0xe1},
		color.NRGBA{R: 0x92, G: 0x6d, B: 0x6d, A: 0xe1},
		color.NRGBA{R: 0x91, G: 0x6e, B: 0x6e, A: 0xe1},
		color.NRGBA{R: 0x8f, G: 0x6f, B: 0x6f, A: 0xe0},
		color.NRGBA{R: 0x8e, G: 0x71, B: 0x71, A: 0xe0},
		color.NRGBA{R: 0x8c, G: 0x72, B: 0x72, A: 0xe0},
		color.NRGBA{R: 0x8c, G: 0x73, B: 0x73, A: 0xdf},
		color.NRGBA{R: 0x8a, G: 0x76, B: 0x76, A: 0xdf},
		color.NRGBA{R: 0x89, G: 0x77, B: 0x77, A: 0xde},
		color.NRGBA{R: 0x88, G: 0x79, B: 0x79, A: 0xde},
		color.NRGBA{R: 0x85, G: 0x7a, B: 0x7a, A: 0xde},
		color.NRGBA{R: 0x85, G: 0x7b, B: 0x7b, A: 0xdd},
		color.NRGBA{R: 0x82, G: 0x7d, B: 0x7d, A: 0xdd},
		color.NRGBA{R: 0x81, G: 0x7e, B: 0x7e, A: 0xdd},
		color.NRGBA{R: 0x7f, G: 0x7f, B: 0x7f, A: 0xdc},
		color.NRGBA{R: 0x7f, G: 0x7f, B: 0x7f, A: 0xd8},
		color.NRGBA{R: 0x80, G: 0x80, B: 0x80, A: 0xd3},
		color.NRGBA{R: 0x7f, G: 0x7f, B: 0x7f, A: 0xce},
		color.NRGBA{R: 0x80, G: 0x80, B: 0x80, A: 0xc9},
		color.NRGBA{R: 0x7f, G: 0x7f, B: 0x7f, A: 0xc4},
		color.NRGBA{R: 0x7f, G: 0x7f, B: 0x7f, A: 0xc0},
		color.NRGBA{R: 0x7f, G: 0x7f, B: 0x7f, A: 0xbb},
		color.NRGBA{R: 0x7f, G: 0x7f, B: 0x7f, A: 0xb6},
		color.NRGBA{R: 0x80, G: 0x80, B: 0x80, A: 0xb1},
		color.NRGBA{R: 0x7f, G: 0x7f, B: 0x7f, A: 0xac},
		color.NRGBA{R: 0x7f, G: 0x7f, B: 0x7f, A: 0xa8},
		color.NRGBA{R: 0x7f, G: 0x7f, B: 0x7f, A: 0xa3},
		color.NRGBA{R: 0x7f, G: 0x7f, B: 0x7f, A: 0x9e},
		color.NRGBA{R: 0x80, G: 0x80, B: 0x80, A: 0x99},
		color.NRGBA{R: 0x7f, G: 0x7f, B: 0x7f, A: 0x94},
		color.NRGBA{R: 0x7f, G: 0x7f, B: 0x7f, A: 0x90},
		color.NRGBA{R: 0x7f, G: 0x7f, B: 0x7f, A: 0x8b},
		color.NRGBA{R: 0x7f, G: 0x7f, B: 0x7f, A: 0x86},
		color.NRGBA{R: 0x80, G: 0x80, B: 0x80, A: 0x81},
		color.NRGBA{R: 0x7f, G: 0x7f, B: 0x7f, A: 0x7c},
		color.NRGBA{R: 0x7f, G: 0x7f, B: 0x7f, A: 0x78},
		color.NRGBA{R: 0x7e, G: 0x7e, B: 0x7e, A: 0x73},
		color.NRGBA{R: 0x7f, G: 0x7f, B: 0x7f, A: 0x6e},
		color.NRGBA{R: 0x81, G: 0x81, B: 0x81, A: 0x69},
		color.NRGBA{R: 0x7f, G: 0x7f, B: 0x7f, A: 0x64},
		color.NRGBA{R: 0x7f, G: 0x7f, B: 0x7f, A: 0x60},
		color.NRGBA{R: 0x7e, G: 0x7e, B: 0x7e, A: 0x5b},
		color.NRGBA{R: 0x7f, G: 0x7f, B: 0x7f, A: 0x56},
		color.NRGBA{R: 0x7e, G: 0x7e, B: 0x7e, A: 0x51},
		color.NRGBA{R: 0x7f, G: 0x7f, B: 0x7f, A: 0x4c},
		color.NRGBA{R: 0x7f, G: 0x7f, B: 0x7f, A: 0x48},
		color.NRGBA{R: 0x7e, G: 0x7e, B: 0x7e, A: 0x43},
		color.NRGBA{R: 0x7f, G: 0x7f, B: 0x7f, A: 0x3e},
		color.NRGBA{R: 0x7d, G: 0x7d, B: 0x7d, A: 0x39},
		color.NRGBA{R: 0x7f, G: 0x7f, B: 0x7f, A: 0x34},
		color.NRGBA{R: 0x7f, G: 0x7f, B: 0x7f, A: 0x30},
		color.NRGBA{R: 0x7d, G: 0x7d, B: 0x7d, A: 0x2b},
		color.NRGBA{R: 0x7f, G: 0x7f, B: 0x7f, A: 0x26},
		color.NRGBA{R: 0x7c, G: 0x7c, B: 0x7c, A: 0x21},
		color.NRGBA{R: 0x7f, G: 0x7f, B: 0x7f, A: 0x1c},
		color.NRGBA{R: 0x7f, G: 0x7f, B: 0x7f, A: 0x18},
		color.NRGBA{R: 0x79, G: 0x79, B: 0x79, A: 0x13},
		color.NRGBA{R: 0x7f, G: 0x7f, B: 0x7f, A: 0xe},
		color.NRGBA{R: 0x71, G: 0x71, B: 0x71, A: 0x9},
		color.NRGBA{R: 0x7f, G: 0x7f, B: 0x7f, A: 0x4},
	}
}