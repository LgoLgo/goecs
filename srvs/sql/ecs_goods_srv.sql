/*
SQLyog Ultimate v8.32 
MySQL - 5.7.36 : Database - ecs_goods_srv
*********************************************************************
*/


/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`ecs_goods_srv` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `ecs_goods_srv`;

/*Table structure for table `banner` */

DROP TABLE IF EXISTS `banner`;

CREATE TABLE `banner` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(3) DEFAULT NULL,
  `update_time` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `is_deleted` tinyint(1) DEFAULT NULL,
  `image` varchar(200) NOT NULL,
  `url` varchar(200) NOT NULL,
  `index` int(11) NOT NULL DEFAULT '1',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `banner` */

/*Table structure for table `brands` */

DROP TABLE IF EXISTS `brands`;

CREATE TABLE `brands` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(3) DEFAULT NULL,
  `update_time` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `is_deleted` tinyint(1) DEFAULT NULL,
  `name` varchar(20) NOT NULL COMMENT '''品牌名称''',
  `logo` varchar(200) NOT NULL DEFAULT '' COMMENT '''品牌图标''',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4;

/*Data for the table `brands` */

insert  into `brands`(`id`,`add_time`,`update_time`,`deleted_at`,`is_deleted`,`name`,`logo`) values (1,NULL,NULL,NULL,0,'顶端',''),(2,NULL,NULL,NULL,0,'nba',''),(3,NULL,NULL,NULL,0,'文果',''),(4,NULL,NULL,NULL,0,'寻真',''),(5,NULL,NULL,NULL,0,'轻恋',''),(6,NULL,NULL,NULL,0,'木石',''),(7,NULL,NULL,NULL,0,'马小二',''),(8,NULL,NULL,NULL,0,'金山湾',''),(9,NULL,NULL,NULL,0,'日日顺','');

/*Table structure for table `category` */

DROP TABLE IF EXISTS `category`;

