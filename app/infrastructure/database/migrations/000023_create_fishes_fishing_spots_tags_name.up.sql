CREATE TABLE fish_fishing_spot_tags (
    id INT AUTO_INCREMENT PRIMARY KEY,
    fishing_spot_id INT NOT NULL,
    fish_id INT NOT NULL,
    time_tag_ids JSON COMMENT '時間帯タグのJSON配列',
    FOREIGN KEY (fishing_spot_id) REFERENCES fishing_spots(id),
    FOREIGN KEY (fish_id) REFERENCES fishes(id)
) COMMENT='魚種と釣り場のタグ情報を管理する中間テーブル';
