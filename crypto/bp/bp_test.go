package bp

import (
	"fmt"
	"testing"
)

func TestBulletProofs(t *testing.T) {
	t.Run("TestParams", func(t *testing.T) {
		params := GenParametersSecp256k1(64)
		fmt.Println(params)
	})
}
