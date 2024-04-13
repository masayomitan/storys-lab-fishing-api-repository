ALTER TABLE tools
ADD material_id INT NOT NULL COMMENT '材料のID' AFTER tool_category_id,
ADD company_id INT NULL COMMENT '会社のID' AFTER material_id,
ADD CONSTRAINT fk_material
    FOREIGN KEY (material_id) REFERENCES materials(ID),
ADD CONSTRAINT fk_company
    FOREIGN KEY (company_id) REFERENCES companies(ID);
