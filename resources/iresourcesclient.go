package resources

type IResourcesClient interface {
	SetResourcesProvider(provider IResourceManager)
}
