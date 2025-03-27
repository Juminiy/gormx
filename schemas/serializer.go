package schemas

type MarshalFunc func(v any) ([]byte, error)
type UnmarshalFunc func(data []byte, v any) error

type JSONType struct {
	m MarshalFunc
	u UnmarshalFunc
}

type XMLType struct {
	m MarshalFunc
	u UnmarshalFunc
}

type YAMLType struct {
	m MarshalFunc
	u UnmarshalFunc
}

type TOMLType struct {
	m MarshalFunc
	u UnmarshalFunc
}

type ProtoBufferType struct {
	m MarshalFunc
	u UnmarshalFunc
}

type MessagePackType struct {
	m MarshalFunc
	u UnmarshalFunc
}
