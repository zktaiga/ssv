package discovery

import (
	"github.com/bloxapp/ssv/utils/format"
	"github.com/ethereum/go-ethereum/p2p/enode"
	"strconv"
)

const (
	// SubnetsCount is the count of subnets in the network
	SubnetsCount = 128
	// ENRKeySubnets is the entry key for saving subnets
	ENRKeySubnets = "subnets"
)

var regPool = format.NewRegexpPool("\\w+:bloxstaking\\.ssv\\.(\\d+)")

// nsToSubnet converts the given topic to subnet
// TODO: return other value than zero upon failure?
func nsToSubnet(ns string) int64 {
	r, done := regPool.Get()
	defer done()
	found := r.FindStringSubmatch(ns)
	if len(found) != 2 {
		return -1
	}
	val, err := strconv.ParseUint(found[1], 10, 64)
	if err != nil {
		return -1
	}
	return int64(val)
}

// isSubnet checks if the given string is a subnet string
func isSubnet(ns string) bool {
	r, done := regPool.Get()
	defer done()
	return r.MatchString(ns)
}

func setSubnetsEntry(node *enode.LocalNode, subnets []bool) error {
	//bl := bitfield.NewBitlist(uint64(SubnetsCount))
	//for i, state := range subnets {
	//	bl.SetBitAt(uint64(i), state)
	//}
	//node.Set(enr.WithEntry(ENRKeySubnets, bl))
	return nil
}

func getSubnetsEntry(node *enode.Node) ([]bool, error) {
	var subnets []bool
	// TODO: fix and unmark
	//bl := bitfield.NewBitlist(uint64(SubnetsCount))
	//err := node.Record().Load(enr.WithEntry(ENRKeySubnets, &bl))
	//if err != nil {
	//	return subnets, err
	//}
	//l := len(bl)
	//if l == 0 {
	//	return nil, errors.New("subnets entry not found")
	//}
	//if l > byteCount(SubnetsCount)+1 || l < byteCount(SubnetsCount)-1 {
	//	return subnets, errors.Errorf("invalid bitvector provided, it has a size of %d", l)
	//}
	//for i := 0; i < SubnetsCount; i++ {
	//	subnets = append(subnets, bl.BitAt(uint64(i)))
	//}
	for i := 0; i < SubnetsCount; i++ {
		subnets = append(subnets, true)
	}
	return subnets, nil
}

//// Determines the number of bytes that are used
//// to represent the provided number of bits.
//func byteCount(bitCount int) int {
//	numOfBytes := bitCount / 8
//	if bitCount%8 != 0 {
//		numOfBytes++
//	}
//	return numOfBytes
//}
