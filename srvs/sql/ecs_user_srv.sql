/*
SQLyog Ultimate v8.32 
MySQL - 5.7.36 : Database - ecs_user_srv
*********************************************************************
*/


/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`ecs_user_srv` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `ecs_user_srv`;

/*Table structure for table `user` */

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(3) DEFAULT NULL,
  `update_time` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `is_deleted` tinyint(1) DEFAULT NULL,
  `mobile` varchar(11) NOT NULL,
  `password` varchar(100) NOT NULL,
  `nick_name` varchar(20) DEFAULT NULL,
  `birthday` datetime DEFAULT NULL,
  `gender` varchar(6) DEFAULT 'male' COMMENT 'female表示女, male表示男',
  `role` int(11) DEFAULT '1' COMMENT '1表示普通用户, 2表示管理员',
  PRIMARY KEY (`id`),
  UNIQUE KEY `mobile` (`mobile`),
  KEY `idx_mobile` (`mobile`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4;

/*Data for the table `user` */

insert  into `user`(`id`,`add_time`,`update_time`,`deleted_at`,`is_deleted`,`mobile`,`password`,`nick_name`,`birthday`,`gender`,`role`) values (1,'2022-07-30 23:40:06.347','2022-07-30 23:40:06.347',NULL,0,'18782222220','$pbkdf2-sha512$RZxJZfad3AlLM7fp$87b688cf87a9f2db0d579e5bf888b90dda3958f8af6e0abca25b4f61a3b0e1e1','bobby0',NULL,'male',1),(2,'2022-07-30 23:40:06.351','2022-07-30 23:40:06.351',NULL,0,'18782222221','$pbkdf2-sha512$zd7iaCYC8IVBAleM$8f70ce66cd3c4a976b7071e675cb542ffb5a59c97f3185d5703d482f6d963b06','bobby1',NULL,'male',1),(3,'2022-07-30 23:40:06.355','2022-07-30 23:40:06.355',NULL,0,'18782222222','$pbkdf2-sha512$4mWaYbKhbmDRD1RN$1e5c0da7d625c09d2da287d19975be3e16445244b5a5f54aab281dc98f111a5b','bobby2',NULL,'male',1),(4,'2022-07-30 23:40:06.357','2022-07-30 23:40:06.357',NULL,0,'18782222223','$pbkdf2-sha512$PTr97agIHk6BGevY$dc9d4ec1282dbee99e97c13e4a38e26e41fec942b0be2f7a9ddb900bb0a385c4','bobby3',NULL,'male',1),(5,'2022-07-30 23:40:06.359','2022-07-30 23:40:06.359',NULL,0,'18782222224','$pbkdf2-sha512$mh3t6VlM97eaU1j1$0593f6232bae46e6d4c81475e83db11b59b0b192c865d3f7c5122b9d5971213b','bobby4',NULL,'male',1),(6,'2022-07-30 23:40:06.361','2022-07-30 23:40:06.361',NULL,0,'18782222225','$pbkdf2-sha512$8w5jIexT7ta2J13j$efff968dc443bf3d94c230f25c4f2117a7f922818ee8620f06653533f9be718c','bobby5',NULL,'male',1),(7,'2022-07-30 23:40:06.364','2022-07-30 23:40:06.364',NULL,0,'18782222226','$pbkdf2-sha512$YQAcPT7AZR5rGhjm$d383db5a4b72a85d7cafb80a73c89080f66e66e8618c1379e65703e3cd680ec6','bobby6',NULL,'male',1),(8,'2022-07-30 23:40:06.366','2022-07-30 23:40:06.366',NULL,0,'18782222227','$pbkdf2-sha512$30fafTGbsR47Svvr$dd5f71f4b3365c1ddd2bc2ba4cd9a2c44486abd16a8aad12619a5e79c535925e','bobby7',NULL,'male',1),(9,'2022-07-30 23:40:06.368','2022-07-30 23:40:06.368',NULL,0,'18782222228','$pbkdf2-sha512$v1hD9nw22fLY5DjB$581dabf005a63defd6bfb2ae8f0cc50bdf16f1c8fed2a1eb204b9ee1279bda01','bobby8',NULL,'male',1),(10,'2022-07-30 23:40:06.370','2022-07-30 23:40:06.370',NULL,0,'18782222229','$pbkdf2-sha512$Lu9vftguecWAWwJp$50bf36b177b0e8b31286a7003238754cd5736ebcdbb48b3c07fbb59476ec9d28','bobby9',NULL,'male',1);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
