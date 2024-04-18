CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL COMMENT '名前',
    name_kana VARCHAR(255) COMMENT '名前のカナ表記',
    role INT COMMENT 'ユーザーの役割ID',
    initial_set BOOLEAN NOT NULL DEFAULT FALSE COMMENT '初期設定が完了しているかどうか',
    tel VARCHAR(255) COMMENT '電話番号',
    address VARCHAR(255) COMMENT '住所',
    sex INT COMMENT '性別（例：1=男性、2=女性、3=その他）',
    birthday DATE COMMENT '生年月日',
    last_login_datetime DATETIME COMMENT '最後のログイン日時',
    email VARCHAR(255) NOT NULL COMMENT 'メールアドレス',
    pass VARCHAR(255) NOT NULL COMMENT 'パスワード',
    created_at DATETIME NOT NULL COMMENT '作成日時',
    updated_at DATETIME NOT NULL COMMENT '更新日時',
    deleted_at DATETIME DEFAULT NULL COMMENT '削除日時'
) COMMENT='ユーザー情報を格納するテーブル';
