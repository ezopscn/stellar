import { useState, useEffect } from 'react';
import { Helmet } from 'react-helmet';
import { TitleSuffix } from '@/common/Text.jsx';
import { App } from 'antd';
import { Card, Row, Col, Tree, Button, Alert, Form, Space, List, Avatar, Modal, Popconfirm } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import { AxiosGet } from '@/utils/Request.jsx';
import { Apis } from '@/common/APIConfig.jsx';
import { GenerateTreeNode } from '@/utils/Tree.jsx';
import { GenerateRecordFormItem } from '@/utils/Form.jsx';
import { ConvertNameIdToLabelValueTree, ConvertNameIdToTitleKeyTree, GetExpandedAllTreeKeys, HasChildren } from '@/utils/Tree.jsx';
import { AxiosPost } from '@/utils/Request.jsx';
import { useSnapshot } from 'valtio';
import { SystemRoleStates } from '@/store/StoreSystemRole.jsx';
import { BackendURL } from '@/common/APIConfig.jsx';

// 页面常量配置
const PageConfig = {
  pageTitle: '部门管理' + TitleSuffix,
  pageHeaderTitle: '部门管理 / DEPARTMENT MANAGEMENT.',
  pageKeyword: '部门管理'
};

// 页面说明组件
const PageDescriptionComponent = () => (
  <>
    <div>由于该系统的特殊性，并不会跟普通的人力资源管理系统一样，所以并不会同步公司所有的组织架构，而是需要手动创建部门。</div>
    <ul>
      <li>部门名称没有唯一性限制，因为不同部门或者项目组下面可能存在相同的部门。</li>
      <li>当删除部门时，部门下面的还存在用户，那么该用户会被移动到未分配部门。</li>
      <li>部门 ID 为 1 和 2 是系统保留的，其中 1 是公司组织架构主体，管理员可以修改它的名称，2 是系统需要特殊预留的未分配部门，不允许调整和创建子部门。</li>
    </ul>
  </>
);

// 编辑表单字段配置
const getUpdateSystemDepartmentFields = (updateSystemDepartment, systemDepartmentSelectTreeData) => [
  { label: '部门 ID', name: 'id', type: 'input', disabled: true },
  { label: '部门名称', name: 'name', type: 'input', rules: [
    { required: true, message: '部门名称不能为空' },
    { max: 30, message: '部门名称长度不能超过30个字符' },
    { min: 3, message: '部门名称长度不能小于3个字符' }
  ],
    disabled: updateSystemDepartment?.id === 2
  },
  { label: '父部门', name: 'parentId', type: 'treeSelect', search: true, tree: true, multiple: false, data: systemDepartmentSelectTreeData, 
    rules: [{ required: true, message: '父部门不能为空' }],
    disabled: updateSystemDepartment?.id === 1 || updateSystemDepartment?.id === 2
  },
  { label: '创建人', name: 'creator', type: 'input', disabled: true },
  { label: '创建时间', name: 'createdAt', type: 'input', disabled: true },
  { label: '更新时间', name: 'updatedAt', type: 'input', disabled: true }
];

