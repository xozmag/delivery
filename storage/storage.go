package storage

import (
	"delivery/configs"
	"delivery/db"
	"delivery/storage/postgres"
	"delivery/storage/repo"
	"log"
)

type Storage interface {
	Admin() repo.IAdminStorage
}

type storage struct {
	adminRepo repo.IAdminStorage
}

// New
func New(cfg *configs.Configuration) *storage {
	postgresDB, err := db.Init(cfg)
	if err != nil {
		log.Fatalf("error connecting to postgres database: %v", err)
	}
	return &storage{
		adminRepo: postgres.NewAdmin(postgresDB),
	}
}

// Admin returns admin repository
func (s storage) Admin() repo.IAdminStorage {
	return s.adminRepo
}
