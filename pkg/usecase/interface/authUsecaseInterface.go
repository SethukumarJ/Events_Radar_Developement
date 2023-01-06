package interfaces

//AuthUseCase is the interface for authentication usecase
type AuthUsecase interface {
	VerifyUser(email string, password string) error
	VerifyAdmin(email, password string) error
}