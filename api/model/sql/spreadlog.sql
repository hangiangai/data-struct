/*
Date: 2019-06-26 13:57:45
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for spreadlog_copy1
-- ----------------------------
DROP TABLE IF EXISTS `spreadlog_copy1`;
CREATE TABLE `spreadlog_copy1`  (
  `createAt` datetime(0) NOT NULL,
  `achievement` int(20) NULL DEFAULT NULL,
  `targetProcedure` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `targetContact` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `targetInfo` varchar(200) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `staffName` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `staffId` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `id` int(20) NOT NULL AUTO_INCREMENT,
  `uuid` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `update` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;


DELIMITER ;;
DROP TRIGGER IF EXISTS `before_insert_spreadlog`;;
CREATE TRIGGER `before_insert_spreadlog`
BEFORE INSERT ON `spreadlog`
FOR EACH ROW
BEGIN
-- 插入时生成UUID和创建时间
	SET new.uuid = uuid();
    SET new.createdat = now();
    SET new.updatedat = now();
END
;;

DROP TRIGGER IF EXISTS `after_update_spreadlog`;;
CREATE TRIGGER `after_update_spreadlog`
BEFORE UPDATE ON `spreadlog`
FOR EACH ROW
BEGIN
-- 更新时改变更新时间
    SET new.updatedat = now();
END
;;

DELIMITER ;
-- ----------------------------
-- Records of spreadlog
-- ----------------------------
