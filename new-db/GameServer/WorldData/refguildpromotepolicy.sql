DROP TABLE IF EXISTS `refguildpromotepolicy`;
CREATE TABLE `refguildpromotepolicy` (
	`Level` INTEGER(1) NOT NULL,
	`Gold` INTEGER NOT NULL,
	`GP` INTEGER NOT NULL,
	PRIMARY KEY(`Level`)
);

INSERT INTO `refguildpromotepolicy` VALUES ('1','500000','0');
INSERT INTO `refguildpromotepolicy` VALUES ('2','3000000','5400');
INSERT INTO `refguildpromotepolicy` VALUES ('3','9000000','50400');
INSERT INTO `refguildpromotepolicy` VALUES ('4','150000000','135000');
INSERT INTO `refguildpromotepolicy` VALUES ('5','21000000','378000');

