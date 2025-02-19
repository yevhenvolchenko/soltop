/*
Copyright © 2025 Yevhen Volchenko eugene.volchenko@gmail.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package solana

import (
	"strings"

	"github.com/gagliardetto/solana-go/rpc"
)

func formatSliceAsStirng(validators []rpc.VoteAccountsResult, nodeNumber int) string {
	// Form and print output string
	// Expected format ("NodePubkey", ..."NodePubkey")
	validatorsString := "("
	for _, v := range validators[:min(len(validators), nodeNumber)] {
		validatorsString = validatorsString + "'" + v.NodePubkey.String() + "',"
	}
	validatorsString = strings.TrimRight(validatorsString, ", ")
	validatorsString = validatorsString + ")"
	return (validatorsString)
}
