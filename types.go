package events_db

import (
	"encoding/json"
	"github.com/MinterTeam/minter-go-node/core/types"
	"github.com/tendermint/go-amino"
	"math/big"
)

func RegisterAminoEvents(codec *amino.Codec) {
	codec.RegisterInterface((*Event)(nil), nil)
	codec.RegisterConcrete(RewardEvent{},
		"minter/RewardEvent", nil)
	codec.RegisterConcrete(SlashEvent{},
		"minter/SlashEvent", nil)
	codec.RegisterConcrete(UnbondEvent{},
		"minter/UnbondEvent", nil)
}

type Event interface{}
type Events []Event

type Role byte

const (
	RoleValidator Role = iota
	RoleDelegator
	RoleDAO
	RoleDevelopers
)

func (r Role) String() string {
	switch r {
	case RoleValidator:
		return "Validator"
	case RoleDelegator:
		return "Delegator"
	case RoleDAO:
		return "DAO"
	case RoleDevelopers:
		return "Developers"
	}

	return "Undefined"
}

type event interface{}
type events []event

type reward struct {
	Role      byte
	AddressID uint32
	Amount    []byte
	PubKeyID  uint16
}

type RewardEvent struct {
	Role            Role
	Address         types.Address
	Amount          []byte
	ValidatorPubKey types.Pubkey
}

func (r RewardEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Role            string `json:"role"`
		Address         string `json:"address"`
		Amount          string `json:"amount"`
		ValidatorPubKey string `json:"validator_pub_key"`
	}{
		Role:            r.Role.String(),
		Address:         r.Address.String(),
		Amount:          big.NewInt(0).SetBytes(r.Amount).String(),
		ValidatorPubKey: r.ValidatorPubKey.String(),
	})
}

func rewardConvert(event *RewardEvent, pubKeyID uint16, addressID uint32) interface{} {
	result := new(reward)
	result.AddressID = addressID
	result.Role = byte(event.Role)
	result.Amount = event.Amount
	result.PubKeyID = pubKeyID
	return result
}

func compileReward(item *reward, pubKey string, address [20]byte) interface{} {
	event := new(RewardEvent)
	copy(event.ValidatorPubKey[:], pubKey)
	copy(event.Address[:], address[:])
	event.Role = Role(item.Role)
	event.Amount = item.Amount
	return event
}

type slash struct {
	AddressID uint32
	Amount    []byte
	Coin      [10]byte
	PubKeyID  uint16
}

type SlashEvent struct {
	Address         types.Address
	Amount          []byte
	Coin            types.CoinSymbol
	ValidatorPubKey types.Pubkey
}

func (s SlashEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Address         string `json:"address"`
		Amount          string `json:"amount"`
		Coin            string `json:"coin"`
		ValidatorPubKey string `json:"validator_pub_key"`
	}{
		Address:         s.Address.String(),
		Amount:          big.NewInt(0).SetBytes(s.Amount).String(),
		Coin:            s.Coin.String(),
		ValidatorPubKey: s.ValidatorPubKey.String(),
	})
}

func convertSlash(event *SlashEvent, pubKeyID uint16, addressID uint32) interface{} {
	result := new(slash)
	result.AddressID = addressID
	copy(result.Coin[:], event.Coin[:])
	result.Amount = event.Amount
	result.PubKeyID = pubKeyID
	return result
}

func compileSlash(item *slash, pubKey string, address [20]byte) interface{} {
	event := new(SlashEvent)
	copy(event.ValidatorPubKey[:], pubKey)
	copy(event.Address[:], address[:])
	copy(event.Coin[:], item.Coin[:])
	event.Amount = item.Amount
	return event
}

type unbond struct {
	AddressID uint32
	Amount    []byte
	Coin      [10]byte
	PubKeyID  uint16
}

type UnbondEvent struct {
	Address         types.Address
	Amount          []byte
	Coin            types.CoinSymbol
	ValidatorPubKey types.Pubkey
}

func (u UnbondEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Address         string `json:"address"`
		Amount          string `json:"amount"`
		Coin            string `json:"coin"`
		ValidatorPubKey string `json:"validator_pub_key"`
	}{
		Address:         u.Address.String(),
		Amount:          big.NewInt(0).SetBytes(u.Amount).String(),
		Coin:            u.Coin.String(),
		ValidatorPubKey: u.ValidatorPubKey.String(),
	})
}

func convertUnbound(event *UnbondEvent, pubKeyID uint16, addressID uint32) interface{} {
	result := new(unbond)
	result.AddressID = addressID
	copy(result.Coin[:], event.Coin[:])
	result.Amount = event.Amount
	result.PubKeyID = pubKeyID
	return result
}

func compileUnbond(item *unbond, pubKey string, address [20]byte) interface{} {
	event := new(UnbondEvent)
	copy(event.ValidatorPubKey[:], pubKey)
	copy(event.Address[:], address[:])
	copy(event.Coin[:], item.Coin[:])
	event.Amount = item.Amount
	return event
}
