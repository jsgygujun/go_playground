CREATE TABLE `users` (
    `id` INT(10) NOT NULL AUTO_INCREMENT COMMENT '主键ID。',
    `sex` TINYINT(1) NOT NULL COMMENT '性别。1:男；2:女；3:其他。',
    `username` VARCHAR(32) NOT NULL COMMENT '用户名。',
    `password` VARCHAR(16) NOT NULL COMMENT '密码。',
    `email` VARCHAR(64) NOT NULL COMMENT '邮箱。',
    `created_at` DATETIME NOT NULL COMMENT '创建时间。',
    `updated_at` DATETIME NOT NULL COMMENT '更新时间。',

    PRIMARY KEY (`id`),
    KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户信息表';