package types

import (
	"github.com/Juminiy/kube/pkg/util"
	"k8s.io/apimachinery/pkg/api/resource"
)

// BinarySize
// use int64 mixed with resource.Quantity
type BinarySize int64

var ErrBinarySizeFromDB = ValueFromDBError("BinarySize")

var ErrBinarySizeFromJSON = ValueFromJSONError("BinarySize")

func (s BinarySize) MarshalJSON() ([]byte, error) {
	binSi := resource.Quantity{Format: resource.BinarySI}
	binSi.Set(int64(s))
	return binSi.MarshalJSON()
}

func (s *BinarySize) UnmarshalJSON(b []byte) error {
	if InValidJSONValue(util.Bytes2StringNoCopy(b)) {
		return nil
	}

	binSi := resource.Quantity{Format: resource.BinarySI}
	if err := binSi.UnmarshalJSON(b); err != nil {
		return err
	} else if binI64, ok := binSi.AsInt64(); ok {
		*s = BinarySize(binI64)
		return nil
	} // much bigger than

	return ErrBinarySizeFromJSON
}

// TODO: replace Ki -> KB, Mi -> MB, Gi -> GB, add none_unit 'B'
// TODO: cannot hold bigger eq than 1Ti size
