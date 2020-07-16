package categories

type PublicCategory struct {
	ID int `json:"id"`
}

func (categories Categories) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(categories))
	for index, category := range categories {
		result[index] = category.Marshall(isPublic)
	}

	return result
}

func (category Category) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicCategory{
			ID: category.ID,
		}
	}

	return category
}
