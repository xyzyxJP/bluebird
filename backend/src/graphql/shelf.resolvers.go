package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.45

import (
	"context"
	"fmt"

	ulid "github.com/oklog/ulid/v2"
	"github.com/xyzyxJP/bluebird/src/model"
)

// CreateShelfItem is the resolver for the createShelfItem field.
func (r *mutationResolver) CreateShelfItem(ctx context.Context, name string, categoryUlid string, tagsUlid []string, locationUlid string, description string) (*ShelfItem, error) {
	ulid := ulid.Make().String()
	var shelfCategory model.ShelfCategory
	if err := r.DB.Where("ulid = ?", categoryUlid).First(&shelfCategory).Error; err != nil {
		return nil, err
	}
	var shelfTags []model.ShelfTag
	var parsedShelfTags []*ShelfTag
	for _, tagUlid := range tagsUlid {
		var shelfTag model.ShelfTag
		if err := r.DB.Where("ulid = ?", tagUlid).First(&shelfTag).Error; err != nil {
			return nil, err
		}
		shelfTags = append(shelfTags, shelfTag)
		parsedShelfTags = append(parsedShelfTags, &ShelfTag{Ulid: shelfTag.Ulid, Name: shelfTag.Name})
	}
	var shelfLocation model.ShelfLocation
	if err := r.DB.Where("ulid = ?", locationUlid).First(&shelfLocation).Error; err != nil {
		return nil, err
	}
	shelfItem := model.ShelfItem{
		Ulid:        ulid,
		Name:        name,
		CategoryID:  shelfCategory.ID,
		LocationID:  shelfLocation.ID,
		Description: description,
	}
	if err := r.DB.Create(&shelfItem).Association("Tags").Append(&shelfTags); err != nil {
		return nil, err
	}
	return &ShelfItem{
		Ulid:        ulid,
		Name:        name,
		Category:    &ShelfCategory{Ulid: categoryUlid},
		Tags:        parsedShelfTags,
		Location:    &ShelfLocation{Ulid: locationUlid},
		Description: description,
	}, nil
}

// UpdateShelfItem is the resolver for the updateShelfItem field.
func (r *mutationResolver) UpdateShelfItem(ctx context.Context, ulid string, name *string, categoryUlid *string, tags []string, locationUlid *string, description *string) (*ShelfItem, error) {
	panic(fmt.Errorf("not implemented: UpdateShelfItem - updateShelfItem"))
}

// DeleteShelfItem is the resolver for the deleteShelfItem field.
func (r *mutationResolver) DeleteShelfItem(ctx context.Context, ulid string) (bool, error) {
	if err := r.DB.Where("ulid = ?", ulid).Delete(&model.ShelfItem{}).Error; err != nil {
		return false, err
	}
	return true, nil
}

// CreateShelfCategory is the resolver for the createShelfCategory field.
func (r *mutationResolver) CreateShelfCategory(ctx context.Context, name string) (*ShelfCategory, error) {
	shelfCategory := model.ShelfCategory{Ulid: ulid.Make().String(), Name: name}
	if err := r.DB.Create(&shelfCategory).Error; err != nil {
		return nil, err
	}
	return &ShelfCategory{Ulid: shelfCategory.Ulid, Name: shelfCategory.Name}, nil
}

// UpdateShelfCategory is the resolver for the updateShelfCategory field.
func (r *mutationResolver) UpdateShelfCategory(ctx context.Context, ulid string, name *string) (*ShelfCategory, error) {
	if err := r.DB.Model(&model.ShelfCategory{}).Where("ulid = ?", ulid).Update("name", name).Error; err != nil {
		return nil, err
	}
	return &ShelfCategory{Ulid: ulid, Name: *name}, nil
}

// DeleteShelfCategory is the resolver for the deleteShelfCategory field.
func (r *mutationResolver) DeleteShelfCategory(ctx context.Context, ulid string) (bool, error) {
	if err := r.DB.Where("ulid = ?", ulid).Delete(&model.ShelfCategory{}).Error; err != nil {
		return false, err
	}
	return true, nil
}

