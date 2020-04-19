// Package jpeg implements a JPEG image encoder based on mozjpeg
package jpeg

import (
	"image"
	"io"
)

// DefaultQuality is the default quality encoding parameter.
const DefaultQuality = 75

type TrellisOptimization int

const (
	// Disable trellis optimization of DC coefficients
	DisableTrellisOptimization TrellisOptimization = iota
	// Tune trellis optimization for PSNR
	TunePSNRTrellisOptimization
	// Tune trellis optimization for PSNR-HVS (default)
	TuneHVSPSNRTrellisOptimization
	// Tune trellis optimization for SSIM
	TuneSSIMTrellisOptimization
	// Tune trellis optimization for MS-SSIM
	TuneMSSSIMTrellisOptimization
)

// Options are the encoding parameters.
type Options struct {
	// Quality ranges from 1 to 100 inclusive, higher is better.
	Quality int
	// Optimize Huffman table (smaller file, but slow compression, enabled by default)
	Optimize bool
	// Create monochrome JPEG file
	GrayScale bool
	// Trellis optimization of DC coefficients (default PSNR-HVS)
	TrellisOptimization TrellisOptimization
}

// Encode writes the Image m to w in JPEG 4:2:0 baseline format with the given
// options. Default parameters are used if a nil *Options is passed.
func Encode(w io.Writer, m image.Image, o *Options) error {

}
