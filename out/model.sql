-- TestModel
DROP TABLE IF EXISTS `test_model`;
CREATE TABLE IF NOT EXISTS `test_model` (
`id` bigint(20) NOT NULL AUTO_INCREMENT,
 `test_string` VARCHAR() DEFAULT NULL COMMENT 'TestString',
 `test_bool` TINYINT(1) DEFAULT NULL COMMENT 'TestBool',
 `test_double` DOUBLE DEFAULT NULL COMMENT 'TestDouble',
 `test_float` FLOAT DEFAULT NULL COMMENT 'TestFloat',
 `test_int32` INTEGER DEFAULT NULL COMMENT 'TestInt32',
 `test_uint32` uint32 DEFAULT NULL COMMENT 'TestUint32',
 `test_int64` BIGINT DEFAULT NULL COMMENT 'TestInt64',
 `test_uint64` uint64 DEFAULT NULL COMMENT 'TestUint64',
 `test_sint32` INTEGER DEFAULT NULL COMMENT 'TestSint32',
 `test_sint64` BIGINT DEFAULT NULL COMMENT 'TestSint64',
 `test_sfixd32` uint32 DEFAULT NULL COMMENT 'TestSfixd32',
 `test_sfix64` uint64 DEFAULT NULL COMMENT 'TestSfix64',
 `test_bytes` VARCHAR() DEFAULT NULL COMMENT 'TestBytes', 
PRIMARY KEY (`id`),
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = 'TestModel';
		