CREATE TABLE fishing_methods_fishes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    fishing_method_id INT NOT NULL,
    fish_id INT NOT NULL,
    FOREIGN KEY (fishing_method_id) REFERENCES fishing_methods(id),
    FOREIGN KEY (fish_id) REFERENCES fishes(id)
) COMMENT='釣り方と魚種の関連情報を格納する中間テーブル';
