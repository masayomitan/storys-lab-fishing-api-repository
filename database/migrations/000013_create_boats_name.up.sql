CREATE TABLE boats (
    id INT AUTO_INCREMENT PRIMARY KEY,
    company_id INT,
    registration_number VARCHAR(255) COMMENT '登録番号',
    boat_type INT COMMENT 'タイプ',
    length FLOAT COMMENT '長さ',
    width FLOAT COMMENT '幅',
    establish DATE COMMENT '創立日',
    max_crew_number INT COMMENT '乗船可能な最大乗員数',
    fishing_gear TEXT COMMENT '使用可能な漁具',
    price INT COMMENT '価格',
    gross_tonnage FLOAT COMMENT '総トン数',
    FOREIGN KEY (company_id) REFERENCES companies(id)
) COMMENT='船舶情報を保存するテーブル';
