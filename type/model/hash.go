package model

// Hash represents a JSON object
type Hash map[string]interface{}

// Map returns a map object casted from the hash
func (h Hash) Map() map[string]interface{} {
	return map[string]interface{}(h)
}
