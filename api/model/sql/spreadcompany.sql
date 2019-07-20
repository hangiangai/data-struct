
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for spreadcompany
-- ----------------------------
DROP TABLE IF EXISTS `spreadcompany`;
CREATE TABLE `spreadcompany`  (
  `id` int(20) NOT NULL,
  `uuid` varchar(50) CHARACTER SET utf8 COLLATE utf8_estonian_ci NOT NULL,
  `companyName` varchar(200) CHARACTER SET utf8 COLLATE utf8_estonian_ci NOT NULL,
  `companyCode` varchar(50) CHARACTER SET utf8 COLLATE utf8_estonian_ci NOT NULL,
  `contact` varchar(100) CHARACTER SET utf8 COLLATE utf8_estonian_ci NOT NULL,
  `legalPerson` varchar(50) CHARACTER SET utf8 COLLATE utf8_estonian_ci NULL DEFAULT NULL,
  `createdAt` datetime(0) NOT NULL,
  `updateAt` datetime(0) NOT NULL,
  PRIMARY KEY (`id`, `uuid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_estonian_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;



DELIMITER ;;
DROP TRIGGER IF EXISTS `before_insert_spreadcompany`;;
CREATE TRIGGER `before_insert_spreadcompany`
BEFORE INSERT ON `spreadcompany`
FOR EACH ROW
BEGIN
-- 插入时生成UUID和创建时间
	SET new.uuid = uuid();
    SET new.createdat = now();
    SET new.updatedat = now();
END
;;

DROP TRIGGER IF EXISTS `after_update_spreadcompany`;;
CREATE TRIGGER `after_update_spreadcompany`
BEFORE UPDATE ON `spreadcompany`
FOR EACH ROW
BEGIN
-- 更新时改变更新时间
    SET new.updatedat = now();
END
;;

DELIMITER ;