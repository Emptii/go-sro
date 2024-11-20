DROP TABLE IF EXISTS `refpackageitem`;
CREATE TABLE `refpackageitem` (
	`EntryFlag` INTEGER(1) NOT NULL,
	`ItemFlag` INTEGER(4) NOT NULL,
	`ID` INTEGER NOT NULL,
	/* temporarily set from 50 to 70 since i am getting 
	'too long' errors on insert */
	`PKItem` VARCHAR(70) NOT NULL,
	`Reserv1` INTEGER(2) NOT NULL,
	`ExpandStr` VARCHAR(16) NOT NULL,
	/* temporarily set from 40 to 60 since i am getting 
	'too long' errors on insert */
	`ItemName` VARCHAR(60) NOT NULL,
	/* temporarily set from 40 to 60 since i am getting 
	'too long' errors on insert */
	`ItemDescMsg` VARCHAR(60) NOT NULL,
	/* temporarily set from 40 to 60 since i am getting 
	'too long' errors on insert */
	`Icon` VARCHAR(60) NOT NULL,
	`E1` INTEGER(1) NOT NULL,
	`ES1` VARCHAR(30) NOT NULL,
	`E2` INTEGER(1) NOT NULL,
	`ES2` VARCHAR(3) NOT NULL,
	`E3` INTEGER(1) NOT NULL,
	`ES3` VARCHAR(30) NOT NULL,
	`E4` INTEGER(1) NOT NULL,
	`ES4` VARCHAR(3) NOT NULL
);
