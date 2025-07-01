package users

type UserInterface interface {
	Create(user UserDTO) error
	Read(id int) (*UserDTO, error)
	Update(user UserDTO) error
	Delete(id int) error
	Exists(email string) (*UserDTO, bool)
	List() ([]UserDTO, error)
}
