CREATE TABLE fishing_spots_to_images (
    id INT AUTO_INCREMENT PRIMARY KEY,
    fishing_spot_id INT NOT NULL,
    image_id INT NOT NULL,
    sort INT DEFAULT 0 COMMENT '昇順/降順',
    is_main BOOLEAN DEFAULT FALSE COMMENT 'メイン画像',
    FOREIGN KEY (fishing_spot_id) REFERENCES fishing_spots(id),
    FOREIGN KEY (image_id) REFERENCES images(id)
) COMMENT='釣り場と画像の中間テーブル';
