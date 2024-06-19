package main

func changeString(a *string) {
	*a = "Goodbye, world!"
}
func changeInt(a *int) {
	*a = 20
}
func changeInt16(a *int16) {
	*a = 19
}
func changeByte(a *byte) {
	*a = 9
}
func changeBool(a *bool) {
	*a = false
}
func changeRune(a *rune) {
	*a = 23
}
func changeFloat64(a *float64) {
	*a = 46.4646
}
func changeFloat(a *float64) {
	*a = 6.28
}

func main() {
	var (
		valString  string  = "string"
		valInt8            = -8
		valInt16   int16   = 16
		valUint8   byte    = 8
		valBool    bool    = true
		valInt32   rune    = 32
		valFloat64 float64 = 64.6464
		valFloat32         = 32.3232
	)
	changeString(&valString)
	changeInt(&valInt8)
	changeInt16(&valInt16)
	changeByte(&valUint8)
	changeBool(&valBool)
	changeRune(&valInt32)
	changeFloat64(&valFloat64)
	changeFloat(&valFloat32)

}
