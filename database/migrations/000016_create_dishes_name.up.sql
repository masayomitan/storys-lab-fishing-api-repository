CREATE TABLE dishes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) COMMENT '名前',
    description TEXT COMMENT '説明',
    ingredients JSON COMMENT '料理に使われる材料',
    kind VARCHAR(255) COMMENT '種類（前菜、温菜、冷菜など）',
    level INT COMMENT '難易度レベル（1〜5の範囲）'
) COMMENT='料理情報を格納するテーブル';
