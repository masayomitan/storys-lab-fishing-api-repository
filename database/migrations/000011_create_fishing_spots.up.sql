CREATE TABLE fishing_spots (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL COMMENT '釣りスポットの名前',
    area_id INT NOT NULL COMMENT 'エリアのID',
    scientific_name VARCHAR(255) COMMENT '魚の学名',
    description TEXT COMMENT '釣りスポットの説明',
    length DECIMAL(10, 3) COMMENT '長さ',
    width DECIMAL(10, 3) COMMENT '幅',
    habitat VARCHAR(255) COMMENT '生息地',
    depth_range FLOAT COMMENT '生息深度の範囲',
    water_temperature_range FLOAT COMMENT '水温',
    conservation_status INT COMMENT '保存状態',
    edibility BOOLEAN COMMENT '食べられるか',
    recommended_fishing_methods INT COMMENT 'おすすめ釣り方法',
    FOREIGN KEY (area_id) REFERENCES areas(ID),
    INDEX idx_area (area_id)
) COMMENT='釣りスポットとその特徴を保存するテーブル';