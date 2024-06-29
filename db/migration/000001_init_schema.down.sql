-- Drop Foreign Keys
ALTER TABLE `photos` DROP FOREIGN KEY `photos_ibfk_1`;
ALTER TABLE `photos` DROP FOREIGN KEY `photos_ibfk_2`;
ALTER TABLE `photos` DROP FOREIGN KEY `photos_ibfk_3`;
ALTER TABLE `photos` DROP FOREIGN KEY `photos_ibfk_4`;
ALTER TABLE `photos` DROP FOREIGN KEY `photos_ibfk_5`;
ALTER TABLE `photos` DROP FOREIGN KEY `photos_ibfk_6`;

ALTER TABLE `comments` DROP FOREIGN KEY `comments_ibfk_1`;
ALTER TABLE `comments` DROP FOREIGN KEY `comments_ibfk_2`;

ALTER TABLE `likes` DROP FOREIGN KEY `likes_ibfk_1`;
ALTER TABLE `likes` DROP FOREIGN KEY `likes_ibfk_2`;

ALTER TABLE `posts` DROP FOREIGN KEY `posts_ibfk_1`;

ALTER TABLE `reminders` DROP FOREIGN KEY `reminders_ibfk_1`;
ALTER TABLE `reminders` DROP FOREIGN KEY `reminders_ibfk_2`;
ALTER TABLE `reminders` DROP FOREIGN KEY `reminders_ibfk_3`;
ALTER TABLE `reminders` DROP FOREIGN KEY `reminders_ibfk_4`;
ALTER TABLE `reminders` DROP FOREIGN KEY `reminders_ibfk_5`;

ALTER TABLE `schedule_detail` DROP FOREIGN KEY `schedule_detail_ibfk_1`;
ALTER TABLE `schedule_detail` DROP FOREIGN KEY `schedule_detail_ibfk_2`;

ALTER TABLE `meal_detail` DROP FOREIGN KEY `meal_detail_ibfk_1`;
ALTER TABLE `meal_detail` DROP FOREIGN KEY `meal_detail_ibfk_2`;

ALTER TABLE `meals` DROP FOREIGN KEY `meals_ibfk_1`;

ALTER TABLE `schedules` DROP FOREIGN KEY `schedules_ibfk_1`;

ALTER TABLE `exersices` DROP FOREIGN KEY `exersices_ibfk_1`;

ALTER TABLE `accounts` DROP FOREIGN KEY `accounts_ibfk_1`;

ALTER TABLE `users` DROP FOREIGN KEY `users_ibfk_1`;
ALTER TABLE `users` DROP FOREIGN KEY `users_ibfk_2`;

-- Drop Tables
DROP TABLE IF EXISTS `photos`;
DROP TABLE IF EXISTS `photo_type`;
DROP TABLE IF EXISTS `comments`;
DROP TABLE IF EXISTS `likes`;
DROP TABLE IF EXISTS `posts`;
DROP TABLE IF EXISTS `reminders`;
DROP TABLE IF EXISTS `reminder_type`;
DROP TABLE IF EXISTS `schedule_detail`;
DROP TABLE IF EXISTS `schedules`;
DROP TABLE IF EXISTS `meal_detail`;
DROP TABLE IF EXISTS `meals`;
DROP TABLE IF EXISTS `dishes`;
DROP TABLE IF EXISTS `exersices`;
DROP TABLE IF EXISTS `exersice_type`;
DROP TABLE IF EXISTS `accounts`;
DROP TABLE IF EXISTS `users`;
DROP TABLE IF EXISTS `roles`;
