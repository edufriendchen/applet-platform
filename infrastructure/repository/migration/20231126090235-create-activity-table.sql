-- migrate:up
CREATE TABLE `activity`  (
                                   `id` bigint NOT NULL,
                                   `title` varchar(128) NULL,
                                   `poster_url` varchar(128) NULL,
                                   `content` tinytext NULL,
                                   `welfare` varchar(128) NULL,
                                   `visit_num` int NOT NULL,
                                   `type` tinyint(1) NOT NULL COMMENT '1- 官方大行动 （社团，个人）\r\n2- 社团 （社团）\r\n3- 个人  （个人）',
                                   `start_time` datetime NOT NULL,
                                   `end_time` datetime NULL,
                                   `status` tinyint(1) NULL,
                                   `created_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
                                   `updated_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
                                   `created_by` int NULL,
                                   `updated_by` int NULL,
                                   PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB;

-- migrate:down
DROP TABLE IF EXISTS `activity`;