const SystemDepartment = () => {
  const { message } = App.useApp();
  // 全局数据，用于父子组件之间数据传递
  const { SystemRoleApis } = useSnapshot(SystemRoleStates);

  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 按钮权限控制
  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 新增按钮权限控制
  const addSystemDepartmentButtonDisabled = !SystemRoleApis.list?.includes(Apis.System.Department.Add.replace(BackendURL, ''));

  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 基础数据
  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 左侧部门树数据
  const [systemDepartmentTreeData, setSystemDepartmentTreeData] = useState([]);
  // 展开的部门树节点
  const [expandedSystemDepartmentKeys, setExpandedSystemDepartmentKeys] = useState([]);
  // 部门选择树数据（下拉选择）
  const [systemDepartmentSelectTreeData, setSystemDepartmentSelectTreeData] = useState([]);

  ///////////////////////////////////////////////////////////////////////////////////////////////////// 
  // 新增部门
  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 新增部门弹窗
  const [addSystemDepartmentModalVisible, setAddSystemDepartmentModalVisible] = useState(false);

  // 新增部门表单
  const [addSystemDepartmentForm] = Form.useForm();

  // 新增部门字段
  const addSystemDepartmentFields = [
    { label: '父部门', name: 'parentId', type: 'treeSelect', search: true, tree: true, multiple: false, data: systemDepartmentSelectTreeData, rules: [{ required: true, message: '父部门不能为空' }] },
    { label: '部门名称', name: 'name', type: 'input', rules: [{ required: true, message: '部门名称不能为空' }, { max: 30, message: '部门名称长度不能超过30个字符' }, { min: 3, message: '部门名称长度不能小于3个字符' }] }
  ];

  // 生成新增部门表单项
  const generateAddSystemDepartmentFormItems = () => {
    return addSystemDepartmentFields?.map((item) => GenerateRecordFormItem(item));
  };

  // 新增部门
  const addSystemDepartmentHandler = async () => {
    try {
      const res = await AxiosPost(Apis.System.Department.Add, addSystemDepartmentForm.getFieldsValue());
      if (res?.code === 200) {
        message.success('新增部门成功');
        setAddSystemDepartmentModalVisible(false);
        getSystemDepartmentDataHandler();
        addSystemDepartmentForm.resetFields();
      } else {
        message.error(res?.message || '新增部门失败');
      }
    } catch (error) {
      console.error(error);
      message.error('服务异常，新增部门失败');
    }
  };

  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 编辑部门详情
  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 编辑部门表单
  const [updateSystemDepartmentForm] = Form.useForm();
  // 当前编辑的部门数据
  const [updateSystemDepartment, setUpdateSystemDepartment] = useState(null);
  // 删除部门按钮状态
  const [deleteSystemDepartmentButtonDisabled, setDeleteSystemDepartmentButtonDisabled] = useState(false);
  // 保存部门按钮状态
  const [saveSystemDepartmentButtonDisabled, setSaveSystemDepartmentButtonDisabled] = useState(false);

  // 获取部门列表数据
  const getSystemDepartmentDataHandler = async () => {
    try {
      const res = await AxiosGet(Apis.System.Department.List);
      if (res?.code === 200 && Array.isArray(res?.data)) {
        const treeData = GenerateTreeNode(res?.data, 0);
        setSystemDepartmentTreeData(ConvertNameIdToTitleKeyTree(treeData));
        setExpandedSystemDepartmentKeys(GetExpandedAllTreeKeys(treeData));
        setSystemDepartmentSelectTreeData(ConvertNameIdToLabelValueTree(treeData));
      } else {
        message.error(res?.message || '获取部门数据失败');
      }
    } catch (error) {
      console.error('获取部门列表错误:', error);
      message.error('服务异常，获取部门列表失败');
    }
  };

  // 获取部门详情
  const getSystemDepartmentDetailHandler = async (selectedKeys) => {
    if (!selectedKeys?.length) return;

    try {
      const res = await AxiosGet(Apis.System.Department.Detail, { id: selectedKeys[0] });
      if (res?.code === 200) {
        const { data } = res;
        setUpdateSystemDepartment(data);

        // 设置表单值
        const [creatorUsername, creatorCnName, creatorEnName, creatorEmail] = data.creator.split(',');
        updateSystemDepartmentForm.setFieldsValue({
          ...data,
          creator: `${creatorCnName} / ${creatorEnName} (${creatorUsername} / ${creatorEmail})`
        });

        // 是否是根部门
        const isRootDepartment = data.id === 1;
        // 是否是未分配部门
        const isUnassignedDepartment = data.id === 2;
        // 是否存在子部门
        const hasChildDepartments = HasChildren(systemDepartmentTreeData, data.id);
        // 设置按钮状态
        setDeleteSystemDepartmentButtonDisabled(isRootDepartment || isUnassignedDepartment || hasChildDepartments);
        setSaveSystemDepartmentButtonDisabled(isUnassignedDepartment);
      } else {
        message.error(res?.message || '获取部门详情失败');
      }
    } catch (error) {
      console.error(error);
      message.error('服务异常，获取部门详情失败');
    }
  };

  // 默认选中根部门
  useEffect(() => {
    if (systemDepartmentTreeData.length > 0) {
      getSystemDepartmentDetailHandler(['1']);
    }
  }, [systemDepartmentTreeData]);

  // 初始化数据
  useEffect(() => {
    getSystemDepartmentDataHandler();
  }, []);



  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 用户部门变更
  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 用户部门变更弹窗
  const [userSystemDepartmentChangeModalVisible, setUserSystemDepartmentChangeModalVisible] = useState(false);

  // 用户部门变更表单
  const [userSystemDepartmentChangeForm] = Form.useForm();

  // 用户部门变更字段
  const userSystemDepartmentChangeFields = [
    { label: '用户', name: 'userId', type: 'input', disabled: true },
    { label: '部门', name: 'departmentId', type: 'treeSelect', search: true, tree: true, multiple: true, data: systemDepartmentSelectTreeData, 
      rules: [{ required: true, message: '部门不能为空' }] 
    }
  ];

  // 生成用户部门变更表单项
  const generateUserSystemDepartmentChangeFormItems = () => {
    return userSystemDepartmentChangeFields?.map((item) => GenerateRecordFormItem(item));
  };

  // 用户部门变更
  const userSystemDepartmentChangeHandler = async () => {
    try {
      const res = await AxiosPost(Apis.System.Department.UpdateUserDepartment, userSystemDepartmentChangeForm.getFieldsValue());
      if (res?.code === 200) {
        message.success('用户部门变更成功');
        setUserSystemDepartmentChangeModalVisible(false);
        userSystemDepartmentChangeForm.resetFields();
        // 重新获取部门详情
        getSystemDepartmentDetailHandler(updateSystemDepartment?.id);
      } else {
        message.error(res?.message || '用户部门变更失败');
      }
    } catch (error) {
      console.error(error);
      message.error('服务异常，用户部门变更失败');
    }
  };

  return (
    <>
      <Helmet>
        <title>{PageConfig.pageTitle}</title>
        <meta name="description" content={PageConfig.pageDesc} />
      </Helmet>

      <div className="admin-page-header">
        <div className="admin-page-title">{PageConfig.pageHeaderTitle}</div>
        <div className="admin-page-desc">
          <PageDescriptionComponent />
        </div>
      </div>

      <div className="admin-page-main">
        <Row>
          <Col span={6} style={{ padding: '10px' }}>
            <Card title="部门列表">
              <Button type="primary" block style={{ marginBottom: '15px' }} icon={<PlusOutlined />} 
                onClick={() => setAddSystemDepartmentModalVisible(true)}
                disabled={addSystemDepartmentButtonDisabled}
              >新增部门</Button>
              <Tree
                defaultExpandAll={true}
                defaultSelectedKeys={['1']}
                showLine={true}
                expandedKeys={expandedSystemDepartmentKeys}
                onExpand={setExpandedSystemDepartmentKeys}
                treeData={systemDepartmentTreeData}
                onSelect={(selectedKeys) => {
                  getSystemDepartmentDetailHandler(selectedKeys);
                }}
              />
            </Card>
          </Col>

          <Col span={6} style={{ padding: '10px 0' }}>
            <Card title="部门用户">
              <List
                className="admin-user-list"
                itemLayout="horizontal"
                dataSource={updateSystemDepartment?.systemUsers}
                split={false}
                renderItem={(item, index) => (
                  <List.Item actions={[
                    <Popconfirm
                      placement="topRight"
                      title="确定要移除该用户吗？"
                      okText="确定"
                      cancelText="取消"
                      okButtonProps={{ style: { backgroundColor: '#ff4d4f', borderColor: '#ff4d4f' } }}
                      onConfirm={() => {}}
                    >
                      <a style={{color: "#ff4d4f"}}>移除</a>
                    </Popconfirm>
                  , <a onClick={() => setUserSystemDepartmentChangeModalVisible(true)}>变更</a>]}>
                    <List.Item.Meta
                      avatar={<Avatar src={item.avatar} />}
                      title={`${item.cnName}（${item.enName}）`}
                      description={item.systemJobPositions?.map(job => `${job.name}`).join(', ')}
                    />
                  </List.Item>
                )}
                pagination={{
                  position: 'bottom',
                  align: 'end',
                  pageSize: 1,
                  size: 'small',
                  hideOnSinglePage: true,
                  total: updateSystemDepartment?.systemUsers?.length,
                  showSizeChanger: true,
                  showQuickJumper: true,
                }}
              />
            </Card>
          </Col>

          <Col span={12} style={{ padding: '10px' }}>
            <Card title="部门详情">
              <Alert message="从菜单树列表任意选择一项后，即可进行编辑修改。" type="info" />
              <div className="admin-tree-edit-form">
                <Form
                  form={updateSystemDepartmentForm}
                  labelCol={{ span: 6 }}
                  wrapperCol={{ span: 18 }}
                  colon={false}
                  autoComplete="off"
                  name="updateSystemDepartmentForm"
                  onFinish={() => {}}
                >
                  {getUpdateSystemDepartmentFields(updateSystemDepartment, systemDepartmentSelectTreeData)
                    .map(item => GenerateRecordFormItem(item))}
                  <Form.Item wrapperCol={{ span: 24 }} style={{ textAlign: 'right' }}>
                    <Row>
                      <Col offset={6} span={18}>
                        <Space>
                          <Button color="danger" variant="outlined" disabled={deleteSystemDepartmentButtonDisabled}>删除部门</Button>
                          <Button type="primary" htmlType="submit" disabled={saveSystemDepartmentButtonDisabled}>保存编辑</Button>
                        </Space>
                      </Col>
                    </Row>
                  </Form.Item>
                </Form>
              </div>
            </Card>
          </Col>
        </Row>
      </div>

      {/* 添加部门 */}
      <Modal
        title="新增部门"
        open={addSystemDepartmentModalVisible}
        onCancel={() => {
          setAddSystemDepartmentModalVisible(false);
          addSystemDepartmentForm.resetFields();
        }}
        maskClosable={false}
        onOk={addSystemDepartmentHandler}
      >
        <Form form={addSystemDepartmentForm} labelCol={{ span: 6 }} wrapperCol={{ span: 18 }} colon={false} name="addSystemDepartmentForm" autoComplete="off">
          {generateAddSystemDepartmentFormItems()}
        </Form>
      </Modal>

      {/* 用户部门变更 */}
      <Modal
        title="用户部门变更"
        open={userSystemDepartmentChangeModalVisible}
        onCancel={() => {
          setUserSystemDepartmentChangeModalVisible(false);
          userSystemDepartmentChangeForm.resetFields();
        }}
        maskClosable={false}
        onOk={userSystemDepartmentChangeHandler}
      >
        <Form form={userSystemDepartmentChangeForm} labelCol={{ span: 6 }} wrapperCol={{ span: 18 }} colon={false} name="userSystemDepartmentChangeForm" autoComplete="off">
          {generateUserSystemDepartmentChangeFormItems()}
        </Form>
      </Modal>
    </>
  );
};

export default SystemDepartment;
