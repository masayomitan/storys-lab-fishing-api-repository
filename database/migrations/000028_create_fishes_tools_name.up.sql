CREATE TABLE fishes_tools (
    id INT AUTO_INCREMENT PRIMARY KEY,
    fish_id INT NOT NULL,
    tool_id INT NOT NULL,
    FOREIGN KEY (fish_id) REFERENCES fishes(id),
    FOREIGN KEY (tool_id) REFERENCES tools(id)
) COMMENT='魚種と道具を管理する中間テーブル';
