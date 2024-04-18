CREATE TABLE fishing_methods_tools (
    id INT AUTO_INCREMENT PRIMARY KEY,
    fishing_method_id INT NOT NULL,
    tool_id INT NOT NULL,
    FOREIGN KEY (fishing_method_id) REFERENCES fishing_methods(id),
    FOREIGN KEY (tool_id) REFERENCES tools(id)
) COMMENT='釣り方と道具の関連情報を格納する中間テーブル';
