package resources

type IResourceManager interface {
	GetResource(name string, resourceType ResourceType) *Resource
}
