package datasource

import "Rein/data/cache"

// DatabaseDepartmentDataStore Implementation of the DepartmentDataStore
type DatabaseDepartmentDataStoreFactory struct {
	departmentCache cache.DepartmentCache
}

// NewDatabaseDepartmentDataStoreFactory Constructor for a new Department Data Store using a database
func NewDatabaseDepartmentDataStoreFactory() DatabaseDepartmentDataStoreFactory {
	return DatabaseDepartmentDataStoreFactory{departmentCache: cache.NewDepartmentCacheImpl()}
}

// Create Build a factory for the datastore
func (d *DatabaseDepartmentDataStoreFactory) Create() DepartmentDataStore {
	departmentDataStore := NewDatabaseDepartmentDataStore(d.departmentCache)
	return departmentDataStore
}
