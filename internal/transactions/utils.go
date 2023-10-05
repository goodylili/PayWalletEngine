package transactions

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	sequence     int64
	sequenceLock sync.Mutex
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// GenerateTransactionRef generates a unique transaction reference.
func GenerateTransactionRef() (string, error) {
	now := time.Now()

	// Lock to safely increment the sequence number
	sequenceLock.Lock()
	sequence++
	sequenceLock.Unlock()

	// Construct the reference using the timestamp, a random number, and the sequence
	// Format: YYYYMMDDHHMMSSmmm-RRRR-SSS
	// YYYYMMDDHHMMSSmmm: Year, Month, Day, Hour, Minute, Second, Millisecond
	// RRRR: Random 4 digits
	// SSS: Sequence number (can be increased in size if needed)
	reference := fmt.Sprintf("%s-%04d-%03d", now.Format("20060102150405.999"), rand.Intn(9999), sequence%1000)

	return reference, nil
}
