package bin

import (
	"encoding/binary"
	"reflect"
	"strings"
)

type fieldTag struct {
	SizeOf          []string
	Skip            bool
	Order           binary.ByteOrder
	Optional        bool
	BinaryExtension bool
}

func parseFieldTag(tag reflect.StructTag) *fieldTag {
	t := &fieldTag{
		Order: binary.LittleEndian,
	}
	tagStr := tag.Get("bin")
	for _, s := range strings.Split(tagStr, " ") {
		if strings.HasPrefix(s, "sizeof=") {
			_, tmp, _ := strings.Cut(s, "=")
			t.SizeOf = strings.Split(tmp, ",")
		} else if s == "big" {
			t.Order = binary.BigEndian
		} else if s == "little" {
			t.Order = binary.LittleEndian
		} else if s == "optional" {
			t.Optional = true
		} else if s == "binary_extension" {
			t.BinaryExtension = true
		} else if s == "-" {
			t.Skip = true
		}
	}
	return t
}
