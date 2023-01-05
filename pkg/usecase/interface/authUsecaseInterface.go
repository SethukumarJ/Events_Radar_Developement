package interfaces

//AuthUseCase is the interface for authentication usecase
type authUsecase interface {
	VerifyUser(email string, password string) error
}