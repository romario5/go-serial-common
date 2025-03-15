package serial

import (
	"errors"
	"io"
)

const (
	CHANNELS_COUNT = 18
)

type SerialProtocol interface {
	ReadPacket(r io.Reader) (any, error)
	WritePacket(w io.Writer, p any) error
}

type GpsPacket struct {
	Latitude    int32
	Longitude   int32
	GroundSpeed uint16
	Heading     uint16
	Altitude    uint16
	Satellites  uint8
}

type BatteryPacket struct {
	AverageCellVoltage uint16 // Value should be divided by 100
}

type ChannelsPacket struct {
	Id         uint32
	Channels   [CHANNELS_COUNT]uint16
	IsFailsafe bool
}

// Packs channels value in to the bytes slice where each
// channel value presented in 11 bits, big endian.
// Only first 16 channels can be packed.
func (p *ChannelsPacket) PackChannels(b []byte) error {
	if len(b) < 22 {
		return errors.New("Length of slice for packing Channels less than 22")
	}

	b[0] = uint8((p.Channels[0] & 0x07FF))
	b[1] = uint8((p.Channels[0]&0x07FF)>>8 | (p.Channels[1]&0x07FF)<<3)
	b[2] = uint8((p.Channels[1]&0x07FF)>>5 | (p.Channels[2]&0x07FF)<<6)
	b[3] = uint8((p.Channels[2] & 0x07FF) >> 2)
	b[4] = uint8((p.Channels[2]&0x07FF)>>10 | (p.Channels[3]&0x07FF)<<1)
	b[5] = uint8((p.Channels[3]&0x07FF)>>7 | (p.Channels[4]&0x07FF)<<4)
	b[6] = uint8((p.Channels[4]&0x07FF)>>4 | (p.Channels[5]&0x07FF)<<7)
	b[7] = uint8((p.Channels[5] & 0x07FF) >> 1)
	b[8] = uint8((p.Channels[5]&0x07FF)>>9 | (p.Channels[6]&0x07FF)<<2)
	b[9] = uint8((p.Channels[6]&0x07FF)>>6 | (p.Channels[7]&0x07FF)<<5)
	b[10] = uint8((p.Channels[7] & 0x07FF) >> 3)
	b[11] = uint8((p.Channels[8] & 0x07FF))
	b[12] = uint8((p.Channels[8]&0x07FF)>>8 | (p.Channels[9]&0x07FF)<<3)
	b[13] = uint8((p.Channels[9]&0x07FF)>>5 | (p.Channels[10]&0x07FF)<<6)
	b[14] = uint8((p.Channels[10] & 0x07FF) >> 2)
	b[15] = uint8((p.Channels[10]&0x07FF)>>10 | (p.Channels[11]&0x07FF)<<1)
	b[16] = uint8((p.Channels[11]&0x07FF)>>7 | (p.Channels[12]&0x07FF)<<4)
	b[17] = uint8((p.Channels[12]&0x07FF)>>4 | (p.Channels[13]&0x07FF)<<7)
	b[18] = uint8((p.Channels[13] & 0x07FF) >> 1)
	b[19] = uint8((p.Channels[13]&0x07FF)>>9 | (p.Channels[14]&0x07FF)<<2)
	b[20] = uint8((p.Channels[14]&0x07FF)>>6 | (p.Channels[15]&0x07FF)<<5)
	b[21] = uint8((p.Channels[15] & 0x07FF) >> 3)

	return nil
}

// Parses given slice of bytes that represents channels values
// stored in 11 bits each, big endian.
// Only first 16 channels can be parsed.
func (p *ChannelsPacket) ParseChannels(b []byte) error {
	if len(b) < 22 {
		return errors.New("Length of slice for packing Channels less than 22")
	}

	p.Channels[0] = uint16(b[0]) | ((uint16(b[1]) << 8) & 0x07FF)
	p.Channels[1] = (uint16(b[1]) >> 3) | ((uint16(b[2]) << 5) & 0x07FF)
	p.Channels[2] = ((uint16(b[2]) >> 6) & 0x07FF) | (uint16(b[3]) << 2) | ((uint16(b[4]) << 10) & 0x07FF)
	p.Channels[3] = (uint16(b[4]) >> 1) | ((uint16(b[5]) << 7) & 0x07FF)
	p.Channels[4] = (uint16(b[5]) >> 4) | ((uint16(b[6]) << 4) & 0x07FF)
	p.Channels[5] = (uint16(b[6]) >> 7) | (uint16(b[7]) << 1) | ((uint16(b[8]) << 9) & 0x07FF)
	p.Channels[6] = (uint16(b[8]) >> 2) | ((uint16(b[9]) << 6) & 0x07FF)
	p.Channels[7] = (uint16(b[9]) >> 5) | ((uint16(b[10]) << 3) & 0x07FF)
	p.Channels[8] = uint16(b[11]) | ((uint16(b[12]) << 8) & 0x07FF)
	p.Channels[9] = (uint16(b[12]) >> 3) | ((uint16(b[13]) << 5) & 0x07FF)
	p.Channels[10] = (uint16(b[13]) >> 6) | (uint16(b[14]) << 2) | ((uint16(b[15]) << 10) & 0x07FF)
	p.Channels[11] = (uint16(b[15]) >> 1) | ((uint16(b[16]) << 7) & 0x07FF)
	p.Channels[12] = (uint16(b[16]) >> 4) | ((uint16(b[17]) << 4) & 0x07FF)
	p.Channels[13] = (uint16(b[17]) >> 7) | (uint16(b[18]) << 1) | ((uint16(b[19]) << 9) & 0x07FF)
	p.Channels[14] = (uint16(b[19]) >> 2) | ((uint16(b[20]) << 6) & 0x07FF)
	p.Channels[15] = (uint16(b[20]) >> 5) | ((uint16(b[21]) << 3) & 0x07FF)

	return nil
}
