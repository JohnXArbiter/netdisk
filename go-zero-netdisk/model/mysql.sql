CREATE TABLE `user`
(
    id         bigint unsigned NOT NULL AUTO_INCREMENT,
    name       varchar(32)     not null default '',
    username   varchar(32)     not null default '',
    password   varchar(32)     not null default '',
    email      varchar(127)    not null default '',
    avatar     varchar(127)    not null default '',
    signature  varchar(100)    not null default '个性签名',
    created_at datetime        not null DEFAULT '',
    updated_at datetime        not null DEFAULT '',
    PRIMARY KEY (id)
);

create table file_folder
(
    id        bigint unsigned not null auto_increment,
    name      varchar(32)     not null default '' comment '文件夹名',
    parent_id bigint unsigned not null default 0 comment '父文件夹id',
    user_id   bigint unsigned not null,
    create_at datetime        not null default '' comment '创建时间',
    update_at datetime        not null default '' comment '更新时间',
    PRIMARY KEY (id)
);

create table file_repository
(
    id        bigint unsigned not null auto_increment,
    name      varchar(255)    not null default '' comment '文件名',
    filename  varchar(1023)   not null default '' comment '存储路径名',
    folder_id bigint unsigned not null default 0 comment '文件夹id',
    bucket    varchar(127)    not null default '' comment '桶名',
    ext       varchar(30)     not null DEFAULT '' COMMENT '文件扩展名',
    hash      varchar(255)    not null default '' comment '哈希值',
    size      bigint          not null default 0 comment '文件大小',
    url       varchar(1023)   not null default '' comment '访问地址',
    status    tinyint         not null default 1 comment '文件状态，1：上传成功，0：待合并',
    del_flag  tinyint         not null default 0 comment '文件删除状态：0：未删除，1：删除',
    create_at datetime        not null default '' comment '创建时间',
    update_at datetime        not null default '' comment '更新时间',
    done_at   datetime        not null default '' comment '大文件合并完成时间',
    PRIMARY KEY (id)
);

create table file_task
(
    id        bigint unsigned not null auto_increment,
    name      varchar(255)    not null default '' comment '文件名',
    file_id   bigint unsigned not null default '' comment '原文件id',
    filename  varchar(1023)   not null default '' comment '存储路径名',
    bucket    varchar(127)    not null default '' comment '桶名',
    ext       varchar(30)     not null DEFAULT '' COMMENT '文件扩展名',
    hash      varchar(255)    not null default '' comment '哈希值',
    size      bigint          not null default 0 comment '文件大小',
    create_at datetime        not null default '' comment '创建时间',
    update_at datetime        not null default '' comment '更新时间',
    PRIMARY KEY (id)
);

create table file_share
(
    id bigint unsigned not null auto_increment

);

