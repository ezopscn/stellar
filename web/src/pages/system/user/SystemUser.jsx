import { useState, useEffect } from 'react';
import { Helmet } from 'react-helmet';
import { Button, Col, Form, Input, Row, Space, Table, Avatar, Tag, Descriptions, TreeSelect, Select, Modal, App, Upload, Dropdown, Popconfirm } from 'antd';
import {
  ClearOutlined,
  DownOutlined,
  SearchOutlined,
  UserAddOutlined,
  ManOutlined,
  WomanOutlined,
  QuestionOutlined,
  DownloadOutlined,
  UploadOutlined,
  InboxOutlined,
  ClockCircleOutlined,
  EditOutlined,
  DeleteOutlined,
  RestOutlined,
  UnorderedListOutlined
} from '@ant-design/icons';
import * as XLSX from 'xlsx';
import { TitleSuffix } from '@/common/Text.jsx';
import { AxiosGet, AxiosPost } from '@/utils/Request.jsx';
import { Apis } from '@/common/APIConfig.jsx';
import APIRequest from '@/common/APIRequest.jsx';
import { SystemRoleStates } from '@/store/StoreSystemRole.jsx';
import { useSnapshot } from 'valtio';
import { BackendURL } from '@/common/APIConfig.jsx';
const { Dragger } = Upload;

