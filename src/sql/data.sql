insert into users (name, nick, email, password)
VALUES
("Michael Jordan", "MJ", "michaeljordan@email.com", "$2a$10$q8G9BFMue67BziowL5.yqOjltcqhiA/Dasn7VR5gIb1hUL21N1Y2i"),
("Oliver Sykes", "BMTH", "oliversykes@email.com", "$2a$10$q8G9BFMue67BziowL5.yqOjltcqhiA/Dasn7VR5gIb1hUL21N1Y2i"),
("Kendrick Lamar", "be humble", "kendriclamar@email.com", "$2a$10$q8G9BFMue67BziowL5.yqOjltcqhiA/Dasn7VR5gIb1hUL21N1Y2i");

insert into followers (user_id, follower_id)
VALUES
(1, 2),
(3, 1),
(1, 3);

-- password = 123456