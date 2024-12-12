import { useState, useEffect } from 'react';
import { Helmet } from 'react-helmet';
import { Button, Col, Form, Input, Row, Space, Table, Avatar, Tag, Descriptions, TreeSelect, Select, Modal, App } from 'antd';
import { ClearOutlined, DownOutlined, SearchOutlined, UserAddOutlined, ManOutlined, WomanOutlined, QuestionOutlined } from '@ant-design/icons';
import { TitleSuffix } from '@/common/Text.jsx';
import { AxiosGet } from '@/utils/Request';
import { Apis } from '@/common/APIConfig';
import APIRequest from '@/common/APIRequest';

const User = () => {
  const { message } = App.useApp(); // 消息提示
  const [form] = Form.useForm(); // 表单
  const pageKeyword = '用户'; // 页面关键词
  const title = '用户管理' + TitleSuffix; // 页面标题
  const pageTitle = '用户管理 / USER MANAGEMENT.'; // 页面标题
  const pageDesc = '用户是系统的核心资产之一，也是许多其它资源的强制依赖，所以对于用户的管理，我提供了以下的要求和建议，请知悉：'; // 页面描述

  /////////////////////////////////////////////////////
  // 搜索栏
  /////////////////////////////////////////////////////
  const defaultExpandItemCount = 8; // 默认展开的搜索项数量
  const [expand, setExpand] = useState(false); // 是否展开更多搜索

  // 页面加载的时候一次性获取依赖的异步数据
  const [roleList, setRoleList] = useState([]); // 下拉菜单的角色列表数据
  const [jobPositionList, setJobPositionList] = useState([]); // 下拉菜单的岗位列表数据
  const [departmentList, setDepartmentList] = useState([]); // 下拉菜单的部门列表数据
  useEffect(() => {
    APIRequest.GetSelectDataList(Apis.System.Role.List, setRoleList);
    APIRequest.GetSelectDataList(Apis.System.JobPosition.List, setJobPositionList);
    APIRequest.GetSelectDataList(Apis.System.Department.List, setDepartmentList, true);
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
  // 默认设置
  const tableDataListAPI = Apis.System.User.List; // 数据列表接口
  const defaultPageSize = 2; // 默认每页显示的数据量
  const defaultPageNumber = 1; // 默认页码
  const defaultTotal = 0; // 默认数据总数
  const isPagination = true; // 是否需要分页

  // 表格：行选择
  const rowSelection = {
    onChange: (selectedRowKeys, selectedRows) => {
      console.log(`selectedRowKeys: ${selectedRowKeys}`, 'selectedRows: ', selectedRows);
    },
    getCheckboxProps: (record) => ({
      disabled: record.name === 'Disabled User',
      name: record.name
    })
  };

  // 表格：列
  const tableColumns = [
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
        const genderIcons = {
          1: <ManOutlined style={{ color: '#165dff' }} />,
          2: <WomanOutlined style={{ color: '#ff4d4f' }} />,
          default: <QuestionOutlined style={{ color: '#999' }} />
        };
        return genderIcons[gender] || genderIcons.default;
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
        const roleColors = {
          '超级管理员': 'magenta',
          '管理员': 'volcano',
          '运维': 'green'
        };
        const color = roleColors[name] || '';
        return (
          <Tag bordered={false} color={color}>
            {name}
          </Tag>
        );
      }
    },
    {
      title: '状态',
      dataIndex: 'status',
      render: (status) => {
        const statusMap = {
          1: { text: '启用', color: 'green' },
          0: { text: '禁用', color: 'red' }
        };
        const { text, color } = statusMap[status] || {};
        return (
          <Tag bordered={false} color={color}>
            {text}
          </Tag>
        );
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

  // 状态数据
  const [tableDataList, setTableDataList] = useState([]); // 数据列表
  const [pageSize, setPageSize] = useState(defaultPageSize); // 每页显示的条数
  const [pageNumber, setPageNumber] = useState(defaultPageNumber); // 当前页码
  const [total, setTotal] = useState(defaultTotal); // 数据总数
  const [filterParams, setFilterParams] = useState({}); // 条件搜索

  // 默认列表加载和监听所有页码变化事件（包括搜索）
  useEffect(() => {
    const params = { ...filterParams, pageSize, pageNumber, isPagination };
    filterDataList(params);
  }, [pageSize, pageNumber, filterParams]);

  // 条件搜索方法
  const filterDataListHandle = (data) => {
    // 初始化页码，避免在搜索结果溢出搜索数据的页码的时候，导致请求参数中带了页码，无法请求到数据
    setPageNumber(defaultPageNumber);
    setPageSize(defaultPageSize);
    setTotal(defaultTotal);
    setFilterParams(data);
  };

  // 条件搜索
  const filterDataList = async (params) => {
    try {
      const res = await AxiosGet(tableDataListAPI, params);
      if (res.code === 200) {
        setTableDataList(res.data.list);
        setPageSize(res.data.pagination.pageSize);
        setPageNumber(res.data.pagination.pageNumber);
        setTotal(res.data.pagination.total);
      } else {
        message.error(res.message);
      }
    } catch (error) {
      console.error('后端服务异常，获取表格数据列表失败');
      message.error(error);
    }
  };

  /////////////////////////////////////////////////////
  // 添加弹窗
  /////////////////////////////////////////////////////
  const [addModalVisible, setAddModalVisible] = useState(false);

  // 添加数据方法
  const addRecordHandle = () => {
    console.log('添加数据');
  };

  return (
    <>
      <Helmet>
        <title>{title}</title>
        <meta name="description" content={title} />
      </Helmet>
      {/* 页面头部介绍 */}
      <div className="admin-page-header">
        <div className="admin-page-title">{pageTitle}</div>
        <div className="admin-page-desc">
          <div>{pageDesc}</div>
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
          <Form form={form} labelCol={{ span: 6 }} wrapperCol={{ span: 18 }} colon={false} name="filterForm" onFinish={filterDataListHandle}>
            <Row gutter={24}>
              {getFilterFields()}
              <Col span={24} key="x" style={{ marginTop: '10px', textAlign: 'right' }}>
                <Space>
                  <Button icon={<SearchOutlined />} htmlType="submit">
                    条件搜索
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
                setAddModalVisible(true);
              }}
            >
              添加{pageKeyword}
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
            columns={tableColumns}
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
            dataSource={tableDataList}
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
        title={'添加' + pageKeyword}
        open={addModalVisible}
        onOk={addRecordHandle}
        onCancel={() => {
          setAddModalVisible(false);
        }}
        width={400}
        maskClosable={false}
        footer={[
          <Button key="submit" block type="primary" onClick={addRecordHandle}>
            添加{pageKeyword}
          </Button>
        ]}
      >
        <Form form={form} labelCol={{ span: 6 }} wrapperCol={{ span: 18 }} colon={false} name="addForm" onFinish={addRecordHandle}>
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
