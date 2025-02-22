CREATE TABLE areas (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) COMMENT 'エリア名称',
    description TEXT NULL COMMENT 'エリア説明',
    prefecture_id INT NOT NULL COMMENT '都道府県のID',
    created_at DATETIME NOT NULL COMMENT '作成日時',
    updated_at DATETIME NOT NULL COMMENT '更新日時',
    deleted_at DATETIME DEFAULT NULL COMMENT '削除日時',
    FOREIGN KEY (prefecture_id) REFERENCES prefectures(id),
    INDEX idx_prefecture (prefecture_id)
) COMMENT='都道府県を保存するテーブル';
