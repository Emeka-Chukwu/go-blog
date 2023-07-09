package tokenpkg

import "time"

// Maker is an interface  for managing tokens
type Maker interface {
	/// CreateToken creates a new token for a specifica username and duration
	CreateToken(id int32, role string, duration time.Duration) (string, *Payload, error)

	/// VerifiyToken checks if the token is valid or not
	VerifiyToken(token string) (*Payload, error)
}
