package pkg

import (
	"encoding/base64"
	"golang.org/x/sys/windows"
	"unsafe"
)

func DecryptData(data []byte) ([]byte, error) {
	// Decode base64 string and remove first 5 bytes
	decoded, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		return nil, err
	}
	decoded = decoded[5:]

	// Create a new DataBlob with the decrypted data
	dataBlob := windows.DataBlob{
		Size: uint32(len(decoded)),
		Data: &decoded[0],
	}

	var outBlob windows.DataBlob

	// Call CryptUnprotectData

	err = windows.CryptUnprotectData(&dataBlob, nil, nil, 0, nil, 0, &outBlob)
	if err != nil {
		return nil, err
	}

	// Copy the decrypted data to a new slice and return it
	outData := make([]byte, outBlob.Size)
	copy(outData, (*[1 << 30]byte)(unsafe.Pointer(outBlob.Data))[:outBlob.Size])
	return outData, nil
}
