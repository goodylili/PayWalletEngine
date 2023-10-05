package transactions

import "github.com/google/uuid"

var base62Chars = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")

func EncodeBase62(data []byte) string {
	result := make([]byte, 0, len(data)*2)
	num := uint64(0)
	bits := uint(0)

	for _, b := range data {
		num = (num << 8) | uint64(b)
		bits += 8

		for bits >= 6 {
			bits -= 6
			index := (num >> bits) & 63
			if index < 0 || index >= uint64(len(base62Chars)) {
				// Just for sanity, this shouldn't happen
				return "INVALID"
			}
			result = append(result, base62Chars[index])
		}
	}

	if bits > 0 {
		num <<= 6 - bits
		index := num & 63
		if index < 0 || index >= uint64(len(base62Chars)) {
			// Just for sanity, this shouldn't happen
			return "INVALID"
		}
		result = append(result, base62Chars[index])
	}

	return string(result)
}

func GenerateTransactionRef() (string, error) {
	// Generate a random UUID
	u, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	// Encode the UUID into base62
	encoded := EncodeBase62(u[:])

	return encoded, nil
}
