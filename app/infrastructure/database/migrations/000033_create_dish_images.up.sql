CREATE TABLE dish_images (
    id INT AUTO_INCREMENT PRIMARY KEY,
    dish_id INT NOT NULL COMMENT '料理ID',
    image_url VARCHAR(255) NOT NULL COMMENT '画像保存URL',
    sort INT DEFAULT 0 COMMENT '昇順/降順',
    is_main BOOLEAN DEFAULT FALSE COMMENT 'メイン画像',
    created_at DATETIME NOT NULL COMMENT '作成日時',
    updated_at DATETIME NOT NULL COMMENT '更新日時',
    deleted_at DATETIME DEFAULT NULL COMMENT '削除日時',
    FOREIGN KEY (dish_id) REFERENCES dishes(id)
);
