CREATE TABLE `users` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `email` varchar(255) NOT NULL,
                         `firstname` varchar(255),
                         `lastname` varchar(255),
                         `role` int,
                         `height` double,
                         `weight` double,
                         `bdf` double,
                         `tdee` double,
                         `calorie` double,
                         `status` int,
                         `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         PRIMARY KEY (`id`)
);

CREATE TABLE `roles` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `name` varchar(255),
                         PRIMARY KEY (`id`)
);
ALTER TABLE `users` ADD FOREIGN KEY (`role`) REFERENCES `roles` (`id`);

CREATE TABLE `accounts` (
                            `id` int NOT NULL AUTO_INCREMENT,
                            `user_id` int NOT NULL,
                            `password` varchar(255),
                            `status` int,
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `unique_user_id` (`user_id`),
                            FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
);

CREATE TABLE `exersice_type` (
                                 `id` int NOT NULL AUTO_INCREMENT,
                                 `name` varchar(255),
                                 `status` int,
                                 PRIMARY KEY (`id`)
);

CREATE TABLE `exersices` (
                             `id` int NOT NULL AUTO_INCREMENT,
                             `exersice_type` int NOT NULL,
                             `name` varchar(255),
                             `description` text,
                             `calorie` double NOT NULL,
                             `rep_serving` int,
                             `time_serving` int,
                             `status` int,
                             PRIMARY KEY (`id`),
                             FOREIGN KEY (`exersice_type`) REFERENCES `exersice_type`(`id`)
);

CREATE TABLE `schedules` (
                             `id` int NOT NULL AUTO_INCREMENT,
                             `user_id` int,
                             `time` timestamp,
                             `calories_burn` double,
                             `status` int,
                             `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                             `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                             PRIMARY KEY (`id`),
                             FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
);

CREATE TABLE `schedule_detail` (
                                   `id` int NOT NULL AUTO_INCREMENT,
                                   `schedule_id` int NOT NULL,
                                   `exersice_id` int NOT NULL,
                                   `rep` int,
                                   `time` int,

                                   PRIMARY KEY (`id`),
                                   UNIQUE KEY `unique_schedule` (`schedule_id`, `exersice_id`),
                                   FOREIGN KEY (`schedule_id`) REFERENCES `schedules`(`id`),
                                   FOREIGN KEY (`exersice_id`) REFERENCES `exersices`(`id`)
);


CREATE TABLE `dishes` (
                          `id` int NOT NULL AUTO_INCREMENT,
                          `name` varchar(255) NOT NULL,
                          `description` text,
                          `calorie` double,
                          `protein` double,
                          `fat` double,
                          `carb` double,
                          `sugar` double,
                          `serving` double,

                          PRIMARY KEY (`id`)
);
CREATE TABLE `meals` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `user_id` int NOT NULL,
                         `description` text,
                         `date` timestamp,
                         `total_calorie` double,
                         PRIMARY KEY (`id`),
                         FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
);


CREATE TABLE `meal_detail` (
                               `id` int NOT NULL AUTO_INCREMENT,
                               `dish_id` int,
                               `meal_id` int,
                               `serving` double,
                               `calorie` double,
                               PRIMARY KEY (`id`),
                               UNIQUE KEY `unique_meal` (`dish_id`, `meal_id`),
                               FOREIGN KEY (`dish_id`) REFERENCES `dishes`(`id`),
                               FOREIGN KEY (`meal_id`) REFERENCES `meals`(`id`)
);

CREATE TABLE `reminders` (
                             `id` int NOT NULL AUTO_INCREMENT,
                             `user_id` int NOT NULL,
                             `description` text,
                             `schedule_id` int,
                             `meal_id` int,
                             `reminder_type` int,
                             `date` timestamp,
                             `status` int,

                             PRIMARY KEY (`id`),
                             FOREIGN KEY (`user_id`) REFERENCES `users`(`id`),
                             FOREIGN KEY (`schedule_id`) REFERENCES `schedules`(`id`),
                             FOREIGN KEY (`meal_id`) REFERENCES `meals`(`id`),
                             FOREIGN KEY (`reminder_type`) REFERENCES `reminder_type`(`id`)
);

CREATE TABLE `reminder_type` (
                                 `id` int NOT NULL AUTO_INCREMENT,
                                 `name` varchar(255),
                                 PRIMARY KEY (`id`)
);

CREATE TABLE `posts` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `title` varchar(255),
                         `body` text,
                         `user_id` int,
                         `status` int,
                         `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         PRIMARY KEY (`id`),
                         FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
);

CREATE TABLE `likes` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `user_id` int,
                         `post_id` int,
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `unique_like` (`user_id`, `post_id`),
                         FOREIGN KEY (`user_id`) REFERENCES `users`(`id`),
                         FOREIGN KEY (`post_id`) REFERENCES `posts`(`id`)
);

