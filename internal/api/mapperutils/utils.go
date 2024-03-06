package mapperutils

import "gorm.io/gorm/schema"

func EntityListToDtoList[E schema.Tabler, D any](entities []E, mappingFunc func(E) D) []D {
	var dtoList []D
	for i := range entities {
		dtoList = append(dtoList, mappingFunc(entities[i]))
	}
	return dtoList
}

func DtoListToEntityList[D any, E schema.Tabler](dtoList []D, id string, mappingFunc func(D, string) E) []E {
	var entities []E
	for i := range dtoList {
		entities = append(entities, mappingFunc(dtoList[i], id))
	}
	return entities
}
func EntityListToEntityList[E schema.Tabler, string, F schema.Tabler](entities []E, id string, mappingFunc func(E, string) F) []F {
	var resEntityList []F
	for i := range entities {
		resEntityList = append(resEntityList, mappingFunc(entities[i], id))
	}
	return resEntityList
}
