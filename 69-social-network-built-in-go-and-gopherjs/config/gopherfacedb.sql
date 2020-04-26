/* *****************************************************************************
// Setup preferences
// ****************************************************************************/
SET NAMES utf8 COLLATE 'utf8_unicode_ci';
SET time_zone = '-07:00';
SET CHARACTER SET utf8;

/* *****************************************************************************
// Remove database (if it already exists)
// ****************************************************************************/
DROP DATABASE IF EXISTS gopherfacedb;

/* *****************************************************************************
// Create new database
// ****************************************************************************/
CREATE DATABASE gopherfacedb DEFAULT CHARSET = utf8 COLLATE = utf8_unicode_ci;
USE gopherfacedb;


/* *****************************************************************************
// Create the table(s)
// ****************************************************************************/

DROP TABLE IF EXISTS user;

CREATE TABLE user (
    id TINYINT(1) UNSIGNED NOT NULL AUTO_INCREMENT,
	username VARCHAR(18) NOT NULL,
	uuid VARCHAR(64) NOT NULL,
    first_name VARCHAR(64) NOT NULL,
    last_name VARCHAR(64) NOT NULL,
    password_hash CHAR(64) NOT NULL,
    email VARCHAR(255) NOT NULL,
    created_ts TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_ts TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	UNIQUE (username),
    PRIMARY KEY (id)
);


DROP TABLE IF EXISTS user_profile;

CREATE TABLE user_profile(
	uuid VARCHAR(64) NOT NULL,
	about VARCHAR(255) NOT NULL DEFAULT "",
	location VARCHAR(64) NOT NULL DEFAULT "",
	interests VARCHAR(255) NOT NULL DEFAULT "",
	profile_image_path VARCHAR(255) NOT NULL DEFAULT "",
    created_ts TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_ts TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	UNIQUE(uuid),
	PRIMARY KEY (uuid)
);


DROP TABLE IF EXISTS friend_relation;

CREATE TABLE friend_relation(
	owner_uuid	VARCHAR(64) NOT NULL,
	friend_uuid	VARCHAR(64) NOT NULL,
    created_ts TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_ts TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	UNIQUE KEY relation (owner_uuid, friend_uuid)
);


DROP TABLE IF EXISTS post;

CREATE TABLE post(
	uuid VARCHAR(64) NOT NULL,
	title VARCHAR(65) NOT NULL,
	body VARCHAR(255) NOT NULL,
	mood INT,
    created_ts TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_ts TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
