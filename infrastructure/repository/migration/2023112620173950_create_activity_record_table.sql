-- migrate:up
CREATE TABLE `activity_record`  (
                                `id` bigint NOT NULL,
                                `activity_id` varchar(128) NOT NULL,
                                `participants` varchar(128) NOT NULL,
                                `type` varchar(128) NOT NULL,
                                `submit` varchar(128) NULL DEFAULT NULL,
                                `status` tinyint(1) NOT NULL,
                                `created_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
                                `updated_at` datetime NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
                                PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB;

-- migrate:down
DROP TABLE IF EXISTS `activity_record`;