/*
 Navicat Premium Data Transfer

 Source Server         : hefulin
 Source Server Type    : MySQL
 Source Server Version : 50740
 Source Host           : localhost:3306
 Source Schema         : qimo

 Target Server Type    : MySQL
 Target Server Version : 50740
 File Encoding         : 65001

 Date: 24/12/2023 17:30:00
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for order_table
-- ----------------------------
DROP TABLE IF EXISTS `order_table`;
CREATE TABLE `order_table`  (
  `r_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `space_number` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `start_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `end_time` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `cost` double NOT NULL DEFAULT 0,
  `status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `user_id_id` int(11) NOT NULL,
  `vehicle_id_id` int(11) NOT NULL,
  PRIMARY KEY (`r_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 21 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of order_table
-- ----------------------------
INSERT INTO `order_table` VALUES (9, '科技1路', 'A-002', '2023-12-14 14:38:13', '2023-12-22 14:54:22', 1.384338e15, '已完成', 4, 10);
INSERT INTO `order_table` VALUES (10, '科技3路', 'A-002', '2023-12-14 14:38:29', '2023-12-22 16:57:22', 1.399066e15, '已完成', 4, 11);
INSERT INTO `order_table` VALUES (14, '科技2路', 'A-002', '2023-12-14 14:40:47', '', 0, '未完成', 6, 15);
INSERT INTO `order_table` VALUES (15, '科技1路', 'A-003', '2023-12-14 14:41:06', '', 0, '未完成', 6, 16);
INSERT INTO `order_table` VALUES (16, '科技1路', 'A-001', '2023-12-15 10:50:32', '2023-12-15 10:51:08', 72000000000, '已完成', 1, 17);
INSERT INTO `order_table` VALUES (18, '科技1路', 'A-001', '2023-12-15 19:12:31', '2023-12-15 19:13:06', 70000000000, '已完成', 8, 19);
INSERT INTO `order_table` VALUES (19, '科技1路', 'A-002', '2023-12-22 19:10:11', '2023-12-22 21:16:11', 15120000000000, '已完成', 4, 21);
INSERT INTO `order_table` VALUES (20, '科技2路', 'A-002', '2023-12-24 12:22:58', '2023-12-24 12:24:10', 144000000000, '已完成', 4, 23);

-- ----------------------------
-- Table structure for parking_space
-- ----------------------------
DROP TABLE IF EXISTS `parking_space`;
CREATE TABLE `parking_space`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `road_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `space_number` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of parking_space
-- ----------------------------
INSERT INTO `parking_space` VALUES (1, '路1', 'A-001', '占用');
INSERT INTO `parking_space` VALUES (2, '路1', 'A-002', '占用');
INSERT INTO `parking_space` VALUES (3, '路1', 'A-003', '占用');
INSERT INTO `parking_space` VALUES (4, '路2', 'A-001', '占用');
INSERT INTO `parking_space` VALUES (5, '路2', 'A-002', '占用');
INSERT INTO `parking_space` VALUES (6, '路2', 'A-003', '占用');
INSERT INTO `parking_space` VALUES (7, '路3', 'A-001', '占用');
INSERT INTO `parking_space` VALUES (8, '路3', 'A-002', '占用');

-- ----------------------------
-- Table structure for road_segment
-- ----------------------------
DROP TABLE IF EXISTS `road_segment`;
CREATE TABLE `road_segment`  (
  `road_id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `code` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`road_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of road_segment
-- ----------------------------
INSERT INTO `road_segment` VALUES (1, '科技1路', '路1', '不可用');
INSERT INTO `road_segment` VALUES (2, '科技2路', '路2', '不可用');
INSERT INTO `road_segment` VALUES (3, '科技3路', '路3', '可用');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `phone` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `role_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, '管理员', '17870047136', 'e10adc3949ba59abbe56e057f20f883e', '2241975694@qq.com', '管理员');
INSERT INTO `users` VALUES (4, 'yu2', '15581218604', 'e10adc3949ba59abbe56e057f20f883e', '2241975694@qq.com', '用户');
INSERT INTO `users` VALUES (6, 'wenjingcheng', '15078007134', 'e10adc3949ba59abbe56e057f20f883e', '2298976791@qq.com', '用户');
INSERT INTO `users` VALUES (8, 'huzhuzai', '110', '5f93f983524def3dca464469d2cf9f3e', '2241975694@qq.com', '用户');
INSERT INTO `users` VALUES (9, '12', '3456789', 'e10adc3949ba59abbe56e057f20f883e', '2241975694@qq.com', '管理员');
INSERT INTO `users` VALUES (10, 'ooooo', '17870057136', 'e10adc3949ba59abbe56e057f20f883e', '2241975694@qq.com', '管理员');
INSERT INTO `users` VALUES (11, '暗淡星辰', '15789007136', 'e10adc3949ba59abbe56e057f20f883e', '2256978994@qq.com', '用户');
INSERT INTO `users` VALUES (12, 'ad', '17899077136', 'e10adc3949ba59abbe56e057f20f883e', '2241975694@qq.com', '巡查员');
INSERT INTO `users` VALUES (13, 'go语言', '17890947158', 'fae0b27c451c728867a567e8c1bb4e53', '2241975694@qq.com', '巡查员');

-- ----------------------------
-- Table structure for vehicle
-- ----------------------------
DROP TABLE IF EXISTS `vehicle`;
CREATE TABLE `vehicle`  (
  `v_id` int(11) NOT NULL AUTO_INCREMENT,
  `plate_number` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `owner_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `status` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  PRIMARY KEY (`v_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 24 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of vehicle
-- ----------------------------
INSERT INTO `vehicle` VALUES (1, '111111', 'yu', '已离开');
INSERT INTO `vehicle` VALUES (3, '333333', 'yu', '已离开');
INSERT INTO `vehicle` VALUES (4, '444444', 'yu', '未离开');
INSERT INTO `vehicle` VALUES (5, '555555', 'yu', '未离开');
INSERT INTO `vehicle` VALUES (6, '666666', 'yu', '未离开');
INSERT INTO `vehicle` VALUES (7, '777777', 'yu', '未离开');
INSERT INTO `vehicle` VALUES (8, '0000000', 'yu', '已离开');
INSERT INTO `vehicle` VALUES (9, '333', 'yu', '未离开');
INSERT INTO `vehicle` VALUES (10, 'a11111', 'yu2', '未离开');
INSERT INTO `vehicle` VALUES (11, 'a33333', 'yu2', '未离开');
INSERT INTO `vehicle` VALUES (12, 'b2222', 'andanxinkong', '未离开');
INSERT INTO `vehicle` VALUES (13, 'b3333', 'andanxinkong', '未离开');
INSERT INTO `vehicle` VALUES (14, 'c1111', 'wenjingcheng', '未离开');
INSERT INTO `vehicle` VALUES (15, 'c7777', 'wenjingcheng', '未离开');
INSERT INTO `vehicle` VALUES (16, 'c99999', 'wenjingcheng', '未离开');
INSERT INTO `vehicle` VALUES (17, '444444444444', 'hefulin', '未离开');
INSERT INTO `vehicle` VALUES (18, '99999', '12', '未离开');
INSERT INTO `vehicle` VALUES (19, '123456', 'huzhuzai', '未离开');
INSERT INTO `vehicle` VALUES (20, '赣-D88888', '贺福林', '未离开');
INSERT INTO `vehicle` VALUES (21, '沪D-99999', 'yu2', '已离开');
INSERT INTO `vehicle` VALUES (23, '6666666', 'yu2', '未离开');

SET FOREIGN_KEY_CHECKS = 1;