// CreateShelfTag is the resolver for the createShelfTag field.
func (r *mutationResolver) CreateShelfTag(ctx context.Context, name string) (*ShelfTag, error) {
	shelfTag := model.ShelfTag{Ulid: ulid.Make().String(), Name: name}
	if err := r.DB.Create(&shelfTag).Error; err != nil {
		return nil, err
	}
	return &ShelfTag{Ulid: shelfTag.Ulid, Name: shelfTag.Name}, nil
}

// UpdateShelfTag is the resolver for the updateShelfTag field.
func (r *mutationResolver) UpdateShelfTag(ctx context.Context, ulid string, name *string) (*ShelfTag, error) {
	if err := r.DB.Model(&model.ShelfTag{}).Where("ulid = ?", ulid).Update("name", name).Error; err != nil {
		return nil, err
	}
	return &ShelfTag{Ulid: ulid, Name: *name}, nil
}

// DeleteShelfTag is the resolver for the deleteShelfTag field.
func (r *mutationResolver) DeleteShelfTag(ctx context.Context, ulid string) (bool, error) {
	if err := r.DB.Where("ulid = ?", ulid).Delete(&model.ShelfTag{}).Error; err != nil {
		return false, err
	}
	return true, nil
}

// CreateShelfLocation is the resolver for the createShelfLocation field.
func (r *mutationResolver) CreateShelfLocation(ctx context.Context, name string) (*ShelfLocation, error) {
	shelfLocation := model.ShelfLocation{Ulid: ulid.Make().String(), Name: name}
	if err := r.DB.Create(&shelfLocation).Error; err != nil {
		return nil, err
	}
	return &ShelfLocation{Ulid: shelfLocation.Ulid, Name: shelfLocation.Name}, nil
}

// UpdateShelfLocation is the resolver for the updateShelfLocation field.
func (r *mutationResolver) UpdateShelfLocation(ctx context.Context, ulid string, name *string) (*ShelfLocation, error) {
	if err := r.DB.Model(&model.ShelfLocation{}).Where("ulid = ?", ulid).Update("name", name).Error; err != nil {
		return nil, err
	}
	return &ShelfLocation{Ulid: ulid, Name: *name}, nil
}

// DeleteShelfLocation is the resolver for the deleteShelfLocation field.
func (r *mutationResolver) DeleteShelfLocation(ctx context.Context, ulid string) (bool, error) {
	if err := r.DB.Where("ulid = ?", ulid).Delete(&model.ShelfLocation{}).Error; err != nil {
		return false, err
	}
	return true, nil
}

// ShelfItems is the resolver for the shelfItems field.
func (r *queryResolver) ShelfItems(ctx context.Context) ([]*ShelfItem, error) {
	shelfItems := []model.ShelfItem{}
	if err := r.DB.Find(&shelfItems).Error; err != nil {
		return nil, err
	}
	var parsedShelfItems []*ShelfItem
	for _, shelfItem := range shelfItems {
		category := model.ShelfCategory{}
		if err := r.DB.Where("id = ?", shelfItem.CategoryID).First(&category).Error; err != nil {
			return nil, err
		}
		tags := []model.ShelfTag{}
		if err := r.DB.Model(&shelfItem).Association("Tags").Find(&tags); err != nil {
			return nil, err
		}
		var parsedTags []*ShelfTag
		for _, tag := range tags {
			parsedTags = append(parsedTags, &ShelfTag{Ulid: tag.Ulid, Name: tag.Name})
		}
		location := model.ShelfLocation{}
		if err := r.DB.Where("id = ?", shelfItem.LocationID).First(&location).Error; err != nil {
			return nil, err
		}
		parsedShelfItems = append(parsedShelfItems, &ShelfItem{
			Ulid:        shelfItem.Ulid,
			Name:        shelfItem.Name,
			Category:    &ShelfCategory{Ulid: category.Ulid, Name: category.Name},
			Tags:        parsedTags,
			Location:    &ShelfLocation{Ulid: location.Ulid, Name: location.Name},
			Description: shelfItem.Description,
		})
	}
	return parsedShelfItems, nil
}

