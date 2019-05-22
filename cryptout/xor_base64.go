package cryptout

import "encoding/base64"

func XOREncrypt(data, key []byte) string {

	keyLen := len(key)
	bytesLen := len(data)

	for i := 0; i < bytesLen; i++ {
		data[i] = data[i] ^ key[i%keyLen]
	}

	return base64.StdEncoding.EncodeToString(data)
}

func XORDecrypt(value string, key []byte) ([]byte, error) {

	bytes, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return nil, err
	}

	keyLen := len(key)
	bytesLen := len(bytes)

	for i := 0; i < bytesLen; i++ {
		bytes[i] = bytes[i] ^ key[i%keyLen]
	}

	return bytes, nil
}
