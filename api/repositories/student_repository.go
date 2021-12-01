package repositories

import (
	"fmt"

	"github.com/kier1021/student-course-app/api/models"
	"github.com/kier1021/student-course-app/databases"

	"github.com/Masterminds/squirrel"
)

type StudentRepository struct {
	db *databases.DBSchool
}

func NewStudentRepository(db *databases.DBSchool) *StudentRepository {
	return &StudentRepository{
		db: db,
	}
}

func (repo *StudentRepository) AddStudent(student *models.Student) (int64, error) {
	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Insert("t_student").
		Columns("student_name", "student_email", "student_phone").
		Values(student.Name, student.Email, student.Phone).
		ToSql()

	query += " RETURNING student_id"

	if err != nil {
		return 0, err
	}

	lastInsertID := 0

	err = repo.db.Conn.QueryRow(query, args...).Scan(&lastInsertID)
	if err != nil {
		return 0, err
	}

	return int64(lastInsertID), nil
}

func (repo *StudentRepository) AddStudentCourses(studentID int, courseIDs []string) error {
	for _, courseID := range courseIDs {
		query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
			Insert("t_student_course_binding").
			Columns("student_id", "course_id").
			Values(studentID, courseID).
			ToSql()

		if err != nil {
			return err
		}

		_, err = repo.db.Conn.Exec(query, args...)
		if err != nil {
			return err
		}
	}

	return nil
}

func (repo *StudentRepository) GetStudents() (students []*models.StudentWithCourse, err error) {
	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("t_student.*", "t_course.course_name").
		From("t_student").
		LeftJoin("t_student_course_binding ON t_student_course_binding.student_id = t_student.student_id").
		LeftJoin("t_course ON t_course.course_id = t_student_course_binding.course_id").
		OrderBy("-t_student.student_id").
		ToSql()

	if err != nil {
		return nil, err
	}

	var (
		studentID    int
		studentName  string
		studentEmail string
		studentPhone string
		courseName   interface{}
	)

	rows, err := repo.db.Conn.Query(query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		fmt.Println(rows)

		err := rows.Scan(&studentID, &studentName, &studentEmail, &studentPhone, &courseName)
		if err != nil {
			return nil, err
		}

		course, _ := courseName.(string)

		students = append(students, &models.StudentWithCourse{
			Student: models.Student{
				ID:    studentID,
				Name:  studentName,
				Email: studentEmail,
				Phone: studentPhone,
			},
			CourseName: course,
		})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return students, nil
}

func (repo *StudentRepository) GetStudentsByCourseID(courseID string) (students []*models.Student, err error) {
	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("t_student.*").
		From("t_student").
		LeftJoin("t_student_course_binding ON t_student_course_binding.student_id = t_student.student_id").
		LeftJoin("t_course ON t_course.course_id = t_student_course_binding.course_id").
		Where(squirrel.Eq{"t_course.course_id": courseID}).
		OrderBy("-t_student.student_id").
		ToSql()

	if err != nil {
		return nil, err
	}

	var (
		studentID    int
		studentName  string
		studentEmail string
		studentPhone string
	)

	rows, err := repo.db.Conn.Query(query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&studentID, &studentName, &studentEmail, &studentPhone)
		if err != nil {
			return nil, err
		}
		students = append(students, &models.Student{
			ID:    studentID,
			Name:  studentName,
			Email: studentEmail,
			Phone: studentPhone,
		})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return students, nil
}

func (repo *StudentRepository) GetStudentByID(id int) (student *models.Student, err error) {
	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("t_student.*").
		From("t_student").
		Where(squirrel.Eq{"t_student.student_id": id}).
		Limit(1).
		ToSql()

	if err != nil {
		return nil, err
	}

	var (
		studentID    int
		studentName  string
		studentEmail string
		studentPhone string
	)

	rows, err := repo.db.Conn.Query(query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var students []*models.Student
	for rows.Next() {
		err := rows.Scan(&studentID, &studentName, &studentEmail, &studentPhone)
		if err != nil {
			return nil, err
		}
		students = append(students, &models.Student{
			ID:    studentID,
			Name:  studentName,
			Email: studentEmail,
			Phone: studentPhone,
		})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	if len(students) != 0 {
		return students[0], nil
	}

	return nil, nil
}

func (repo *StudentRepository) GetStudentsByStudentAndCourseID(studID int, courseID string) (students []*models.Student, err error) {
	query, args, err := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar).
		Select("t_student.*").
		From("t_student").
		LeftJoin("t_student_course_binding ON t_student_course_binding.student_id = t_student.student_id").
		LeftJoin("t_course ON t_course.course_id = t_student_course_binding.course_id").
		Where(squirrel.Eq{"t_course.course_id": courseID, "t_student.student_id": studID}).
		OrderBy("-t_student.student_id").
		ToSql()

	if err != nil {
		return nil, err
	}

	var (
		studentID    int
		studentName  string
		studentEmail string
		studentPhone string
	)

	rows, err := repo.db.Conn.Query(query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&studentID, &studentName, &studentEmail, &studentPhone)
		if err != nil {
			return nil, err
		}
		students = append(students, &models.Student{
			ID:    studentID,
			Name:  studentName,
			Email: studentEmail,
			Phone: studentPhone,
		})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return students, nil
}
