package miniQueue

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

const testSize = 10000

func TestMiniQueue(test *testing.T) {
	type (
		Req struct {
			A int
			B int
		}

		Resp struct {
			C int
		}
	)

	test.Run("singleTest", func(t *testing.T) {
		m := NewMiniQueue[Req, Resp](10)

		m.Consume(func(v *Req) (res *Resp, err error) {
			res = &Resp{
				C: v.A + v.B,
			}
			return
		})

		req := &Req{
			A: 1,
			B: 2,
		}

		resp, err := m.Publish(req, 1)

		if err != nil {
			t.Error(err)
		}

		assert.Equal(t, req.A+req.B, resp.C)
	})

	test.Run("multiTest", func(t *testing.T) {
		m := NewMiniQueue[Req, Resp](10)

		m.Consume(func(v *Req) (res *Resp, err error) {
			res = &Resp{
				C: v.A + v.B,
			}
			return
		})

		for i := 0; i < testSize; i++ {
			go func(fn func(data *Req, id ...int) (*Resp, error)) {
				req := &Req{
					A: rand.Intn(10000),
					B: rand.Intn(10000),
				}

				resp, err := fn(req, rand.Intn(1000000))
				if err != nil {
					panic(err)
				}

				assert.Equal(t, req.A+req.B, resp.C)

			}(m.Publish)
		}
	})

}
