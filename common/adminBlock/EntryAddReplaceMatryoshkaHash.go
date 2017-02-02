package adminBlock

import (
	"bytes"
	"fmt"

	"github.com/FactomProject/factomd/common/constants"
	"github.com/FactomProject/factomd/common/interfaces"
	"github.com/FactomProject/factomd/common/primitives"
)

type AddReplaceMatryoshkaHash struct {
	IdentityChainID interfaces.IHash
	MHash           interfaces.IHash
}

var _ interfaces.Printable = (*AddReplaceMatryoshkaHash)(nil)
var _ interfaces.BinaryMarshallable = (*AddReplaceMatryoshkaHash)(nil)
var _ interfaces.IABEntry = (*AddReplaceMatryoshkaHash)(nil)

func (e *AddReplaceMatryoshkaHash) String() string {
	callTime := time.Now().UnixNano()
	defer entryAddReplaceMatryoshkaHashString.Observe(float64(time.Now().UnixNano() - callTime))	
	var out primitives.Buffer
	out.WriteString(fmt.Sprintf("    E: %35s -- %17s %8x %12s %8s",
		"AddReplaceMatryoshkaHash",
		"IdentityChainID", e.IdentityChainID.Bytes()[3:5],
		"MHash", e.MHash.String()[:8]))
	return (string)(out.DeepCopyBytes())
}

func (m *AddReplaceMatryoshkaHash) Type() byte {
	callTime := time.Now().UnixNano()
	defer entryAddReplaceMatryoshkaHashType.Observe(float64(time.Now().UnixNano() - callTime))	
	return constants.TYPE_ADD_MATRYOSHKA
}

func (c *AddReplaceMatryoshkaHash) UpdateState(state interfaces.IState) error {
	callTime := time.Now().UnixNano()
	defer entryAddReplaceMatryoshkaHashUpdateState.Observe(float64(time.Now().UnixNano() - callTime))	
	state.UpdateAuthorityFromABEntry(c)
	return nil
}

func NewAddReplaceMatryoshkaHash(identityChainID interfaces.IHash, mHash interfaces.IHash) *AddReplaceMatryoshkaHash {
	callTime := time.Now().UnixNano()
	defer entryAddReplaceMatryoshkaHashNewAddReplaceMatryoshkaHash.Observe(float64(time.Now().UnixNano() - callTime))	
	e := new(AddReplaceMatryoshkaHash)
	e.IdentityChainID = identityChainID
	e.MHash = mHash
	return e
}

func (e *AddReplaceMatryoshkaHash) MarshalBinary() (data []byte, err error) {
	callTime := time.Now().UnixNano()
	defer entryAddReplaceMatryoshkaHashMarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))	
	var buf primitives.Buffer

	buf.Write([]byte{e.Type()})
	buf.Write(e.IdentityChainID.Bytes())
	buf.Write(e.MHash.Bytes())

	return buf.DeepCopyBytes(), nil
}

func (e *AddReplaceMatryoshkaHash) UnmarshalBinaryData(data []byte) (newData []byte, err error) {
	callTime := time.Now().UnixNano()
	defer entryAddReplaceMatryoshkaHashUnmarshalBinaryData.Observe(float64(time.Now().UnixNano() - callTime))	
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Error unmarshalling Add Replace Matryoshka Hash: %v", r)
		}
	}()
	newData = data
	if newData[0] != e.Type() {
		return nil, fmt.Errorf("Invalid Entry type")
	}

	newData = newData[1:]
	e.IdentityChainID = new(primitives.Hash)
	newData, err = e.IdentityChainID.UnmarshalBinaryData(newData)
	if err != nil {
		return
	}
	e.MHash = new(primitives.Hash)
	newData, err = e.MHash.UnmarshalBinaryData(newData)
	if err != nil {
		return
	}

	return
}

func (e *AddReplaceMatryoshkaHash) UnmarshalBinary(data []byte) (err error) {
	callTime := time.Now().UnixNano()
	defer entryAddReplaceMatryoshkaHashUnmarshalBinary.Observe(float64(time.Now().UnixNano() - callTime))	
	_, err = e.UnmarshalBinaryData(data)
	return
}

func (e *AddReplaceMatryoshkaHash) JSONByte() ([]byte, error) {
	callTime := time.Now().UnixNano()
	defer entryAddReplaceMatryoshkaHashJSONByte.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSON(e)
}

func (e *AddReplaceMatryoshkaHash) JSONString() (string, error) {
	callTime := time.Now().UnixNano()
	defer entryAddReplaceMatryoshkaHashJSONString.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSONString(e)
}

func (e *AddReplaceMatryoshkaHash) JSONBuffer(b *bytes.Buffer) error {
	callTime := time.Now().UnixNano()
	defer entryAddReplaceMatryoshkaHashJSONBuffer.Observe(float64(time.Now().UnixNano() - callTime))	
	return primitives.EncodeJSONToBuffer(e, b)
}

func (e *AddReplaceMatryoshkaHash) IsInterpretable() bool {
	return false
}

func (e *AddReplaceMatryoshkaHash) Interpret() string {
	return ""
}

func (e *AddReplaceMatryoshkaHash) Hash() interfaces.IHash {
	callTime := time.Now().UnixNano()
	defer entryAddReplaceMatryoshkaHashHash.Observe(float64(time.Now().UnixNano() - callTime))	
	bin, err := e.MarshalBinary()
	if err != nil {
		panic(err)
	}
	return primitives.Sha(bin)
}
