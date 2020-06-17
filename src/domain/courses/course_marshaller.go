package courses

type PublicCourse struct {
	ID          int    `json:"id"`
	CategoryID  int    `json:"category_id"`
	DateCreated string `json:"date_created"`
}

func (courses Courses) Marshall(isPublic bool) []interface{} {
	result := make([]interface{}, len(courses))
	for index, course := range courses {
		result[index] = course.Marshall(isPublic)
	}

	return result
}

func (course Course) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicCourse{
			ID:          course.ID,
			CategoryID:  course.CategoryID,
			DateCreated: course.DateCreated,
		}
	}

	return course
}
