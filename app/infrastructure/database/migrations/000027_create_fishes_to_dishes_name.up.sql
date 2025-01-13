CREATE TABLE fishes_to_dishes (
    id INT AUTO_INCREMENT PRIMARY KEY,
    fish_id INT NOT NULL,
    dish_id INT NOT NULL,
    FOREIGN KEY (fish_id) REFERENCES fishes(id),
    FOREIGN KEY (dish_id) REFERENCES dishes(id)
) COMMENT='魚種と料理の関連情報を格納する中間テーブル';
