import { useState, useEffect } from 'react';
import { Helmet } from 'react-helmet';
import { TitleSuffix } from '@/common/Text.jsx';
import { App } from 'antd';
import { Card, Row, Col, Tree, Button, Alert } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import { AxiosGet } from '@/utils/Request.jsx';
import { Apis } from '@/common/APIConfig.jsx';
import { GenerateDepartmentTree } from '@/utils/Tree.jsx';

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

  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 部门列表
  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 部门树数据
  const [departmentTreeData, setDepartmentTreeData] = useState([]);
  // 展开的部门节点
  const [expandedDepartmentKeys, setExpandedDepartmentKeys] = useState([]);

  // 获取树形结构所有节点的 key
  const getAllTreeKeys = (data) => {
    let keys = [];
    if (!Array.isArray(data)) return keys;
    data.forEach(item => {
      if (item?.key) keys.push(item.key);
      if (item?.children?.length > 0) {
        keys = keys.concat(getAllTreeKeys(item.children));
      }
    });
    return keys;
  };

  // 获取部门树数据
  const getDepartmentTreeData = async () => {
    try {
      const res = await AxiosGet(Apis.System.Department.List);
      if (res?.code === 200 && Array.isArray(res.data)) {
        const treeData = GenerateDepartmentTree(res.data, 0);
        setDepartmentTreeData(treeData);
        setExpandedDepartmentKeys(getAllTreeKeys(treeData));
      } else {
        message.error(res?.message || '获取部门数据失败');
      }
    } catch (error) {
      console.error('获取部门列表错误:', error);
      message.error('服务异常，获取部门列表失败');
    }
  };
  useEffect(() => {
    getDepartmentTreeData();
  }, []);

  // 处理展开/收起操作
  const onExpandDepartmentTree = (newExpandedKeys) => {
    setExpandedDepartmentKeys(newExpandedKeys);
  };

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
              <Button type="primary" block style={{ marginBottom: '15px' }} icon={<PlusOutlined />}>新增部门</Button>
              <Tree
                defaultExpandAll={true}
                showLine={true}
                expandedKeys={expandedDepartmentKeys}
                onExpand={onExpandDepartmentTree}
                treeData={departmentTreeData}
              />
            </Card>
          </Col>
          <Col span={18} style={{ padding: '10px', paddingLeft: '0' }}>
            <Card title="部门详情">
              <Alert message="从菜单树列表任意选择一项后，即可进行编辑修改。" type="error" />
            </Card>
          </Col>
        </Row>
      </div>
    </>
  );
};

export default SystemDepartment;
