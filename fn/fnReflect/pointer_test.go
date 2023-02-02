package fnReflect

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPointer(test *testing.T) {

	test.Run("IsPointer()", func(t *testing.T) {
		type Str struct {
			Name string
		}

		v := Str{
			Name: "Ciao",
		}

		p := &Str{
			Name: "Lee",
		}

		assert.Equal(t, true, IsStruct(v))
		assert.Equal(t, true, IsStruct(p))

		sp := "hello"
		assert.Equal(t, false, IsStruct(sp))
		assert.Equal(t, false, IsStruct(&sp))
	})

}
