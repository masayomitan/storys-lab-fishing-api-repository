CREATE TABLE fishing_methods_fishing_spots (
    id INT AUTO_INCREMENT PRIMARY KEY,
    fishing_method_id INT NOT NULL,
    fishing_spot_id INT NOT NULL,
    FOREIGN KEY (fishing_method_id) REFERENCES fishing_methods(id),
    FOREIGN KEY (fishing_spot_id) REFERENCES fishing_spots(id)
) COMMENT='釣り方と釣り場の関連情報を格納する中間テーブル';
