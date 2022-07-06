package osuhelper

func ModeToStr(mode int8) string {
	if mode == 3 {
		return "mania"
	}
	if mode == 2 {
		return "ctb"
	}
	if mode == 1 {
		return "taiko"
	}
	return "std"
}
