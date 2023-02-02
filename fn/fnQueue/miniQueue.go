package miniQueue

import (
	"fmt"
	"github.com/bamtols/fn-go/fn/fnParams"
)

type (
	MiniQueue[REQ, RESP any] struct {
		consumer FnMiniQueueConsumer[REQ, RESP]
		queue    []miniQueue[REQ, RESP]
	}

	FnMiniQueueConsumer[REQ any, RESP any] func(v *REQ) (res *RESP, err error)
	IMiniQueue[REQ any, RESP any]          struct {
		Id        int
		RespQueue chan OMiniQueue[RESP]
		Data      *REQ
	}

	OMiniQueue[RESP any] struct {
		Resp *RESP
		Err  error
	}

	miniQueue[REQ, RESP any] chan IMiniQueue[REQ, RESP]
)

func NewMiniQueue[REQ any, RESP any](queueSize int) *MiniQueue[REQ, RESP] {
	queue := make([]miniQueue[REQ, RESP], 0)

	for i := 0; i < queueSize; i++ {
		queue = append(queue, make(chan IMiniQueue[REQ, RESP]))
	}

	return &MiniQueue[REQ, RESP]{
		consumer: nil,
		queue:    queue,
	}
}

func (x *MiniQueue[REQ, RESP]) Consume(fn FnMiniQueueConsumer[REQ, RESP]) {
	if x.consumer != nil {
		panic(fmt.Errorf("already consumed queue"))
	}

	x.consumer = fn

	for i := range x.queue {
		go func(q chan IMiniQueue[REQ, RESP], queueNumber int) {

			for d := range q {

				resp, err := x.consumer(d.Data)
				d.RespQueue <- OMiniQueue[RESP]{
					Resp: resp,
					Err:  err,
				}
				close(d.RespQueue)
			}
		}(x.queue[i], i)
	}
}

func (x *MiniQueue[REQ, RESP]) Publish(data *REQ, id ...int) (*RESP, error) {
	iQueueData := IMiniQueue[REQ, RESP]{
		Id:        fnParams.Pick(id),
		RespQueue: make(chan OMiniQueue[RESP]),
		Data:      data,
	}

	x.queue[x.hash(iQueueData.Id)] <- iQueueData
	resp := <-iQueueData.RespQueue
	return resp.Resp, resp.Err
}

func (x *MiniQueue[REQ, RESP]) hash(id int) int {
	return id % len(x.queue)
}
