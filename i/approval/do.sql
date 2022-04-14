drop table approval_faform;
CREATE TABLE `approval_faform` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(150) NOT NULL,
  `en_name` varchar(150) NOT NULL,
  `type` varchar(150) NOT NULL,
  `serial_number` int NOT NULL,
  `f_approval_id` int NOT NULL,
  `del_tag` smallint NOT NULL,
  PRIMARY KEY (`id`)
);
drop table approval_selectoptions;
CREATE TABLE `approval_selectoptions` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `value` varchar(120) NOT NULL,
  `form_id` int NOT NULL,
  `del_tag` smallint NOT NULL,
  PRIMARY KEY (`id`)
);
drop table approval_fanode;
CREATE TABLE `approval_fanode` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(150) NOT NULL,
  `need_approver` tinyint(1) NOT NULL,
  `node_type` varchar(100) NOT NULL,
  `user_id` varchar(150),
  `approval_type` varchar(150) NOT NULL,

  `serial_number` int NOT NULL,
  `f_approval_id` int NOT NULL,
  `del_tag` smallint NOT NULL,
  PRIMARY KEY (`id`)
);
drop table approval_fapproval;
CREATE TABLE `approval_fapproval` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `approval_name` varchar(250) NOT NULL,
  `approval_code` varchar(250) NOT NULL,
  `open_id` varchar(250) NOT NULL,
  `job_name` varchar(250) NOT NULL,
  `subscribe` smallint NOT NULL,
  `descriptions` varchar(250) NOT NULL DEFAULT "",
  `del_tag` smallint NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `approval_name` (`approval_name`),
  UNIQUE KEY `approval_code` (`approval_code`),
  UNIQUE KEY `open_id` (`open_id`),
  UNIQUE KEY `job_name` (`job_name`)
);
drop table approval_ficomment;
CREATE TABLE `approval_ficomment` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_id` varchar(150) NOT NULL,
  `comment` varchar(250) NOT NULL,
  `create_time` varchar(150) NOT NULL,
  `f_instance_id` int NOT NULL,
  `del_tag` smallint NOT NULL,
  PRIMARY KEY (`id`)
);
drop table approval_fiform;
CREATE TABLE `approval_fiform` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(250) NOT NULL,
  `en_name` varchar(250) NOT NULL,
  `type` varchar(50) NOT NULL DEFAULT 'input',
  `f_instance_id` int NOT NULL,
  `del_tag` smallint NOT NULL,
  `value` varchar(250) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
);
drop table approval_finstancevalue;
CREATE TABLE `approval_finstancevalue` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `instance_code` varchar(250) NOT NULL,
  `status` varchar(100) NOT NULL,
  `user_id` varchar(150) NOT NULL,
  `descriptions` varchar(250) NOT NULL,
  `job_number` int NOT NULL,
  `f_approval_id` int NOT NULL,
  `del_tag` smallint NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `instance_code` (`instance_code`)
);
drop table approval_fitask;
CREATE TABLE `approval_fitask` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(150) NOT NULL,
  `user_id` varchar(150) NOT NULL,
  `status` varchar(150) NOT NULL,
  `type` varchar(150) NOT NULL,
  `comment` varchar(250) NOT NULL,
  `start_time` varchar(150) NOT NULL,
  `end_time` varchar(150) NOT NULL,
  `f_node_id` int NOT NULL,
  `f_instance_id` int NOT NULL,
  `serial_number` int NOT NULL,
  `del_tag` smallint NOT NULL,
  PRIMARY KEY (`id`)
);
drop table approval_fiviewers;
CREATE TABLE `approval_fiviewers` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `type` varchar(100) NOT NULL,
  `user_id` varchar(150) NOT NULL,
  `f_instance_id` int NOT NULL,
  `del_tag` smallint NOT NULL,
  PRIMARY KEY (`id`)
);