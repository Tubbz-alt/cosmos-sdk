package distribution

import (
	"github.com/cosmos/cosmos-sdk/x/distribution/client"
	"github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
)

// nolint

const (
	DefaultParamspace                = keeper.DefaultParamspace
	ModuleName                       = types.ModuleName
	StoreKey                         = types.StoreKey
	RouterKey                        = types.RouterKey
	QuerierRoute                     = types.QuerierRoute
	ProposalTypeCommunityPoolSpend   = types.ProposalTypeCommunityPoolSpend
	QueryParams                      = types.QueryParams
	QueryValidatorOutstandingRewards = types.QueryValidatorOutstandingRewards
	QueryValidatorCommission         = types.QueryValidatorCommission
	QueryValidatorSlashes            = types.QueryValidatorSlashes
	QueryDelegationRewards           = types.QueryDelegationRewards
	QueryDelegatorTotalRewards       = types.QueryDelegatorTotalRewards
	QueryDelegatorValidators         = types.QueryDelegatorValidators
	QueryWithdrawAddr                = types.QueryWithdrawAddr
	QueryCommunityPool               = types.QueryCommunityPool
	ParamCommunityTax                = types.ParamCommunityTax
	ParamBaseProposerReward          = types.ParamBaseProposerReward
	ParamBonusProposerReward         = types.ParamBonusProposerReward
	ParamWithdrawAddrEnabled         = types.ParamWithdrawAddrEnabled
	TypeMsgFundCommunityPool         = types.TypeMsgFundCommunityPool
)

