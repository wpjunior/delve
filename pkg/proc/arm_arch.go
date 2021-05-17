package proc

// ARMArch returns an initialized ARM64
// struct.
func ARMArch(goos string) *Arch {
	// TODO: understand arm details
	return &Arch{
		Name: "arm",
	}
}
