package resources

//ResourceType is used as type for resources - it's enum
type ResourceType int

const (
	//Graphics is used as resource type for graphics
	Graphics ResourceType = iota
	//Sound is used as resource type for sound
	Sound
	//Other is used as resource tyoe for other
	Other
)
