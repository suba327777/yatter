CREATE TABLE `account` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NOT NULL UNIQUE,
  `password_hash` varchar(255) NOT NULL,
  `display_name` varchar(255),
  `avatar` text,
  `header` text,
  `note` text,
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);

CREATE TABLE `status`(
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `account_id` bigint(20) NOT NULL,
  `content` text NOT NULL,
  `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  CONSTRAINT `fk_statuses_account_id` FOREIGN KEY (`account_id`) REFERENCES `account`(`id`)
);

CREATE TABLE `relationship`(
  follower_id bigint(20) NOT NULL,
  followed_id bigint(20) NOT NULL
  PRIMARY KEY (`follower_id`, `followed_id`),
  CONSTRAINT `fk_relationship_follower_id` FOREIGN KEY (`follower_id`) REFERENCES `account`(`id`),
  CONSTRAINT `fk_relationship_followed_id` FOREIGN KEY (`followed_id`) REFERENCES `account`(`id`),
);

CREATE TABLE `attachment`(
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `type` varchar(255),
  `url` text ,
  `description` text,
  PRIMARY KEY (`id`)
);

CREATE TABLE `attachment_binding`(
  attachment_id bigint(20) NOT NULL,
  status_id bigint(20) NOT NULL,
  PRIMARY KEY (`attachment_id`, `status_id`),
  CONSTRAINT `fk_attachment_binding_attachment_id` FOREIGN KEY (`attachment_id`) REFERENCES `attachment`(`id`),
  CONSTRAINT `fk_attachment_binding_status_id` FOREIGN KEY (`status_id`) REFERENCES `status`(`id`)
);
