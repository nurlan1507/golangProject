# golangProject
This is the backend part of my website testApp, Frontend : https://github.com/nurlan1507/testApp-frontend



```postgresql
create table pending_users(
    teacher_id serial primary key ,
    email varchar(40),
    username varchar(40)
);


create table invited_tokens(
    invitation_id serial primary key ,
    teacher_id int not null, constraint fk_tokens_teached_id foreign key(teacher_id) references pending_users(teacher_id) ON DELETE CASCADE ,
    token text
);
alter table invited_tokens add constraint invited_teacher_fk foreign key(teacher_id) references pending_users(teacher_id)


create table refreshTokens(
    user_id serial not null,
    refresh_token text,
    expires date not null
);


create table users(
    id serial not null primary key,
    username varchar(21),
    password text,
    refreshToken text default null,
    role varchar(10)
)

create table test (
                      id serial unique not null primary key ,
                      title varchar(24),
                      description text,
                      author_id int not null ,
                      created_at date,
                      start_at date,
                     expires_at date
);
create table question(
    question_id serial unique not null primary key ,
    description text,
    question_type question_type,
    question_order int,
    test_id int not null,
    image text
);

create table answer(
    answer_id serial unique not null primary key ,
    value text,
    correct bool,
    question_id int not null
);
alter table question
    alter column point TYPE int
alter table question
    alter column question_type TYPE varchar(20)
    
alter table answer add constraint answer_question_fk foreign key (question_id) references question(question_id);

alter table question add constraint  question_test_id_fk foreign key(test_id) references test(id);

alter table users drop column group_name;
alter table users add column group_id int;
create table groups(
    id serial primary key ,
    name varchar(20)
);
insert into groups(id,name)values(1,'SE2101');
insert into groups(id,name)values(2,'SE2103');
insert into groups(id,name)values(3,'SE2103');
create table groups_students(
                                student_id int not null ,
                                group_id int not null ,
                                foreign key (student_id) references users(id),
                                foreign key (group_id) references groups(id)
);
alter table test drop column created_at;

alter table test drop column expires_at;
alter table test drop column start_at;
```

alter table question add column correctAnswer varchar(40);
