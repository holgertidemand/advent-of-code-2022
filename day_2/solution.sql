create table day_2
(
    opponent varchar(10),
    me       varchar(10)
);

create table answer_day_2
(
    name  varchar(10),
    value int
);


create table rules_day_2
(
    me       varchar(1),
    opponent varchar(1),
    score    int
);

insert into rules_day_2
    (opponent, me, score)
values ('A', 'X', 4), # Draw Rock
       ('B', 'X', 1), # Loose
       ('C', 'X', 7), # Win
       ('A', 'Y', 8), # Win Paper
       ('B', 'Y', 5), # Draw
       ('C', 'Y', 2), # Loose
       ('A', 'Z', 3), # Loose Scissors
       ('B', 'Z', 9), # Draw
       ('C', 'Z', 6); # Win

DROP PROCEDURE IF EXISTS Day2;
CREATE PROCEDURE Day2()
BEGIN
    DECLARE counter,current, answer1 INT DEFAULT 0;

    WHILE counter < (SELECT COUNT(*) FROM day_2)
        DO
        SET current = 0;
        select score
            from day_2
                     JOIN rules_day_2
                          on day_2.me = rules_day_2.me and
                             day_2.opponent = rules_day_2.opponent
            limit counter, 1
            INTO current;

            SET answer1 = answer1 + current;
            SET counter = counter + 1;
        END WHILE;

    insert into answers_day_2 (name, value) values ('First', answer1);
End;
CALL Day2();