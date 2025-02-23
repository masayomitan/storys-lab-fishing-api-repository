package model

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"storys-lab-fishing-api/app/domain"
)

func (a PrefectureSQL) Find(ctx context.Context) ([]domain.Prefecture, error) {
	var json = make([]domain.Prefecture, 0)
	var prefectures = make([]domain.Prefecture, 0)

	if err := a.db.FindPrefectureORM(ctx, a.tableName, &json); err != nil {
		return []domain.Prefecture{}, errors.Wrap(err, "error listing prefectures")
	}

	for _, json := range json {
		var prefecture = domain.NewPrefecture(
			json.ID,
			json.Name,
			json.NameKana,
			json.ImageUrl,
	
			convertAreas(json.Areas),
		)
		prefectures = append(prefectures, prefecture)
	}


	return prefectures, nil
}

func (a PrefectureSQL) FindOne(ctx context.Context, id int) (domain.Prefecture, error) {
	var prefJSON = domain.Prefecture{}
	fmt.Println(id)

	if err := a.db.FindOnePrefectureORM(ctx, a.tableName, id, &prefJSON); err != nil {
		return domain.Prefecture{}, errors.Wrap(err, "error listing prefectures")
	}

	var pref = domain.NewPrefecture(
		prefJSON.ID,
		prefJSON.Name,
		prefJSON.NameKana,
		prefJSON.ImageUrl,

		convertAreas(prefJSON.Areas),
	)
	fmt.Println(pref)

	return pref, nil
}

func (a PrefectureSQL) FindByAdmin(ctx context.Context) ([]domain.Prefecture, error) {
	var json = make([]domain.Prefecture, 0)
	var prefectures = make([]domain.Prefecture, 0)

	if err := a.db.FindPrefectureORM(ctx, a.tableName, &json); err != nil {
		return []domain.Prefecture{}, errors.Wrap(err, "error listing prefectures")
	}

	for _, json := range json {
		var prefecture = domain.NewPrefecture(
			json.ID,
			json.Name,
			json.NameKana,
			json.ImageUrl,
	
			convertAreas(json.Areas),
		)
		prefectures = append(prefectures, prefecture)
	}


	return prefectures, nil
}

func (a PrefectureSQL) FindOneByAdmin(ctx context.Context, id int) (domain.Prefecture, error) {
	var prefJSON = domain.Prefecture{}
	fmt.Println(id)

	if err := a.db.FindOnePrefectureORM(ctx, a.tableName, id, &prefJSON); err != nil {
		return domain.Prefecture{}, errors.Wrap(err, "error listing prefectures")
	}

	var pref = domain.NewPrefecture(
		prefJSON.ID,
		prefJSON.Name,
		prefJSON.NameKana,
		prefJSON.ImageUrl,

		convertAreas(prefJSON.Areas),
	)
	fmt.Println(pref)

	return pref, nil
}

func (ga *GormAdapter) FindPrefectureORM(ctx context.Context, table string, result interface{}) error {
	return ga.DB.Table(table).
		Where("deleted_at IS NULL").
		Order("id asc").
		Preload("Areas").
		// Preload("Areas.Images").
		Find(result).Error
}


func (ga *GormAdapter) FindOnePrefectureORM(ctx context.Context, table string, id int, result interface{}) error {
	return ga.DB.Table(table).Where("id = ?", id).
		Where("deleted_at IS NULL").
		Order("id asc").
		Preload("Areas").
		// Preload("Areas.Images").
		First(result).Error
}

