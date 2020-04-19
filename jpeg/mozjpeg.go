package jpeg

//#include <jpeglib.h>
//
//static void go_init_color_space(jpeg_compress_struct *cinfo)
//{
//  cinfo->in_color_space = JCS_RGB;
//}
import "C"

import (
	"unsafe"
)

type JPEGCompress struct {
	CJPEGCompressStruct *C.struct_jpeg_compress_struct
}

func (jpeg *JPEGCompress) SetDefaults() {
	C.go_init_color_space((*C.struct_jpeg_compress_struct)(jpeg.CJPEGCompressStruc))
	C.jpeg_set_defaults(jpeg.CJPEGCompressStruct)
}

func (jpeg *JPEGCompress) JPEGMemDest() ([]byte, int, error) {
	var outBuffer []byte
	var outSize uint32
	C.jpeg_mem_dest((*C.struct_jpeg_compress_struct)(jpeg.CJPEGCompressStruct), unsafe.Pointer(uintptr(outBuffer)), &(C.ulong(outSize)))
	return outBuffer, outSize, nil
}
