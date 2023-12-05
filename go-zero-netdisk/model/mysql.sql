CREATE TABLE `user`
(
    id         BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    name       VARCHAR(32)     NOT NULL                        DEFAULT '',
    username   VARCHAR(32)     NOT NULL                        DEFAULT '',
    password   VARCHAR(32)     NOT NULL                        DEFAULT '',
    email      VARCHAR(127)    NOT NULL                        DEFAULT '',
    avatar     VARCHAR(127)    NOT NULL                        DEFAULT '',
    signature  VARCHAR(100)    NOT NULL                        DEFAULT '个性签名',
    created_at DATETIME        NOT NULL                        DEFAULT '',
    updated_at DATETIME        NOT NULL ON UPDATE CURRENT_TIME DEFAULT '',
    PRIMARY KEY (id)
);

-- file_folder
CREATE TABLE file_folder
(
    id        BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    name      VARCHAR(32)     NOT NULL                        DEFAULT '' COMMENT '文件夹名',
    parent_id BIGINT UNSIGNED NOT NULL                        DEFAULT 0 COMMENT '父文件夹id',
    user_id   BIGINT UNSIGNED NOT NULL,
    create_at DATETIME        NOT NULL                        DEFAULT '' COMMENT '创建时间',
    update_at DATETIME        NOT NULL ON UPDATE CURRENT_TIME DEFAULT '' COMMENT '更新时间',
    PRIMARY KEY (id)
);

CREATE TABLE file_task
(
    id        BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    name      VARCHAR(255)    NOT NULL                        DEFAULT '' COMMENT '文件名',
    file_id   BIGINT UNSIGNED NOT NULL                        DEFAULT '' COMMENT '原文件id',
    filename  VARCHAR(255)    NOT NULL                        DEFAULT '' COMMENT '存储路径名',
    bucket    VARCHAR(255)    NOT NULL                        DEFAULT '' COMMENT '桶名',
    ext       VARCHAR(64)     NOT NULL                        DEFAULT '' COMMENT '文件扩展名',
    hash      VARCHAR(255)    NOT NULL                        DEFAULT '' COMMENT '哈希值',
    size      BIGINT          NOT NULL                        DEFAULT 0 COMMENT '文件大小',
    create_at DATETIME        NOT NULL                        DEFAULT '' COMMENT '创建时间',
    update_at DATETIME        NOT NULL ON UPDATE CURRENT_TIME DEFAULT '' COMMENT '更新时间',
    PRIMARY KEY (id)
);

-- file_netdisk
CREATE TABLE file_netdisk
(
    id            BIGINT AUTO_INCREMENT PRIMARY KEY,
    create_at     DATETIME     NOT NULL                        DEFAULT '',
    update_at     DATETIME     NOT NULL ON UPDATE CURRENT_TIME DEFAULT '',
    delete_at     DATETIME     NOT NULL                        DEFAULT '',
    user_id       BIGINT       NOT NULL                        DEFAULT 0 COMMENT '用户id',
    repository_id BIGINT       NOT NULL                        DEFAULT 0 COMMENT '实际存储id',
    folder_id     BIGINT       NOT NULL                        DEFAULT 0 COMMENT '文件夹id',
    name          VARCHAR(255) NOT NULL                        DEFAULT 0 COMMENT '用户视角文件名',
    url           VARCHAR(255) NOT NULL                        DEFAULT 0 COMMENT '访问地址',
    status        TINYINT      NOT NULL                        DEFAULT 0 COMMENT '文件状态，1：上传成功（小文件为1），0：待合并（大文件）',
    done_at       VARCHAR(255) NOT NULL                        DEFAULT 0 COMMENT '大文件合并完成时间',
    del_flag      TINYINT      NOT NULL                        DEFAULT 0 COMMENT '文件删除状态：0：未删除，1：删除（回收站）'
);

-- file_repository
CREATE TABLE file_repository
(
    id        BIGINT       NOT NULL                        DEFAULT 0 AUTO_INCREMENT PRIMARY KEY,
    create_at DATETIME     NOT NULL                        DEFAULT '',
    update_at DATETIME     NOT NULL ON UPDATE CURRENT_TIME DEFAULT '',
    delete_at DATETIME     NOT NULL                        DEFAULT '',
    bucket    VARCHAR(255) NOT NULL                        DEFAULT '' COMMENT '桶名',
    ext       VARCHAR(64)  NOT NULL                        DEFAULT '' COMMENT '文件扩展名',
    filename  VARCHAR(255) NOT NULL                        DEFAULT '' COMMENT '存储路径名',
    hash      VARCHAR(255) NOT NULL                        DEFAULT '' COMMENT '哈希值',
    name      VARCHAR(255) NOT NULL                        DEFAULT '' COMMENT '实际文件名',
    size      BIGINT       NOT NULL                        DEFAULT 0 COMMENT '文件大小',
    url       VARCHAR(255) NOT NULL                        DEFAULT '' COMMENT '访问地址',
    status    TINYINT      NOT NULL                        DEFAULT 1 COMMENT '文件状态，1：上传成功，0：待合并',
    done_at   VARCHAR(255) NOT NULL                        DEFAULT '' COMMENT '大文件合并完成时间'
);

-- file_uploading
CREATE TABLE file_uploading
(
    id            BIGINT NOT NULL DEFAULT 0 AUTO_INCREMENT PRIMARY KEY,
    create_at     DATETIME,
    update_at     DATETIME,
    delete_at     DATETIME,
    netdisk_id    BIGINT,
    repository_id BIGINT,
    chunk_num     INT
);

create table file_share
(
    id bigint unsigned not null auto_increment

);