CREATE TABLE `comments` (
                            `id` int NOT NULL AUTO_INCREMENT,
                            `body` text,
                            `user_id` int,
                            `post_id` int,
                            `status` int,
                            `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                            `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

                            PRIMARY KEY (`id`),
                            FOREIGN KEY (`user_id`) REFERENCES `users`(`id`),
                            FOREIGN KEY (`post_id`) REFERENCES `posts`(`id`)
);
CREATE TABLE `photo_type` (
                              `id` int NOT NULL AUTO_INCREMENT,
                              `name` varchar(255),

                              PRIMARY KEY (`id`)
);
CREATE TABLE `photos` (
                          `id` int NOT NULL AUTO_INCREMENT,
                          `photo_type` int,
                          `url` varchar(255),
                          `post_id` int,
                          `comment_id` int,
                          `exersice_id` int,
                          `dish_id` int,
                          `user_id` int,

                          PRIMARY KEY (`id`),
                          FOREIGN KEY (`photo_type`) REFERENCES `photo_type`(`id`),
                          FOREIGN KEY (`user_id`) REFERENCES `users`(`id`),
                          FOREIGN KEY (`post_id`) REFERENCES `posts`(`id`),
                          FOREIGN KEY (`comment_id`) REFERENCES `comments`(`id`),
                          FOREIGN KEY (`dish_id`) REFERENCES `dishes`(`id`),
                          FOREIGN KEY (`exersice_id`) REFERENCES `exersices`(`id`)
);

