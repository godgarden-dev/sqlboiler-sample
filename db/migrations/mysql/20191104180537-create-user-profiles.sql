
-- +migrate Up
CREATE TABLE IF NOT EXISTS `user_profiles` (
  `user_id` int(11)  PRIMARY KEY NOT NULL,
  `prefecture` VARCHAR(255) NOT NULL,
  `gender` VARCHAR(255) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  CONSTRAINT `user_profiles_user` FOREIGN KEY (`user_id`) REFERENCES users (`id`)
)
ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- +migrate Down
DROP TABLE `user_profiles`;