// ShelfItem is the resolver for the shelfItem field.
func (r *queryResolver) ShelfItem(ctx context.Context, ulid string) (*ShelfItem, error) {
	shelfItem := model.ShelfItem{}
	if err := r.DB.Where("ulid = ?", ulid).First(&shelfItem).Error; err != nil {
		return nil, err
	}
	category := model.ShelfCategory{}
	if err := r.DB.Where("id = ?", shelfItem.CategoryID).First(&category).Error; err != nil {
		return nil, err
	}
	tags := []model.ShelfTag{}
	if err := r.DB.Model(&shelfItem).Association("Tags").Find(&tags); err != nil {
		return nil, err
	}
	var parsedTags []*ShelfTag
	for _, tag := range tags {
		parsedTags = append(parsedTags, &ShelfTag{Ulid: tag.Ulid, Name: tag.Name})
	}
	location := model.ShelfLocation{}
	if err := r.DB.Where("id = ?", shelfItem.LocationID).First(&location).Error; err != nil {
		return nil, err
	}
	return &ShelfItem{
		Ulid:        shelfItem.Ulid,
		Name:        shelfItem.Name,
		Category:    &ShelfCategory{Ulid: category.Ulid, Name: category.Name},
		Tags:        parsedTags,
		Location:    &ShelfLocation{Ulid: location.Ulid, Name: location.Name},
		Description: shelfItem.Description,
	}, nil
}

// ShelfCategories is the resolver for the shelfCategories field.
func (r *queryResolver) ShelfCategories(ctx context.Context) ([]*ShelfCategory, error) {
	shelfCategories := []*model.ShelfCategory{}
	if err := r.DB.Find(&shelfCategories).Error; err != nil {
		return nil, err
	}
	var parsedShelfCategories []*ShelfCategory
	for _, shelfCategory := range shelfCategories {
		parsedShelfCategories = append(parsedShelfCategories, &ShelfCategory{Ulid: shelfCategory.Ulid, Name: shelfCategory.Name})
	}
	return parsedShelfCategories, nil
}

// ShelfCategory is the resolver for the shelfCategory field.
func (r *queryResolver) ShelfCategory(ctx context.Context, ulid string) (*ShelfCategory, error) {
	shelfCategory := model.ShelfCategory{}
	if err := r.DB.Where("ulid = ?", ulid).First(&shelfCategory).Error; err != nil {
		return nil, err
	}
	return &ShelfCategory{Ulid: shelfCategory.Ulid, Name: shelfCategory.Name}, nil
}

// ShelfTags is the resolver for the shelfTags field.
func (r *queryResolver) ShelfTags(ctx context.Context) ([]*ShelfTag, error) {
	shelfTags := []model.ShelfTag{}
	if err := r.DB.Find(&shelfTags).Error; err != nil {
		return nil, err
	}
	var parsedShelfTags []*ShelfTag
	for _, shelfTag := range shelfTags {
		parsedShelfTags = append(parsedShelfTags, &ShelfTag{Ulid: shelfTag.Ulid, Name: shelfTag.Name})
	}
	return parsedShelfTags, nil
}

// ShelfTag is the resolver for the shelfTag field.
func (r *queryResolver) ShelfTag(ctx context.Context, ulid string) (*ShelfTag, error) {
	shelfTag := model.ShelfTag{}
	if err := r.DB.Where("ulid = ?", ulid).First(&shelfTag).Error; err != nil {
		return nil, err
	}
	return &ShelfTag{Ulid: shelfTag.Ulid, Name: shelfTag.Name}, nil
}

// ShelfLocations is the resolver for the shelfLocations field.
func (r *queryResolver) ShelfLocations(ctx context.Context) ([]*ShelfLocation, error) {
	shelfLocations := []model.ShelfLocation{}
	if err := r.DB.Find(&shelfLocations).Error; err != nil {
		return nil, err
	}
	var parsedShelfLocations []*ShelfLocation
	for _, shelfLocation := range shelfLocations {
		parsedShelfLocations = append(parsedShelfLocations, &ShelfLocation{Ulid: shelfLocation.Ulid, Name: shelfLocation.Name})
	}
	return parsedShelfLocations, nil
}

// ShelfLocation is the resolver for the shelfLocation field.
func (r *queryResolver) ShelfLocation(ctx context.Context, ulid string) (*ShelfLocation, error) {
	shelfLocation := model.ShelfLocation{}
	if err := r.DB.Where("ulid = ?", ulid).First(&shelfLocation).Error; err != nil {
		return nil, err
	}
	return &ShelfLocation{Ulid: shelfLocation.Ulid, Name: shelfLocation.Name}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
