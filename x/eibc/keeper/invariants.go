package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	commontypes "github.com/dymensionxyz/dymension/v3/x/common/types"
	"github.com/dymensionxyz/dymension/v3/x/eibc/types"
)

const (
	demandOrderCountInvariantName = "demand-order-count"
)

// RegisterInvariants registers the bank module invariants
func RegisterInvariants(ir sdk.InvariantRegistry, k Keeper) {
	ir.RegisterRoute(types.ModuleName, "demand-order-count", DemandOrderCountInvariant(k))
	ir.RegisterRoute(types.ModuleName, "underlying-packet-exist", UnderlyingPacketExistInvariant(k))
}

// AllInvariants runs all invariants of the x/streamer module.
func AllInvariants(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		res, stop := DemandOrderCountInvariant(k)(ctx)
		if stop {
			return res, stop
		}
		res, stop = UnderlyingPacketExistInvariant(k)(ctx)
		if stop {
			return res, stop
		}
		return "", false
	}
}

func DemandOrderCountInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var (
			broken bool
			msg    string
		)
		allDemandOrders, err := k.ListAllDemandOrders(ctx)
		if err != nil {
			msg += fmt.Sprintf("list all demand orders failed: %v\n", err)
			broken = true
		}
		pendingDemandOrders, err := k.ListDemandOrders(ctx, commontypes.Status_PENDING)
		if err != nil {
			msg += fmt.Sprintf("list pending demand orders failed: %v\n", err)
			broken = true
		}
		revertedDemandOrders, err := k.ListDemandOrders(ctx, commontypes.Status_REVERTED)
		if err != nil {
			msg += fmt.Sprintf("list reverted demand orders failed: %v\n", err)
			broken = true
		}
		finalizedDemandOrders, err := k.ListDemandOrders(ctx, commontypes.Status_FINALIZED)
		if err != nil {
			msg += fmt.Sprintf("list finalized demand orders failed: %v\n", err)
			broken = true
		}
		// Validate the count of demand orders is equal to the sum of demand orders in all statuses
		if len(allDemandOrders) != len(pendingDemandOrders)+len(revertedDemandOrders)+len(finalizedDemandOrders) {
			msg += fmt.Sprintf("demand orders count mismatch: all(%d) != pending(%d) + reverted(%d) + finalized(%d)\n",
				len(allDemandOrders), len(pendingDemandOrders), len(revertedDemandOrders), len(finalizedDemandOrders))
			broken = true
		}
		return sdk.FormatInvariant(types.ModuleName, demandOrderCountInvariantName, msg), broken
	}
}

func UnderlyingPacketExistInvariant(k Keeper) sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		var (
			broken bool
			msg    string
		)
		allDemandOrders, err := k.ListAllDemandOrders(ctx)
		if err != nil {
			msg += fmt.Sprintf("list all demand orders failed: %v\n", err)
			broken = true
		}
		for _, demandOrder := range allDemandOrders {
			// Get the underlying packet for the demand order
			_, err := k.GetRollappPacket(ctx, demandOrder.TrackingPacketKey)
			if err != nil {
				msg += fmt.Sprintf("underlying packet for demand order %s not found: %v\n", demandOrder.Id, err)
				broken = true
				break
			}
		}
		return sdk.FormatInvariant(types.ModuleName, "underlying-packet-exist", msg), broken
	}
}
