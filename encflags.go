package evencflags

//F defines event structure or field encryption "status" flags.
type F uint16

const NoneF F = iota

// rest of the flags starting from 1
const (
	CltIPencF      = (F)(1) << iota // client ip is encrypted
	SrvIPencF                       // server ip is encrypted
	CallIDencF                      // call-id is encrypted
	URIencF                         // uri is encrypted
	ReasonAencF                     // reason attr. is enc
	CountryISOencF                  // country iso is enc
	CityIDencF                      // city id is enc
	UAencF                          // user-agent is enc
	IpcipherF                       // ipcipher is used
)

const (
	SrcIPencF = CltIPencF // alias for client ip
	DstIPencF = SrcIPencF // alias for server ip
)
