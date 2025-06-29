DROP TABLE IF EXISTS `SRO_SHARD`.`MASTERIES`;
CREATE TABLE  `SRO_SHARD`.`MASTERIES` (
  `FK_CHAR` int(11) NOT NULL,
  `ID` int NOT NULL,
  `LEVEL` int NOT NULL,
  KEY `FK_CHAR_idx` (`FK_CHAR`),
  CONSTRAINT `FK_MASTERIES_CHAR` FOREIGN KEY (`FK_CHAR`) REFERENCES `SRO_SHARD`.`CHAR` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
