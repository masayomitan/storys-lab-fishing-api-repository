ALTER TABLE tools
ADD COLUMN maker VARCHAR(255) COMMENT 'メーカー名' after price,
ADD COLUMN recommend TINYINT UNSIGNED COMMENT 'おすすめ度 (5段階)' after maker,
ADD COLUMN easy_fishing TINYINT UNSIGNED COMMENT '使いやすさ (5段階)' after recommend,
ADD COLUMN description TEXT COMMENT '説明' after easy_fishing;
