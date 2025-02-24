package keeper

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/palomachain/paloma/x/valset/types"
	"github.com/palomachain/paloma/x/valset/types/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestJailingInactiveValidators(t *testing.T) {
	k, ms, ctx := newValsetKeeper(t)
	ctx = ctx.WithBlockHeight(1000)

	valBuild := func(id int, toBeJailed bool) (*mocks.StakingValidatorI, sdk.ValAddress) {
		val := sdk.ValAddress(fmt.Sprintf("validator_%d", id))
		vali := mocks.NewStakingValidatorI(t)
		stake := sdk.DefaultPowerReduction
		ms.StakingKeeper.On("Validator", mock.Anything, val).Return(vali)
		vali.On("IsJailed").Return(false)
		vali.On("GetConsensusPower", k.powerReduction).Return(stake.Int64())
		vali.On("IsBonded").Return(true)
		vali.On("GetOperator").Return(val)
		vali.On("GetStatus").Return(stakingtypes.Bonded)
		consAddr := sdk.ConsAddress(val)
		if toBeJailed {
			vali.On("GetConsAddr").Return(consAddr, nil)
			ms.SlashingKeeper.On("Jail", mock.Anything, consAddr)
			ms.SlashingKeeper.On("JailUntil", mock.Anything, mock.Anything, mock.Anything)
		} else {
			err := k.KeepValidatorAlive(ctx.WithBlockHeight(ctx.BlockHeight()-(cJailingDefaultKeepAliveBlockHeight/2)), val, "v1.4.0")
			require.NoError(t, err)
		}
		return vali, val
	}

	v1, a1 := valBuild(1, false)
	v2, a2 := valBuild(2, false)
	v3, a3 := valBuild(3, true)
	v4, a4 := valBuild(4, true)
	newUnjailed, _ := valBuild(5, false)

	k.unjailedSnapshotStore(ctx).Set([]byte(cUnjailedSnapshotStoreKey), bytes.Join([][]byte{a1, a2, a3, a4}, []byte(",")))
	ms.StakingKeeper.On("IterateValidators", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		callback := args.Get(1).(func(int64, stakingtypes.ValidatorI) bool)
		callback(0, v1)
		callback(0, v2)
		callback(0, v3)
		callback(0, v4)
		callback(0, newUnjailed)
	}).Return(false)

	err := k.JailInactiveValidators(ctx)
	require.NoError(t, err)
}

func TestCanAcceptKeepAlive(t *testing.T) {
	valAddr := sdk.ValAddress("testvalidator")

	testcases := []struct {
		name               string
		inputPigeonVersion string
		setup              func(services mockedServices)
		expectedErr        error
	}{
		{
			name:               "validator not found",
			inputPigeonVersion: "v1.4.0",
			setup: func(services mockedServices) {
				services.StakingKeeper.On("Validator", mock.Anything, mock.Anything).Return(nil)
			},
			expectedErr: ErrValidatorWithAddrNotFound.Format(valAddr.String()),
		},
		{
			name:               "pigeon version too low",
			inputPigeonVersion: "v1.3.100",
			setup: func(services mockedServices) {
				services.StakingKeeper.On("Validator", mock.Anything, mock.Anything).Return(mocks.NewStakingValidatorI(t))
			},
			expectedErr: ErrValidatorPigeonOutOfDate.Format(valAddr.String(), "v1.3.100", "v1.4.0"),
		},
		{
			name:               "pigeon version equal, validator found",
			inputPigeonVersion: "v1.4.0",
			setup: func(services mockedServices) {
				services.StakingKeeper.On("Validator", mock.Anything, mock.Anything).Return(mocks.NewStakingValidatorI(t))
			},
		},
		{
			name:               "pigeon version major higher, validator found",
			inputPigeonVersion: "v2.0.0",
			setup: func(services mockedServices) {
				services.StakingKeeper.On("Validator", mock.Anything, mock.Anything).Return(mocks.NewStakingValidatorI(t))
			},
		},
		{
			name:               "pigeon version patch higher, validator found",
			inputPigeonVersion: "v1.4.1",
			setup: func(services mockedServices) {
				services.StakingKeeper.On("Validator", mock.Anything, mock.Anything).Return(mocks.NewStakingValidatorI(t))
			},
		},
		{
			name:               "pigeon version minor higher, validator found",
			inputPigeonVersion: "v1.5.0",
			setup: func(services mockedServices) {
				services.StakingKeeper.On("Validator", mock.Anything, mock.Anything).Return(mocks.NewStakingValidatorI(t))
			},
		},
	}

	for _, tt := range testcases {
		t.Run(tt.name, func(t *testing.T) {
			k, ms, ctx := newValsetKeeper(t)
			tt.setup(ms)

			actualErr := k.CanAcceptKeepAlive(ctx, valAddr, tt.inputPigeonVersion)
			require.Equal(t, tt.expectedErr, actualErr)
		})
	}
}

