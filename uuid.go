package uuid

type UUID [16]byte

func (u *UUID) variantRFC4122() {
	u[8] = (u[8] & 0x3f) | 0x80
}
