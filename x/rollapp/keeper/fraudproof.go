package keeper

import (
	fraudtypes "github.com/cosmos/cosmos-sdk/baseapp"
)

func (k *Keeper) VerifyFraudProof(fp fraudtypes.FraudProof) error {
	err := k.fraudProofVerifier.InitFromFraudProof(&fp)
	if err != nil {
		return err
	}

	err = k.fraudProofVerifier.VerifyFraudProof(&fp)
	if err != nil {
		return err
	}

	return nil
}