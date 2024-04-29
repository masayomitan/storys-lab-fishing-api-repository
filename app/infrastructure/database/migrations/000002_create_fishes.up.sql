CREATE TABLE fishes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) COMMENT '魚の名前',
    family_name VARCHAR(255) COMMENT '科の名前',
    scientific_name VARCHAR(255) COMMENT '学名',
    fish_category_id INT NOT NULL COMMENT '魚種別 外部ID',
    description TEXT COMMENT '説明',
    length DECIMAL(10, 3) COMMENT '長さ',
    weight DECIMAL(10, 3) COMMENT '重さ',
    habitat VARCHAR(255) COMMENT '生息地',
    depth_range FLOAT COMMENT '生息深度の範囲',
    water_temperature_range FLOAT COMMENT '生息水温の範囲',
    conservation_status VARCHAR(255) COMMENT '保存状態',
    created_at DATETIME NOT NULL COMMENT '作成日時',
    updated_at DATETIME NOT NULL COMMENT '更新日時',
    deleted_at DATETIME DEFAULT NULL COMMENT '削除日時',
    FOREIGN KEY (fish_category_id) REFERENCES fish_categories(ID),
    INDEX idx_fish_category (fish_category_id)
) COMMENT='魚情報を保存するテーブル';
