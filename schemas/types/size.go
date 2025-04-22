package types

import (
	"encoding/json"
	"k8s.io/apimachinery/pkg/api/resource"
)

// BinarySize
// use int64 mixed with resource.Quantity
type BinarySize int64

var ErrBinarySizeFromDB = ValueFromDBError("BinarySize")

var ErrBinarySizeFromJSON = ValueFromJSONError("BinarySize")

func (s BinarySize) MarshalJSON() ([]byte, error) {
	if s >= 0 {
		binSi := resource.NewQuantity(int64(s), resource.BinarySI)
		binSi.Set(int64(s))
		return binSi.MarshalJSON()
	}
	return json.Marshal(nil)
}

func (s *BinarySize) UnmarshalJSON(b []byte) error {
	if InValidJSON(b) {
		return nil
	}

	var bStr string
	if err := json.Unmarshal(b, &bStr); err == nil {
		if qtt, err := resource.ParseQuantity(bStr); err == nil {
			*s = BinarySize(qtt.Value())
			return nil
		}
	}
	// same as upper
	/*qtt := resource.NewQuantity(0, resource.BinarySI)
	if err := qtt.UnmarshalJSON(b); err == nil {
		*s = BinarySize(qtt.Value())
		return nil
	} else if binI64, ok := qtt.AsInt64(); ok {
		*s = BinarySize(binI64)
		return nil
	}*/
	return ErrBinarySizeFromJSON
}
