package solana

import (
	"math/rand"
	"testing"
	"time"

	sol "github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
)

// Function to generate random strings
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

// Helper function to generate sol.PublicKey from random string
func generateRandomPublicKey() sol.PublicKey {
	randomString := generateRandomString(32) // 32 characters for public key
	var publicKey sol.PublicKey
	publicKey.Set(randomString)
	return publicKey
}

func TestFormatSliceAsString(t *testing.T) {
	// Sample validators with random PublicKey
	validators := []rpc.VoteAccountsResult{
		{NodePubkey: generateRandomPublicKey()},
		{NodePubkey: generateRandomPublicKey()},
		{NodePubkey: generateRandomPublicKey()},
	}

	tests := []struct {
		name           string
		validators     []rpc.VoteAccountsResult
		nodeNumber     int
		expectedResult string
	}{
		{
			name:           "Test with 3 validators and nodeNumber 2",
			validators:     validators,
			nodeNumber:     2,
			expectedResult: "('" + validators[0].NodePubkey.String() + "','" + validators[1].NodePubkey.String() + "')",
		},
		{
			name:           "Test with 3 validators and nodeNumber 5 (more than available)",
			validators:     validators,
			nodeNumber:     5,
			expectedResult: "('" + validators[0].NodePubkey.String() + "','" + validators[1].NodePubkey.String() + "','" + validators[2].NodePubkey.String() + "')",
		},
		{
			name:           "Test with 0 validators",
			validators:     []rpc.VoteAccountsResult{},
			nodeNumber:     1,
			expectedResult: "()",
		},
		{
			name:           "Test with 1 validator and nodeNumber 1",
			validators:     validators[:1], // only 1 validator
			nodeNumber:     1,
			expectedResult: "('" + validators[0].NodePubkey.String() + "')",
		},
		{
			name:           "Test with nodeNumber 0 (should return empty parentheses)",
			validators:     validators,
			nodeNumber:     0,
			expectedResult: "()",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Run the function
			result := formatSliceAsStirng(tt.validators, tt.nodeNumber)

			// Check if the result matches the expected output
			if result != tt.expectedResult {
				t.Errorf("expected %s, got %s", tt.expectedResult, result)
			}
		})
	}
}
