
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for salastatistics
-- ----------------------------
DROP TABLE IF EXISTS `salastatistics`;
CREATE TABLE `salastatistics`  (
  `id` int(200) NOT NULL AUTO_INCREMENT,
  `uuid` varchar(50) CHARACTER SET utf8 COLLATE utf8_estonian_ci NOT NULL,
  `staffname` varchar(100) CHARACTER SET utf8 COLLATE utf8_estonian_ci NOT NULL,
  `staffid` varchar(100) CHARACTER SET utf8 COLLATE utf8_estonian_ci NOT NULL,
  `createdat` datetime(0) NOT NULL,
  `updateat` datetime(0) NOT NULL,
  PRIMARY KEY (`id`, `uuid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_estonian_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;



DELIMITER ;;
DROP TRIGGER IF EXISTS `before_insert_salastatistics`;;
CREATE TRIGGER `before_insert_salastatistics`
BEFORE INSERT ON `salastatistics`
FOR EACH ROW
BEGIN
-- 插入时生成UUID和创建时间
	SET new.uuid = uuid();
    SET new.createdat = now();
    SET new.updatedat = now();
END
;;

DROP TRIGGER IF EXISTS `after_update_salastatistics`;;
CREATE TRIGGER `after_update_salastatistics`
BEFORE UPDATE ON `salastatistics`
FOR EACH ROW
BEGIN
-- 更新时改变更新时间
    SET new.updatedat = now();
END
;;

DELIMITER ;