# The School App
## A REST API that demonstrate the relationship between the student and its courses using Golang

### How to run
```
go run main.go
```

### Table Structure

```SQL
CREATE DATABASE db_school;

CREATE TABLE t_student(
   student_id       SERIAL PRIMARY KEY,
   student_name     VARCHAR(100) NOT NULL,
   student_email    VARCHAR(100) NOT NULL,
   student_phone    VARCHAR(100) NOT NULL
)

CREATE TABLE t_course(
   course_id        VARCHAR(100) PRIMARY KEY,
   course_name      VARCHAR(100) NOT NULL,
   professor_name   VARCHAR(100) NOT NULL,
   description      VARCHAR(100) NOT NULL
)

CREATE TABLE t_student_course_binding(
    id              SERIAL PRIMARY KEY,
    student_id      int,
    course_id       VARCHAR(100)
)
```

### Configs
You can edit the database configuration inside the databases/db_school.go file.
The table structure is located in school.sql
```
Host = localhost
Port = 8080
Database = PostgreSQL
```

### Endpoints
#### Get students with its courses
```
GET localhost:8080/students
```

#### Get Students by course ID
```
GET localhost:8080/course/student?course_id=COURSE-100
```

#### Add Student
```
POST localhost:8080/student
```
```
application/json:
{
    "student_name": "John Doe",
    "student_email": "john@email.com",
    "student_phone": "1234-567-890"
}
```

#### Add course
```
POST localhost:8080/course
```
```
application/json:
{
    "course_id": "SCIENCE-100",
    "course_name": "Science 100",
    "professor_name": "Thomas Blake",
    "description": "All about science"
}
```

#### Delete course
````
DELETE localhost:8080/course
````
```
application/json:
{
    "course_id": "SCIENCE-100"
}
```

#### Add courses to student
```
localhost:8080/courses/student
```
````
{
    "student_id": 5,
    "course_ids": ["MATH-100"]
}
```


