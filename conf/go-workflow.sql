/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 50018
Source Host           : localhost:3306
Source Database       : go-workflow

Target Server Type    : MYSQL
Target Server Version : 50018
File Encoding         : 65001

Date: 2020-07-25 22:11:13
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for bytearry
-- ----------------------------
DROP TABLE IF EXISTS `bytearry`;
CREATE TABLE `bytearry` (
  `id` bigint(20) NOT NULL auto_increment COMMENT '主键id',
  `name` varchar(64) collate utf8_bin default NULL COMMENT '名称',
  `key` varchar(64) collate utf8_bin default NULL COMMENT '流程定义key',
  `version` bigint(2) default NULL COMMENT '版本',
  `deployment_id` bigint(64) default NULL COMMENT '流程部署id',
  `bytes` longblob COMMENT '资源文件',
  PRIMARY KEY  (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Table structure for deployment
-- ----------------------------
DROP TABLE IF EXISTS `deployment`;
CREATE TABLE `deployment` (
  `id` bigint(20) NOT NULL auto_increment COMMENT '主键id',
  `name` varchar(64) collate utf8_bin default NULL COMMENT '名称',
  `key` varchar(64) collate utf8_bin default NULL COMMENT '流程定义key',
  `version` bigint(2) default NULL COMMENT '版本',
  `tenantI_id` bigint(64) default NULL COMMENT '租户id',
  `deploy_time` timestamp NOT NULL default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP COMMENT '部署时间',
  PRIMARY KEY  (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Table structure for identity_link
-- ----------------------------
DROP TABLE IF EXISTS `identity_link`;
CREATE TABLE `identity_link` (
  `id` bigint(20) NOT NULL auto_increment COMMENT '主键id',
  `type` varchar(64) collate utf8_bin default NULL COMMENT '类型',
  `task_id` varchar(64) collate utf8_bin default NULL COMMENT '任务ID',
  `proc_inst_id` bigint(2) default NULL COMMENT '流程实例id',
  `group_id` bigint(64) default NULL COMMENT '组id',
  `user_id` timestamp NOT NULL default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP COMMENT '用户id',
  PRIMARY KEY  (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Table structure for process_instance
-- ----------------------------
DROP TABLE IF EXISTS `process_instance`;
CREATE TABLE `process_instance` (
  `id` bigint(20) NOT NULL auto_increment COMMENT '主键id',
  `key` varchar(64) collate utf8_bin default NULL COMMENT '流程定义key',
  `name` varchar(64) collate utf8_bin default NULL COMMENT '名称',
  `version` bigint(2) default NULL COMMENT '版本',
  `business_key` varchar(64) collate utf8_bin default NULL,
  `tenant_id` varchar(64) collate utf8_bin default NULL COMMENT '租户id',
  `deployment_id` bigint(20) default NULL COMMENT '流程部署id',
  `start_time` timestamp NOT NULL default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP COMMENT '启动时间',
  `start_user_id` varchar(64) collate utf8_bin default NULL COMMENT '启动用户',
  `process_define_id` bigint(20) default NULL,
  PRIMARY KEY  (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Table structure for task
-- ----------------------------
DROP TABLE IF EXISTS `task`;
CREATE TABLE `task` (
  `id` bigint(20) NOT NULL auto_increment COMMENT '主键id',
  `task_define_key` varchar(64) collate utf8_bin default NULL COMMENT '任务key',
  `task_define_name` varchar(64) collate utf8_bin default NULL COMMENT '任务名称',
  `version` bigint(2) default NULL COMMENT '版本',
  `tenant_id` varchar(64) collate utf8_bin default NULL COMMENT '租户id',
  `deployment_id` bigint(64) default NULL COMMENT '流程部署id',
  `start_time` timestamp NOT NULL default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP COMMENT '启动时间',
  `assignee` varchar(64) collate utf8_bin default NULL COMMENT '审批人',
  `proc_inst_id` bigint(64) default NULL COMMENT '实例iD',
  PRIMARY KEY  (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Table structure for variable
-- ----------------------------
DROP TABLE IF EXISTS `variable`;
CREATE TABLE `variable` (
  `id` bigint(20) NOT NULL auto_increment COMMENT '主键id',
  `version` bigint(2) default NULL COMMENT '版本',
  `task_id` varchar(64) collate utf8_bin default NULL COMMENT '任务ID',
  `proc_inst_id` bigint(2) default NULL COMMENT '流程实例id',
  `name` varchar(64) collate utf8_bin default NULL COMMENT '变量名',
  `type` varchar(10) collate utf8_bin default NULL COMMENT '类型',
  `number` int(11) default NULL COMMENT '类型',
  `date` datetime default NULL COMMENT '时间类型',
  `float` double default NULL COMMENT 'double',
  `text` varchar(4000) collate utf8_bin default NULL COMMENT '字符',
  `blob` longblob COMMENT '字符串',
  PRIMARY KEY  (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Table structure for hi_identity_link
-- ----------------------------
DROP TABLE IF EXISTS `hi_identity_link`;
CREATE TABLE `hi_identity_link` (
  `id` bigint(20) NOT NULL auto_increment COMMENT '主键id',
  `type` varchar(64) collate utf8_bin default NULL COMMENT '类型',
  `task_id` varchar(64) collate utf8_bin default NULL COMMENT '任务ID',
  `proc_inst_id` bigint(2) default NULL COMMENT '流程实例id',
  `group_id` bigint(64) default NULL COMMENT '组id',
  `user_id` timestamp NOT NULL default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP COMMENT '用户id',
  PRIMARY KEY  (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Table structure for hi_process_instance
-- ----------------------------
DROP TABLE IF EXISTS `hi_process_instance`;
CREATE TABLE `hi_process_instance` (
  `id` bigint(20) NOT NULL auto_increment COMMENT '主键id',
  `key` varchar(64) collate utf8_bin default NULL COMMENT '流程定义key',
  `name` varchar(64) collate utf8_bin default NULL COMMENT '名称',
  `version` bigint(2) default NULL COMMENT '版本',
  `business_key` varchar(64) collate utf8_bin default NULL,
  `tenant_id` varchar(64) collate utf8_bin default NULL COMMENT '租户id',
  `deployment_id` bigint(20) default NULL COMMENT '流程部署id',
  `start_time` timestamp NULL default NULL COMMENT '启动时间',
  `end_time` timestamp NULL default NULL COMMENT '结束时间',
  `start_user_id` varchar(64) collate utf8_bin default NULL COMMENT '启动用户',
  `process_define_id` bigint(20) default NULL,
  `proc_inst_id` bigint(64) default NULL COMMENT '实例id',
  PRIMARY KEY  (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Table structure for hi_task
-- ----------------------------
DROP TABLE IF EXISTS `hi_task`;
CREATE TABLE `hi_task` (
  `id` bigint(20) NOT NULL auto_increment COMMENT '主键id',
  `task_define_key` varchar(64) collate utf8_bin default NULL COMMENT '任务key',
  `task_define_name` varchar(64) collate utf8_bin default NULL COMMENT '任务名称',
  `version` bigint(2) default NULL COMMENT '版本',
  `tenant_id` varchar(64) collate utf8_bin default NULL COMMENT '租户id',
  `deployment_id` bigint(64) default NULL COMMENT '流程部署id',
  `start_time` timestamp NULL default NULL COMMENT '启动时间',
  `end_time` timestamp NULL default NULL COMMENT '结束时间',
  `assignee` varchar(64) collate utf8_bin default NULL COMMENT '审批人',
  `proc_inst_id` bigint(64) default NULL COMMENT '实例id',
  `task_id` bigint(64) default NULL COMMENT '运行任务id',
  PRIMARY KEY  (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;


-- ----------------------------
-- Table structure for hi_variable
-- ----------------------------
DROP TABLE IF EXISTS `hi_variable`;
CREATE TABLE `hi_variable` (
  `id` bigint(20) NOT NULL auto_increment COMMENT '主键id',
  `version` bigint(2) default NULL COMMENT '版本',
  `task_id` varchar(64) collate utf8_bin default NULL COMMENT '任务ID',
  `proc_inst_id` bigint(2) default NULL COMMENT '流程实例id',
  `name` varchar(64) collate utf8_bin default NULL COMMENT '变量名',
  `type` varchar(10) collate utf8_bin default NULL COMMENT '类型',
  `number` int(11) default NULL COMMENT '类型',
  `date` datetime default NULL COMMENT '时间类型',
  `float` double default NULL COMMENT 'double',
  `text` varchar(4000) collate utf8_bin default NULL COMMENT '字符',
  `blob` longblob COMMENT '字符串',
  PRIMARY KEY  (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Table structure for hi_actinst
-- ----------------------------
DROP TABLE IF EXISTS `hi_actinst`;
CREATE TABLE `hi_actinst` (
  `id` bigint(20) NOT NULL auto_increment COMMENT '主键id',
  `task_id` varchar(64) collate utf8_bin default NULL COMMENT '任务ID',
  `proc_inst_id` bigint(2) default NULL COMMENT '流程实例id',
  `act_id` varchar(64) collate utf8_bin default NULL COMMENT '实例Id',
  `act_name` varchar(64) collate utf8_bin default NULL COMMENT '实例名称',
  `act_type` varchar(10) collate utf8_bin default NULL COMMENT '实例类型',
  `start_time` timestamp NULL default NULL COMMENT '启动时间',
  `end_time` timestamp NULL default NULL COMMENT '结束时间',
  `assignee` varchar(64) collate utf8_bin default NULL COMMENT '审批人',
  `start_user_id` varchar(64) collate utf8_bin default NULL COMMENT '启动用户',
  `process_define_id` bigint(20) default NULL,
  `tenant_id` varchar(64) collate utf8_bin default NULL COMMENT '租户id',
  PRIMARY KEY  (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

