package mysql

import (
	"context"
	"log"
	"strconv"

	"github.com/nashrull/averin/shared/config"
	siswa "github.com/nashrull/averin/siswa"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlRepository struct {
	db     *gorm.DB
	log    *log.Logger
	config config.Config
}

// Create implements siswa.UserRepository.
func (r *mysqlRepository) Create(ctx context.Context, payload siswa.Siswa) (insertedID uint, err error) {
	result := r.db.WithContext(ctx).Create(&payload)
	if result.Error != nil {
		return payload.ID, err
	}
	r.log.Print("user has been created ", payload.ID)
	return payload.ID, nil
}

// Delete implements siswa.UserRepository.
func (r *mysqlRepository) Delete(ctx context.Context, id string) (err error) {
	// delete user
	tx := r.db.WithContext(ctx).Delete(siswa.Siswa{}, id)
	if tx.Error != nil {
		r.log.Println("error delete user ", id, " \t", err.Error())
		return err
	}
	return nil
}

// List implements siswa.UserRepository.
func (r *mysqlRepository) List(ctx context.Context, payload siswa.Siswa) (response []siswa.Siswa, err error) {
	err = r.db.WithContext(ctx).Find(&response).Error
	if err != nil {
		return response, err
	}
	return response, nil
}

// Update implements siswa.UserRepository.
func (r *mysqlRepository) Save(ctx context.Context, id string, payload siswa.Siswa) (err error) {
	ui, _ := strconv.Atoi(id)
	payload.ID = uint(ui)

	result := r.db.WithContext(ctx).Save(&payload)
	return result.Error
}

func NewMysqlRepository(log *log.Logger, config config.Config) (response siswa.SiswaRepository, err error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		// Replace with your MySQL configuration
		DriverName: config.DB.Driver,
		DSN:        config.DB.CreateConnection(),
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = db.AutoMigrate(&siswa.Siswa{})
	if err != nil {
		log.Println(err.Error())
		return response, err
	}
	db.Logger.LogMode(1)
	log.Println("Auto Migrate siswas success")
	return &mysqlRepository{db: db, log: log}, nil
}
