package types

import (
	"fmt"
	"strings"
)

func ParseNftId(nftId string) (Nft, error) {
	// split the nftId into its components
	data := strings.Split(nftId, "/")
	if len(data) != 3 {
		return Nft{}, fmt.Errorf("invalid nftId: %s", nftId)
	}

	return Nft{
		ChainId:      data[0],
		ContractAddr: NormalizeHexAddress(data[1]),
		TokenId:      NormalizeHexAddress(data[2]),
	}, nil
}

func (nft Nft) FormatString() string {
	return fmt.Sprintf("%s/%s/%s", nft.ChainId, nft.ContractAddr, nft.TokenId)
}
