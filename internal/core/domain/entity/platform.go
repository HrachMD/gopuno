package entity

type Platform int

const (
	_ Platform = iota
	PlatformIOS
	PlatformAndroid
	PlatformHuawei
	PlatformWEB
)

func (p Platform) Int() int {
	return int(p)
}
