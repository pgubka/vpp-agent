// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package vpp1901_test

import (
	"testing"

	ifModel "github.com/ligato/vpp-agent/api/models/vpp/interfaces"
	"github.com/ligato/vpp-agent/plugins/vpp/binapi/vpp1901/interfaces"
	. "github.com/onsi/gomega"
)

func TestSetRxMode(t *testing.T) {
	ctx, ifHandler := ifTestSetup(t)
	defer ctx.TeardownTestCtx()

	ctx.MockVpp.MockReply(&interfaces.SwInterfaceSetRxModeReply{})

	err := ifHandler.SetRxMode(1, &ifModel.Interface_RxModeSettings{
		RxMode:       ifModel.Interface_RxModeSettings_DEFAULT,
		QueueId:      1,
		QueueIdValid: 2,
	})

	Expect(err).To(BeNil())
	vppMsg, ok := ctx.MockChannel.Msg.(*interfaces.SwInterfaceSetRxMode)
	Expect(ok).To(BeTrue())
	Expect(vppMsg.SwIfIndex).To(BeEquivalentTo(1))
	Expect(vppMsg.Mode).To(BeEquivalentTo(4))
	Expect(vppMsg.QueueID).To(BeEquivalentTo(1))
	Expect(vppMsg.QueueIDValid).To(BeEquivalentTo(2))
}

func TestSetRxModeError(t *testing.T) {
	ctx, ifHandler := ifTestSetup(t)
	defer ctx.TeardownTestCtx()

	ctx.MockVpp.MockReply(&interfaces.SwInterfaceSetRxMode{})

	err := ifHandler.SetRxMode(1, &ifModel.Interface_RxModeSettings{
		RxMode:       ifModel.Interface_RxModeSettings_DEFAULT,
		QueueId:      1,
		QueueIdValid: 2,
	})

	Expect(err).ToNot(BeNil())
}

func TestSetRxModeRetval(t *testing.T) {
	ctx, ifHandler := ifTestSetup(t)
	defer ctx.TeardownTestCtx()

	ctx.MockVpp.MockReply(&interfaces.SwInterfaceSetRxModeReply{
		Retval: 1,
	})

	err := ifHandler.SetRxMode(1, &ifModel.Interface_RxModeSettings{
		RxMode:       ifModel.Interface_RxModeSettings_DEFAULT,
		QueueId:      1,
		QueueIdValid: 2,
	})

	Expect(err).ToNot(BeNil())
}
