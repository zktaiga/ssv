package goclient

import (
	"time"

	"github.com/attestantio/go-eth2-client/spec"
	"github.com/attestantio/go-eth2-client/spec/altair"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	"github.com/pkg/errors"
)

// GetSyncMessageBlockRoot returns beacon block root for sync committee
func (gc *goClient) GetSyncMessageBlockRoot(slot phase0.Slot) (phase0.Root, spec.DataVersion, error) {
	// Wait a 1/3 into the slot.
	gc.waitOneThirdOrValidBlock(slot)

	reqStart := time.Now()
	root, err := gc.client.BeaconBlockRoot(gc.ctx, "head")
	if err != nil {
		return phase0.Root{}, DataVersionNil, err
	}
	if root == nil {
		return phase0.Root{}, DataVersionNil, errors.New("root is nil")
	}

	metricsSyncCommitteeDataRequest.Observe(time.Since(reqStart).Seconds())

	return *root, spec.DataVersionAltair, nil
}

// SubmitSyncMessage submits a signed sync committee msg
func (gc *goClient) SubmitSyncMessage(msg *altair.SyncCommitteeMessage) error {
	if err := gc.client.SubmitSyncCommitteeMessages(gc.ctx, []*altair.SyncCommitteeMessage{msg}); err != nil {
		return err
	}
	return nil
}
