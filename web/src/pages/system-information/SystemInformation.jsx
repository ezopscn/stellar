import { Helmet } from 'react-helmet';
import { TitleSuffix } from '@/common/Text.jsx';
import { Descriptions, Card, App, Tag } from 'antd';
import { AxiosGet } from '@/utils/Request.jsx';
import { useState, useEffect } from 'react';
import { Apis } from '@/common/APIConfig.jsx';
import { GithubOutlined } from '@ant-design/icons';
import { Table } from 'antd';
import { NODE_ROLE_MAP } from '@/common/GlobalConfig.jsx';
// 页面常量设置
const PageConfig = {
  // 页面标题
  pageTitle: '系统信息' + TitleSuffix,
  // 页面顶部标题
  pageHeaderTitle: '系统信息 / SYSTEM INFORMATION.',
  // 页面关键词
  pageKeyword: '系统信息'
};

// 页面说明组件
const PageDescriptionComponent = () => {
  return (
    <>
      <div>欢迎来到系统信息页面，在这里您可以查看到系统的一些基础运行信息，主要包括：</div>
      <ul>
        <li>Stellar 系统前后端版本，构建相关信息。</li>
        <li>Stellar 各个角色运行的实例相关信息。</li>
      </ul>
    </>
  );
};

const SystemInformation = () => {
  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 全局配置
  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 消息提示
  const { message } = App.useApp();

  // 获取系统信息
  const [systemInformation, setSystemInformation] = useState({});
  useEffect(() => {
    const getSystemInformationHandler = async () => {
      try {
        const res = await AxiosGet(Apis.Public.Information);
        if (res.code === 200) {
          setSystemInformation(res.data);
        } else {
          message.error(res.message);
        }
      } catch (error) {
        console.error(error);
        message.error('获取系统信息失败');
      }
    };
    getSystemInformationHandler();
  }, []);

  const systemBasicInformation = [
    {
      key: '0',
      label: '系统名称',
      children: systemInformation?.systemProjectName
    },
    {
      key: '1',
      label: '开发者信息',
      children: `${systemInformation?.systemDeveloperName} <${systemInformation?.systemDeveloperEmail}>`
    },
    {
      key: '2',
      label: 'Go 版本',
      children: systemInformation?.systemGoVersion
    },
    {
      key: '3',
      label: '系统版本',
      children: systemInformation?.systemVersion
    },
    {
      key: '4',
      label: '项目地址',
      children: (
        <a href="https://github.com/ezopscn/stellar" target="_blank">
          <GithubOutlined /> https://github.com/ezopscn/stellar
        </a>
      )
    }
  ];

  // 获取节点信息
  const [nodeInformation, setNodeInformation] = useState([]);
  useEffect(() => {
    const getNodeInformationHandler = async () => {
      try {
        const res = await AxiosGet(Apis.Public.NodeInformation);
        if (res.code === 200) {
          setNodeInformation(res.data);
        } else {
          message.error(res.message);
        }
      } catch (error) {
        console.error(error);
        message.error('获取节点信息失败');
      }
    };
    getNodeInformationHandler();
  }, []);

  // 节点信息表格列
  const columns = [
    {
      title: '节点名称',
      dataIndex: 'name',
      key: 'name',
      width: '300px'
    },
    {
      title: '角色列表',
      dataIndex: 'roles',
      key: 'roles',
      render: (roles) => {
        return roles.map((role, idx) => {
          const text = NODE_ROLE_MAP[role]?.text;
          const color = NODE_ROLE_MAP[role]?.color;
          return text ? <Tag key={role + idx} color={color}>{text}</Tag> : null;
        });
      }
    }
  ];

  return (
    <>
      {/* 页面 header */}
      <Helmet>
        <title>{PageConfig.pageTitle}</title>
        <meta name="description" content={PageConfig.pageDesc} />
      </Helmet>
      {/* 页面头部介绍 */}
      <div className="admin-page-header">
        <div className="admin-page-title">{PageConfig.pageHeaderTitle}</div>
        <div className="admin-page-desc">
          <PageDescriptionComponent />
        </div>
      </div>
      {/* 页面主体 */}
      <div className="admin-page-main">
        {/* 版本构建信息 */}
        <Card className="admin-page-card">
          <Descriptions column={1} title="版本构建信息" items={systemBasicInformation} />
        </Card>

        {/* 所有节点列表 */}
        <div className="admin-table-item">
          <div className="admin-table-item-title">所有节点列表</div>
          <Table dataSource={nodeInformation} size="small" tableLayout="auto" pagination={false} columns={columns} rowKey="name" />
        </div>
      </div>
    </>
  );
};

export default SystemInformation;
