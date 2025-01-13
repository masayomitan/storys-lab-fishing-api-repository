CREATE TABLE images (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) COMMENT '画像名称',
    image_url VARCHAR(255) NOT NULL COMMENT '画像保存URL',
    s3_url VARCHAR(255) NOT NULL COMMENT '画像保存S3URL',
    image_type INT DEFAULT 0 COMMENT '画像タイプ 例 1:魚、2:釣り場',
    created_at DATETIME NOT NULL COMMENT '作成日時',
    updated_at DATETIME NOT NULL COMMENT '更新日時',
    deleted_at DATETIME DEFAULT NULL COMMENT '削除日時'
);
