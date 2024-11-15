package packet

import (
	"go.minekube.com/gate/pkg/edition/java/proto/util"
	"go.minekube.com/gate/pkg/edition/java/proto/version"
	"go.minekube.com/gate/pkg/gate/proto"
	"io"
)

type KeepAlive struct {
	RandomID int64
}

func (k *KeepAlive) Encode(c *proto.PacketContext, wr io.Writer) error {
	if c.Protocol.GreaterEqual(version.Minecraft_1_12_2) {
		return util.WriteInt64(wr, k.RandomID)
	} else if c.Protocol.GreaterEqual(version.Minecraft_1_8) {
		return util.WriteVarInt(wr, int(k.RandomID))
	}
	return util.WriteInt32(wr, int32(k.RandomID))
}

func (k *KeepAlive) Decode(c *proto.PacketContext, rd io.Reader) (err error) {
	if c.Protocol.GreaterEqual(version.Minecraft_1_12_2) {
		k.RandomID, err = util.ReadInt64(rd)
	} else if c.Protocol.GreaterEqual(version.Minecraft_1_8) {
		var id int
		id, err = util.ReadVarInt(rd)
		k.RandomID = int64(id)
	} else {
		var id int32
		id, err = util.ReadInt32(rd)
		k.RandomID = int64(id)
	}
	return
}

var _ proto.Packet = (*KeepAlive)(nil)
