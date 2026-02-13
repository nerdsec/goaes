//go:build goexperiment.runtimesecret

package internal

import "runtime/secret"

// SecretDo runs f inside secret mode, which zeroes registers, stack, and
// heap allocations made by f once they become unreachable.
func SecretDo(f func()) {
	secret.Do(f)
}

func SecretEnabled() bool {
	return secret.Enabled()
}
