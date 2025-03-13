CREATE TABLE fishing_spots_to_tags (
    id INT AUTO_INCREMENT PRIMARY KEY,
    fishing_spot_id INT NOT NULL,
    tag_id INT NOT NULL,
    FOREIGN KEY (fishing_spot_id) REFERENCES fishing_spots(id),
    FOREIGN KEY (tag_id) REFERENCES tags(id)
) COMMENT='釣り場とタグの中間テーブル';
