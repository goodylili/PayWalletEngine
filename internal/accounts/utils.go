package accounts

import (
	crypto "crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/google/uuid"
	"math/big"
	"math/rand"
	"time"
)

// GenerateAccountNumber generates a unique 10-digit account number
func GenerateAccountNumber() (int64, error) {
	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Int63n(1e10)
	uid, err := uuid.NewRandomFromReader(crypto.Reader)
	if err != nil {
		return 0, err
	}
	hash := sha256.Sum256([]byte(fmt.Sprintf("%d%v", randomInt, uid.String())))
	hashInt := big.NewInt(0)
	hashInt.SetBytes(hash[:])
	accountNumber := hashInt.Mod(hashInt, big.NewInt(1e10)).Int64()
	return accountNumber, nil
}
