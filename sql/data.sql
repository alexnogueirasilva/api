USE devbook;

INSERT INTO devbook.users (name, nickname, email, password)
VALUES ("Alex Nogueira", "alex", "alex@devaction.com.br",
        "$2a$10$DkAZqNIvDVoz.k9dVnc/fugsfgTsluZJ85fiMc2BxcHn8i42xP.0."),
       ("Bruna Gabriela", "bruna", "bruna@devactioncombr",
        "$2a$10$DkAZqNIvDVoz.k9dVnc/fugsfgTsluZJ85fiMc2BxcHn8i42xP.0."),
       ("Gabriela", "gabriela", "gabriela@devactioncombr",
        "$2a$10$DkAZqNIvDVoz.k9dVnc/fugsfgTsluZJ85fiMc2BxcHn8i42xP.0.");

INSERT INTO followers (user_id, follower_id)
VALUES (1, 2),
       (1, 3),
       (2, 1),
       (3, 1);


truncate table publications;
INSERT INTO publications
    (title, content, author_id)
VALUES ("Primeira publicação",
        "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla euismod, nisl vitae ultricies ultricies, nunc nunc aliquam nunc, vitae a",
        1),
       ("Segunda publicação",
        "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla euismod, nisl vitae ultricies ultricies, nunc nunc aliquam nunc, vitae a",
        1),
       ("Terceira publicação",
        "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla euismod, nisl vitae ultricies ultricies, nunc nunc aliquam nunc, vitae a",
        2),
       ("Quarta publicação",
        "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla euismod, nisl vitae ultricies ultricies, nunc nunc aliquam nunc, vitae a",
        2),
       ("Quinta publicação",
        "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla euismod, nisl vitae ultricies ultricies, nunc nunc aliquam nunc, vitae a",
        3),
       ("Sexta publicação",
        "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nulla euismod, nisl vitae ultricies ultricies, nunc nunc aliquam nunc, vitae a",
        3);