/*
SQLyog Ultimate v8.32 
MySQL - 5.7.36 : Database - ecs_userop_srv
*********************************************************************
*/


/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`ecs_userop_srv` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `ecs_userop_srv`;

/*Table structure for table `address` */

DROP TABLE IF EXISTS `address`;

CREATE TABLE `address` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(3) DEFAULT NULL,
  `update_time` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `is_deleted` tinyint(1) DEFAULT NULL,
  `user` int(11) DEFAULT NULL,
  `province` varchar(10) DEFAULT NULL,
  `city` varchar(10) DEFAULT NULL,
  `district` varchar(20) DEFAULT NULL,
  `address` varchar(100) DEFAULT NULL,
  `signer_name` varchar(20) DEFAULT NULL,
  `signer_mobile` varchar(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_address_user` (`user`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `address` */

/*Table structure for table `leavingmessages` */

DROP TABLE IF EXISTS `leavingmessages`;

CREATE TABLE `leavingmessages` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(3) DEFAULT NULL,
  `update_time` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `is_deleted` tinyint(1) DEFAULT NULL,
  `user` int(11) DEFAULT NULL,
  `message_type` int(11) DEFAULT NULL COMMENT '留言类型: 1(留言),2(投诉),3(询问),4(售后),5(求购)',
  `subject` varchar(100) DEFAULT NULL,
  `message` longtext,
  `file` varchar(200) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_leavingmessages_user` (`user`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `leavingmessages` */

/*Table structure for table `userfav` */

DROP TABLE IF EXISTS `userfav`;

CREATE TABLE `userfav` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(3) DEFAULT NULL,
  `update_time` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `is_deleted` tinyint(1) DEFAULT NULL,
  `user` int(11) DEFAULT NULL,
  `goods` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_goods` (`user`,`goods`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `userfav` */

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
