CREATE TABLE fishing_methods (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL COMMENT '名称',
    description TEXT COMMENT '説明',
    difficulty_level INT COMMENT '難易度レベル'
) COMMENT='釣り方に関する情報を格納するテーブル';
