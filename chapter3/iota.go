package chapter3

// weeday generator
type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

// flag generator
type Flags uint

const (
	FlagUp Flags = 1 << iota
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func IsUp(v Flags) bool { return v&FlagUp == FlagUp }

// ^ is Bitwise XOR (exclusive OR) operator - if bits are different, result is 1. If bits are same, result is 0
// 1101 ^ 1010 = 0111
// ^ is also Bitwise NOT unary operator - turns 1 to 0 and 0 to 1
// & is Bitwise AND operator - if both bits are 1, result is 1, otherwise 0
// &^ is Bitwise CLEAR operator - a &^ b will clear common bits between a and b
// &^= is Bitwise CLEAR assignment operator
// *v &^= FlagUp is same as *v = *v &^ FlagUp
// This is used to toggle the flag
func TurnDown(v *Flags) { *v &^= FlagUp }

func SetBroadcaset(v *Flags) { *v |= FlagBroadcast }

func IsCast(v Flags) bool { return v&(FlagBroadcast|FlagMulticast) != 0 }

// powers of 1024
// Note constants are untyped int type
const (
	_   = 1 << (10 * iota) // 1 << 1
	KiB                    // 1 << 10; 2^10 or 1024 (~10^3) Kibibyte
	MiB                    // 1 << 20; 2^20 or 1048576 (~10^6) Mebibyte
	GiB                    // 1 << 30; 2^30 or 1073741824 (~10^9) Gibibyte
	TiB                    // 1 << 40; 2^40 or 1099511627776 (~10^12) (exceeds 1 << 32) Tebibyte
	PiB                    // 1 << 50; 2^50 or (~10^15) Pebibyte
	EiB                    // 1 << 60; 2^60 or (~10^18) Exbibyte
	ZiB                    // 1 << 70; 2^70 or (~10^21) (exceeds 1 << 64) Zebibyte
	YiB                    // 1 << 80; 2^80 or (~10^24) Yobibyte
)

// powers of 1000
// Note - we can't use iota here
// Note - constants are untyped float type
const (
	KB = 1e3
	MB = 1e6
	GB = 1e9
	TB = 1e12
	PB = 1e15
	EB = 1e18
	ZB = 1e21
	YB = 1e24
)
