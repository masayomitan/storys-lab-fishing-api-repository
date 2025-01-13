CREATE TABLE fishes_to_images (
    id INT AUTO_INCREMENT PRIMARY KEY,
    fish_id INT NOT NULL,
    image_id INT NOT NULL,
    sort INT DEFAULT 0 COMMENT '昇順/降順',
    is_main BOOLEAN DEFAULT FALSE COMMENT 'メイン画像',
    FOREIGN KEY (fish_id) REFERENCES fishes(id),
    FOREIGN KEY (image_id) REFERENCES images(id)
) COMMENT='魚と画像の中間テーブル';
