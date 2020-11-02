package resources

//ResourceMap is map used for resources
type ResourceMap struct {
	Resources map[string]*Resource
}

//InitResourceMap is used to init resource map
func (resourceMap *ResourceMap) InitResourceMap() {
	resourceMap.Resources = make(map[string]*Resource)
}