var (
	// functions aliases
	RegisterInvariants                         = keeper.RegisterInvariants
	AllInvariants                              = keeper.AllInvariants
	NonNegativeOutstandingInvariant            = keeper.NonNegativeOutstandingInvariant
	CanWithdrawInvariant                       = keeper.CanWithdrawInvariant
	ReferenceCountInvariant                    = keeper.ReferenceCountInvariant
	ModuleAccountInvariant                     = keeper.ModuleAccountInvariant
	NewKeeper                                  = keeper.NewKeeper
	GetValidatorOutstandingRewardsAddress      = keeper.GetValidatorOutstandingRewardsAddress
	GetDelegatorWithdrawInfoAddress            = keeper.GetDelegatorWithdrawInfoAddress
	GetDelegatorStartingInfoAddresses          = keeper.GetDelegatorStartingInfoAddresses
	GetValidatorHistoricalRewardsAddressPeriod = keeper.GetValidatorHistoricalRewardsAddressPeriod
	GetValidatorCurrentRewardsAddress          = keeper.GetValidatorCurrentRewardsAddress
	GetValidatorAccumulatedCommissionAddress   = keeper.GetValidatorAccumulatedCommissionAddress
	GetValidatorSlashEventAddressHeight        = keeper.GetValidatorSlashEventAddressHeight
	GetValidatorOutstandingRewardsKey          = keeper.GetValidatorOutstandingRewardsKey
	GetDelegatorWithdrawAddrKey                = keeper.GetDelegatorWithdrawAddrKey
	GetDelegatorStartingInfoKey                = keeper.GetDelegatorStartingInfoKey
	GetValidatorHistoricalRewardsPrefix        = keeper.GetValidatorHistoricalRewardsPrefix
	GetValidatorHistoricalRewardsKey           = keeper.GetValidatorHistoricalRewardsKey
	GetValidatorCurrentRewardsKey              = keeper.GetValidatorCurrentRewardsKey
	GetValidatorAccumulatedCommissionKey       = keeper.GetValidatorAccumulatedCommissionKey
	GetValidatorSlashEventPrefix               = keeper.GetValidatorSlashEventPrefix
	GetValidatorSlashEventKeyPrefix            = keeper.GetValidatorSlashEventKeyPrefix
	GetValidatorSlashEventKey                  = keeper.GetValidatorSlashEventKey
	ParamKeyTable                              = keeper.ParamKeyTable
	HandleCommunityPoolSpendProposal           = keeper.HandleCommunityPoolSpendProposal
	NewQuerier                                 = keeper.NewQuerier
	MakeTestCodec                              = keeper.MakeTestCodec
	CreateTestInputDefault                     = keeper.CreateTestInputDefault
	CreateTestInputAdvanced                    = keeper.CreateTestInputAdvanced
	RegisterCodec                              = types.RegisterCodec
	NewDelegatorStartingInfo                   = types.NewDelegatorStartingInfo
	ErrEmptyDelegatorAddr                      = types.ErrEmptyDelegatorAddr
	ErrEmptyWithdrawAddr                       = types.ErrEmptyWithdrawAddr
	ErrEmptyValidatorAddr                      = types.ErrEmptyValidatorAddr
	ErrEmptyDelegationDistInfo                 = types.ErrEmptyDelegationDistInfo
	ErrNoValidatorDistInfo                     = types.ErrNoValidatorDistInfo
	ErrNoValidatorExists                       = types.ErrNoValidatorExists
	ErrNoDelegationExists                      = types.ErrNoDelegationExists
	ErrNoValidatorCommission                   = types.ErrNoValidatorCommission
	ErrSetWithdrawAddrDisabled                 = types.ErrSetWithdrawAddrDisabled
	ErrBadDistribution                         = types.ErrBadDistribution
	ErrInvalidProposalAmount                   = types.ErrInvalidProposalAmount
	ErrEmptyProposalRecipient                  = types.ErrEmptyProposalRecipient
	InitialFeePool                             = types.InitialFeePool
	NewGenesisState                            = types.NewGenesisState
	DefaultGenesisState                        = types.DefaultGenesisState
	ValidateGenesis                            = types.ValidateGenesis
	NewMsgSetWithdrawAddress                   = types.NewMsgSetWithdrawAddress
	NewMsgWithdrawDelegatorReward              = types.NewMsgWithdrawDelegatorReward
	NewMsgWithdrawValidatorCommission          = types.NewMsgWithdrawValidatorCommission
	MsgFundCommunityPool                       = types.NewMsgFundCommunityPool
	NewCommunityPoolSpendProposal              = types.NewCommunityPoolSpendProposal
	NewQueryValidatorOutstandingRewardsParams  = types.NewQueryValidatorOutstandingRewardsParams
	NewQueryValidatorCommissionParams          = types.NewQueryValidatorCommissionParams
	NewQueryValidatorSlashesParams             = types.NewQueryValidatorSlashesParams
	NewQueryDelegationRewardsParams            = types.NewQueryDelegationRewardsParams
	NewQueryDelegatorParams                    = types.NewQueryDelegatorParams
	NewQueryDelegatorWithdrawAddrParams        = types.NewQueryDelegatorWithdrawAddrParams
	NewQueryDelegatorTotalRewardsResponse      = types.NewQueryDelegatorTotalRewardsResponse
	NewDelegationDelegatorReward               = types.NewDelegationDelegatorReward
	NewValidatorHistoricalRewards              = types.NewValidatorHistoricalRewards
	NewValidatorCurrentRewards                 = types.NewValidatorCurrentRewards
	InitialValidatorAccumulatedCommission      = types.InitialValidatorAccumulatedCommission
	NewValidatorSlashEvent                     = types.NewValidatorSlashEvent

	// variable aliases
	FeePoolKey                           = keeper.FeePoolKey
	ProposerKey                          = keeper.ProposerKey
	ValidatorOutstandingRewardsPrefix    = keeper.ValidatorOutstandingRewardsPrefix
	DelegatorWithdrawAddrPrefix          = keeper.DelegatorWithdrawAddrPrefix
	DelegatorStartingInfoPrefix          = keeper.DelegatorStartingInfoPrefix
	ValidatorHistoricalRewardsPrefix     = keeper.ValidatorHistoricalRewardsPrefix
	ValidatorCurrentRewardsPrefix        = keeper.ValidatorCurrentRewardsPrefix
	ValidatorAccumulatedCommissionPrefix = keeper.ValidatorAccumulatedCommissionPrefix
	ValidatorSlashEventPrefix            = keeper.ValidatorSlashEventPrefix
	ParamStoreKeyCommunityTax            = keeper.ParamStoreKeyCommunityTax
	ParamStoreKeyBaseProposerReward      = keeper.ParamStoreKeyBaseProposerReward
	ParamStoreKeyBonusProposerReward     = keeper.ParamStoreKeyBonusProposerReward
	ParamStoreKeyWithdrawAddrEnabled     = keeper.ParamStoreKeyWithdrawAddrEnabled
	TestAddrs                            = keeper.TestAddrs
	ModuleCdc                            = types.ModuleCdc
	EventTypeSetWithdrawAddress          = types.EventTypeSetWithdrawAddress
	EventTypeRewards                     = types.EventTypeRewards
	EventTypeCommission                  = types.EventTypeCommission
	EventTypeWithdrawRewards             = types.EventTypeWithdrawRewards
	EventTypeWithdrawCommission          = types.EventTypeWithdrawCommission
	EventTypeProposerReward              = types.EventTypeProposerReward
	AttributeKeyWithdrawAddress          = types.AttributeKeyWithdrawAddress
	AttributeKeyValidator                = types.AttributeKeyValidator
	AttributeValueCategory               = types.AttributeValueCategory
	ProposalHandler                      = client.ProposalHandler
)

