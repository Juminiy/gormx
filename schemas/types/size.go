package types

import (
	"encoding/json"
	"github.com/docker/go-units"
	"k8s.io/apimachinery/pkg/api/resource"
	"strings"
)

var ErrBinarySizeFromDB = ValueFromDBError("BinarySize")

var ErrBinarySizeFromJSON = ValueFromJSONError("BinarySize")

var ErrNBinarySizeFromJSON = ValueFromJSONError("NBinarySize")

var ErrHBinarySizeFromJSON = ValueFromJSONError("HBinarySize")

var ErrKBinarySizeFromJSON = ValueFromJSONError("KBinarySize")

// BinarySize
//  1. range: can hold [0,1<<63)B
//  2. stand:
//
// - 50GiB(normalized) | 50GB(humanRead) | 50Gi(kubernetes)
// - 10MiB(normalized) | 10MB(humanRead) | 10Mi(kubernetes)
// 3. BinarySize can receive stringFormat: HBinarySize, intFormat: int64
type BinarySize HBinarySize

func (s BinarySize) MarshalJSON() ([]byte, error) {
	return (HBinarySize)(s).MarshalJSON()
}

func (s *BinarySize) UnmarshalJSON(b []byte) error {
	var i64v int64
	if err := json.Unmarshal(b, &i64v); err == nil {
		*s = BinarySize(i64v)
		return nil
	}
	return (*HBinarySize)(s).UnmarshalJSON(b)
}

// NBinarySize
// 1. fullName: NormalizedBinarySize
// 2. representation: 50PiB, 33GiB, 10MiB, 100B
// 3. referredFrom: units.RAMInBytes, units.BytesSize
type NBinarySize int64

func (s NBinarySize) MarshalJSON() ([]byte, error) {
	if s >= 0 {
		return json.Marshal(units.BytesSize(float64(s)))
	}
	return json.Marshal(nil)
}

func (s *NBinarySize) UnmarshalJSON(b []byte) error {
	if InValidJSON(b) {
		return nil
	}

	var bStr string
	if err := json.Unmarshal(b, &bStr); err == nil {
		if i64v, err := units.RAMInBytes(bStr); err == nil {
			*s = NBinarySize(i64v)
			return nil
		}
	}
	return ErrNBinarySizeFromJSON
}

// HBinarySize
// 1. fullName: HumanReadBinarySize
// 2. representation: 50PB, 33GB, 10MB, 100B
type HBinarySize int64

func (s HBinarySize) MarshalJSON() ([]byte, error) {
	if s >= 0 {
		return json.Marshal(strings.Replace(units.BytesSize(float64(s)), "i", "", 1))
	}
	return json.Marshal(nil)
}

func (s *HBinarySize) UnmarshalJSON(b []byte) error {
	if InValidJSON(b) {
		return nil
	}

	var bStr string
	if err := json.Unmarshal(b, &bStr); err == nil {
		if i64v, err := units.RAMInBytes(bStr); err == nil {
			*s = HBinarySize(i64v)
			return nil
		}
	}
	return ErrHBinarySizeFromJSON
}

// KBinarySize
// 1. fullName: KubernetesBinarySize
// 2. representation: 50Pi, 33Gi, 10Mi, 100
// 3. referredFrom: resource.Quantity
type KBinarySize int64

func (s KBinarySize) MarshalJSON() ([]byte, error) {
	if s >= 0 {
		binSi := resource.NewQuantity(int64(s), resource.BinarySI)
		binSi.Set(int64(s))
		return binSi.MarshalJSON()
	}
	return json.Marshal(nil)
}

func (s *KBinarySize) UnmarshalJSON(b []byte) error {
	if InValidJSON(b) {
		return nil
	}

	var bStr string
	if err := json.Unmarshal(b, &bStr); err == nil {
		if qtt, err := resource.ParseQuantity(bStr); err == nil {
			*s = KBinarySize(qtt.Value())
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
	return ErrKBinarySizeFromJSON
}
