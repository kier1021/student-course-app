package repositories

import (
	"github.com/Masterminds/squirrel"
	"github.com/kier1021/student-course-app/api/models"
	"github.com/kier1021/student-course-app/databases"
)

type CourseRepository struct {
	db *databases.DBSchool
}

func NewCourseRepository(db *databases.DBSchool) *CourseRepository {
	return &CourseRepository{
		db: db,
	}
}

func (repo *CourseRepository) AddCourse(course *models.Course) error {
	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Insert("t_course").
		Columns("course_id", "course_name", "professor_name", "description").
		Values(course.ID, course.Name, course.ProfessorName, course.Description).
		ToSql()

	if err != nil {
		return err
	}

	_, err = repo.db.Conn.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (repo *CourseRepository) DeleteCourse(courseID string) error {
	// Delete in t_course
	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Delete("t_course").
		Where(squirrel.Eq{"course_id": courseID}).
		ToSql()

	if err != nil {
		return err
	}

	_, err = repo.db.Conn.Exec(query, args...)
	if err != nil {
		return err
	}

	// Delete in t_student_course_binding
	query, args, err = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Delete("t_student_course_binding").
		Where(squirrel.Eq{"course_id": courseID}).
		ToSql()

	if err != nil {
		return err
	}

	_, err = repo.db.Conn.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (repo *CourseRepository) GetCourseByID(id string) (course *models.Course, err error) {
	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("t_course.*").
		From("t_course").
		Where(squirrel.Eq{"t_course.course_id": id}).
		Limit(1).
		ToSql()

	if err != nil {
		return nil, err
	}

	var (
		courseID      string
		courseName    string
		professorName string
		description   string
	)

	rows, err := repo.db.Conn.Query(query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var courses []*models.Course
	for rows.Next() {
		err := rows.Scan(&courseID, &courseName, &professorName, &description)
		if err != nil {
			return nil, err
		}
		courses = append(courses, &models.Course{
			ID:            courseID,
			Name:          courseName,
			ProfessorName: professorName,
			Description:   description,
		})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	if len(courses) != 0 {
		return courses[0], nil
	}

	return nil, nil
}
