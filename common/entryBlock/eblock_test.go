package entryBlock_test

import (
	. "github.com/FactomProject/factomd/common/entryBlock"
	"github.com/FactomProject/factomd/common/primitives"
	"testing"

	"encoding/hex"
)

func TestEBlockMarshal(t *testing.T) {
	eb := newTestingEntryBlock()

	t.Log(eb)
	p, err := eb.MarshalBinary()
	if err != nil {
		t.Error(err)
	}

	eb2 := NewEBlock()
	if err := eb2.UnmarshalBinary(p); err != nil {
		t.Error(err)
	}
	t.Log(eb2)
	p2, err := eb2.MarshalBinary()
	if err != nil {
		t.Error(err)
	}

	if string(p) != string(p2) {
		t.Logf("eb1 = %x\n", p)
		t.Logf("eb2 = %x\n", p2)
		t.Fail()
	}
}

func TestAddEBEntry(t *testing.T) {
	eb := newTestingEntryBlock()
	e := newEntry()
	if err := eb.AddEBEntry(e); err != nil {
		t.Error(err)
	}
	t.Log(eb)
	p, err := eb.MarshalBinary()
	if err != nil {
		t.Error(err)
	}

	eb2 := NewEBlock()
	if err := eb2.UnmarshalBinary(p); err != nil {
		t.Error(err)
	}
	t.Log(eb2)
}

func byteof(b byte) []byte {
	r := make([]byte, 0, 32)
	for i := 0; i < 32; i++ {
		r = append(r, b)
	}
	return r
}

func TestEntryBlockMisc(t *testing.T) {
	e := newEntryBlock()
	hash, err := e.Hash()

	if err != nil {
		t.Error(err)
	}
	if hash.String() != "1ec4c9a52ede96e57f855efc8cb1475e4a449773bad7a5b9a8b9abf4c683a1da" {
		t.Fail()
	}
	hash, err = e.KeyMR()

	if err != nil {
		t.Error(err)
	}
	if hash.String() != "a9fc0b656430d8bf71d180760b0b352c08f45a55a8cf157383613484587b4d21" {
		t.Fail()
	}
}

func newTestingEntryBlock() *EBlock {
	// build an EBlock for testing
	eb := NewEBlock()
	hash := primitives.NewZeroHash()
	hash.SetBytes(byteof(0x11))
	eb.Header.SetChainID(hash)

	hash = primitives.NewZeroHash()
	hash.SetBytes(byteof(0x22))
	eb.Header.SetBodyMR(hash)

	hash = primitives.NewZeroHash()
	hash.SetBytes(byteof(0x33))
	eb.Header.SetPrevKeyMR(hash)

	hash = primitives.NewZeroHash()
	hash.SetBytes(byteof(0x44))
	eb.Header.SetPrevFullHash(hash)

	eb.Header.SetEBSequence(5)
	eb.Header.SetDBHeight(6)
	eb.Header.SetEntryCount(7)
	ha := primitives.NewZeroHash()
	ha.SetBytes(byteof(0xaa))
	hb := primitives.NewZeroHash()
	hb.SetBytes(byteof(0xbb))
	eb.Body.EBEntries = append(eb.Body.EBEntries, ha)
	eb.AddEndOfMinuteMarker(0xcc)
	eb.Body.EBEntries = append(eb.Body.EBEntries, hb)

	return eb
}

func newEntryBlock() *EBlock {
	e := NewEBlock()
	entryStr := "4bf71c177e71504032ab84023d8afc16e302de970e6be110dac20adbf9a1974625f25d9375533b44505964af993212ef7c13314736b2c76a37c73571d89d8b21c6180f7430677d46d93a3e17b68e6a25dc89ecc092cee1459101578859f7f6969d171a092a1d04f067d55628b461c6a106b76b4bc860445f87b0052cdc5f2bfd000002d800001b080000000272d72e71fdee4984ecb30eedcc89cb171d1f5f02bf9a8f10a8b2cfbaf03efe1c0000000000000000000000000000000000000000000000000000000000000001"
	h, err := hex.DecodeString(entryStr)
	if err != nil {
		panic(err)
	}
	err = e.UnmarshalBinary(h)
	if err != nil {
		panic(err)
	}
	return e
}
