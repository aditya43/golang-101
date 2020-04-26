-- MySQL dump 10.13  Distrib 5.7.18, for Linux (x86_64)
--
-- Host: localhost    Database: gopherfacedb
-- ------------------------------------------------------
-- Server version	5.7.18

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `friend_relation`
--

DROP TABLE IF EXISTS `friend_relation`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `friend_relation` (
  `owner_uuid` varchar(64) COLLATE utf8_unicode_ci NOT NULL,
  `friend_uuid` varchar(64) COLLATE utf8_unicode_ci NOT NULL,
  `created_ts` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_ts` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  UNIQUE KEY `relation` (`owner_uuid`,`friend_uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `friend_relation`
--

LOCK TABLES `friend_relation` WRITE;
/*!40000 ALTER TABLE `friend_relation` DISABLE KEYS */;
INSERT INTO `friend_relation` VALUES ('0d364cb6-7284-04ec-a7c5-5d794a035d61','4496dcd7-d29f-8153-56fa-dae61af5f4c2','2017-09-17 03:56:06','2017-09-17 03:56:06'),('0d364cb6-7284-04ec-a7c5-5d794a035d61','6f2a919b-b626-9402-2a5d-c290cf14afb3','2017-09-17 03:12:38','2017-09-17 03:12:38'),('6f2a919b-b626-9402-2a5d-c290cf14afb3','4496dcd7-d29f-8153-56fa-dae61af5f4c2','2017-09-17 03:09:15','2017-09-17 03:09:15');
/*!40000 ALTER TABLE `friend_relation` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `post`
--

DROP TABLE IF EXISTS `post`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `post` (
  `uuid` varchar(64) COLLATE utf8_unicode_ci NOT NULL,
  `title` varchar(65) COLLATE utf8_unicode_ci NOT NULL,
  `body` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `mood` int(11) DEFAULT NULL,
  `created_ts` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_ts` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `post`
--

LOCK TABLES `post` WRITE;
/*!40000 ALTER TABLE `post` DISABLE KEYS */;
INSERT INTO `post` VALUES ('4496dcd7-d29f-8153-56fa-dae61af5f4c2','Golang is awesome!','Hi! My name is ThreeD Gopher and this is my very first post on the Interwebs!\n\nI like programming in Go on the front-end and the back-end and I like making new gopher pals also!\n\nHit me up!',1,'2017-09-17 03:07:52','2017-09-17 03:07:52'),('6f2a919b-b626-9402-2a5d-c290cf14afb3','Welcome ThreeD!','Hey ThreeD!\n\nWelcome to GopherFace. Glad you could finally join the party!',8,'2017-09-17 03:10:18','2017-09-17 03:10:18'),('0d364cb6-7284-04ec-a7c5-5d794a035d61','I was busy merging code.','I was just going about my business merging code, but I had to look at what was going on in the GopherFace Feed!',9,'2017-09-17 03:12:03','2017-09-17 03:12:03');
/*!40000 ALTER TABLE `post` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` tinyint(1) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(18) COLLATE utf8_unicode_ci NOT NULL,
  `uuid` varchar(64) COLLATE utf8_unicode_ci NOT NULL,
  `first_name` varchar(64) COLLATE utf8_unicode_ci NOT NULL,
  `last_name` varchar(64) COLLATE utf8_unicode_ci NOT NULL,
  `password_hash` char(64) COLLATE utf8_unicode_ci NOT NULL,
  `email` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  `created_ts` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_ts` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'3DGopher','4496dcd7-d29f-8153-56fa-dae61af5f4c2','ThreeD','Gopher','d754ed9f64ac293b10268157f283ee23256fb32a4f8dedb25c8446ca5bcb0bb3','3dgopher@test.com','2017-09-16 03:16:51','2017-09-16 03:16:51'),(2,'Molly','6f2a919b-b626-9402-2a5d-c290cf14afb3','Molly','Gopher','d754ed9f64ac293b10268157f283ee23256fb32a4f8dedb25c8446ca5bcb0bb3','mollygopher@test.com','2017-09-16 03:23:44','2017-09-16 03:23:44'),(3,'Wintermute','0d364cb6-7284-04ec-a7c5-5d794a035d61','Wintermute','Gopher','d754ed9f64ac293b10268157f283ee23256fb32a4f8dedb25c8446ca5bcb0bb3','wintermute@test.com','2017-09-16 03:24:49','2017-09-16 03:24:49'),(4,'Case','935d0561-cebb-85ac-ab81-62685aee98a6','Case','Gopher','d754ed9f64ac293b10268157f283ee23256fb32a4f8dedb25c8446ca5bcb0bb3','case@test.com','2017-09-16 03:25:23','2017-09-16 03:25:23'),(5,'Linda','839964f4-140d-03c6-547c-9143092cef77','Linda','Gopher','d754ed9f64ac293b10268157f283ee23256fb32a4f8dedb25c8446ca5bcb0bb3','linda@test.com','2017-09-16 03:26:24','2017-09-16 03:26:24'),(6,'Riviera','0199063b-9936-ef19-04f3-7e38448e2b9c','Riviera','Gopher','d754ed9f64ac293b10268157f283ee23256fb32a4f8dedb25c8446ca5bcb0bb3','riviera@test.com','2017-09-16 03:27:10','2017-09-16 03:27:10'),(7,'Flatline','bc2589e1-130b-ff25-9c8d-c6459417270a','Flatline','Gopher','d754ed9f64ac293b10268157f283ee23256fb32a4f8dedb25c8446ca5bcb0bb3','flatline@test.com','2017-09-16 03:27:51','2017-09-16 03:27:51'),(8,'Jane','9e9e7389-6c07-3f4f-6404-665d3f53929b','Jane','Gopher','d754ed9f64ac293b10268157f283ee23256fb32a4f8dedb25c8446ca5bcb0bb3','jane@test.com','2017-09-16 03:28:45','2017-09-16 03:28:45'),(9,'Aerol','ea6c4fff-5471-3784-e91b-1e645482adde','Aerol','Gopher','d754ed9f64ac293b10268157f283ee23256fb32a4f8dedb25c8446ca5bcb0bb3','aerol@test.com','2017-09-16 03:29:29','2017-09-16 03:29:29'),(10,'Marie','6f7c4c58-a074-a8ed-8599-29699edbffac','Marie','Gopher','d754ed9f64ac293b10268157f283ee23256fb32a4f8dedb25c8446ca5bcb0bb3','marie@test.com','2017-09-16 03:31:10','2017-09-16 03:31:10');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user_profile`
--

DROP TABLE IF EXISTS `user_profile`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_profile` (
  `uuid` varchar(64) COLLATE utf8_unicode_ci NOT NULL,
  `about` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `location` varchar(64) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `interests` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `profile_image_path` varchar(255) COLLATE utf8_unicode_ci NOT NULL DEFAULT '',
  `created_ts` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_ts` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`uuid`),
  UNIQUE KEY `uuid` (`uuid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_profile`
--

LOCK TABLES `user_profile` WRITE;
/*!40000 ALTER TABLE `user_profile` DISABLE KEYS */;
INSERT INTO `user_profile` VALUES ('0199063b-9936-ef19-04f3-7e38448e2b9c','I am an actor and performer.','Straylight','Acting','/static/uploads/images/b2dcbd58-9489-132e-b393-9965aa40d146_thumb.png','2017-09-16 03:48:28','2017-09-16 03:49:08'),('0d364cb6-7284-04ec-a7c5-5d794a035d61','I enjoy creating plans and making calculations.','Matrix','Merging Code','/static/uploads/images/2c2ef747-38b9-9968-8215-ce28099a8b63_thumb.png','2017-09-16 03:49:19','2017-09-16 03:51:01'),('4496dcd7-d29f-8153-56fa-dae61af5f4c2','I like to program in Go on both the back-end and the front-end.','Los Angeles','Go, Web Development','/static/uploads/images/3cc3af2b-9fb2-86b5-05b5-80e32a08b9d6_thumb.png','2017-09-16 03:18:48','2017-09-16 22:38:33'),('6f2a919b-b626-9402-2a5d-c290cf14afb3','I am a multi-talented individual.','Chiba City','Getting things done.','/static/uploads/images/0ef62347-cead-bfbf-6d23-7cdd84f8b080_thumb.png','2017-09-16 03:44:54','2017-09-16 03:48:09'),('6f7c4c58-a074-a8ed-8599-29699edbffac','I\'m into artificial intelligence.','Matrix','AI','/static/uploads/images/6f3196b0-ccdc-c335-3984-f1e6b206e6a1_thumb.png','2017-09-16 03:42:43','2017-09-16 03:43:48'),('839964f4-140d-03c6-547c-9143092cef77','I like to analyze logic.','Matrix','Coding','/static/uploads/images/5a4cd16c-93f1-5bff-cefb-5173e6d8f941_thumb.png','2017-09-16 03:39:39','2017-09-16 03:42:24'),('935d0561-cebb-85ac-ab81-62685aee98a6','Don\'t mind me, I\'m just finding my way through cyberspace.','Chiba City','Solving complex cases.','/static/uploads/images/38a52ac9-7704-60e7-b690-abfec7309162_thumb.png','2017-09-16 03:34:06','2017-09-16 03:35:13'),('9e9e7389-6c07-3f4f-6404-665d3f53929b','I love to party and have fun.','Straylight','Cloning','/static/uploads/images/f79db6c9-3189-20c0-f0fc-bbe5e5f57335_thumb.png','2017-09-16 03:37:22','2017-09-16 03:39:19'),('bc2589e1-130b-ff25-9c8d-c6459417270a','I\'m a console cowboy!','ROM','System Programming','/static/uploads/images/c868ac58-d7b6-f0b4-93f1-4d7249f55226_thumb.png','2017-09-16 03:35:34','2017-09-16 03:37:07'),('ea6c4fff-5471-3784-e91b-1e645482adde','Just another Zionite tagging along with Case and Molly.','Villa Straylight','Piloting, Crafts','/static/uploads/images/605f6d17-d3f2-9090-ada7-27e72d1ebf72_thumb.png','2017-09-16 03:31:53','2017-09-16 03:33:51');
/*!40000 ALTER TABLE `user_profile` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-09-17 12:50:10
