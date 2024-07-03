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
                         `deleted_at` TIMESTAMP NULL DEFAULT NULL,
                         PRIMARY KEY (`id`)
);

CREATE TABLE `roles` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `name` varchar(255),
                         `status` int,
                         `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         `deleted_at` TIMESTAMP NULL DEFAULT NULL,
                         PRIMARY KEY (`id`)
);
ALTER TABLE `users` ADD FOREIGN KEY (`role`) REFERENCES `roles` (`id`);

CREATE TABLE `accounts` (
                            `id` int NOT NULL AUTO_INCREMENT,
                            `user_id` int NOT NULL,
                            `password` varchar(255),
                            `status` int,
                            `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                            `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            `deleted_at` TIMESTAMP NULL DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `unique_user_id` (`user_id`),
                            FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
);

CREATE TABLE `exersice_type` (
                                 `id` int NOT NULL AUTO_INCREMENT,
                                 `name` varchar(255),
                                 `status` int,
                                 `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                 `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                 `deleted_at` TIMESTAMP NULL DEFAULT NULL,
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
                             `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                             `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                             `deleted_at` TIMESTAMP NULL DEFAULT NULL,
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
                             `deleted_at` TIMESTAMP NULL DEFAULT NULL,
                             PRIMARY KEY (`id`),
                             FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
);

CREATE TABLE `schedule_detail` (
                                   `id` int NOT NULL AUTO_INCREMENT,
                                   `schedule_id` int NOT NULL,
                                   `exersice_id` int NOT NULL,
                                   `rep` int,
                                   `time` int,
                                   `status` int,
                                   `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                   `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                   `deleted_at` TIMESTAMP NULL DEFAULT NULL,
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
                          `status` int,
                          `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                          `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          `deleted_at` TIMESTAMP NULL DEFAULT NULL,
                          PRIMARY KEY (`id`)
);
CREATE TABLE `meals` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `user_id` int NOT NULL,
                         `description` text,
                         `date` timestamp,
                         `total_calorie` double,
                         `status` int,
                         `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         `deleted_at` TIMESTAMP NULL DEFAULT NULL,
                         PRIMARY KEY (`id`),
                         FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
);


CREATE TABLE `meal_detail` (
                               `id` int NOT NULL AUTO_INCREMENT,
                               `dish_id` int,
                               `meal_id` int,
                               `serving` double,
                               `calorie` double,
                               `status` int,
                               `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                               `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                               `deleted_at` TIMESTAMP NULL DEFAULT NULL,
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
                             `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                             `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                             `deleted_at` TIMESTAMP NULL DEFAULT NULL,
                             PRIMARY KEY (`id`),
                             FOREIGN KEY (`user_id`) REFERENCES `users`(`id`),
                             FOREIGN KEY (`schedule_id`) REFERENCES `schedules`(`id`),
                             FOREIGN KEY (`meal_id`) REFERENCES `meals`(`id`),
                             FOREIGN KEY (`reminder_type`) REFERENCES `reminder_type`(`id`)
);

CREATE TABLE `reminder_type` (
                                 `id` int NOT NULL AUTO_INCREMENT,
                                 `name` varchar(255),
                                 `status` int,
                                 `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                                 `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                                 `deleted_at` TIMESTAMP NULL DEFAULT NULL,
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
                         `deleted_at` TIMESTAMP NULL DEFAULT NULL,
                         PRIMARY KEY (`id`),
                         FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
);

CREATE TABLE `likes` (
                         `id` int NOT NULL AUTO_INCREMENT,
                         `user_id` int,
                         `post_id` int,
                         `status` int,
                         `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         `deleted_at` TIMESTAMP NULL DEFAULT NULL,
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
                            `deleted_at` TIMESTAMP NULL DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            FOREIGN KEY (`user_id`) REFERENCES `users`(`id`),
                            FOREIGN KEY (`post_id`) REFERENCES `posts`(`id`)
);
CREATE TABLE `photo_type` (
                              `id` int NOT NULL AUTO_INCREMENT,
                              `name` varchar(255),
                              `status` int,
                              `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                              `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                              `deleted_at` TIMESTAMP NULL DEFAULT NULL,
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
                          `status` int,
                          `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                          `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                          `deleted_at` TIMESTAMP NULL DEFAULT NULL,
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


--   ALTER TABLE `photos` ADD FOREIGN KEY (`photo_type`) REFERENCES `photo_type` (`id`);

--   ALTER TABLE `photos` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--   ALTER TABLE `photos` ADD FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`);

--   ALTER TABLE `photos` ADD FOREIGN KEY (`comment_id`) REFERENCES `comments` (`id`);

--   ALTER TABLE `photos` ADD FOREIGN KEY (`dish_id`) REFERENCES `dishes` (`id`);

--   ALTER TABLE `photos` ADD FOREIGN KEY (`exersice_id`) REFERENCES `exersices` (`id`);

--   ALTER TABLE `comments` ADD FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`);

--   ALTER TABLE `likes` ADD FOREIGN KEY (`post_id`) REFERENCES `posts` (`id`);

--   ALTER TABLE `posts` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--   ALTER TABLE `comments` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--   ALTER TABLE `likes` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--   ALTER TABLE `users` ADD FOREIGN KEY (`id`) REFERENCES `reminders` (`user_id`);

--   ALTER TABLE `reminders` ADD FOREIGN KEY (`reminder_type`) REFERENCES `reminder_type` (`id`);

--   ALTER TABLE `schedules` ADD FOREIGN KEY (`id`) REFERENCES `reminders` (`schedule_id`);

--   ALTER TABLE `meals` ADD FOREIGN KEY (`id`) REFERENCES `reminders` (`meal_id`);

--   ALTER TABLE `schedule_detail` ADD FOREIGN KEY (`exersice_id`) REFERENCES `exersices` (`id`);

--   ALTER TABLE `schedule_detail` ADD FOREIGN KEY (`schedule_id`) REFERENCES `schedules` (`id`);

--   ALTER TABLE `schedules` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--   ALTER TABLE `exersices` ADD FOREIGN KEY (`exersice_type`) REFERENCES `exersice_type` (`id`);

--   ALTER TABLE `meal_detail` ADD FOREIGN KEY (`dish_id`) REFERENCES `dishes` (`id`);

--   ALTER TABLE `meal_detail` ADD FOREIGN KEY (`meal_id`) REFERENCES `meals` (`id`);

--   ALTER TABLE `meals` ADD FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);


--   ALTER TABLE `users` ADD FOREIGN KEY (`id`) REFERENCES `accounts` (`user_id`);

