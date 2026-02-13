//go:build !goexperiment.runtimesecret

package internal

// SecretDo calls f directly when the runtimesecret experiment is not enabled.
// Manual Clear calls still provide best-effort memory zeroing in this case.
func SecretDo(f func()) {
	f()
}

func SecretEnabled() bool {
	return false
}
