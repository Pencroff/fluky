package fluky

import "github.com/stretchr/testify/mock"

type RngMock struct {
	mock.Mock
}

func (r *RngMock) Uint32() uint32 {
	//TODO implement me
	panic("implement me")
}

func (r *RngMock) Int31() int32 {
	//TODO implement me
	panic("implement me")
}

func (r *RngMock) Int() int {
	//TODO implement me
	panic("implement me")
}

func (r *RngMock) Int63n(n int64) int64 {
	//TODO implement me
	panic("implement me")
}

func (r *RngMock) Int31n(n int32) int32 {
	//TODO implement me
	panic("implement me")
}

func (r *RngMock) Intn(n int) int {
	//TODO implement me
	panic("implement me")
}

func (r *RngMock) Float32() float32 {
	//TODO implement me
	panic("implement me")
}

func (r *RngMock) Perm(n int) []int {
	//TODO implement me
	panic("implement me")
}

func (r *RngMock) Shuffle(n int, swap func(i int, j int)) {
	//TODO implement me
	panic("implement me")
}

func (r *RngMock) Read(p []byte) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (r *RngMock) Seed(v int64) {
	r.Called(v)
}

func (r *RngMock) Uint64() uint64 {
	return r.Called().Get(0).(uint64)
}

func (r *RngMock) Int63() int64 {
	return r.Called().Get(0).(int64)
}

func (r *RngMock) Float64() float64 {
	return r.Called().Get(0).(float64)
}