CREATE TABLE conversations (
                               conversation_id INT AUTO_INCREMENT PRIMARY KEY,
                               created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE conversationparticipants (
                                          conversation_id INT,
                                          user_id INT,
                                          UNIQUE (conversation_id, user_id),
                                          FOREIGN KEY (conversation_id) REFERENCES conversations(conversation_id),
                                          FOREIGN KEY (user_id) REFERENCES users(id)
);
CREATE TABLE messages (
                          message_id INT AUTO_INCREMENT PRIMARY KEY,
                          conversation_id INT,
                          sender_id INT,
                          content TEXT,
                          timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          deleted_by JSON,
                          FOREIGN KEY (conversation_id) REFERENCES conversations(conversation_id),
                          FOREIGN KEY (sender_id) REFERENCES users(id)
);
ALTER TABLE users ADD COLUMN jwt_secret VARCHAR(255);

CREATE TABLE password_reset_tokens (
                                       id SERIAL PRIMARY KEY,
                                       user_id INT NOT NULL,
                                       token VARCHAR(255) NOT NULL,
                                       expires_at TIMESTAMP NOT NULL,
                                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                       updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                       FOREIGN KEY (user_id) REFERENCES users(id)
);


DELIMITER $$

CREATE PROCEDURE sp_total_calorie_burn(schedule_id INT)
BEGIN
    DECLARE total_calories DOUBLE DEFAULT 0;

    -- Tổng calo đốt cháy cho mỗi chi tiết lịch trình
SELECT SUM(
               CASE
                   WHEN sd.rep IS NOT NULL AND e.rep_serving IS NOT NULL THEN e.calorie * sd.rep / e.rep_serving
                   WHEN sd.time IS NOT NULL AND e.time_serving IS NOT NULL THEN e.calorie * sd.time / e.time_serving
                   ELSE 0
                   END
       ) INTO total_calories
FROM schedule_detail sd
         JOIN exersices e ON sd.exersice_id = e.id
WHERE sd.schedule_id = schedule_id;

-- Cập nhật tổng calo trong bảng schedules
UPDATE schedules
SET calories_burn = total_calories
WHERE id = schedule_id;
END$$

DELIMITER ;



DELIMITER $$

CREATE TRIGGER before_exersice_update
    AFTER UPDATE ON exersices
    FOR EACH ROW
BEGIN
    DECLARE done INT DEFAULT 0;
    DECLARE sched_id INT;
    DECLARE cur CURSOR FOR
    SELECT s.id
    FROM schedules s
             JOIN schedule_detail sd ON s.id = sd.schedule_id
    WHERE sd.exersice_id = NEW.id;

    DECLARE CONTINUE HANDLER FOR NOT FOUND SET done = 1;

    OPEN cur;

    read_loop: LOOP
        FETCH cur INTO sched_id;
        IF done THEN
            LEAVE read_loop;
END IF;
CALL sp_total_calorie_burn(sched_id);
END LOOP;

CLOSE cur;
END$$

DELIMITER ;
DELIMITER $$

CREATE TRIGGER after_insert_schedule_detail
    AFTER INSERT ON schedule_detail
    FOR EACH ROW
BEGIN
    CALL sp_total_calorie_burn(NEW.schedule_id);
    END$$
    DELIMITER ;
    DELIMITER $$

    CREATE TRIGGER after_delete_schedule_detail
        AFTER DELETE ON schedule_detail
        FOR EACH ROW
    BEGIN
        -- Gọi stored procedure để cập nhật tổng calo cho lịch trình sau khi xóa chi tiết lịch trình
        CALL sp_total_calorie_burn(OLD.schedule_id);
        END$$

        DELIMITER ;
    DELIMITER $$

        CREATE TRIGGER after_update_schedule_detail
            AFTER UPDATE ON schedule_detail
            FOR EACH ROW
        BEGIN
            -- Gọi stored procedure để cập nhật tổng calo cho lịch trình sau khi thêm chi tiết lịch trình
            CALL sp_total_calorie_burn(NEW.schedule_id);
            END$$
            DELIMITER ;


DELIMITER $$

            CREATE PROCEDURE sp_update_total_calorie()
            BEGIN
    -- Cập nhật tổng calo cho từng meal_id
            UPDATE meals m
                JOIN (
                SELECT md.meal_id, SUM(d.calorie * md.serving / d.serving) AS total_calorie
                FROM meal_detail md
                JOIN dishes d ON md.dish_id = d.id
                GROUP BY md.meal_id
                ) AS calorie_summary ON m.id = calorie_summary.meal_id
                SET m.total_calorie = calorie_summary.total_calorie;
            END$$

            DELIMITER ;



DELIMITER $$

            -- Stored Procedure để cập nhật tổng calo cho các bữa ăn chứa dish_id cụ thể
            CREATE PROCEDURE sp_update_total_calorie_by_dish(dish_id INT)
            BEGIN
    -- Cập nhật tổng calo cho từng meal_id chứa dish_id cụ thể
            UPDATE meals m
                JOIN (
                SELECT md.meal_id, SUM(d.calorie * md.serving / d.serving) AS total_calorie
                FROM meal_detail md
                JOIN dishes d ON md.dish_id = d.id
                WHERE md.dish_id = dish_id
                GROUP BY md.meal_id
                ) AS calorie_summary ON m.id = calorie_summary.meal_id
                SET m.total_calorie = calorie_summary.total_calorie;
            END$$

            -- Trigger cho sự kiện UPDATE trên bảng dishes
            CREATE TRIGGER after_update_dish
                AFTER UPDATE ON dishes
                FOR EACH ROW
            BEGIN
                -- Gọi stored procedure để cập nhật tổng calo cho các bữa ăn chứa dish_id cụ thể
                CALL sp_update_total_calorie_by_dish(NEW.id);
                END$$

                DELIMITER ;
DELIMITER $$

                CREATE PROCEDURE sp_update_total_calorie_by_meal(meal_id INT)
                BEGIN
    -- Cập nhật tổng calo cho meal_id cụ thể
                UPDATE meals m
                    JOIN (
                    SELECT md.meal_id, SUM(d.calorie * md.serving / d.serving) AS total_calorie
                    FROM meal_detail md
                    JOIN dishes d ON md.dish_id = d.id
                    WHERE md.meal_id = meal_id
                    GROUP BY md.meal_id
                    ) AS calorie_summary ON m.id = calorie_summary.meal_id
                    SET m.total_calorie = calorie_summary.total_calorie;
                END$$

                DELIMITER ;

DELIMITER $$

                -- Trigger cho sự kiện INSERT trên bảng meal_detail
                CREATE TRIGGER after_insert_meal_detail
                    AFTER INSERT ON meal_detail
                    FOR EACH ROW
                BEGIN
                    -- Gọi stored procedure để cập nhật tổng calo cho bữa ăn chứa chi tiết vừa thêm
                    CALL sp_update_total_calorie_by_meal(NEW.meal_id);
                    END$$

                    -- Trigger cho sự kiện UPDATE trên bảng meal_detail
                    CREATE TRIGGER after_update_meal_detail
                        AFTER UPDATE ON meal_detail
                        FOR EACH ROW
                    BEGIN
                        -- Gọi stored procedure để cập nhật tổng calo cho bữa ăn chứa chi tiết vừa cập nhật
                        CALL sp_update_total_calorie_by_meal(NEW.meal_id);
                        END$$

                        -- Trigger cho sự kiện DELETE trên bảng meal_detail
                        CREATE TRIGGER after_delete_meal_detail
                            AFTER DELETE ON meal_detail
                            FOR EACH ROW
                        BEGIN
                            -- Gọi stored procedure để cập nhật tổng calo cho bữa ăn chứa chi tiết vừa xóa
                            CALL sp_update_total_calorie_by_meal(OLD.meal_id);
                            END$$

                            DELIMITER ;
