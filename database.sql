

DROP TABLE IF EXISTS `games`;

CREATE TABLE `games` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` text NOT NULL,
  `year` int(11) NOT NULL,
  `company` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8;


LOCK TABLES `games` WRITE;

INSERT INTO `games` (`id`, `name`, `year`, `company`) VALUES (4,'Elden Ring',2022,'Bandai Namco'),(8,'Division 2',2019,'Ubisoft'),(9,'Rayman Legends',2013,'Ubisoft'),(10,'Dark Souls 3',2016,'Bandai Namco'),(11,'Disney Magical World',2013,'Bandai Namco'),(12,'Assassin\'s Creed Odyssey',2018,'Ubisoft'),(13,'Bloodborne',2015,'FromSoftware'),(14,'The Elder Scrolls Online',2014,'ZeniMax'),(15,'Just Cause 3',2015,'Avalanche'),(16,'Prototype 2',2012,'Radical Entertainment'),(17,'Prototype',2009,'Radical Entertainment'),(18,'Roblox',2006,'Roblox Corporation'),(19,'Ghost of Tsushima',2020,'Sucker Punch Productions');



