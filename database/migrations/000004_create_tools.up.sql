CREATE TABLE tools (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) COMMENT '釣り道具の名前',
    tool_category_id INT NOT NULL COMMENT '釣り道具カテゴリのID',
    size VARCHAR(255) COMMENT 'サイズ',
    weight FLOAT COMMENT '重量',
    durability INT COMMENT '耐久性',
    tool_usage TEXT COMMENT '使用用途',
    price INT COMMENT '価格',
    created_at DATETIME NOT NULL COMMENT '作成日時',
    updated_at DATETIME NOT NULL COMMENT '更新日時',
    deleted_at DATETIME DEFAULT NULL COMMENT '削除日時',
    FOREIGN KEY (tool_category_id) REFERENCES tool_categories(ID)
) COMMENT='釣り道具情報を保存するテーブル';