func TestUpdateGracePeriod(t *testing.T) {
	k, ms, ctx := newValsetKeeper(t)
	ctx = ctx.WithBlockHeight(1000)

	valBuild := func(id int) (*mocks.StakingValidatorI, sdk.ValAddress) {
		val := sdk.ValAddress(fmt.Sprintf("validator_%d", id))
		vali := mocks.NewStakingValidatorI(t)
		vali.On("IsJailed").Return(false)
		vali.On("GetOperator").Return(val)
		return vali, val
	}

	putSn := func(ctx sdk.Context, vals ...[]byte) {
		k.unjailedSnapshotStore(ctx).Set([]byte(cUnjailedSnapshotStoreKey), bytes.Join(vals, []byte(",")))
	}

	v1, a1 := valBuild(1)
	v2, a2 := valBuild(2)
	v3, a3 := valBuild(3)

	teardown := func() {
		k.unjailedSnapshotStore(ctx).Delete([]byte(cUnjailedSnapshotStoreKey))
		for _, v := range []sdk.ValAddress{a1, a2, a3} {
			k.gracePeriodStore(ctx).Delete(v)
		}
	}

	ms.StakingKeeper.On("IterateValidators", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		callback := args.Get(1).(func(int64, stakingtypes.ValidatorI) bool)
		callback(0, v1)
		callback(0, v2)
		callback(0, v3)
	}).Return(false)

	t.Run("with no elements in past snapshot", func(t *testing.T) {
		t.Cleanup(teardown)
		k.UpdateGracePeriod(ctx)
		for _, v := range []sdk.ValAddress{a1, a2, a3} {
			x := int64(sdk.BigEndianToUint64(k.gracePeriodStore(ctx).Get(v)))
			require.Equal(t, ctx.BlockHeight(), x)
		}
	})

	t.Run("with all elements in past snapshot", func(t *testing.T) {
		t.Cleanup(teardown)
		putSn(ctx, a1, a2, a3)
		k.UpdateGracePeriod(ctx)
		for _, v := range []sdk.ValAddress{a1, a2, a3} {
			x := k.gracePeriodStore(ctx).Get(v)
			require.Nil(t, x)
		}
	})

	t.Run("with some elements in past snapshot", func(t *testing.T) {
		t.Cleanup(teardown)
		putSn(ctx, a1, a3)
		k.UpdateGracePeriod(ctx)
		for _, v := range []sdk.ValAddress{a1, a3} {
			x := k.gracePeriodStore(ctx).Get(v)
			require.Nil(t, x)
		}
		x := int64(sdk.BigEndianToUint64(k.gracePeriodStore(ctx).Get(a2)))
		require.Equal(t, ctx.BlockHeight(), x)
	})
}

