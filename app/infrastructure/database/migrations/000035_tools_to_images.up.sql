CREATE TABLE tools_to_images (
    id INT AUTO_INCREMENT PRIMARY KEY,
    tool_id INT NOT NULL,
    image_id INT NOT NULL,
    sort INT DEFAULT 0 COMMENT '昇順/降順',
    is_main BOOLEAN DEFAULT FALSE COMMENT 'メイン画像',
    FOREIGN KEY (tool_id) REFERENCES tools(id),
    FOREIGN KEY (image_id) REFERENCES images(id)
) COMMENT='エリアと画像の中間テーブル';
