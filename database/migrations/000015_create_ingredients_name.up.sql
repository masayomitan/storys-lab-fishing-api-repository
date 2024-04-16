CREATE TABLE ingredients (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL COMMENT '食材名',
    allergens VARCHAR(255) COMMENT 'アレルゲン情報',
    shelf_Life INT COMMENT '保存可能期間（日数）',
    storage_Method VARCHAR(255) COMMENT '保存方法',
    origin VARCHAR(255) COMMENT '原産国または産地',
    supplier VARCHAR(255) COMMENT '仕入れ先やブランド',
    unit VARCHAR(255) COMMENT '単位',
    season VARCHAR(255) COMMENT '旬や入手可能な時期'
) COMMENT='食材情報を格納するテーブル';