const SystemUser = () => {
  const { message } = App.useApp(); // 消息提示
  const [form] = Form.useForm(); // 表单
  const [updateForm] = Form.useForm(); // 添加这行

  // 页面信息
  const title = '用户管理' + TitleSuffix; // 页面标题
  const pageKeyword = '用户'; // 页面关键词

  // 全局
  const { SystemRoleApis } = useSnapshot(SystemRoleStates);

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

  // 定义筛选列表和数据限制，支持类型：input、select、treeSelect
  const filterFields = [
    {
      label: '用户名',
      name: 'username',
      placeholder: '使用用户名��行搜索',
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
      type: 'treeSelect',
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
    filterFields.slice(0, expand ? filterFields.length : defaultExpandItemCount).forEach((item, index) => {
      children.push(
        <Col span={6} key={item.label}>
          <Form.Item name={item.name} label={item.label} rules={item.rules}>
            {item.type === 'input' ? (
              <Input placeholder={item.placeholder} allowClear={true} />
            ) : item.type === 'treeSelect' ? (
              <TreeSelect multiple={item.multiple} placeholder={item.placeholder} treeData={item.data} showSearch={item.search} treeNodeFilterProp="label" allowClear={true} treeDefaultExpandAll />
            ) : item.type === 'select' ? (
              <Select mode={item.multiple ? 'multiple' : 'default'} placeholder={item.placeholder} options={item.data} showSearch={item.search} optionFilterProp="label" allowClear={true} />
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

  // 修改用户状态
  const modifyStatusHandle = async (recordId, operate) => {
    try {
      const res = await AxiosPost(Apis.System.User.StatusModify, { id: recordId, operate: operate });
      if (res.code === 200) {
        message.success('操作成功');
        if (pageNumber > 1) {
          setPageNumber(pageNumber - 1);
        } else {
          window.location.reload();
        }
      } else {
        message.error(res.message);
      }
    } catch (error) {
      console.error('后端服务异常，修改用户状态失败');
      message.error(error);
    }
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
          default: <QuestionOutlined style={{ color: '#999999' }} />
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
          超级管理员: 'magenta',
          管理员: 'volcano',
          运维: 'green'
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
      title: '创建时间',
      dataIndex: 'createdAt'
    },
    {
      title: '操作',
      key: 'action',
      fixed: 'right',
      render: (_, record) => (
        <>
          <Button
            color="primary"
            variant="link"
            icon={<EditOutlined />}
            disabled={!SystemRoleApis.list?.includes(Apis.System.User.Update.replace(BackendURL, ''))}
            onClick={() => {
              setUpdateModalVisible(true);
              setUpdateRecord(record);
            }}
          >
            编辑
          </Button>
          {record.status === 1 ? (
            // 系统内置超级管理员账户不允许禁用
            <Popconfirm
              placement="topRight"
              title="确定要禁用该用户吗？"
              // description="禁用后该用户将无法登录系统"
              okText="确定"
              okButtonProps={{
                style: {
                  backgroundColor: '#ff4d4f',
                  borderColor: '#ff4d4f'
                }
              }}
              cancelText="取消"
              onConfirm={() => {
                modifyStatusHandle(record.id, 'disable');
              }}
            >
              <Button color="danger" variant="link" icon={<DeleteOutlined />} disabled={!SystemRoleApis.list?.includes(Apis.System.User.StatusModify.replace(BackendURL, '')) || record.id === 1}>
                禁用
              </Button>
            </Popconfirm>
          ) : (
            <Popconfirm
              placement="topRight"
              title="确定要启用该用户吗？"
              okText="确定"
              okButtonProps={{
                style: {
                  backgroundColor: '#52c41a',
                  borderColor: '#52c41a'
                }
              }}
              cancelText="取消"
              onConfirm={() => {
                modifyStatusHandle(record.id, 'enable');
              }}
            >
              <Button color="success" variant="link" icon={<RestOutlined />} disabled={!SystemRoleApis.list?.includes(Apis.System.User.StatusModify.replace(BackendURL, ''))}>
                启用
              </Button>
            </Popconfirm>
          )}
        </>
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
  const [multiAddModalVisible, setMultiAddModalVisible] = useState(false);
  const [mutiAddHistoryVisible, setMutiAddHistoryVisible] = useState(false);

  // 定义添加数据的字段
  const addRecordFields = [
    {
      label: '用���名',
      name: 'username',
      placeholder: '请输入用户名',
      type: 'input',
      rules: [
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
      ]
    },
    {
      label: '密码',
      name: 'password',
      placeholder: '请输入密码',
      type: 'inputPassword',
      rules: [
        {
          required: true,
          message: '密码不能为空'
        },
        {
          max: 30,
          message: '密码长度不能超过30个字符'
        },
        {
          min: 8,
          message: '密码长度不能小于8个字符'
        },
        {
          pattern: /^[A-Za-z0-9@$!%*?&]+$/,
          message: '密码只能包含大小写字母、数字和特殊字符（@$!%*?&）'
        }
      ]
    },
    {
      label: '中文名',
      name: 'cnName',
      placeholder: '请输入中文名',
      type: 'input',
      rules: [
        {
          required: true,
          message: '中文名不能为空'
        },
        {
          max: 30,
          message: '中文名长度不能超过30个字符'
        },
        {
          min: 2,
          message: '中文名长度不能小于2个字符'
        },
        {
          pattern: /^[\u4E00-\u9FA5]+$/,
          message: '中文名只能包含中文'
        }
      ]
    },
    {
      label: '英文名',
      name: 'enName',
      placeholder: '请输入英文名',
      type: 'input',
      rules: [
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
      ]
    },
    {
      label: '邮箱',
      name: 'email',
      placeholder: '请输入邮箱',
      type: 'input',
      rules: [
        {
          required: true,
          message: '邮箱不能为空'
        },
        {
          type: 'email',
          message: '邮箱格式不正确'
        }
      ]
    },
    {
      label: '手机号',
      name: 'phone',
      placeholder: '请输入手机号',
      type: 'input',
      rules: [
        {
          required: true,
          message: '手机号不能为空'
        },
        {
          pattern: /^1[3-9]\d{9}$/,
          message: '手机号格式不正确'
        }
      ]
    },
    {
      label: '隐藏手机号',
      name: 'hidePhone',
      type: 'select',
      search: false,
      tree: false,
      multiple: false,
      data: [
        { label: '是', value: 1 },
        { label: '否', value: 0 }
      ],
      rules: [
        {
          required: true,
          message: '隐藏手机号不能为空'
        }
      ]
    },
    {
      label: '性别',
      name: 'gender',
      type: 'select',
      search: false,
      tree: false,
      multiple: false,
      data: [
        { label: '男', value: 1 },
        { label: '女', value: 2 },
        { label: '未知', value: 3 }
      ],
      rules: [
        {
          required: true,
          message: '性别不能为空'
        }
      ]
    },
    {
      label: '部门',
      name: 'systemDepartments',
      type: 'treeSelect',
      search: true,
      tree: true,
      multiple: true,
      data: departmentList,
      rules: [
        {
          required: true,
          message: '部门不能为空'
        }
      ]
    },
    {
      label: '岗位',
      name: 'systemJobPositions',
      type: 'select',
      search: true,
      tree: false,
      multiple: true,
      data: jobPositionList,
      rules: [
        {
          required: true,
          message: '岗位不能为空'
        }
      ]
    },
    {
      label: '角色',
      name: 'systemRole',
      type: 'select',
      search: true,
      tree: false,
      multiple: false,
      data: roleList,
      rules: [
        {
          required: true,
          message: '角色不能为空'
        }
      ]
    },
    {
      label: '个人介绍',
      name: 'description',
      placeholder: '请输入个人介绍',
      type: 'textarea'
    }
  ];

  // 获取添加数据字段
  const getAddRecordFields = () => {
    const children = [];
    addRecordFields.forEach((item) => {
      children.push(
        <Form.Item key={item.name} label={item.label} name={item.name} rules={item.rules}>
          {item.type === 'input' ? (
            <Input placeholder={item.placeholder} allowClear={true} />
          ) : item.type === 'inputPassword' ? (
            <Input.Password placeholder={item.placeholder} allowClear={true} />
          ) : item.type === 'select' ? (
            <Select mode={item.multiple ? 'multiple' : 'default'} options={item.data} showSearch={item.search} optionFilterProp="label" placeholder={item.placeholder} allowClear={true} />
          ) : item.type === 'treeSelect' ? (
            <TreeSelect multiple={item.multiple} placeholder={item.placeholder} treeData={item.data} showSearch={item.search} treeNodeFilterProp="label" allowClear={true} treeDefaultExpandAll />
          ) : item.type === 'textarea' ? (
            <Input.TextArea placeholder={item.placeholder} allowClear={true} />
          ) : null}
        </Form.Item>
      );
    });
    return children;
  };

  // 添加数据方法
  const addRecordHandle = async (data) => {
    try {
      const res = await AxiosPost(Apis.System.User.Add, data);
      if (res.code === 200) {
        message.success('用户添加成功');
        if (pageNumber > 1) {
          setPageNumber(pageNumber - 1); // 返回上一页，则相当于刷新页面
        } else {
          window.location.reload();
        }
      } else {
        message.error(res.message);
      }
    } catch (error) {
      message.error('添加失败：' + error);
    }
  };

  /////////////////////////////////////////////////////
  // 批量导入
  /////////////////////////////////////////////////////
  const [mutiAddFileList, setMutiAddFileList] = useState([]);
  const [mutiAddRecordList, setMutiAddRecordList] = useState([]);
  const [mutiAddRecordBtnDisabled, setMutiAddRecordBtnDisabled] = useState(true);

  // 批量导入记录表格列
  const mutiAddHistoryColumns = [
    {
      title: '任务ID',
      dataIndex: 'taskId'
    },
    {
      title: '执行人',
      dataIndex: 'executor'
    },
    {
      title: '记录数',
      dataIndex: 'recordCount'
    },
    {
      title: '成功数',
      dataIndex: 'successCount'
    },
    {
      title: '失败数',
      dataIndex: 'failCount'
    },
    {
      title: '任务状态',
      dataIndex: 'status',
      render: (_, record) => {
        const statusMap = {
          1: { text: '执行中', color: 'gray' },
          2: { text: '执行成功', color: 'green' },
          3: { text: '执行失败', color: 'red' }
        };
        return <Tag color={statusMap[record.taskStatus].color}>{statusMap[record.taskStatus].text}</Tag>;
      }
    },
    {
      title: '任务开始时间',
      dataIndex: 'taskStartTime'
    },
    {
      title: '任务结束时间',
      dataIndex: 'taskEndTime'
    },
    {
      title: '操作',
      key: 'action',
      fixed: 'right',
      render: (_, record) => (
        <Space size="middle">
          <a>
            <UnorderedListOutlined /> 查看明细
          </a>
        </Space>
      )
    }
  ];

  // 监听 mutiAddRecordList 的变化，如果列表不为空，则按钮可以点击
  useEffect(() => {
    if (mutiAddRecordList.length > 0) {
      setMutiAddRecordBtnDisabled(false);
    } else {
      setMutiAddRecordBtnDisabled(true);
    }
  }, [mutiAddRecordList]);

  // 批量添加方法
  const mutiAddRecordHandle = async () => {
    try {
      if (mutiAddRecordList.length === 0) {
        message.error('未获取到导入数据，请检测文件格式是否正确');
        return;
      }
      console.log(mutiAddRecordList);
      const res = await AxiosPost(Apis.System.User.MutiAdd, mutiAddRecordList);
      if (res.code === 200) {
        message.success('批量导入任务已创建，请稍后查看任务状态');
      } else {
        message.error(res.message);
      }
    } catch (error) {
      message.error('批量导入失败：' + error);
    }
    // 清理上传文件
    setMultiAddModalVisible(false);
    setMutiAddFileList([]);
    setMutiAddRecordList([]);
  };

  // 选择文件后的操作
  const mutiAddRecordProps = {
    name: 'file',
    multiple: false,
    maxCount: 1,
    fileList: mutiAddFileList,
    beforeUpload(file) {
      // 阻止自动上传
      // 判断如果文件不是 xlsx 后缀，则不进行上传
      if (!file.name.endsWith('.xlsx')) {
        message.error('请根据用户批量添加模板上传 Excel 文件');
      }
      return false;
    },
    onChange(info) {
      // 获取文件列表，然后读取文件
      setMutiAddFileList(info.fileList);
      // 如果文件列表为空，则清空列表
      if (info.fileList.length === 0) {
        setMutiAddRecordList([]);
        setMutiAddFileList([]);
        return;
      }

      // 使用 xlsx 读取文件
      const file = info.fileList[0]?.originFileObj;
      const reader = new FileReader();
      reader.readAsArrayBuffer(file);
      reader.onload = (e) => {
        const data = new Uint8Array(e.target.result);
        // 为了避免因为单元格格式问题导致和后端数据类型不一致的情况，所有都使用字符串
        const workbook = XLSX.read(data, { type: 'array', raw: false });
        const sheetNameList = workbook.SheetNames;
        const json = sheetNameList.map((sheet) => {
          return XLSX.utils.sheet_to_json(workbook.Sheets[sheet], {
            raw: false,
            defval: '' // 设置空单元格的默认值为空字符串
          });
        });
        // 前 4 条数据都是模板数据，不需要
        const jsonData = json[0].slice(4);
        setMutiAddRecordList(jsonData);
      };
    } // 单文件不需要再定义 drop，删除也是 change 的一种
  };

  /////////////////////////////////////////////////////
  // 编辑弹窗
  /////////////////////////////////////////////////////
  const [updateModalVisible, setUpdateModalVisible] = useState(false);
  const [updateRecord, setUpdateRecord] = useState(null);

  // 定义编辑数据的字段
  const updateRecordFields = [
    {
      label: 'ID',
      name: 'id',
      type: 'input',
      value: updateRecord?.id,
      hidden: true
    },
    {
      label: '用户名',
      name: 'username',
      placeholder: '请输入用户名',
      type: 'input',
      value: updateRecord?.username,
      rules: [
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
      ]
    },
    {
      label: '中文名',
      name: 'cnName',
      placeholder: '请输入中文名',
      type: 'input',
      value: updateRecord?.cnName,
      rules: [
        {
          required: true,
          message: '中文名不能为空'
        },
        {
          max: 30,
          message: '中文名长度不能超过30个字符'
        },
        {
          min: 2,
          message: '中文名长度不能小于2个字符'
        },
        {
          pattern: /^[\u4E00-\u9FA5]+$/,
          message: '中文名只能包含中文'
        }
      ]
    },
    {
      label: '英文名',
      name: 'enName',
      placeholder: '请输入英文名',
      type: 'input',
      value: updateRecord?.enName,
      rules: [
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
      ]
    },
    {
      label: '邮箱',
      name: 'email',
      placeholder: '请输入邮箱',
      type: 'input',
      value: updateRecord?.email,
      rules: [
        {
          required: true,
          message: '邮箱不能为空'
        },
        {
          type: 'email',
          message: '邮箱格式不正确'
        }
      ]
    },
    {
      label: '手机号',
      name: 'phone',
      placeholder: '请输入手机号',
      type: 'input',
      value: updateRecord?.phone,
      rules: [
        {
          required: true,
          message: '手机号不能为空'
        },
        {
          pattern: /^1[3-9]\d{9}$/,
          message: '手机号格式不正确'
        }
      ]
    },
    {
      label: '隐藏手机号',
      name: 'hidePhone',
      type: 'select',
      search: false,
      tree: false,
      multiple: false,
      value: updateRecord?.hidePhone,
      data: [
        { label: '是', value: 1 },
        { label: '否', value: 0 }
      ],
      rules: [
        {
          required: true,
          message: '隐藏手机号不能为空'
        }
      ]
    },
    {
      label: '性别',
      name: 'gender',
      type: 'select',
      search: false,
      tree: false,
      multiple: false,
      value: updateRecord?.gender,
      data: [
        { label: '男', value: 1 },
        { label: '女', value: 2 },
        { label: '未知', value: 3 }
      ],
      rules: [
        {
          required: true,
          message: '性别不能为空'
        }
      ]
    },
    {
      label: '部门',
      name: 'systemDepartments',
      type: 'treeSelect',
      search: true,
      tree: true,
      multiple: true,
      data: departmentList,
      // 遍历数据，获取 Id 列表
      value: updateRecord?.systemDepartments?.map((item) => item.id),
      rules: [
        {
          required: true,
          message: '部门不能为空'
        }
      ]
    },
    {
      label: '岗位',
      name: 'systemJobPositions',
      type: 'select',
      search: true,
      tree: false,
      multiple: true,
      data: jobPositionList,
      // 遍历数据，获取 Id 列表
      value: updateRecord?.systemJobPositions?.map((item) => item.id),
      rules: [
        {
          required: true,
          message: '岗位不能为空'
        }
      ]
    },
    {
      label: '角色',
      name: 'systemRole',
      type: 'select',
      search: true,
      tree: false,
      multiple: false,
      data: roleList,
      // 遍历数据，获取 Id 列表
      value: updateRecord?.systemRole?.id,
      rules: [
        {
          required: true,
          message: '角色不能为空'
        }
      ]
    },
    {
      label: '个人介绍',
      name: 'description',
      placeholder: '请输入个人介绍',
      type: 'textarea',
      value: updateRecord?.description
    }
  ];

  // 获取编辑数据字段
  const getUpdateRecordFields = () => {
    const children = [];
    updateRecordFields.forEach((item) => {
      children.push(
        <Form.Item key={item.name} label={item.label} name={item.name} rules={item.rules} hidden={item.hidden} initialValue={item.value}>
          {item.type === 'input' ? (
            <Input placeholder={item.placeholder} allowClear={true} disabled={item.disabled} />
          ) : item.type === 'inputPassword' ? (
            <Input.Password placeholder={item.placeholder} allowClear={true} disabled={item.disabled} />
          ) : item.type === 'select' ? (
            <Select
              mode={item.multiple ? 'multiple' : 'default'}
              options={item.data}
              showSearch={item.search}
              optionFilterProp="label"
              placeholder={item.placeholder}
              allowClear={true}
              disabled={item.disabled}
            />
          ) : item.type === 'treeSelect' ? (
            <TreeSelect
              multiple={item.multiple}
              placeholder={item.placeholder}
              treeData={item.data}
              showSearch={item.search}
              treeNodeFilterProp="label"
              allowClear={true}
              treeDefaultExpandAll
              disabled={item.disabled}
            />
          ) : item.type === 'textarea' ? (
            <Input.TextArea placeholder={item.placeholder} allowClear={true} disabled={item.disabled} />
          ) : null}
        </Form.Item>
      );
    });
    return children;
  };

  // 编辑数据方法
  const updateRecordHandle = async (data) => {
    try {
      const res = await AxiosPost(Apis.System.User.Edit, data);
      if (res.code === 200) {
        message.success('用户编辑成功');
        if (pageNumber > 1) {
          setPageNumber(pageNumber - 1); // 返回上一页，则相当于刷新页面
        } else {
          window.location.reload();
        }
      } else {
        message.error(res.message);
      }
    } catch (error) {
      message.error('添加失败：' + error);
    }
  };

  // 监听 updateRecord 的变化，并在变化时重置表单
  useEffect(() => {
    if (updateRecord && updateModalVisible) {
      updateForm.setFieldsValue({
        id: updateRecord.id,
        username: updateRecord.username,
        cnName: updateRecord.cnName,
        enName: updateRecord.enName,
        email: updateRecord.email,
        phone: updateRecord.phone,
        hidePhone: updateRecord.hidePhone,
        gender: updateRecord.gender,
        systemDepartments: updateRecord.systemDepartments?.map((item) => item.id),
        systemJobPositions: updateRecord.systemJobPositions?.map((item) => item.id),
        systemRole: updateRecord.systemRole?.id,
        description: updateRecord.description
      });
    }
  }, [updateRecord, updateModalVisible]);

  /////////////////////////////////////////////////////
  // 批量操作菜单
  /////////////////////////////////////////////////////
  const [mutiOperateKey, setMutiOperateKey] = useState('');
  // 表格：行选择
  const rowSelection = {
    selectedRowKeys: mutiOperateKey,
    onChange: (selectedRowKeys, selectedRows) => {
      setMutiOperateKey(selectedRowKeys);
    },
    getCheckboxProps: (record) => ({
      // 系统超级管理员账户不允许选择
      disabled: record.id === 1
    })
  };

  // 批量操作菜单
  const mutiOperateMenuProps = {
    items: [
      {
        label: '批量禁用',
        key: 'disable',
        danger: true,
        disabled: mutiOperateKey.length === 0
      },
      {
        label: '批量启用',
        key: 'enable',
        danger: false,
        disabled: mutiOperateKey.length === 0
      }
    ],
    onClick: (key) => {
      if (mutiOperateKey.length > 0) {
        const mutiModifyStatusHandle = async () => {
          const req = {
            ids: mutiOperateKey,
            operate: key.key
          };
          try {
            const res = await AxiosPost(Apis.System.User.StatusMutiModify, req);
            if (res.code === 200) {
              message.success('批量操作成功');
              setMutiOperateKey([]);
              if (pageNumber > 1) {
                setPageNumber(pageNumber - 1); // 返回上一页，则相当于刷新页面
              } else {
                window.location.reload();
              }
            } else {
              message.error(res.message);
            }
          } catch (error) {
            message.error('批量操作失败：' + error);
          }
        };
        mutiModifyStatusHandle();
      }
    }
  };

  return (
    <>
      {/* 页面 header */}
      <Helmet>
        <title>{title}</title>
        <meta name="description" content={title} />
      </Helmet>
      {/* 页面头部介绍 */}
      <div className="admin-page-header">
        <div className="admin-page-title">用户管理 / USER MANAGEMENT.</div>
        <div className="admin-page-desc">
          <div>用户是系统的核心资产之一，也是许多其它资源的强制依赖，所以对于用户的管理，我提供了以下的要求和建议，请知悉：</div>
          <ul>
            <li>系统内置的超级管理员账户涉及到系统的基础逻辑判断，不允许从数据库中物理删除，用户在初始化完成之后需要对其基础信息进行修改，保障系统安全性。</li>
            <li>为了保障数据的安全性和可靠性，系统将强制使用禁用用户来替代删除用户，禁用用户将无法登录系统，但是数据仍然保留，可以随时恢复。</li>
            <li>针对某些特殊的用户，例如老板、高管等，我们建议隐藏其联系方式，保护个人隐私。</li>
          </ul>
        </div>
      </div>
      {/* 页面主体 */}
      <div className="admin-page-main">
        {/* 搜索栏 */}
        <div className="admin-page-search">
          <Form form={form} labelCol={{ span: 6 }} wrapperCol={{ span: 18 }} colon={false} name="filterForm" onFinish={filterDataListHandle} autoComplete="off">
            <Row gutter={24}>
              {getFilterFields()}
              <Col span={24} key="x" style={{ marginTop: '10px', textAlign: 'right' }}>
                <Space>
                  <Button icon={<SearchOutlined />} htmlType="submit">
                    条件搜索
                  </Button>
                  <Button icon={<ClearOutlined />} onClick={() => form.resetFields()}>
                    清理条件
                  </Button>
                  {filterFields.length > defaultExpandItemCount && (
                    <a onClick={() => setExpand(!expand)}>
                      <DownOutlined rotate={expand ? 180 : 0} /> {expand ? '收起条件' : '展开更多'}
                    </a>
                  )}
                </Space>
              </Col>
            </Row>
          </Form>
        </div>
        {/* 表格 */}
        <div className="admin-page-list">
          <div className="admin-page-btn-group">
            <Space>
              <Button type="primary" icon={<UserAddOutlined />} onClick={() => setAddModalVisible(true)} disabled={!SystemRoleApis.list?.includes(Apis.System.User.Add.replace(BackendURL, ''))}>
                添加{pageKeyword}
              </Button>
              <Button icon={<UploadOutlined />} onClick={() => setMultiAddModalVisible(true)} disabled={!SystemRoleApis.list?.includes(Apis.System.User.MutiAdd.replace(BackendURL, ''))}>
                批量导入
              </Button>
              <Dropdown menu={mutiOperateMenuProps} disabled={!SystemRoleApis.list?.includes(Apis.System.User.StatusMutiModify.replace(BackendURL, ''))}>
                <Button>
                  <Space>
                    <DownOutlined />
                    批量操作
                  </Space>
                </Button>
              </Dropdown>
            </Space>
            <Space style={{ float: 'right' }}>
              <Button
                icon={<DownloadOutlined />}
                onClick={() => {
                  window.open('/template/用户批量添加模板.xlsx');
                }}
              >
                模板下载
              </Button>
              <Button icon={<ClockCircleOutlined />} onClick={() => setMutiAddHistoryVisible(true)}>
                导入记录
              </Button>
            </Space>
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
                    children: record.creator.split(',')[1] + ' / ' + record.creator.split(',')[2] + ' ( ' + record.creator.split(',')[0] + ' / ' + record.creator.split(',')[3] + ' )'
                  },
                  {
                    label: '个人介绍',
                    children: record.description
                  },
                  {
                    label: '更新时间',
                    children: record.updatedAt
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
            scroll={{ x: 'max-content' }}
          />
        </div>
      </div>

      {/* 用户添加弹窗 */}
      <Modal title={'添加' + pageKeyword} open={addModalVisible} onOk={addRecordHandle} onCancel={() => setAddModalVisible(false)} width={400} maskClosable={false} footer={null}>
        <Form form={form} labelCol={{ span: 6 }} wrapperCol={{ span: 18 }} colon={false} name="addRecordForm" onFinish={addRecordHandle} autoComplete="off">
          {getAddRecordFields()}
          <Form.Item wrapperCol={{ span: 24 }}>
            <Button type="primary" htmlType="submit" block>
              添加{pageKeyword}
            </Button>
          </Form.Item>
        </Form>
      </Modal>

      {/* 用户编辑弹窗 */}
      <Modal
        title={'编辑' + pageKeyword}
        open={updateModalVisible}
        onCancel={() => {
          setUpdateModalVisible(false);
          updateForm.resetFields(); // 关闭时重置表单，避免数据不更新
        }}
        width={400}
        maskClosable={false}
        footer={null}
      >
        <Form form={updateForm} labelCol={{ span: 6 }} wrapperCol={{ span: 18 }} colon={false} name="updateRecordForm" onFinish={updateRecordHandle} autoComplete="off">
          {getUpdateRecordFields()}
          <Form.Item wrapperCol={{ span: 24 }}>
            <Button type="primary" htmlType="submit" block>
              保存编辑
            </Button>
          </Form.Item>
        </Form>
      </Modal>

      {/* 批量导入弹窗 */}
      <Modal
        title={'批量导入' + pageKeyword}
        open={multiAddModalVisible}
        onCancel={() => {
          setMultiAddModalVisible(false);
          setMutiAddRecordList([]);
          setMutiAddFileList([]);
        }}
        width={800}
        maskClosable={false}
        footer={null}
      >
        <Dragger {...mutiAddRecordProps}>
          <p className="ant-upload-drag-icon">
            <InboxOutlined />
          </p>
          <p className="ant-upload-text">点击或者拖拽批量导入模板文件到此区域完成创建</p>
          <p className="ant-upload-hint">为了数据安全，只支持单个文件模板导入，严格禁止上传公司数据或者其它违规文件。</p>
        </Dragger>
        <div style={{ marginTop: '10px' }}>
          <Button disabled={mutiAddRecordBtnDisabled} type="primary" htmlType="submit" block onClick={mutiAddRecordHandle}>
            批量导入
          </Button>
        </div>
      </Modal>

      {/* 批量导入记录 */}
      <Modal title={'批量导入任务'} open={mutiAddHistoryVisible} onCancel={() => setMutiAddHistoryVisible(false)} width={1200} maskClosable={false} footer={null}>
        <Table
          // 表格布局大小
          size="small"
          // 表格布局方式，支持 fixed、auto
          tableLayout="auto"
          // 表格列
          columns={mutiAddHistoryColumns}
          dataSource={[]}
          // 行唯一标识
          rowKey="id"
          // 表格滚动，目的是为了最后一列固定
          scroll={{ x: 'max-content' }}
        />
      </Modal>
    </>
  );
};

export default SystemUser;