CREATE TABLE `category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(3) DEFAULT NULL,
  `update_time` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `is_deleted` tinyint(1) DEFAULT NULL,
  `name` varchar(20) NOT NULL,
  `parent_category_id` int(11) DEFAULT NULL,
  `level` int(11) NOT NULL DEFAULT '1' COMMENT '''1为1级类目，2为2级...''',
  `is_tab` tinyint(1) NOT NULL DEFAULT '0' COMMENT '''能否展示在Tab栏''',
  PRIMARY KEY (`id`),
  KEY `fk_category_sub_category` (`parent_category_id`),
  CONSTRAINT `fk_category_sub_category` FOREIGN KEY (`parent_category_id`) REFERENCES `category` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3030 DEFAULT CHARSET=utf8mb4;

/*Data for the table `category` */

insert  into `category`(`id`,`add_time`,`update_time`,`deleted_at`,`is_deleted`,`name`,`parent_category_id`,`level`,`is_tab`) values (1001,NULL,NULL,NULL,0,'新鲜水果',0,1,0),(1002,NULL,NULL,NULL,0,'新鲜蔬菜',0,1,0),(1003,NULL,NULL,NULL,0,'化妆品',0,1,0),(2001,NULL,NULL,NULL,0,'苹果',1001,2,0),(2002,NULL,NULL,NULL,0,'香蕉',1001,2,0),(2003,NULL,NULL,NULL,0,'凤梨',1001,2,0),(2004,NULL,NULL,NULL,0,'空心菜',1002,2,0),(2005,NULL,NULL,NULL,0,'黄花菜',1002,2,0),(2006,NULL,NULL,NULL,0,'白菜',1002,2,0),(2007,NULL,NULL,NULL,0,'口红',1003,2,0),(2008,NULL,NULL,NULL,0,'香水',1003,2,0),(2009,NULL,NULL,NULL,0,'护肤品',1003,2,0),(3001,NULL,NULL,NULL,0,'红富士',2001,3,0),(3002,NULL,NULL,NULL,0,'华圣果业',2001,3,0),(3003,NULL,NULL,NULL,0,'潘苹果',2001,3,0),(3004,NULL,NULL,NULL,0,'佳农',2002,3,0),(3005,NULL,NULL,NULL,0,'都乐',2002,3,0),(3006,NULL,NULL,NULL,0,'果迎鲜',2002,3,0),(3007,NULL,NULL,NULL,0,'甘福园',2003,3,0),(3008,NULL,NULL,NULL,0,'百果园',2003,3,0),(3009,NULL,NULL,NULL,0,'鲜蜂堆',2003,3,0),(3011,NULL,NULL,NULL,0,'沿海',2004,3,0),(3012,NULL,NULL,NULL,0,'南方',2004,3,0),(3013,NULL,NULL,NULL,0,'北方',2004,3,0),(3014,NULL,NULL,NULL,0,'绿色',2005,3,0),(3015,NULL,NULL,NULL,0,'蓝色',2005,3,0),(3016,NULL,NULL,NULL,0,'黄色',2005,3,0),(3017,NULL,NULL,NULL,0,'冬季',2006,3,0),(3018,NULL,NULL,NULL,0,'夏季',2006,3,0),(3019,NULL,NULL,NULL,0,'春季',2006,3,0),(3021,NULL,NULL,NULL,0,'Dior',2007,3,0),(3022,NULL,NULL,NULL,0,'香奈儿',2007,3,0),(3023,NULL,NULL,NULL,0,'YSL',2007,3,0),(3024,NULL,NULL,NULL,0,'让巴杜',2008,3,0),(3025,NULL,NULL,NULL,0,'恩加罗',2008,3,0),(3026,NULL,NULL,NULL,0,'龙芳',2008,3,0),(3027,NULL,NULL,NULL,0,'相宜本草',2009,3,0),(3028,NULL,NULL,NULL,0,'百雀羚',2009,3,0),(3029,NULL,NULL,NULL,0,'妮维雅',2009,3,0);

/*Table structure for table `goods` */

DROP TABLE IF EXISTS `goods`;

CREATE TABLE `goods` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(3) DEFAULT NULL,
  `update_time` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `is_deleted` tinyint(1) DEFAULT NULL,
  `category_id` int(11) NOT NULL,
  `brands_id` int(11) NOT NULL,
  `on_sale` tinyint(1) NOT NULL DEFAULT '0' COMMENT '''是否上架''',
  `ship_free` tinyint(1) NOT NULL DEFAULT '0' COMMENT '''是否免运费''',
  `is_new` tinyint(1) NOT NULL DEFAULT '0' COMMENT '''是否新品''',
  `is_hot` tinyint(1) NOT NULL DEFAULT '0' COMMENT '''是否热卖商品''',
  `name` varchar(50) NOT NULL,
  `goods_sn` varchar(50) NOT NULL COMMENT '''商家的内部编号''',
  `click_num` int(11) NOT NULL DEFAULT '0' COMMENT '''点击数''',
  `sold_num` int(11) NOT NULL DEFAULT '0' COMMENT '''销售量''',
  `fav_num` int(11) NOT NULL DEFAULT '0' COMMENT '''收藏数''',
  `market_price` float NOT NULL COMMENT '''商品价格''',
  `shop_price` float NOT NULL COMMENT '''实际价格''',
  `goods_brief` varchar(100) NOT NULL COMMENT '''商品简介''',
  `goods_front_image` varchar(200) NOT NULL COMMENT '''商品展示图''',
  PRIMARY KEY (`id`),
  KEY `fk_goods_category` (`category_id`),
  KEY `fk_goods_brands` (`brands_id`),
  CONSTRAINT `fk_goods_brands` FOREIGN KEY (`brands_id`) REFERENCES `brands` (`id`),
  CONSTRAINT `fk_goods_category` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4;

/*Data for the table `goods` */

insert  into `goods`(`id`,`add_time`,`update_time`,`deleted_at`,`is_deleted`,`category_id`,`brands_id`,`on_sale`,`ship_free`,`is_new`,`is_hot`,`name`,`goods_sn`,`click_num`,`sold_num`,`fav_num`,`market_price`,`shop_price`,`goods_brief`,`goods_front_image`) values (1,NULL,NULL,NULL,NULL,1001,1,0,0,0,0,'芒果','',0,0,0,20.1,12.3,'',''),(2,NULL,NULL,NULL,NULL,1002,5,0,0,0,0,'大白菜','',0,0,0,11.2,6.1,'',''),(3,NULL,NULL,NULL,NULL,1003,3,0,0,0,0,'卸妆水','',0,0,0,98.9,58.8,'',''),(4,NULL,NULL,NULL,NULL,2001,9,0,0,0,0,'香山苹果','',0,0,0,33.7,21.2,'',''),(5,NULL,NULL,NULL,NULL,2002,8,0,0,0,0,'巴西香蕉','',0,0,0,10.3,6.7,'',''),(6,NULL,NULL,NULL,NULL,2003,6,0,0,0,0,'台湾凤梨','',0,0,0,39.2,28.7,'',''),(7,NULL,NULL,NULL,NULL,2009,3,0,0,0,0,'粉底液','',0,0,0,198.9,158.8,'',''),(8,NULL,NULL,NULL,NULL,2007,3,0,0,0,0,'迪奥口红','',0,0,0,1168.8,699.9,'','');

/*Table structure for table `goodscategorybrand` */

DROP TABLE IF EXISTS `goodscategorybrand`;

CREATE TABLE `goodscategorybrand` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `add_time` datetime(3) DEFAULT NULL,
  `update_time` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `is_deleted` tinyint(1) DEFAULT NULL,
  `category_id` int(11) DEFAULT NULL,
  `brands_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_category_brand` (`category_id`,`brands_id`),
  KEY `fk_goodscategorybrand_brands` (`brands_id`),
  CONSTRAINT `fk_goodscategorybrand_brands` FOREIGN KEY (`brands_id`) REFERENCES `brands` (`id`),
  CONSTRAINT `fk_goodscategorybrand_category` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4;

/*Data for the table `goodscategorybrand` */

insert  into `goodscategorybrand`(`id`,`add_time`,`update_time`,`deleted_at`,`is_deleted`,`category_id`,`brands_id`) values (1,NULL,NULL,NULL,NULL,2001,1),(2,NULL,NULL,NULL,NULL,2002,2),(3,NULL,NULL,NULL,NULL,2003,3),(4,NULL,NULL,NULL,NULL,2004,4),(5,NULL,NULL,NULL,NULL,2005,5),(6,NULL,NULL,NULL,NULL,2006,6),(7,NULL,NULL,NULL,NULL,2007,7),(8,NULL,NULL,NULL,NULL,2008,8),(9,NULL,NULL,NULL,NULL,2009,9);

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
