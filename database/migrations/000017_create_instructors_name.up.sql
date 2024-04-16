CREATE TABLE instructors (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL COMMENT '名前',
    name_kana VARCHAR(255) COMMENT 'カナ名',
    email VARCHAR(255) NOT NULL COMMENT 'メールアドレス',
    tel VARCHAR(255) COMMENT '電話番号',
    pass VARCHAR(255) NOT NULL COMMENT 'パスワード',
    initial_setup INT COMMENT '初期設定状態',
    prefecture_id INT COMMENT '都道府県ID',
    introduced TEXT COMMENT '紹介文',
    can_edit_events BOOLEAN NOT NULL DEFAULT FALSE COMMENT 'イベント作成/編集権限',
    FOREIGN KEY (prefecture_id) REFERENCES prefectures(id),
    INDEX idx_prefecture (prefecture_id)
) COMMENT='インストラクターテーブル';
