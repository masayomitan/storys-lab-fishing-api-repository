CREATE TABLE prefectures (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) COMMENT '都道府県',
    name_kana VARCHAR(255) COMMENT '都道府県(カナ)',
    image_url VARCHAR(255) NOT NULL COMMENT '画像保存URL',
    created_at DATETIME NOT NULL COMMENT '作成日時',
    updated_at DATETIME NOT NULL COMMENT '更新日時',
    deleted_at DATETIME DEFAULT NULL COMMENT '削除日時'
) COMMENT='都道府県を保存するテーブル';
