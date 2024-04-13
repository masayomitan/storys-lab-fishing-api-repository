CREATE TABLE tool_categories (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) COMMENT '道具種類の名前',
    description TEXT COMMENT '詳細な説明',
    created_at DATETIME NOT NULL COMMENT '作成日時',
    updated_at DATETIME NOT NULL COMMENT '更新日時',
    deleted_at DATETIME DEFAULT NULL COMMENT '削除日時'
) COMMENT='道具種別情報を保存するテーブル';
