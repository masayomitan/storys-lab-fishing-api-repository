ALTER TABLE companies
ADD prefecture_id INT NOT NULL COMMENT '都道府県ID' AFTER establish,
ADD CONSTRAINT fk_prefecture FOREIGN KEY (prefecture_id) REFERENCES prefectures(id),
ADD INDEX idx_prefecture (prefecture_id);
