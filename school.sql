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