package e2e_test

import (
	"context"
	"testing"

	"github.com/omni-network/omni/lib/cchain"
	"github.com/omni-network/omni/lib/cchain/provider"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	e2e "github.com/cometbft/cometbft/test/e2e/pkg"

	"github.com/stretchr/testify/require"
)

// TestApprovedAttestations tests that all halo instances contain approved attestations
// for at least half of all the source chain blocks.
func TestApprovedAttestations(t *testing.T) {
	t.Parallel()
	testNode(t, func(t *testing.T, network netconf.Network, node *e2e.Node, portals []Portal) {
		t.Helper()
		client, err := node.Client()
		require.NoError(t, err)
		cprov := provider.NewABCIProvider(client, network.ID, nil)

		ctx := context.Background()
		for _, portal := range portals {
			height, err := portal.Client.BlockNumber(ctx)
			require.NoError(t, err)

			atts, err := fetchAllAtts(ctx, cprov, portal.Chain.ID, portal.Chain.DeployHeight)
			require.NoError(t, err)

			totalBlocks := height - portal.Chain.DeployHeight
			require.GreaterOrEqual(t, len(atts), int(totalBlocks/2)) // Assert that at least half of the blocks are approved
		}

		// Ensure at least one (genesis) consensus chain approved attestation
		consChain, ok := network.OmniConsensusChain()
		require.True(t, ok)
		atts, err := fetchAllAtts(ctx, cprov, consChain.ID, consChain.DeployHeight)
		require.NoError(t, err)
		require.NotEmpty(t, len(atts))
	})
}

// TestApprovedAttestations tests that all halo instances contain approved attestations
// for at least half of all the source chain blocks.
func TestApprovedValUpdates(t *testing.T) {
	t.Parallel()
	testNode(t, func(t *testing.T, network netconf.Network, node *e2e.Node, portals []Portal) {
		t.Helper()
		ctx := context.Background()

		// See if this node has a validator update
		var hasUpdate bool
		var power int64
		for _, powers := range node.Testnet.ValidatorUpdates {
			for n, p := range powers {
				if n.Name == node.Name {
					hasUpdate = true
					power = p
				}
			}
		}

		if !hasUpdate {
			return // Nothing to test
		}

		client, err := node.Client()
		require.NoError(t, err)
		cprov := provider.NewABCIProvider(client, network.ID, nil)

		addr, err := k1util.PubKeyToAddress(node.PrivvalKey.PubKey())
		require.NoError(t, err)

		t.Logf("Looking for validator update: %s, %d", addr.Hex(), power)

		consChain, ok := network.OmniConsensusChain()
		require.True(t, ok)

		valSetID := consChain.DeployHeight
		for {
			vals, ok, err := cprov.ValidatorSet(ctx, valSetID)
			require.NoError(t, err)
			require.True(t, ok, "Validator update not found in any set: node=%s, valSetID=%v", node.Name, valSetID)

			for _, val := range vals {
				t.Logf("Got validator update set=%d: %s, %d", valSetID, val.Address.Hex(), val.Power)
				if val.Address == addr && val.Power == power {
					return // Validator update found and confirmed.
				}
			}

			// If update not found, there must be more sets
			valSetID++
		}
	})
}

func fetchAllAtts(ctx context.Context, cprov cchain.Provider, chainID, from uint64) ([]xchain.Attestation, error) {
	var resp []xchain.Attestation
	for {
		atts, err := cprov.AttestationsFrom(ctx, chainID, from)
		if err != nil {
			return nil, err
		}
		if len(atts) == 0 { // No more attestation to fetch
			break
		}
		resp = append(resp, atts...)

		// Update the from height to fetch the next batch of attestation
		from = atts[len(atts)-1].BlockHeight + 1
	}

	return resp, nil
}
