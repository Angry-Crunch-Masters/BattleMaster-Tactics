package basic

//EntityManager is entity manager
type EntityManager struct {
	creators map[EntityType]IEntityCreator
	entities []IEntity
}

func CreateEntityManager() *EntityManager {
	manager := &EntityManager{}
	manager.creators = make(map[EntityType]IEntityCreator)
	return manager
}

func (manager *EntityManager) AddEntityCreator(creator IEntityCreator, entityType EntityType) {
	if _, ok := manager.creators[entityType]; !ok {
		manager.creators[entityType] = creator
	}
}

func (manager *EntityManager) AddEntity(entityType EntityType, x, y int, resource string, data interface{}) IEntity {
	if _, ok := manager.creators[entityType]; ok {
		entity := manager.creators[entityType].CreateEntity(x, y, resource, data)
		manager.entities = append(manager.entities, entity)
		return entity
	}
	return nil
}
