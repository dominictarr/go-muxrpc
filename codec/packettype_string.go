// Code generated by "stringer -type PacketType"; DO NOT EDIT

package codec

import "fmt"

const _PacketType_name = "BufferStringJSON"

var _PacketType_index = [...]uint8{0, 6, 12, 16}

func (i PacketType) String() string {
	if i >= PacketType(len(_PacketType_index)-1) {
		return fmt.Sprintf("PacketType(%d)", i)
	}
	return _PacketType_name[_PacketType_index[i]:_PacketType_index[i+1]]
}