package pkg_interfaces

type PasswordHasher interface {
	GenerateHash(password string) (string, error)
	CheckPassword(password string, hashedPassword string) error
}
