import { useState, useEffect } from 'react';
import { Helmet } from 'react-helmet';
import { Button, Col, Form, Input, Row, Space, Table, App, Avatar, Tag, Descriptions, TreeSelect, Select, Modal, Radio } from 'antd';
import { ClearOutlined, DownOutlined, SearchOutlined, UserAddOutlined, ManOutlined, WomanOutlined, QuestionOutlined, SelectOutlined } from '@ant-design/icons';
import { TitleSuffix } from '@/common/Text.jsx';
import { AxiosGet } from '@/utils/Request';
import { Apis } from '@/common/APIConfig';
import { GenerateSelectTree } from '@/common/GenerateTree';

const User = () => {
  // 消息提示
  const { message } = App.useApp();
  // 表单
  const [form] = Form.useForm();

  /////////////////////////////////////////////////////
  // Page Header 信息
  /////////////////////////////////////////////////////
  const title = '用户中心' + TitleSuffix;

  /////////////////////////////////////////////////////
  // 搜索栏
  /////////////////////////////////////////////////////
  const defaultExpandItemCount = 8; // 默认展开的搜索项数量
  const [expand, setExpand] = useState(false); // 是否展开更多搜索
  const [roleList, setRoleList] = useState([]); // 下拉菜单的角色列表数据
  const [jobPositionList, setJobPositionList] = useState([]); // 下拉菜单的岗位列表数据
  const [departmentList, setDepartmentList] = useState([]); // 下拉菜单的部门列表数据

  // 页面加载的时候一次性获取依赖的异步数据
  useEffect(() => {
    // 通过 tree 参数来区分是否是树结构
    const fetchList = async (api, setter, tree = false) => {
      try {
        const res = await AxiosGet(api);
        if (res.code === 200) {
          if (tree) {
            const treeData = GenerateSelectTree(res.data, 0);
            setter(treeData);
          } else {
            setter(
              res.data.map((item) => ({
                label: item.name,
                value: item.id
              }))
            );
          }
        } else {
          message.error(res.message);
        }
      } catch (error) {
        console.log(`后端服务异常，获取接口：${api} 列表失败`, error);
      }
    };
    fetchList(Apis.System.Role.List, setRoleList);
    fetchList(Apis.System.JobPosition.List, setJobPositionList);
    fetchList(Apis.System.Department.List, setDepartmentList, true);
  }, []);

  // 定义筛选列表和数据限制，支持类型：input、select、checkbox
  const filterFields = [
    {
      label: '用户名',
      name: 'username',
      placeholder: '使用用户名进行搜索',
      type: 'input',
      rules: [
        {
          message: '用户名长度不能超过30个字符',
          max: 30
        }
      ]
    },
    {
      label: '姓名',
      name: 'name',
      placeholder: '使用中文名或者英文名进行搜索',
      type: 'input',
      rules: [
        {
          message: '姓名长度不能超过30个字符',
          max: 30
        }
      ]
    },
    {
      label: '邮箱',
      name: 'email',
      placeholder: '使用邮箱地址进行搜索',
      type: 'input',
      rules: [
        {
          message: '邮箱长度不能超过30个字符',
          max: 50
        }
      ]
    },
    {
      label: '手机号',
      name: 'phone',
      placeholder: '使用手机号码进行搜索',
      type: 'input',
      rules: [
        {
          message: '手机号长度不能超过15个字符',
          max: 15
        }
      ]
    },
    {
      label: '状态',
      name: 'status',
      placeholder: '选择用户状态进行搜索',
      type: 'select',
      search: false,
      tree: false,
      multiple: false,
      data: [
        {
          label: '启用',
          value: 1
        },
        {
          label: '禁用',
          value: 0
        }
      ],
      rules: []
    },
    {
      label: '性别',
      name: 'gender',
      placeholder: '选择性别进行搜索',
      type: 'select',
      search: false,
      tree: false,
      multiple: false,
      data: [
        {
          label: '男',
          value: 1
        },
        {
          label: '女',
          value: 2
        },
        {
          label: '未知',
          value: 3
        }
      ],
      rules: []
    },
    {
      label: '部门',
      name: 'department',
      placeholder: '选择部门进行搜索',
      type: 'select',
      search: true,
      tree: true,
      multiple: false,
      data: departmentList,
      rules: []
    },
    {
      label: '岗位',
      name: 'jobPosition',
      placeholder: '选择岗位进行搜索',
      type: 'select',
      search: true,
      tree: false,
      multiple: false,
      data: jobPositionList,
      rules: []
    },
    {
      label: '角色',
      name: 'role',
      placeholder: '选择角色进行搜索',
      type: 'select',
      search: true,
      tree: false,
      multiple: false,
      data: roleList,
      rules: []
    }
  ];

  // 获取搜索栏字段
  const getFilterFields = () => {
    const children = []; // 子元素

    // 生成搜索表单
    filterFields.slice(0, expand ? filterFields.length : defaultExpandItemCount).forEach((field, index) => {
      children.push(
        <Col span={6} key={field.label}>
          <Form.Item name={field.name} label={field.label} rules={field.rules}>
            {field.type === 'input' ? (
              <Input placeholder={field.placeholder} allowClear={true} autoComplete="off" />
            ) : field.type === 'select' ? (
              field.tree ? (
                <TreeSelect
                  multiple={field.multiple}
                  placeholder={field.placeholder}
                  treeData={field.data}
                  showSearch={field.search}
                  treeNodeFilterProp="label"
                  allowClear={true}
                  treeDefaultExpandAll
                />
              ) : (
                <Select mode={field.multiple ? 'multiple' : 'default'} placeholder={field.placeholder} options={field.data} showSearch={field.search} optionFilterProp="label" allowClear={true} />
              )
            ) : field.type === 'checkbox' ? (
              <Checkbox.Group options={field.data} />
            ) : null}
          </Form.Item>
        </Col>
      );
    });
    return children;
  };

  /////////////////////////////////////////////////////
  // 表格数据
  /////////////////////////////////////////////////////
  // 表格列
  const columns = [
    {
      title: '头像',
      dataIndex: 'avatar',
      render: (avatar) => <Avatar src={avatar} />
    },
    {
      title: '中文名',
      dataIndex: 'cnName',
      minWidth: 80
    },
    {
      title: '英文名',
      dataIndex: 'enName',
      minWidth: 80
    },
    {
      title: '性别',
      dataIndex: 'gender',
      minWidth: 50,
      render: (gender) => {
        if (gender === 1) {
          return <ManOutlined style={{ color: '#165dff' }} />;
        } else if (gender === 2) {
          return <WomanOutlined style={{ color: '#ff4d4f' }} />;
        } else {
          return <QuestionOutlined style={{ color: '#999' }} />;
        }
      }
    },
    {
      title: '用户名',
      dataIndex: 'username',
      minWidth: 80
    },
    {
      title: '邮箱',
      dataIndex: 'email'
    },
    {
      title: '手机号',
      dataIndex: 'phone'
    },
    {
      title: '部门',
      dataIndex: 'systemDepartments',
      minWidth: 100,
      render: (systemDepartments) => {
        return systemDepartments?.map((department) => department.name).join(',');
      }
    },
    {
      title: '岗位',
      dataIndex: 'systemJobPositions',
      minWidth: 120,
      render: (systemJobPositions) => {
        return systemJobPositions?.map((jobPosition) => jobPosition.name).join(',');
      }
    },
    {
      title: '角色名称',
      dataIndex: ['systemRole', 'name'],
      render: (name) => {
        if (name === '超级管理员') {
          return (
            <Tag bordered={false} color="magenta">
              {name}
            </Tag>
          );
        } else if (name === '管理员') {
          return (
            <Tag bordered={false} color="volcano">
              {name}
            </Tag>
          );
        } else if (name === '运维') {
          return (
            <Tag bordered={false} color="green">
              {name}
            </Tag>
          );
        } else {
          return <Tag bordered={false}>{name}</Tag>;
        }
      }
    },
    {
      title: '状态',
      dataIndex: 'status',
      render: (status) => {
        if (status === 1) {
          return (
            <Tag bordered={false} color="green">
              启用
            </Tag>
          );
        } else {
          return (
            <Tag bordered={false} color="red">
              禁用
            </Tag>
          );
        }
      }
    },
    {
      title: '操作',
      key: 'action',
      fixed: 'right',
      render: (_, record) => (
        <Space size="middle">
          <a>编辑</a>
          {record.status === 1 ? <a style={{ color: '#ff4d4f' }}>禁用</a> : <a>启用</a>}
        </Space>
      )
    }
  ];

  // 表格行选择
  const rowSelection = {
    onChange: (selectedRowKeys, selectedRows) => {
      console.log(`selectedRowKeys: ${selectedRowKeys}`, 'selectedRows: ', selectedRows);
    },
    getCheckboxProps: (record) => ({
      disabled: record.name === 'Disabled User',
      name: record.name
    })
  };

  // 获取用户列表
  const [userList, setUserList] = useState([]);
  // 默认每页显示的数据量
  const defaultPageSize = 2;
  // 默认页码
  const defaultPageNumber = 1;
  // 是否需要分页
  const isPagination = true;
  // 默认数据总数
  const defaultTotal = 0;
  // 每页显示的条数
  const [pageSize, setPageSize] = useState(defaultPageSize);
  // 当前页码
  const [pageNumber, setPageNumber] = useState(defaultPageNumber);
  // 数据总数
  const [total, setTotal] = useState(defaultTotal);
  // 条件搜索用户
  const [userFilterParams, setUserFilterParams] = useState({});
  // 默认用户列表加载和监听所有页码变化事件（包括搜索）
  useEffect(() => {
    const params = { ...userFilterParams, pageSize, pageNumber, isPagination };
    filterUserList(params);
  }, [pageSize, pageNumber, userFilterParams]);

  // 条件搜索用户方法
  const filterUserListHandle = (data) => {
    // 初始化页码，避免在搜索结果溢出搜索数据的页码的时候，导致请求参数中带了页码，无法请求到数据
    setPageNumber(defaultPageNumber);
    setPageSize(defaultPageSize);
    setTotal(defaultTotal);
    setUserFilterParams(data);
  };

  // 条件搜索用户
  const filterUserList = async (params) => {
    console.log('查询条件: ', params);
    try {
      const res = await AxiosGet(Apis.System.User.List, params);
      if (res.code === 200) {
        setUserList(res.data.list);
        setPageSize(res.data.pagination.pageSize);
        setPageNumber(res.data.pagination.pageNumber);
        setTotal(res.data.pagination.total);
      } else {
        message.error(res.message);
      }
    } catch (error) {
      console.error('后端服务异常，获取用户列表失败', error);
    }
  };

  /////////////////////////////////////////////////////
  // 用户添加弹窗
  /////////////////////////////////////////////////////
  const [userModalVisible, setUserModalVisible] = useState(false);

  // 添加用户
  const addUserHandle = () => {
    console.log('添加用户');
  };

  return (
    <>
      <Helmet>
        <title>{title}</title>
        <meta name="description" content={title} />
      </Helmet>
      {/* 页面头部介绍 */}
      <div className="admin-page-header">
        <div className="admin-page-title">用户中心 / USER MANAGEMENT.</div>
        <div className="admin-page-desc">
          <div>用户是系统的核心资产之一，也是许多其它资源的强制依赖，所以对于用户的管理，我提供了以下的要求和建议，请知悉：</div>
          <ul>
            <li>出于数据安全考虑，系统将强制使用禁用用户来替代删除用户，以此保证数据的可靠性和稳定性。</li>
            <li>针对某些特殊的用户，例如老板、高管等，我们建议隐藏其联系方式，保护个人隐私。</li>
          </ul>
        </div>
      </div>
      {/* 页面主体 */}
      <div className="admin-page-main">
        {/* 搜索栏 */}
        <div className="admin-page-search">
          <Form form={form} labelCol={{ span: 6 }} wrapperCol={{ span: 18 }} colon={false} name="userFilterForm" onFinish={filterUserListHandle}>
            <Row gutter={24}>
              {getFilterFields()}
              <Col span={24} key="x" style={{ marginTop: '10px', textAlign: 'right' }}>
                <Space>
                  <Button icon={<SearchOutlined />} htmlType="submit">
                    搜索用户
                  </Button>
                  <Button
                    icon={<ClearOutlined />}
                    onClick={() => {
                      form.resetFields();
                    }}
                  >
                    清理条件
                  </Button>
                  <a
                    onClick={() => {
                      setExpand(!expand);
                    }}
                  >
                    <DownOutlined rotate={expand ? 180 : 0} /> {expand ? '收起条件' : '展开更多'}
                  </a>
                </Space>
              </Col>
            </Row>
          </Form>
        </div>
        {/* 表格 */}
        <div className="admin-page-list">
          <div className="admin-page-btn-group">
            <Button
              icon={<UserAddOutlined />}
              onClick={() => {
                setUserModalVisible(true);
              }}
            >
              添加用户
            </Button>
          </div>
          <Table
            // 表格布局大小
            size="small"
            // 表格布局方式，支持 fixed、auto
            tableLayout="auto"
            // 表格行选择
            rowSelection={{
              type: 'checkbox',
              ...rowSelection
            }}
            // 表格列
            columns={columns}
            // 表格展开信息
            expandable={{
              expandedRowRender: (record) => {
                const items = [
                  {
                    label: '用户创建者',
                    children: record.creator.split(',')[0] + ' / ' + record.creator.split(',')[1]
                  },
                  {
                    label: '个人介绍',
                    children: record.description
                  },
                  {
                    label: '最后登录 IP',
                    children: record.lastLoginIP
                  },
                  {
                    label: '最后登录时间',
                    children: record.lastLoginTime
                  }
                ];
                return <Descriptions column={1} items={items} />;
              },
              rowExpandable: (record) => record.name !== 'Not Expandable'
            }}
            dataSource={userList}
            // 行唯一标识
            rowKey="id"
            // 表格分页
            pagination={{
              pageSize: pageSize,
              current: pageNumber,
              total: total,
              showTotal: (total) => `总共 ${total} 条记录`,
              // hideOnSinglePage: true,
              showSizeChanger: true,
              showQuickJumper: true,
              onChange: (page, pageSize) => {
                setPageNumber(page);
                setPageSize(pageSize);
              }
            }}
            // 表格滚动，目的是为了最后一列固定
            scroll={{
              x: 'max-content'
            }}
          />
        </div>
      </div>

      {/* 用户添加弹窗 */}
      <Modal
        title="添加用户"
        open={userModalVisible}
        onOk={addUserHandle}
        onCancel={() => {
          setUserModalVisible(false);
        }}
        width={400}
        maskClosable={false}
        footer={[
          <Button key="submit" block type="primary" onClick={addUserHandle}>
            添加
          </Button>
        ]}
      >
        <Form form={form} labelCol={{ span: 6 }} wrapperCol={{ span: 18 }} colon={false} name="userAddForm" onFinish={addUserHandle}>
          <Form.Item
            name="username"
            label="用户名"
            rules={[
              {
                required: true,
                message: '用户名不能为空'
              },
              {
                pattern: /^[a-zA-Z][a-zA-Z0-9_]*$/,
                message: '用户名只能以字母开头，且只能包含字母、数字和下划线'
              },
              {
                max: 30,
                message: '用户名长度不能超过30个字符'
              },
              {
                min: 3,
                message: '用户名长度不能小于3个字符'
              }
            ]}
          >
            <Input />
          </Form.Item>
          <Form.Item
            name="password"
            label="密码"
            rules={[
              {
                required: true,
                message: '密码不能为空'
              },
              {
                max: 30,
                message: '长度不能超过30个字符'
              },
              {
                min: 8,
                message: '长度不能小于8个字符'
              },
              {
                pattern: /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{6,}$/,
                message: '必须包含至少一个字母、数字和特殊字符'
              }
            ]}
          >
            <Input type="password" />
          </Form.Item>
          <Form.Item
            name="repassword"
            label="确认密码"
            dependencies={['password']}
            rules={[
              {
                required: true,
                message: '确认密码不能为空'
              },
              {
                max: 30,
                message: '长度不能超过30个字符'
              },
              {
                min: 8,
                message: '长度不能小于8个字符'
              },
              {
                pattern: /^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{6,}$/,
                message: '必须包含至少一个字母、数字和特殊字符'
              }
            ]}
          >
            <Input type="password" />
          </Form.Item>
          <Form.Item
            name="cnName"
            label="中文名"
            rules={[
              {
                required: true,
                message: '中文名不能为空'
              },
              {
                pattern: /^[^\d\W]+$/,
                message: '中文名只能包含汉字'
              },
              {
                max: 30,
                message: '中文名长度不能超过30个字符'
              },
              {
                min: 2,
                message: '中文名长度不能小于2个字符'
              }
            ]}
          >
            <Input />
          </Form.Item>
          <Form.Item
            name="enName"
            label="英文名"
            rules={[
              {
                required: true,
                message: '英文名不能为空'
              },
              {
                pattern: /^[a-zA-Z]+$/,
                message: '英文名只能包含字母'
              },
              {
                max: 30,
                message: '英文名长度不能超过30个字符'
              },
              {
                min: 2,
                message: '英文名长度不能小于2个字符'
              }
            ]}
          >
            <Input />
          </Form.Item>
          <Form.Item
            name="email"
            label="邮箱"
            rules={[
              {
                type: 'email',
                message: '邮箱格式不正确'
              },
              {
                required: true,
                message: '邮箱不能为空'
              }
            ]}
          >
            <Input />
          </Form.Item>
          <Form.Item
            name="gender"
            label="性别"
            rules={[
              {
                required: true,
                message: '性别不能为空'
              }
            ]}
          >
            <Select>
              <Select.Option value={1}>男</Select.Option>
              <Select.Option value={2}>女</Select.Option>
              <Select.Option value={3}>未知</Select.Option>
            </Select>
          </Form.Item>
          <Form.Item
            name="department"
            label="部门"
            rules={[
              {
                required: true,
                message: '部门不能为空'
              }
            ]}
          >
            <TreeSelect multiple placeholder="选择部门" treeData={departmentList} showSearch treeNodeFilterProp="label" allowClear={true} treeDefaultExpandAll />
          </Form.Item>
          <Form.Item
            name="jobPosition"
            label="岗位"
            rules={[
              {
                required: true,
                message: '岗位不能为空'
              }
            ]}
          >
            <Select mode="multiple" placeholder="选择岗位" options={jobPositionList} showSearch optionFilterProp="label" allowClear={true} />
          </Form.Item>
        </Form>
      </Modal>
    </>
  );
};

export default User;
