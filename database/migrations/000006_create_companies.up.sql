CREATE TABLE companies (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) COMMENT '会社名',
    name_kana VARCHAR(255) COMMENT '会社名カナ',
    zip_code INT COMMENT '郵便番号',
    address VARCHAR(255) COMMENT '住所',
    building_name VARCHAR(255) COMMENT '建物名',
    tel VARCHAR(255) COMMENT '電話番号',
    establish DATE COMMENT '設立日',
    created_at DATETIME NOT NULL COMMENT '作成日時',
    updated_at DATETIME NOT NULL COMMENT '更新日時',
    deleted_at DATETIME DEFAULT NULL COMMENT '削除日時'
) COMMENT='会社情報を保存するテーブル';