func TestJailBackoff(t *testing.T) {
	k, ms, ctx := newValsetKeeper(t)
	ctx = ctx.WithBlockHeight(1000).WithBlockTime(time.Date(2020, 1, 1, 12, 30, 0, 0, time.UTC))

	valBuild := func(id int) (*mocks.StakingValidatorI, sdk.ValAddress) {
		valAddr := sdk.ValAddress(fmt.Sprintf("validator_%d", id))
		val := mocks.NewStakingValidatorI(t)
		val.On("IsJailed").Return(false)
		val.On("IsBonded").Return(true)
		val.On("GetConsensusPower", k.powerReduction).Return(int64(10000))
		return val, valAddr
	}

	v1, _ := valBuild(1)
	v2, _ := valBuild(2)
	v3, _ := valBuild(3)

	ms.StakingKeeper.On("IterateValidators", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		callback := args.Get(1).(func(int64, stakingtypes.ValidatorI) bool)
		callback(0, v1)
		callback(0, v2)
		callback(0, v3)
	}).Return(false).Maybe()

	// Should reset jail backoff if recovered
	t.Run("with non-recovering validator", func(t *testing.T) {
		for i, v := range jailSentences {
			t.Run(fmt.Sprintf("[%d] Paloma should increase the jail sentence with every occurrence", i), func(t *testing.T) {
				val, valAddr := valBuild(10 + i)
				consAddr := sdk.ConsAddress(valAddr)
				ms.StakingKeeper.On("Validator", mock.Anything, valAddr).Return(val)
				val.On("IsBonded").Unset()
				val.On("GetConsensusPower", k.powerReduction).Unset()
				val.On("GetConsensusPower", k.powerReduction).Return(int64(100))
				val.On("GetConsAddr").Return(consAddr, nil)
				ms.SlashingKeeper.On("Jail", mock.Anything, consAddr)
				ms.SlashingKeeper.On("JailUntil", mock.Anything, consAddr, ctx.BlockTime().Add(v))
				if i > 0 {
					k.jailLog.Set(ctx, consAddr, &types.JailRecord{
						Address:  consAddr.Bytes(),
						Duration: jailSentences[i-1],
						JailedAt: ctx.BlockTime().Add(-1 * jailSentences[i-1]),
					})
				}
				err := k.Jail(ctx, valAddr, "foobar")
				require.NoError(t, err)
			})
		}

		t.Run("Paloma should cap the jail sentence at max sentence level", func(t *testing.T) {
			val, valAddr := valBuild(30)
			consAddr := sdk.ConsAddress(valAddr)
			ms.StakingKeeper.On("Validator", mock.Anything, valAddr).Return(val)
			val.On("IsBonded").Unset()
			val.On("GetConsensusPower", k.powerReduction).Unset()
			val.On("GetConsensusPower", k.powerReduction).Return(int64(100))
			val.On("GetConsAddr").Return(consAddr, nil)
			ms.SlashingKeeper.On("Jail", mock.Anything, consAddr)
			ms.SlashingKeeper.On("JailUntil", mock.Anything, consAddr, ctx.BlockTime().Add(jailSentences[len(jailSentences)-1]))
			k.jailLog.Set(ctx, consAddr, &types.JailRecord{
				Address:  consAddr.Bytes(),
				Duration: jailSentences[len(jailSentences)-1],
				JailedAt: ctx.BlockTime().Add(-1 * jailSentences[len(jailSentences)-1]),
			})
			err := k.Jail(ctx, valAddr, "foobar")
			require.NoError(t, err)
		})
	})

	// Should reset jail backoff if recovered
	t.Run("with recovered validator", func(t *testing.T) {
		for i := range jailSentences {
			t.Run(fmt.Sprintf("[%d] Paloma should reset the jail sentence, no matter the last duration", i), func(t *testing.T) {
				val, valAddr := valBuild(50 + i)
				consAddr := sdk.ConsAddress(valAddr)
				ms.StakingKeeper.On("Validator", mock.Anything, valAddr).Return(val)
				val.On("IsBonded").Unset()
				val.On("GetConsensusPower", k.powerReduction).Unset()
				val.On("GetConsensusPower", k.powerReduction).Return(int64(100))
				val.On("GetConsAddr").Return(consAddr, nil)
				ms.SlashingKeeper.On("Jail", mock.Anything, consAddr)
				ms.SlashingKeeper.On("JailUntil", mock.Anything, consAddr, ctx.BlockTime().Add(jailSentences[0]))
				if i > 0 {
					k.jailLog.Set(ctx, consAddr, &types.JailRecord{
						Address:  consAddr.Bytes(),
						Duration: jailSentences[i-1],
						JailedAt: ctx.BlockTime().Add(-1 * jailSentences[i-1]).Add(-1 * time.Hour * 48),
					})
				}
				err := k.Jail(ctx, valAddr, "foobar")
				require.NoError(t, err)
			})
		}
	})
}
