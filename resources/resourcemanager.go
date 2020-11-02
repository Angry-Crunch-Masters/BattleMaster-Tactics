package resources

//ResourceManager is used as manager for various resources
type ResourceManager struct {
	resourcesStack map[ResourceType]*ResourceMap
}

//InitResourceManager inits resource manager
func (manager *ResourceManager) InitResourceManager() {
	manager.resourcesStack = make(map[ResourceType]*ResourceMap)
}

func (manager *ResourceManager) addResourceMap(resourceType ResourceType) {
	manager.resourcesStack[resourceType] = &ResourceMap{}
	manager.resourcesStack[resourceType].InitResourceMap()
}

//AddResource adds resource
func (manager *ResourceManager) AddResource(item interface{}, name string, resourceType ResourceType) {
	if _, ok := manager.resourcesStack[resourceType]; !ok {
		manager.addResourceMap(resourceType)
	}
	manager.resourcesStack[resourceType].Resources[name] = &Resource{Object: item}
}

//GetResource is used to get resource from manager
func (manager *ResourceManager) GetResource(name string, resourceType ResourceType) *Resource {
	if _, ok := manager.resourcesStack[resourceType]; ok {
		if _, ok = manager.resourcesStack[resourceType].Resources[name]; ok {
			return manager.resourcesStack[resourceType].Resources[name]
		}
	}
	return nil
}
