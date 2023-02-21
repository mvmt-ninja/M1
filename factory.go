// Copyright (C) 2023, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package vm

import (
	"github.com/ava-labs/avalanchego/snow"
	"github.com/ava-labs/avalanchego/vms"

	"github.com/ava-labs/indexvm/controller"
)

var _ vms.Factory = &Factory{}

type Factory struct{}

func (*Factory) New(*snow.Context) (interface{}, error) {
	return controller.New(), nil
}
