create table day_1
(
    id       int primary key auto_increment,
    calories int
);

create table answer_day_1
(
    name  varchar(12),
    value int
);

DROP PROCEDURE IF EXISTS Day1;
CREATE PROCEDURE Day1()
BEGIN
    DECLARE length, counter, answer1,answer2, currentSum, current INT DEFAULT 0;
    SELECT COUNT(*) FROM day_1 INTO length;

    drop table if exists temp;
    create temporary table temp
    (
        totalCalories int
    );

    WHILE counter < length
        DO
            select calories from day_1 limit counter, 1 INTO current;

            if current is null THEN
                insert into temp (totalCalories) values (currentSum);

                SET currentSum = 0;
            end if;

            if current is not null THEN
                SET currentSum = currentSum + current;
            end if;

            SET counter = counter + 1;
        END WHILE;

    select totalCalories
    from temp
    order by totalCalories desc
    limit 1
    INTO answer1;

    select sum(totalCalories)
    from (select *
          from temp
          order by totalCalories desc
          limit 3)
             as a
    INTO answer2;

    insert into answers_day_1 (name, value) values ('First', answer1);
    insert into answers_day_1 (name, value) values ('Second', answer2);
End;
CALL Day1();
