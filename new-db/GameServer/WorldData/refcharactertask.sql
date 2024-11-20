DROP TABLE IF EXISTS `refcharactertask`;
CREATE TABLE `refcharactertask` (
	`refCharID` INTEGER NOT NULL,
	`refName` VARCHAR(50) NOT NULL,
	`TaskID` INTEGER NOT NULL,
	`MinLv` INTEGER(3) NOT NULL
);
