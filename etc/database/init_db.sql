use smartapi;

CREATE TABLE `user` (
  `id` bigint(32) NOT NULL AUTO_INCREMENT,
  `username` varchar(255) CHARACTER SET latin1 DEFAULT "" COMMENT '用户名',
  `password` varchar(255) CHARACTER SET latin1 DEFAULT "" COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

