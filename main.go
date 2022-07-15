package main

import (
	"fmt"
	"strconv"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

// ...
func main() {
	// Replace this with the address of the user's wallet
	// 548364400416034343698204186575808495617
	// var prime1, _ = new(big.Int).SetString("548364400416034343698204186575808495617", 10)
	// var prime1, _ = math.ParseBig256("548364400416034343698204186575808495617")
	var prime1, boo = math.ParseBig256("54836440041603434369820418575808495611")
	if !boo {
		fmt.Print("parse big256 error")
	}
	var abc = math.HexOrDecimal256(*prime1)
	salt := "0x5FC8d32690cc91D4c39d9d3abcBD16989F875707"
	privateKey, er2 := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if er2 != nil {
		fmt.Print(er2)
	}
	signerData := apitypes.TypedData{
		Types: apitypes.Types{
			"BuyOrder": []apitypes.Type{
				{Name: "id", Type: "uint256"},
				{Name: "tokenid", Type: "uint256"},
				{Name: "nfttype", Type: "bool"},
			},
			"EIP712Domain": []apitypes.Type{
				{Name: "name", Type: "string"},
				{Name: "version", Type: "string"},
				{Name: "chainId", Type: "uint256"},
				{Name: "verifyingContract", Type: "address"},
			},
		},
		PrimaryType: "BuyOrder",
		Domain: apitypes.TypedDataDomain{
			Name:              "XP721",
			Version:           "1",
			ChainId:           math.NewHexOrDecimal256(31337),
			VerifyingContract: salt,
		},
		Message: apitypes.TypedDataMessage{
			"id":      &abc,
			"tokenid": math.NewHexOrDecimal256(7),
			"nfttype": true,
		},
	}

	domainSeparator, err1 := signerData.HashStruct("EIP712Domain", signerData.Domain.Map())
	if err1 != nil {
		fmt.Print(err1)
	}
	typedDataHash, err := signerData.HashStruct(signerData.PrimaryType, signerData.Message)
	if err != nil {
		fmt.Print("err", err)
	}
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	hash := crypto.Keccak256Hash(rawData)
	// fmt.Print(hash)
	signature, err3 := crypto.Sign(hash.Bytes(), privateKey)
	if err3 != nil {
		fmt.Print(err3)
	}
	a := hexutil.Encode(signature)
	fmt.Println(hexutil.Encode(signature)) // 0x789a80053e4927d0a898db8e065e948f5cf086e32f9ccaa54c1908e22ac430c62621578113ddbb62
	r := fmt.Sprintf(a[2:66])
	fmt.Print(r, "\n")
	s := fmt.Sprintf(a[66:130])
	fmt.Print(s, "\n")
	v := fmt.Sprintf(a[130:132])
	aa, err := strconv.ParseInt(v, 10, 64)
	fmt.Print(aa+27, "\n")
	// fmt.Print(a[66:130], "\n")
	// fmt.Print(a[130:132], "\n")
}
