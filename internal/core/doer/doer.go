package doer

type Doer interface {
	Do(string) (string, error)

	Finish()
}
