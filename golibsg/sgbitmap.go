package golibsg

/*
#include <stdlib.h>
#include <sg/sgbitmap.h>
#cgo LDFLAGS: -lsg
*/
import "C"

import (
	"unsafe"
	"fmt"
)

type SgBitmap struct {
	ptr unsafe.Pointer

	// We need to keep this alive as the bmp is collected in the file's
	// destructor
	parent *SgFile
}

func createBitmap(ptr *C.struct_SgBitmap, parent *SgFile) *SgBitmap {
	bmp := &SgBitmap{unsafe.Pointer(ptr), parent}
	return bmp
}

func (b *SgBitmap) Filename() string {
	return C.GoString(C.sg_get_bitmap_filename((*C.struct_SgBitmap)(b.ptr)))
}

func (b *SgBitmap) Comment() string {
	return C.GoString(C.sg_get_bitmap_comment((*C.struct_SgBitmap)(b.ptr)))
}

func (b *SgBitmap) Width() uint32 {
	return uint32(C.sg_get_bitmap_width((*C.struct_SgBitmap)(b.ptr)))
}

func (b *SgBitmap) Height() uint32 {
	return uint32(C.sg_get_bitmap_height((*C.struct_SgBitmap)(b.ptr)))
}

func (b *SgBitmap) Images() ([]*SgImage, error) {
	n := int(C.sg_get_bitmap_image_count((*C.struct_SgImage)(b.ptr)))
	imgs := make([]*SgImage, n)
	for i := 0; i < n; i++ {
		img := C.sg_get_bitmap_image((*C.struct_SgImage)(b.ptr), C.int(i))
		if img == nil {
			return []*SgImage{}, fmt.Errorf("Image %v was nil", i)
		}
		imgs[i] = createImage(img, b.parent)
	}
	return imgs, nil
}