CREATE TABLE time_tags (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) COMMENT '名称',
    created_at DATETIME NOT NULL COMMENT '作成日時',
    updated_at DATETIME NOT NULL COMMENT '更新日時',
    deleted_at DATETIME DEFAULT NULL COMMENT '削除日時'
) COMMENT='釣り時間帯を保存するテーブル';
