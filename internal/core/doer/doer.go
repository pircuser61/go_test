package doer

type Doer interface {
	DoIt(string) (string, error)

	Finish()
}
