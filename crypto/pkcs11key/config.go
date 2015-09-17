package pkcs11key

// Config contains configuration information required to use a PKCS
// #11 key.
type Config struct {
	Module          string
	SlotDescription string
	TokenLabel      string
	PIN             string
	PrivateKeyLabel string
}
