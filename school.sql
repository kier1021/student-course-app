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

INSERT INTO t_student(student_name, student_email, student_phone) VALUES ('test', 'test@email.com', '1234-567-890');
INSERT INTO t_course(course_id, course_name, professor_name, description) VALUES ('ENG-100', 'ENGLISH 1', 'James Einstein', 'All about english');
INSERT INTO t_student_course_binding(student_id, course_id) VALUES (1, 'ENG-100');

INSERT INTO t_student(student_name, student_email, student_phone) VALUES ('test2', 'test2@email.com', '1234-567-890');
INSERT INTO t_course(course_id, course_name, professor_name, description) VALUES ('MATH-100', 'ENGLISH 1', 'Henry Newton', 'All about math');
INSERT INTO t_student_course_binding(student_id, course_id) VALUES (1, 'MATH-100');