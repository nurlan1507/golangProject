# golangProject
This is the backend part of my website testApp, Frontend : https://github.com/nurlan1507/testApp-frontend



```
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

alter table answer add constraint answer_question_fk foreign key (question_id) references question(question_id);

alter table question add constraint  question_test_id_fk foreign key(test_id) references test(id);

```

alter table question add column correctAnswer varchar(40);
