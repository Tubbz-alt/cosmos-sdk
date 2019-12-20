package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Governance message types and routes
const (
	TypeMsgDeposit        = "deposit"
	TypeMsgVote           = "vote"
	TypeMsgSubmitProposal = "submit_proposal"
)

var _, _ sdk.Msg = MsgDeposit{}, MsgVote{}

// MsgSubmitProposal defines a message to create a governance proposal with a
// given content and initial deposit
//type MsgSubmitProposal struct {
//	Content        Content        `json:"content" yaml:"content"`
//	InitialDeposit sdk.Coins      `json:"initial_deposit" yaml:"initial_deposit"` //  Initial deposit paid by sender. Must be strictly positive
//	Proposer       sdk.AccAddress `json:"proposer" yaml:"proposer"`               //  Address of the proposer
//}

type MsgSubmitProposal interface {
	sdk.Msg
	GetContent() Content
	GetInitialDeposit() sdk.Coins
	GetProposer() sdk.AccAddress
}

// NewMsgSubmitProposal creates a new MsgSubmitProposal instance
//func NewMsgSubmitProposal(content Content, initialDeposit sdk.Coins, proposer sdk.AccAddress) MsgSubmitProposal {
//	return MsgSubmitProposal{content, initialDeposit, proposer}
//}

// Route implements Msg
func (msg MsgSubmitProposalBase) Route() string { return RouterKey }

// Type implements Msg
func (msg MsgSubmitProposalBase) Type() string { return TypeMsgSubmitProposal }

// ValidateBasic implements Msg
func (msg MsgSubmitProposalBase) ValidateBasic() sdk.Error {
	//if msg.Content == nil {
	//	return ErrInvalidProposalContent(DefaultCodespace, "missing content")
	//}
	if msg.Proposer.Empty() {
		return sdk.ErrInvalidAddress(msg.Proposer.String())
	}
	initialDeposit := sdk.Coins(msg.IntialDeposit)
	if !initialDeposit.IsValid() {
		return sdk.ErrInvalidCoins(initialDeposit.String())
	}
	if initialDeposit.IsAnyNegative() {
		return sdk.ErrInvalidCoins(initialDeposit.String())
	}
	//if !IsValidProposalType(msg.Content.ProposalType()) {
	//	return ErrInvalidProposalType(DefaultCodespace, msg.Content.ProposalType())
	//}
	//
	//return msg.Content.ValidateBasic()
	return nil
}

// String implements the Stringer interface
//func (msg MsgSubmitProposal) String() string {
//	return fmt.Sprintf(`Submit Proposal Message:
//  Content:         %s
//  Initial Deposit: %s
//`, msg.Content.String(), msg.InitialDeposit)
//}

// GetSignBytes implements Msg
func (msg MsgSubmitProposalBase) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgSubmitProposalBase) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Proposer}
}

func (m MsgSubmitProposalBase) GetInitialDeposit() sdk.Coins {
	return m.IntialDeposit
}

func (m MsgSubmitProposalBase) GetProposer() sdk.AccAddress {
	return m.Proposer
}



//// MsgDeposit defines a message to submit a deposit to an existing proposal
//type MsgDeposit struct {
//	ProposalID uint64         `json:"proposal_id" yaml:"proposal_id"` // ID of the proposal
//	Depositor  sdk.AccAddress `json:"depositor" yaml:"depositor"`     // Address of the depositor
//	Amount     sdk.Coins      `json:"amount" yaml:"amount"`           // Coins to add to the proposal's deposit
//}

// NewMsgDeposit creates a new MsgDeposit instance
func NewMsgDeposit(depositor sdk.AccAddress, proposalID uint64, amount sdk.Coins) MsgDeposit {
	return MsgDeposit{proposalID, depositor, amount}
}

// Route implements Msg
func (msg MsgDeposit) Route() string { return RouterKey }

// Type implements Msg
func (msg MsgDeposit) Type() string { return TypeMsgDeposit }

// ValidateBasic implements Msg
func (msg MsgDeposit) ValidateBasic() sdk.Error {
	if msg.Depositor.Empty() {
		return sdk.ErrInvalidAddress(msg.Depositor.String())
	}
	amount := sdk.Coins(msg.Amount)
	if !amount.IsValid() {
		return sdk.ErrInvalidCoins(amount.String())
	}
	if amount.IsAnyNegative() {
		return sdk.ErrInvalidCoins(amount.String())
	}

	return nil
}

// String implements the Stringer interface
func (msg MsgDeposit) String() string {
	return fmt.Sprintf(`Deposit Message:
  Depositer:   %s
  Proposal ID: %d
  Amount:      %s
`, msg.Depositor, msg.ProposalID, msg.Amount)
}

// GetSignBytes implements Msg
func (msg MsgDeposit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgDeposit) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Depositor}
}

// MsgVote defines a message to cast a vote
//type MsgVote struct {
//	ProposalID uint64         `json:"proposal_id" yaml:"proposal_id"` // ID of the proposal
//	Voter      sdk.AccAddress `json:"voter" yaml:"voter"`             //  address of the voter
//	Option     VoteOption     `json:"option" yaml:"option"`           //  option from OptionSet chosen by the voter
//}

// NewMsgVote creates a message to cast a vote on an active proposal
func NewMsgVote(voter sdk.AccAddress, proposalID uint64, option VoteOption) MsgVote {
	return MsgVote{proposalID, voter, option}
}

// Route implements Msg
func (msg MsgVote) Route() string { return RouterKey }

// Type implements Msg
func (msg MsgVote) Type() string { return TypeMsgVote }

// ValidateBasic implements Msg
func (msg MsgVote) ValidateBasic() sdk.Error {
	if msg.Voter.Empty() {
		return sdk.ErrInvalidAddress(msg.Voter.String())
	}
	if !ValidVoteOption(msg.Option) {
		return ErrInvalidVote(DefaultCodespace, msg.Option.String())
	}

	return nil
}

// String implements the Stringer interface
func (msg MsgVote) String() string {
	return fmt.Sprintf(`Vote Message:
  Proposal ID: %d
  Option:      %s
`, msg.ProposalID, msg.Option)
}

// GetSignBytes implements Msg
func (msg MsgVote) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements Msg
func (msg MsgVote) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Voter}
}
