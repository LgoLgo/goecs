/*
SQLyog Ultimate v8.32 
MySQL - 5.7.36 : Database - ecs_inventory_srv
*********************************************************************
*/


/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`ecs_inventory_srv` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `ecs_inventory_srv`;

/*Table structure for table `inventory` */

DROP TABLE IF EXISTS `inventory`;

CREATE TABLE `inventory` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(3) DEFAULT NULL,
  `update_time` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `is_deleted` tinyint(1) DEFAULT NULL,
  `goods` int(11) DEFAULT NULL,
  `stocks` int(11) DEFAULT NULL,
  `version` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_inventory_goods` (`goods`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4;

/*Data for the table `inventory` */

insert  into `inventory`(`id`,`add_time`,`update_time`,`deleted_at`,`is_deleted`,`goods`,`stocks`,`version`) values (1,'2022-08-12 23:45:15.423','2022-08-12 23:45:15.423',NULL,0,1,100,0),(2,'2022-08-12 23:45:15.427','2022-08-12 23:45:15.427',NULL,0,2,100,0),(3,'2022-08-12 23:45:15.431','2022-08-12 23:45:15.431',NULL,0,3,100,0),(4,'2022-08-12 23:45:15.434','2022-08-12 23:45:15.434',NULL,0,4,100,0),(5,'2022-08-12 23:45:15.438','2022-08-12 23:45:15.438',NULL,0,5,100,0),(6,'2022-08-12 23:45:15.441','2022-08-12 23:45:15.441',NULL,0,6,100,0),(7,'2022-08-12 23:45:15.443','2022-08-12 23:45:15.443',NULL,0,7,100,0),(8,'2022-08-12 23:45:15.446','2022-08-12 23:45:15.446',NULL,0,8,100,0);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
