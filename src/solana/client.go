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
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/gagliardetto/solana-go/rpc"
	"golang.org/x/time/rate"
)

var _ Client = (*client)(nil)

// Client interface for an Solana RPC client
type Client interface {
	GetValidators(nodeNumber int)
}

// Client implementation for Solana RPC
type client struct {
	requester rpc.Client
}

func NewClient(uri string) Client {
	return &client{
		requester: *rpc.NewWithCustomRPCClient(rpc.NewWithLimiter(uri, rate.Every(time.Second), 5)),
	}
}

func (c *client) GetValidators(nodeNumber int) {
	//TODO: implement context
	//TODO: implement logging
	validators, err := c.requester.GetVoteAccounts(context.TODO(), &rpc.GetVoteAccountsOpts{})
	if err != nil {
		fmt.Println("Cannot connect to RPC server")
		panic(err)
	}

	// Sort the slice by age in descending order
	sort.Slice(validators.Current, func(i, j int) bool {
		return validators.Current[i].ActivatedStake > validators.Current[j].ActivatedStake
	})
	fmt.Println(formatSliceAsStirng(validators.Current, nodeNumber))
}
