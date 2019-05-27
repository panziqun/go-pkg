drop table `permission`;
CREATE TABLE `permission` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(16) NOT NULL COMMENT '权限名',
    `method` varchar(8) NOT NULL DEFAULT '' COMMENT '请求方式 GET POST ... 不填代表所有',
    `route` varchar(64) NOT NULL COMMENT '路由 例如/user/list /user/*',
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT '权限表';

drop table `role`;
CREATE TABLE `role` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(16) NOT NULL COMMENT '角色名',
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT '角色表';

drop table `admin`;
CREATE TABLE `admin` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `name` varchar(16) NOT NULL COMMENT '用户名',
    `pwd` varchar(32) NOT NULL COMMENT '密码',
    `email` varchar(32) NOT NULL DEFAULT '' COMMENT '邮件',
    `tel` varchar(16) NOT NULL DEFAULT '' COMMENT '电话',
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT '管理员表';

drop table `role_permission`;
CREATE TABLE `role_permission` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `role_id` int(11) NOT NULL COMMENT '角色ID',
    `permission_id` int(11) NOT NULL COMMENT '权限ID',
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT '角色-权限关系表';

drop table `user_role`;
CREATE TABLE `user_role` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `user_id` int(11) NOT NULL COMMENT '用户ID',
    `role_id` int(11) NOT NULL COMMENT '角色ID',
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT '用户-角色关系表';