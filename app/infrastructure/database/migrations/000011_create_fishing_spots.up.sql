CREATE TABLE fishing_spots (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL COMMENT '釣りスポットの名前',
    area_id INT NOT NULL COMMENT 'エリアのID',
    description TEXT COMMENT '釣りスポットの説明',
    recommended_fishing_methods INT COMMENT 'おすすめ釣り方法',
    created_at DATETIME NOT NULL COMMENT '作成日時',
    updated_at DATETIME NOT NULL COMMENT '更新日時',
    deleted_at DATETIME DEFAULT NULL COMMENT '削除日時',
    FOREIGN KEY (area_id) REFERENCES areas(id),
    INDEX idx_area (area_id)
) COMMENT='釣りスポットとその特徴を保存するテーブル';
