package main

type Header struct {
	ID      uint16
	QR      bool
	Opcode  Opcode
	AA      bool
	TC      bool
	RD      bool
	RA      bool
	Z       bool
	Rcode   Rcode
	Qdcount uint16
	Ancount uint16
	Nscount uint16
	Arcount uint16
}

//go:generate stringer -type=Opcode -linecomment
type Opcode uint16

const (
	OpcodeQuery  Opcode = 0 // QUERY
	OpcodeStatus Opcode = 2 // STATUS
	OpcodeNotify Opcode = 4 // NOTIFY
	OpcodeUpdate Opcode = 5 // UPDATE
	OpcodeDso    Opcode = 6 // DSO
)

// For Header/OPT
//
//go:generate stringer -type=Rcode -linecomment
type Rcode uint16

const (
	RcodeNoError   Rcode = 0  // NOERROR
	RcodeFormErr   Rcode = 1  // FORMERR
	RcodeServFail  Rcode = 2  // SERVFAIL
	RcodeNXDomain  Rcode = 3  // NXDOMAIN
	RcodeNotImp    Rcode = 4  // NOTIMP
	RcodeRefused   Rcode = 5  // REFUSED
	RcodeYXDomain  Rcode = 6  // YXDOMAIN
	RcodeYXRRSet   Rcode = 7  // YXRRSET
	RcodeNXRRSet   Rcode = 8  // NXRRSET
	RcodeNotAuth   Rcode = 9  // NOTAUTH
	RcodeNotZone   Rcode = 10 // NOTZONE
	RcodeDsoTypeNi Rcode = 11 // DSOTYPENI
	RcodeBadVers   Rcode = 16 // BADVERS
	RcodeBadCookie Rcode = 23 // BADCOOKIE
)

// For the Error field in TSIG/TKEY RR
//
//go:generate stringer -type=TSIGError -linecomment
type TSIGError uint16

const (
	TSIGBadSig   TSIGError = 16 // BADSIG
	TSIGBadKey   TSIGError = 17 // BADKEY
	TSIGBadTime  TSIGError = 18 // BADTIME
	TSIGBadMode  TSIGError = 19 // BADMODE
	TSIGBadName  TSIGError = 20 // BADNAME
	TSIGBadAlg   TSIGError = 21 // BADALG
	TSIGBadTrunc TSIGError = 22 // BADTRUNC
)

func (r Rcode) HeaderBits() uint8   { return uint8(r & 0x0F) }
func (r Rcode) ExtendedBits() uint8 { return uint8(r >> 4) }