type (
	Hooks                                  = keeper.Hooks
	Keeper                                 = keeper.Keeper
	DelegatorStartingInfo                  = types.DelegatorStartingInfo
	FeePool                                = types.FeePool
	DelegatorWithdrawInfo                  = types.DelegatorWithdrawInfo
	ValidatorOutstandingRewardsRecord      = types.ValidatorOutstandingRewardsRecord
	ValidatorAccumulatedCommissionRecord   = types.ValidatorAccumulatedCommissionRecord
	ValidatorHistoricalRewardsRecord       = types.ValidatorHistoricalRewardsRecord
	ValidatorCurrentRewardsRecord          = types.ValidatorCurrentRewardsRecord
	DelegatorStartingInfoRecord            = types.DelegatorStartingInfoRecord
	ValidatorSlashEventRecord              = types.ValidatorSlashEventRecord
	GenesisState                           = types.GenesisState
	MsgSetWithdrawAddress                  = types.MsgSetWithdrawAddress
	MsgWithdrawDelegatorReward             = types.MsgWithdrawDelegatorReward
	MsgWithdrawValidatorCommission         = types.MsgWithdrawValidatorCommission
	CommunityPoolSpendProposal             = types.CommunityPoolSpendProposal
	QueryValidatorOutstandingRewardsParams = types.QueryValidatorOutstandingRewardsParams
	QueryValidatorCommissionParams         = types.QueryValidatorCommissionParams
	QueryValidatorSlashesParams            = types.QueryValidatorSlashesParams
	QueryDelegationRewardsParams           = types.QueryDelegationRewardsParams
	QueryDelegatorParams                   = types.QueryDelegatorParams
	QueryDelegatorWithdrawAddrParams       = types.QueryDelegatorWithdrawAddrParams
	QueryDelegatorTotalRewardsResponse     = types.QueryDelegatorTotalRewardsResponse
	DelegationDelegatorReward              = types.DelegationDelegatorReward
	ValidatorHistoricalRewards             = types.ValidatorHistoricalRewards
	ValidatorCurrentRewards                = types.ValidatorCurrentRewards
	ValidatorAccumulatedCommission         = types.ValidatorAccumulatedCommission
	ValidatorSlashEvent                    = types.ValidatorSlashEvent
	ValidatorSlashEvents                   = types.ValidatorSlashEvents
	ValidatorOutstandingRewards            = types.ValidatorOutstandingRewards
)
