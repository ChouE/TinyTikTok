CREATE DATABASE
IF
    NOT EXISTS tiktok DEFAULT CHARACTER 
	SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;
USE tiktok;
CREATE TABLE `sys_user` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(32) DEFAULT '' COMMENT '用户名',
    `password` varchar(255) COMMENT '密码',
    `created_at` datetime COMMENT '创建时间',
    `updated_at` datetime COMMENT '更新时间',
    `deleted_at` datetime COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户管理';

CREATE TABLE `sys_video` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `uid` int(10) COMMENT '用户id',
    `play_url` varchar(255) COMMENT '视频播放地址',
    `cover_url` varchar(255) COMMENT '视频封面地址',
    `title` varchar(255) COMMENT '视频标题',
    `created_at` datetime COMMENT '创建时间',
    `updated_at` datetime COMMENT '更新时间',
    `deleted_at` datetime COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='视频管理';

CREATE TABLE `sys_comment` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `uid` int(10) COMMENT '用户id',
    `vid` int(10) COMMENT '视频id',
    `content` varchar(255) COMMENT '评论内容',
    `created_at` datetime COMMENT '创建时间',
    `updated_at` datetime COMMENT '更新时间',
    `deleted_at` datetime COMMENT '删除时间',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户管理';