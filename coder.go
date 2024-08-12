package gobinstruct

import (
	"errors"
	"io"
	"reflect"
)

var ErrMisType = errors.New("mistype")

type (
	Encoder []func(w io.Writer, data any, info *Encoder_ExtraInfo) error
	Decoder []func(r io.Reader, data any, info *Decoder_ExtraInfo) error

	Encoder_ExtraInfo struct {
		Encoder   *Encoder
		StructTag reflect.StructTag
	}

	Decoder_ExtraInfo struct {
		Decoder   *Decoder
		StructTag reflect.StructTag
	}
)

var (
	LE_Encoder = append(Encoder{
		Base_LE_Encoder,
	}, Special_Encoder...)

	LE_Decoder = append(Decoder{
		Base_LE_Decoder,
	}, Special_Decoder...)
)

var (
	BE_Encoder = append(Encoder{
		Base_BE_Encoder,
	}, Special_Encoder...)

	BE_Decoder = append(Decoder{
		Base_BE_Decoder,
	}, Special_Decoder...)
)

var (
	Special_Encoder = Encoder{
		ByteArray_Encoder,
		Array_Encoder,
		Slice_Encoder,
		Struct_Encoder,
	}

	Special_Decoder = Decoder{
		ByteArray_Decoder,
		Array_Decoder,
		Slice_Decoder,
		Struct_Decoder,
	}
)
