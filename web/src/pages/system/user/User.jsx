import { Button, Col, Form, Input, Row, Select, Space, Table, App } from 'antd';
import { ClearOutlined, DownOutlined, SearchOutlined, UserAddOutlined, ManOutlined, WomanOutlined, QuestionOutlined } from '@ant-design/icons';
import { useState, useEffect } from 'react';
import { TitleSuffix } from '@/common/Text.jsx';
import { Helmet } from 'react-helmet';
import { AxiosGet } from '@/utils/Request';
import { Apis } from '@/common/APIConfig';
import { Avatar, Tag } from 'antd';
import { Descriptions } from 'antd';

const { Option } = Select;

const User = () => {
  // 消息提示
  const { message } = App.useApp();
  const title = '用户中心' + TitleSuffix;

  // 每页数据量
  const [pageSize, setPageSize] = useState(1);

  // 获取用户列表
  const [userList, setUserList] = useState([]);
  useEffect(() => {
    const getUserList = async () => {
      const res = await AxiosGet(Apis.System.User.List);
      if (res.code === 200) {
        setUserList(res.data);
      } else {
        message.error(res.message);
      }
    };
    getUserList();
  }, []);

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
      dataIndex: 'department',
      minWidth: 100,
    },
    {
      title: '岗位',
      dataIndex: 'jobPosition',
      minWidth: 120,
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
      title: '角色关键字',
      dataIndex: ['systemRole', 'keyword'],
      render: (keyword) => {
        if (keyword === 'SuperAdministrator') {
          return <Tag bordered={false} color="magenta">{keyword}</Tag>;
        } else if (keyword === 'Administrator') {
          return <Tag bordered={false} color="volcano">{keyword}</Tag>;
        } else if (keyword === 'DevOps') {
          return <Tag bordered={false} color='green'>{keyword}</Tag>;
        } else {
          return <Tag bordered={false}>{keyword}</Tag>;
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

  // 用户详情展开
  const userDetail = [
    {
      key: '1',
      label: 'UserName',
      children: 'Zhou Maomao',
    },
    {
      key: '2',
      label: 'Telephone',
      children: '1810000000',
    },
    {
      key: '3',
      label: 'Live',
      children: 'Hangzhou, Zhejiang',
    },
    {
      key: '4',
      label: 'Remark',
      children: 'empty',
    },
    {
      key: '5',
      label: 'Address',
      children: 'No. 18, Wantang Road, Xihu District, Hangzhou, Zhejiang, China',
    },
  ];

  // 表格行选择
  const rowSelection = {
    onChange: (selectedRowKeys, selectedRows) => {
      console.log(`selectedRowKeys: ${selectedRowKeys}`, 'selectedRows: ', selectedRows);
    },
    getCheckboxProps: (record) => ({
      disabled: record.name === 'Disabled User',
      // Column configuration not to be checked
      name: record.name
    })
  };

  const [form] = Form.useForm();
  const [expand, setExpand] = useState(false);

  const getFields = () => {
    const count = expand ? 10 : 7;
    const children = [];
    for (let i = 0; i < count; i++) {
      children.push(
        <Col span={6} key={i}>
          {i % 3 !== 1 ? (
            <Form.Item
              name={`field-${i}`}
              label={`Field ${i}`}
              rules={[
                {
                  required: true,
                  message: 'Input something!'
                }
              ]}
            >
              <Input placeholder='placeholder' />
            </Form.Item>
          ) : (
            <Form.Item
              name={`field-${i}`}
              label={`Field ${i}`}
              rules={[
                {
                  required: true,
                  message: 'Select something!'
                }
              ]}
              initialValue='1'
            >
              <Select>
                <Option value='1'>111</Option>
                <Option value='2'>222</Option>
              </Select>
            </Form.Item>
          )}
        </Col>
      );
    }
    return children;
  };
  const onFinish = (values) => {
    console.log('Received values of form: ', values);
  };

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
          <Form form={form} name='advanced_search' onFinish={onFinish}>
            <Row gutter={24}>
              {getFields()}
              <Col span={6} key='x' style={{ marginTop: '10px' }}>
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
            size='small'
            tableLayout='auto'
            rowSelection={{
              type: 'checkbox',
              ...rowSelection
            }}
            columns={columns}
            expandable={{
              expandedRowRender: (record) => (
                <Descriptions column={1} items={userDetail} />
              ),
              rowExpandable: (record) => record.name !== 'Not Expandable'
            }}
            dataSource={userList}
            rowKey='id'
            pagination={{
              pageSize: pageSize,
              showSizeChanger: true,
              showQuickJumper: true,
              onChange: (page, pageSize) => {
                setPageSize(pageSize);
              }
            }}
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