package ui

type BlockInterface interface {
	// Content
	// Content() string
	// SetContent(content string) BlockInterface

	// Children
	Children() []BlockInterface
	SetChildren(children []BlockInterface) BlockInterface
	AddChild(child BlockInterface) BlockInterface
	AddChildren(children []BlockInterface) BlockInterface

	// ID
	ID() string
	SetID(id string) BlockInterface

	// Parameters
	Parameters() map[string]string
	Parameter(key string) string
	SetParameter(key string, value string) BlockInterface
	SetParameters(map[string]string) BlockInterface

	// Type
	Type() string
	SetType(blockType string) BlockInterface

	// Serialization
	ToMap() map[string]interface{}
	ToJson() (string, error)
	ToJsonObject() blockJsonObject
	ToJsonPretty() (string, error)
}
