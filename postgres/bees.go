package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/jmagoga/new-equimper-go-graphql/graph/model"
)

type BeesRepo struct {
	DB *pg.DB
}

func (b *BeesRepo) GetBees() ([]*model.Bee, error) {
	var bees []*model.Bee
	err := b.DB.Model(&bees).Order("id").Select()
	if err != nil {
		return nil, err
	}

	return bees, nil
}

func (b *BeesRepo) CreateBee(bee *model.Bee) (*model.Bee, error) {
	_, err := b.DB.Model(bee).Returning("*").Insert()

	return bee, err
}

func (b * BeesRepo) GetBeeById(id string) (*model.Bee, error) {
	var bee model.Bee
	err := b.DB.Model(&bee).Where("id = ?", id).First()
	return &bee, err
}

func (b * BeesRepo) UpdateBee(bee *model.Bee) (*model.Bee, error) {
	_, err := b.DB.Model(bee).Where("id = ?", bee.ID).Update()
	return bee, err
}

func (b *BeesRepo) Delete(bee *model.Bee) error {
	_, err := b.DB.Model(bee).Where("id = ?",bee.ID).Delete()
	return err

}

// to get only ONE bee.

/* func (b *BeesRepo) GetBeeById(id string) (*model.Bee, error) {
	var bee model.Bee
	err := b.DB.Model(&bee).Where( condition: "id = ?", id).First()
	if err != nil {
		return nil, err
	}

	return &bee, nil
} */
