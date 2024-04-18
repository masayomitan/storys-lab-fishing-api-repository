CREATE TABLE admins (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL COMMENT '名前',
    name_kana VARCHAR(255) COMMENT '名前のカナ表記',
    email VARCHAR(255) NOT NULL COMMENT 'メールアドレス',
    pass VARCHAR(255) NOT NULL COMMENT 'パスワード',
    role INT COMMENT '役割',
    created_at DATETIME NOT NULL COMMENT '作成日時',
    updated_at DATETIME NOT NULL COMMENT '更新日時',
    deleted_at DATETIME DEFAULT NULL COMMENT '削除日時'
) COMMENT='管理者テーブル';
