CREATE TABLE areas_to_images (
    id INT AUTO_INCREMENT PRIMARY KEY,
    area_id INT NOT NULL,
    image_id INT NOT NULL,
    sort INT DEFAULT 0 COMMENT '昇順/降順',
    is_main BOOLEAN DEFAULT FALSE COMMENT 'メイン画像',
    FOREIGN KEY (area_id) REFERENCES areas(id),
    FOREIGN KEY (image_id) REFERENCES images(id)
) COMMENT='エリアと画像の中間テーブル';
