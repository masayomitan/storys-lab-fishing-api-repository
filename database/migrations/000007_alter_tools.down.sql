ALTER TABLE `story_fishing_db`.`tools` 
DROP FOREIGN KEY `fk_material`,
DROP FOREIGN KEY `fk_company`;
ALTER TABLE `story_fishing_db`.`tools` 
DROP COLUMN `company_id`,
DROP COLUMN `material_id`,
DROP INDEX `fk_company`,
DROP INDEX `fk_material`;
