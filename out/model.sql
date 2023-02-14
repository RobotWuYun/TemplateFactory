-- User
DROP TABLE IF EXISTS `user`;
CREATE TABLE IF NOT EXISTS `user` (
`id` bigint(20) NOT NULL AUTO_INCREMENT,
 `name` varchar DEFAULT NULL COMMENT 'Name',
 `age_id` varchar DEFAULT NULL COMMENT 'AgeId', 
PRIMARY KEY (`id`),
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'User';
		