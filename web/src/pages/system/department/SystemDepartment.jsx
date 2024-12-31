import { Helmet } from 'react-helmet';
import { TitleSuffix } from '@/common/Text.jsx';
import { App } from 'antd';
import { Card, Row, Col, Tree } from 'antd';

// 页面常量设置
const PageConfig = {
  // 页面标题
  pageTitle: '部门管理' + TitleSuffix,
  // 页面顶部标题
  pageHeaderTitle: '部门管理 / DEPARTMENT MANAGEMENT.',
  // 页面关键词
  pageKeyword: '部门管理'
};

// 页面说明组件
const PageDescriptionComponent = () => {
  return (
    <>
      <div>由于该系统的特殊性，并不会跟普通的人力资源管理系统一样，所以并不会同步公司所有的组织架构，而是需要手动创建部门。</div>
      <ul>
        <li>部门名称没有唯一性限制，因为不同部门或者项目组下面可能存在相同的部门。</li>
        <li>部门领导非必须，但还是建议设置。</li>
      </ul>
    </>
  );
};

const SystemDepartment = () => {
  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 全局配置
  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 消息提示
  const { message } = App.useApp();

  // 部门数据
  const departmentTreeData = [
    {
      title: '某科技公司',
      key: '1',
      children: [
        {
          title: '研发中心',
          key: '1-1',
          children: [
            {
              title: '后台开发部',
              key: '1-1-1',
            },
            {
              title: '前端开发部',
              key: '1-1-2',
            },
            {
              title: '测试部',
              key: '1-1-3',
            },
            {
              title: '运维部',
              key: '1-1-4',
            },
          ],
        },
        {
          title: '产品中心',
          key: '1-2',
          children: [
            {
              title: '产品部',
              key: '1-2-1',
            },
          ],
        },
      ],
    },
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
        <Row>
          <Col span={6} style={{ padding: '10px' }}>
            <Card title="部门列表">
              <Tree
                defaultExpandAll
                showLine={true}
                treeData={departmentTreeData}
              />
            </Card>
          </Col>
          <Col span={18} style={{ padding: '10px', paddingLeft: '0' }}>
            <Card title="部门详情">
              2
            </Card>
          </Col>
        </Row>
      </div>
    </>
  );
};

export default SystemDepartment;
