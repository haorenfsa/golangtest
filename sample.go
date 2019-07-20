package golangtest

//go:generate mockgen -source ./sample.go -destination=./sample_mock.go -package golangtest

type dependency interface {
	foo() error
}

type toTest struct {
	dependency
}

func (t *toTest) toTest() error {
	err := t.dependency.foo()
	if err != nil {
		return err
	}
	return nil
}
