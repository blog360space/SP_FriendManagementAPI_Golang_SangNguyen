ALTER DATABASE friendmanagement CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

-- Migration table
CREATE TABLE `__EFMigrationsHistory` (
    `MigrationId` varchar(150) NOT NULL,
    `ProductVersion` varchar(32) NOT NULL,
    PRIMARY KEY (`MigrationId`)
);

-- User table
CREATE TABLE User
(
	ID INT auto_increment,
	Username nvarchar(50) not null,
	constraint Users_pk
		primary key (ID)
);

create unique index User_Username_uindex
	on User (Username);

ALTER TABLE User CONVERT TO CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;


-- User_Friend table
create table User_Friend
(
	ID int auto_increment,
	ToUserID int not null,
	FromUserID int not null,
	constraint User_Friend_pk
		primary key (ID)
);

-- Subscribe_User table
create table Subscribe_User
(
	ID int auto_increment,
	Requestor int not null,
	Target int not null,
    Status nvarchar(50) not null,
	constraint Subscribe_User_pk
		primary key (ID)
);


-- INSERT INTO `__EFMigrationsHistory` (`MigrationId`, `ProductVersion`)
-- VALUES ('20191203071427_initSchema', '2.1.14-servicing-32113');

