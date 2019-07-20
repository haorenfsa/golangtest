package golangtest

import (
	"testing"

	"github.com/golang/mock/gomock"
)

type toTestMocks struct {
	ctl *gomock.Controller
	dep *Mockdependency
}

func setupToTestMocks(t *testing.T) *toTestMocks {
	ret := new(toTestMocks)
	ret.ctl = gomock.NewController(t)
	ret.dep = NewMockdependency(ret.ctl)
	return ret
}

func (t *toTestMocks) teardown() {
	t.ctl.Finish()
}

func Test_toTest_toTest(t *testing.T) {
	tests := []struct {
		name   string
		mock   func(*toTestMocks)
		assert func(ret error)
	}{
		// TODO: Add test cases.
		{
			name: "test sample",
			mock: func(mocks *toTestMocks) {
				mocks.dep.EXPECT().foo()
			},
			assert: func(ret error) {},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mocks := setupToTestMocks(t)
			defer mocks.teardown()
			toTest := &toTest{
				dependency: mocks.dep,
			}
			if tt.mock != nil {
				tt.mock(mocks)
			}
			err := toTest.toTest()
			if tt.assert != nil {
				tt.assert(err)
			}
		})
	}
}
