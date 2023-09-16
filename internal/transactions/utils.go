package transactions

import (
	crypto "crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/google/uuid"
	"math/big"
	"math/rand"
	"strings"
	"time"
)

// GenerateTransactionID generates a unique 16-digit account number
func GenerateTransactionID() (int64, error) {
	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Int63n(1e16)
	uid, err := uuid.NewRandomFromReader(crypto.Reader)
	if err != nil {
		return 0, err
	}
	hash := sha256.Sum256([]byte(fmt.Sprintf("%d%v", randomInt, uid.String())))
	hashInt := big.NewInt(0)
	hashInt.SetBytes(hash[:])
	accountNumber := hashInt.Mod(hashInt, big.NewInt(1e16)).Int64()

	return accountNumber, nil
}

func GenerateTransactionRef(transactionID int64) (string, error) {
	// Convert the transaction ID to a namespace UUID
	namespaceUUID := uuid.NewSHA1(uuid.NameSpaceDNS, []byte(fmt.Sprintf("%d", transactionID)))

	// Generate a UUID using the namespace
	u := uuid.NewSHA1(namespaceUUID, []byte("transaction"))

	// Convert UUID to string and take the first 16 characters
	ref := strings.ReplaceAll(u.String(), "-", "")[:16]

	return ref, nil
}
