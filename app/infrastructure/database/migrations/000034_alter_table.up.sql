ALTER TABLE tides
ADD COLUMN area_id INT NOT NULL COMMENT 'エリアのID' after prefecture_id,
ADD CONSTRAINT fk_area FOREIGN KEY (area_id) REFERENCES areas(id),
ADD INDEX idx_area (area_id);
