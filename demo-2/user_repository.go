package main

import "gorm.io/gorm"

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindAll() []string {
	// 这里为了简化，直接返回假数据
	return []string{"Alice", "Bob", "Charlie"}
}
