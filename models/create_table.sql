CREATE TABLE users (
                       `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
                       `user_id` BIGINT(20) NOT NULL COMMENT '用户唯一标识符，使用UUID格式',
                       `username` VARCHAR(255) NOT NULL COMMENT '用户的登录用户名，必须唯一',
                       `password` VARCHAR(255) NOT NULL COMMENT '用户的密码，存储加密后的字符串',
                       `email` VARCHAR(255) UNIQUE COMMENT '用户的电子邮件地址，唯一',
                       `first_name` VARCHAR(255) COMMENT '用户的名字',
                       `last_name` VARCHAR(255) COMMENT '用户的姓氏',
                       `phone` VARCHAR(20) COMMENT '用户的电话号码',
                       `is_active` BOOLEAN DEFAULT TRUE COMMENT '标识用户账户是否处于激活状态',
                       `last_login` DATETIME COMMENT '记录用户最后一次登录的时间',
                       `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '记录用户账户的创建时间',
                       `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '记录用户信息最后更新的时间',
                       `is_deleted` BOOLEAN DEFAULT FALSE COMMENT '标识用户是否已被软删除',
                       `deleted_at` DATETIME COMMENT '记录用户被软删除的时间',
                       PRIMARY KEY (`id`),  -- 将主键设置为id
                       UNIQUE KEY `idx_user_id` (`user_id`),  -- user_id作为唯一键
                       UNIQUE KEY `idx_username` (`username`),
                       UNIQUE KEY `idx_email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;