package simulation

import (
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"github.com/cosmos/cosmos-sdk/x/mock/simulation"
)

// SimulateMsgSetWithdrawAddress
func SimulateMsgSetWithdrawAddress(m auth.AccountKeeper, k distribution.Keeper) simulation.Operation {
	handler := distribution.NewHandler(k)
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simulation.Account, event simulation.EventFn) (
		action string, fOp []simulation.FutureOperation, err error) {

		accountOrigin := simulation.RandomAcc(r, accs)
		accountDestination := simulation.RandomAcc(r, accs)
		msg := distribution.NewMsgSetWithdrawAddress(accountOrigin.Address, accountDestination.Address)

		if msg.ValidateBasic() != nil {
			return "", nil, fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		ctx, write := ctx.CacheContext()
		result := handler(ctx, msg)
		if result.IsOK() {
			write()
		}

		event("distribution/MsgSetWithdrawAddress", result.IsOK())

		action = fmt.Sprintf("TestMsgSetWithdrawAddress: ok %v, msg %s", result.IsOK(), msg.GetSignBytes())
		return action, nil, nil
	}
}

// SimulateMsgWithdrawDelegatorRewardsAll
func SimulateMsgWithdrawDelegatorRewardsAll(m auth.AccountKeeper, k distribution.Keeper) simulation.Operation {
	handler := distribution.NewHandler(k)
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simulation.Account, event simulation.EventFn) (
		action string, fOp []simulation.FutureOperation, err error) {

		account := simulation.RandomAcc(r, accs)
		msg := distribution.NewMsgWithdrawDelegatorRewardsAll(account.Address)

		if msg.ValidateBasic() != nil {
			return "", nil, fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		ctx, write := ctx.CacheContext()
		result := handler(ctx, msg)
		if result.IsOK() {
			write()
		}

		event("distribution/MsgWithdrawDelegatorRewardsAll", result.IsOK())

		action = fmt.Sprintf("TestMsgWithdrawDelegatorRewardsAll: ok %v, msg %s", result.IsOK(), msg.GetSignBytes())
		return action, nil, nil
	}
}

// SimulateMsgWithdrawDelegatorReward
func SimulateMsgWithdrawDelegatorReward(m auth.AccountKeeper, k distribution.Keeper) simulation.Operation {
	handler := distribution.NewHandler(k)
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simulation.Account, event simulation.EventFn) (
		action string, fOp []simulation.FutureOperation, err error) {

		delegatorAccount := simulation.RandomAcc(r, accs)
		validatorAccount := simulation.RandomAcc(r, accs)
		msg := distribution.NewMsgWithdrawDelegatorReward(delegatorAccount.Address, sdk.ValAddress(validatorAccount.Address))

		if msg.ValidateBasic() != nil {
			return "", nil, fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		ctx, write := ctx.CacheContext()
		result := handler(ctx, msg)
		if result.IsOK() {
			write()
		}

		event("distribution/MsgWithdrawDelegatorReward", result.IsOK())

		action = fmt.Sprintf("TestMsgWithdrawDelegatorReward: ok %v, msg %s", result.IsOK(), msg.GetSignBytes())
		return action, nil, nil
	}
}

// SimulateMsgWithdrawValidatorRewardsAll
func SimulateMsgWithdrawValidatorRewardsAll(m auth.AccountKeeper, k distribution.Keeper) simulation.Operation {
	handler := distribution.NewHandler(k)
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context,
		accs []simulation.Account, event simulation.EventFn) (
		action string, fOp []simulation.FutureOperation, err error) {

		account := simulation.RandomAcc(r, accs)
		msg := distribution.NewMsgWithdrawValidatorRewardsAll(sdk.ValAddress(account.Address))

		if msg.ValidateBasic() != nil {
			return "", nil, fmt.Errorf("expected msg to pass ValidateBasic: %s", msg.GetSignBytes())
		}

		ctx, write := ctx.CacheContext()
		result := handler(ctx, msg)
		if result.IsOK() {
			write()
		}

		event("distribution/MsgWithdrawValidatorRewardsAll/", result.IsOK())

		action = fmt.Sprintf("TestMsgWithdrawValidatorRewardsAll: ok %v, msg %s", result.IsOK(), msg.GetSignBytes())
		return action, nil, nil
	}
}
