package interactor

// UseCase Interface use to create UseCases
type UseCase interface {
	Execute(ch chan interface{})
}
