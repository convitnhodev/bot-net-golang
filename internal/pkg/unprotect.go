package pkg

import (
	"github.com/go-ole/go-ole"
	"golang.org/x/sys/windows"
	"unsafe"
)

func UnprotectData(secretKey []byte) ([]byte, error) {
	dataBlobIn := windows.DataBlob{
		Data: &secretKey[0],
		Size: uint32(len(secretKey)),
	}
	var dataBlobOut windows.DataBlob
	if err := ole.CoInitializeEx(0, ole.COINIT_APARTMENTTHREADED|ole.COINIT_SPEED_OVER_MEMORY); err != nil {
		return nil, err
	}
	defer ole.CoUninitialize()

	ret, _, err := windows.NewLazySystemDLL("Crypt32.dll").NewProc("CryptUnprotectData").Call(
		uintptr(unsafe.Pointer(&dataBlobIn)),
		0,
		0,
		0,
		0,
		0,
		uintptr(unsafe.Pointer(&dataBlobOut)),
	)
	if ret == 0 {
		return nil, err
	}

	defer windows.LocalFree(windows.Handle(uintptr(unsafe.Pointer(dataBlobOut.Data))))
	//return windows.ByteSlice(dataBlobOut.Data[:dataBlobOut.Size]), nil
	outData := make([]byte, dataBlobOut.Size)

	return outData, nil
}
