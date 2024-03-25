package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"gopkg.in/yaml.v2"
)

var _ paramtypes.ParamSet = (*Params)(nil)

var (
	// KeyEpochIdentifier is the key for the epoch identifier
	KeyEpochIdentifier = []byte("EpochIdentifier")
	// KeyTimeoutFee is the key for the timeout fee
	KeyTimeoutFee = []byte("TimeoutFee")
	// KeyErrAckFee is the key for the error acknowledgement fee
	KeyErrAckFee = []byte("TimeoutFee")
)

const (
	defaultEpochIdentifier = "hour"
	defaultTimeoutFee      = "0.0015"
	defaultErrAckFee       = "0.0015" // TODO(danwt): check
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(epochIdentifier string) Params {
	return Params{
		EpochIdentifier: epochIdentifier,
		TimeoutFee:      sdk.MustNewDecFromStr(defaultTimeoutFee),
		ErrackFee:       sdk.MustNewDecFromStr(defaultErrAckFee),
	}
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return NewParams(defaultEpochIdentifier)
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyEpochIdentifier, &p.EpochIdentifier, func(_ interface{}) error { return nil }),
		paramtypes.NewParamSetPair(KeyTimeoutFee, &p.TimeoutFee, validateTimeoutFee),
		paramtypes.NewParamSetPair(KeyErrAckFee, &p.ErrackFee, validateErrAckFee),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {
	// TODO(danwt): need to validate fees again?
	return nil
}

// String implements the Stringer interface.
func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateTimeoutFee(i interface{}) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v.IsNil() {
		return fmt.Errorf("invalid global pool params: %+v", i)
	}
	if v.IsNegative() {
		return ErrNegativeTimeoutFee
	}

	if v.GTE(sdk.OneDec()) {
		return ErrTooMuchTimeoutFee
	}

	return nil
}

func validateErrAckFee(i any) error {
	v, ok := i.(sdk.Dec)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v.IsNil() {
		return fmt.Errorf("invalid global pool params: %+v", i)
	}
	if v.IsNegative() {
		return ErrNegativeErrAckFee
	}

	if v.GTE(sdk.OneDec()) {
		return ErrTooMuchErrAckFee
	}

	return nil
}
