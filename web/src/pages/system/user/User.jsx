import { useState, useEffect } from 'react';
import { Helmet } from 'react-helmet';
import { Button, Col, Form, Input, Row, Space, Table, App, Avatar, Tag, Descriptions, TreeSelect, Select } from 'antd';
import { ClearOutlined, DownOutlined, SearchOutlined, UserAddOutlined, ManOutlined, WomanOutlined, QuestionOutlined } from '@ant-design/icons';
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
  // 更多搜索
  const [expand, setExpand] = useState(false);

  // 获取角色、岗位、部门列表
  const [roleList, setRoleList] = useState([]);
  const [jobPositionList, setJobPositionList] = useState([]);
  const [departmentList, setDepartmentList] = useState([]);
  useEffect(() => {
    const fetchList = async (api, setter) => {
      try {
        const res = await AxiosGet(api);
        if (res.code === 200) {
          // 部门树结构需要单独处理
          if (api === Apis.System.Department.List) {
            const treeData = GenerateSelectTree(res.data, 0);
            setter(treeData);
          } else {
            setter(res.data.map(item => ({
              label: item.name,
              value: item.id,
            })));
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
    fetchList(Apis.System.Department.List, setDepartmentList);
  }, []);

  // 列表数据
  const filterFields = [
    {
      label: '用户名',
      name: 'username',
      placeholder: '使用用户名进行搜索',
      type: 'input',
      rules: [{
        message: '用户名长度不能超过30个字符',
        max: 30,
      }],
    },
    {
      label: '姓名',
      name: 'name',
      placeholder: '使用中文名，英文名进行搜索',
      type: 'input',
      rules: [{
        message: '姓名长度不能超过30个字符',
        max: 30,
      }],
    },
    {
      label: '邮箱',
      name: 'email',
      placeholder: '使用邮箱地址进行搜索',
      type: 'input',
      rules: [{
        message: '邮箱长度不能超过30个字符',
        max: 50,
      }],
    },
    {
      label: '手机号',
      name: 'phone',
      placeholder: '使用手机号码进行搜索',
      type: 'input',
      rules: [{
        message: '手机号长度不能超过15个字符',
        max: 15,
      }],
    },
    {
      label: '状态',
      name: 'status',
      placeholder: '选择用户状态进行搜索',
      type: 'select',
      data: [
        {
          label: '启用',
          value: 1,
        },
        {
          label: '禁用',
          value: 0,
        },
      ],
      rules: []
    },
    {
      label: '性别',
      name: 'gender',
      placeholder: '选择性别进行搜索',
      type: 'select',
      data: [
        {
          label: '男',
          value: 1,
        },
        {
          label: '女',
          value: 2,
        },
        {
          label: '未知',
          value: 3,
        },
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
      data: departmentList,
      rules: []
    },
    {
      label: '岗位',
      name: 'jobPosition',
      placeholder: '选择岗位进行搜索',
      type: 'select',
      search: true,
      data: jobPositionList,
      rules: []
    },
    {
      label: '角色',
      name: 'role',
      placeholder: '选择角色进行搜索',
      type: 'select',
      search: true,
      data: roleList,
      rules: []
    },
  ];

  // 获取搜索栏字段
  const getFilterFields = () => {
    const expandWidth = 7; // 展开宽度
    const children = [];   // 子元素

    // 生成搜索表单
    filterFields.slice(0, expand ? filterFields.length : expandWidth).forEach((field, index) => {
      children.push(
        <Col span={6} key={field.label}>
          <Form.Item
            name={field.name}
            label={field.label}
            rules={field.rules}
          >
            {field.type === "input" ? <Input placeholder={field.placeholder} autoComplete='off' /> : field.tree ?
              <TreeSelect placeholder={field.placeholder} treeData={field.data} showSearch={field.search} treeNodeFilterProp='label' treeDefaultExpandAll /> :
              <Select placeholder={field.placeholder} options={field.data} showSearch={field.search} optionFilterProp='label' />}
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
      minWidth: 80,
    },
    {
      title: '英文名',
      dataIndex: 'enName',
      minWidth: 80,
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
      minWidth: 80,
    },
    {
      title: '邮箱',
      dataIndex: 'email',
    },
    {
      title: '手机号',
      dataIndex: 'phone',
    },
    {
      title: '部门',
      dataIndex: 'systemDepartments',
      minWidth: 100,
      render: (systemDepartments) => {
        return systemDepartments?.map(department => department.name).join(',');
      }
    },
    {
      title: '岗位',
      dataIndex: 'systemJobPositions',
      minWidth: 120,
      render: (systemJobPositions) => {
        return systemJobPositions?.map(jobPosition => jobPosition.name).join(',');
      }
    },
    {
      title: '角色名称',
      dataIndex: ['systemRole', 'name'],
      render: (name) => {
        if (name === '超级管理员') {
          return <Tag bordered={false} color="magenta">{name}</Tag>;
        } else if (name === '管理员') {
          return <Tag bordered={false} color="volcano">{name}</Tag>;
        } else if (name === '运维') {
          return <Tag bordered={false} color='green'>{name}</Tag>;
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
          return <Tag bordered={false} color="green">启用</Tag>;
        } else {
          return <Tag bordered={false} color="red">禁用</Tag>;
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
      ),
    },
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
  useEffect(() => {
    const getUserList = async () => {
      try {
        const res = await AxiosGet(Apis.System.User.List);
        if (res.code === 200) {
          setUserList(res.data);
        } else {
          message.error(res.message);
        }
      } catch (error) {
        console.log('后端服务异常，获取用户列表失败', error);
      }
    };
    getUserList();
  }, []);

  // 条件搜索用户
  const filterUserList = async (userFilterParams) => {
    console.log('查询条件: ', userFilterParams);
    try {
      const res = await AxiosGet(Apis.System.User.List, userFilterParams);
      if (res.code === 200) {
        setUserList(res.data);
      } else {
        message.error(res.message);
      }
    } catch (error) {
      console.log('后端服务异常，获取用户列表失败', error);
    }
  };

  // 每页数据量
  const [pageSize, setPageSize] = useState(2);

  return (
    <>
      <Helmet>
        <title>{title}</title>
        <meta name='description' content={title} />
      </Helmet>
      <div className='admin-page-header'>
        <div className='admin-page-title'>用户中心 / USER MANAGEMENT.</div>
        <div className='admin-page-desc'>
          <div>用户是系统的核心资产之一，也是许多其它资源的强制依赖，所以对于用户的管理，我提供了以下的要求和建议，请知悉：</div>
          <ul>
            <li>出于数据安全考虑，系统将强制使用禁用用户来替代删除用户，以此保证数据的可靠性和稳定性。</li>
            <li>针对某些特殊的用户，例如老板、高管等，我们建议隐藏其联系方式，保护个人隐私。</li>
          </ul>
        </div>
      </div>
      <div className='admin-page-main'>
        <div className='admin-page-search'>
          <Form form={form} labelCol={{ span: 6 }} wrapperCol={{ span: 18 }} colon={false} name='userFilterForm' onFinish={filterUserList}>
            <Row gutter={24}>
              {getFilterFields()}
              <Col span={6} key='x' style={{ marginTop: '10px', textAlign: 'right' }}>
                <Space>
                  <Button icon={<SearchOutlined />} htmlType='submit'>搜索用户</Button>
                  <Button icon={<ClearOutlined />} onClick={() => {
                    form.resetFields();
                  }}>清理条件</Button>
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
        <div className='admin-page-list'>
          <div className='admin-page-btn-group'>
            <Button icon={<UserAddOutlined />}>添加用户</Button>
          </div>
          <Table
            // 表格布局大小
            size='small'
            // 表格布局方式，支持 fixed、auto
            tableLayout='auto'
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
                    children: record.creator.split(',')[0] + ' / ' + record.creator.split(',')[1],
                  },
                  {
                    label: '个人介绍',
                    children: record.description,
                  },
                  {
                    label: '最后登录 IP',
                    children: record.lastLoginIP,
                  },
                  {
                    label: '最后登录时间',
                    children: record.lastLoginTime,
                  },
                ]
                return <Descriptions column={1} items={items} />
              },
              rowExpandable: (record) => record.name !== 'Not Expandable'
            }}
            dataSource={userList}
            // 行唯一标识
            rowKey='id'
            // 表格分页
            pagination={{
              pageSize: pageSize,
              showSizeChanger: true,
              showQuickJumper: true,
              onChange: (page, pageSize) => {
                setPageSize(pageSize);
              }
            }}
            // 表格滚动，目的是为了最后一列固定
            scroll={{
              x: 'max-content',
            }}
          />
        </div>
      </div>
    </>
  );
};

export default User;