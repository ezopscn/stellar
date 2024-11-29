package data

import (
	"stellar/common"
	"stellar/model"
)

// Logo 图片前缀
var logoUrlPrefix = "/assets/images/logo/"

// 数据源类型
var datasourceTypes = []model.DatasourceType{
	{
		Id:          1,
		Name:        "MySQL",
		Description: "MySQL 是一个关系型数据库管理系统，由瑞典 MySQL AB 公司开发，被 Oracle 收购，是最流行的数据库管理系统之一。",
		Logo:        logoUrlPrefix + "mysql.png",
	},
	{
		Id:          2,
		Name:        "ClickHouse",
		Description: "Clickhouse 是一个高性能的列式存储分析数据库管理系统，由俄罗斯搜索引擎公司 Yandex 开发。",
		Logo:        logoUrlPrefix + "clickhouse.png",
	},
	{
		Id:          3,
		Name:        "MongoDB",
		Description: "MongoDB 是一个高性能、开源、无模式的文档型数据库，由 C++ 编写，旨在为 WEB 应用提供可扩展的高性能数据存储解决方案。",
		Logo:        logoUrlPrefix + "mongodb.png",
	},
	{
		Id:          4,
		Name:        "Redis",
		Description: "Redis 是一个高性能的键值对存储系统，由 Salvatore Sanfilippo 开发。",
		Logo:        logoUrlPrefix + "redis.png",
	},
	{
		Id:          5,
		Name:        "PostgreSQL",
		Description: "PostgreSQL 是一款高级企业级开源关系数据库，由 PostgreSQL Global Development Group 开发。",
		Logo:        logoUrlPrefix + "postgresql.png",
	},
	{
		Id:          6,
		Name:        "Elasticsearch",
		Description: "Elasticsearch 是一个分布式、RESTful 风格的搜索和数据分析引擎，由 Elastic 公司开发。",
		Logo:        logoUrlPrefix + "elasticsearch.png",
	},
	{
		Id:          7,
		Name:        "Kafka",
		Description: "Kafka 是一个分布式流处理平台，由 Apache 软件基金会开发。",
		Logo:        logoUrlPrefix + "kafka.png",
	},
	{
		Id:          8,
		Name:        "HBase",
		Description: "HBase 是一个分布式、面向列的非关系型数据库，由 Apache 软件基金会开发。",
		Logo:        logoUrlPrefix + "hbase.png",
	},
	{
		Id:          9,
		Name:        "Hive",
		Description: "Hive 是一个开源的数据仓库系统，由 Facebook 开发。",
		Logo:        logoUrlPrefix + "hive.png",
	},
	{
		Id:          10,
		Name:        "Oracle",
		Description: "Oracle 是一个关系型数据库管理系统，由 Oracle 公司开发。",
		Logo:        logoUrlPrefix + "oracle.png",
	},
}

// 初始化数据源类型
func InitDatasourceType() {
	common.MySQLDB.Exec("TRUNCATE TABLE datasource_type")
	common.MySQLDB.Create(&datasourceTypes)
}
