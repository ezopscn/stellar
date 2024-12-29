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
import { DATA_STATUS_MAP, GENDER_MAP, FORM_ITEM_COMPONENT_MAP, TRUE_FALSE_MAP } from '@/common/GlobalConfig.jsx';

const { Dragger } = Upload;

// 页面常量设置
const PageConfig = {
  // 默认页码
  defaultPageNumber: 1,
  // 默认每页显示的数据量
  defaultPageSize: 2,
  // 默认数据总数
  defaultTotal: 0,
  // 默认是否需要分页
  defaultIsPagination: true,
  // 页面标题
  pageTitle: '用户管理' + TitleSuffix,
  // 页面顶部标题
  pageHeaderTitle: '用户管理 / USER MANAGEMENT.',
  // 页面关键词
  pageKeyword: '用户',
  // 默认搜索展开字段数量
  defaultFilterExpandItemCount: 8,
  // 表格数据接口地址
  tableDataListAPI: Apis.System.User.List
};

// 页面说明组件
const PageDescriptionComponent = () => {
  return (
    <>
      <div>用户是系统的核心资产之一，也是许多其它资源的强制依赖，所以对于用户的管理，我提供了以下的要求和建议，请知悉：</div>
      <ul>
        <li>系统内置的超级管理员账户涉及到系统的基础逻辑判断，不允许从数据库中物理删除，用户在初始化完成之后需要对其基础信息进行修改，保障系统安全性。</li>
        <li>为了保障数据的安全性和可靠性，系统将强制使用禁用用户来替代删除用户，禁用用户将无法登录系统，但是数据仍然保留，可以随时恢复。</li>
        <li>针对某些特殊的用户，例如老板、高管等，我们建议隐藏其联系方式，保护个人隐私。</li>
      </ul>
    </>
  );
};

// 生成性别图标
const generateGenderIcon = (gender) => {
  const genderIcons = {
    1: <ManOutlined style={{ color: '#165dff' }} />,
    2: <WomanOutlined style={{ color: '#ff4d4f' }} />,
    default: <QuestionOutlined style={{ color: '#999999' }} />
  };
  return genderIcons[gender] || genderIcons.default;
};

// 生成角色标签
const generateRoleTag = (name) => {
  const roleColors = {
    超级管理员: 'magenta',
    管理员: 'volcano',
    运维: 'green'
  };
  const color = roleColors[name] || '';
  return <Tag bordered={false} color={color}>{name}</Tag>;
};

// 生成状态标签
const generateStatusTag = (status) => {
  const statusMap = {
    1: { text: '启用', color: 'green' },
    0: { text: '禁用', color: 'red' }
  };
  const { text, color } = statusMap[status] || {};
  return <Tag bordered={false} color={color}>{text}</Tag>;
};

// 生成任务状态标签
const generateTaskStatusTag = (status) => {
  const statusMap = {
    1: { text: '执行中', color: 'gray' },
    2: { text: '执行成功', color: 'green' },
    3: { text: '执行失败', color: 'red' }
  };
  const { text, color } = statusMap[status] || {};
  return <Tag bordered={false} color={color}>{text}</Tag>;
};

// 用户管理
const SystemUser = () => {
  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 全局配置
  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 消息提示
  const { message } = App.useApp();
  // 全局数据，用于父子组件之间数据传递
  const { SystemRoleApis } = useSnapshot(SystemRoleStates);

  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 页面基础数据
  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 下拉菜单的角色列表数据
  const [systemRoleList, setSystemRoleList] = useState([]);
  // 下拉菜单的岗位列表数据
  const [systemJobPositionList, setSystemJobPositionList] = useState([]);
  // 下拉菜单的部门列表数据
  const [systemDepartmentList, setSystemDepartmentList] = useState([]);
  
  // 页面加载的时候一次性获取依赖的异步数据
  useEffect(() => {
    // 获取角色列表
    APIRequest.GetSelectDataList(Apis.System.Role.List, setSystemRoleList);
    // 获取岗位列表
    APIRequest.GetSelectDataList(Apis.System.JobPosition.List, setSystemJobPositionList);
    // 获取部门列表（树形结构）
    APIRequest.GetSelectDataList(Apis.System.Department.List, setSystemDepartmentList, true);
  }, []);

  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 状态：请求和筛选列表
  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 数据列表 
  const [tableRecordList, setTableRecordList] = useState([]); 
  // 每页显示的数据条数
  const [pageSize, setPageSize] = useState(PageConfig.defaultPageSize); 
  // 当前页码
  const [pageNumber, setPageNumber] = useState(PageConfig.defaultPageNumber); 
  // 数据总数
  const [total, setTotal] = useState(PageConfig.defaultTotal);
  // 是否需要分页
  const [isPagination, setIsPagination] = useState(PageConfig.defaultIsPagination);
  // 条件搜索参数
  const [filterRecordParams, setFilterRecordParams] = useState({}); 

  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 按钮权限控制
  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 添加按钮权限控制
  const addRecordButtonDisabled = !SystemRoleApis.list?.includes(Apis.System.User.Add.replace(BackendURL, ''));
  // 修改按钮权限控制
  const updateRecordButtonDisabled = !SystemRoleApis.list?.includes(Apis.System.User.Update.replace(BackendURL, ''));
  // 状态修改按钮权限控制
  const modifyRecordStatusButtonDisabled = !SystemRoleApis.list?.includes(Apis.System.User.ModifyStatus.replace(BackendURL, ''));
  // 批量添加按钮权限控制
  const multiAddRecordButtonDisabled = !SystemRoleApis.list?.includes(Apis.System.User.MultiAdd.replace(BackendURL, ''));
  // 批量状态修改按钮权限控制
  const multiModifyRecordStatusButtonDisabled = !SystemRoleApis.list?.includes(Apis.System.User.MultiModifyStatus.replace(BackendURL, ''));

  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 公共方法
  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 传入元素，生成表单项通用方法
  const generateRecordFormItem = (item) => {
    const commonProps = {
      allowClear: true,
      placeholder: item.placeholder,
      hidden: item.hidden,
      value: item.value
    };
    const componentProps = {
      input: {},
      inputPassword: {},
      textarea: {},
      treeSelect: {
        treeData: item.data,
        showSearch: item.search,
        treeNodeFilterProp: 'label',
        multiple: item.multiple
      },
      select: {
        options: item.data,
        showSearch: item.search,
        optionFilterProp: 'label',
        mode: item.multiple ? 'multiple' : 'default'
      }
    };
    return (
      <Form.Item key={item.name} name={item.name} label={item.label} rules={item.rules} hidden={item.hidden}>
        {FORM_ITEM_COMPONENT_MAP[item.type]?.({
          ...commonProps,
          ...componentProps[item.type]
        })}
      </Form.Item>  
    );
  };

  // 刷新当前页面数据，而不是刷新整个页面
  const refreshCurrentPage = () => {
    const params = { ...filterRecordParams, pageSize, pageNumber, isPagination };
    filterRecordListRequest(params);
  };

  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 默认表格数据列表和数据搜索
  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 搜索表单
  const [filterRecordForm] = Form.useForm();
  
  // 是否展开更多搜索
  const [expandFilterRecordItems, setExpandFilterRecordItems] = useState(false);

  // 用户操作按钮组
  const actionButtonGroup = (record) => {
    return (
      <>
        <Button color="primary" variant="link" icon={<EditOutlined />} 
          disabled={updateRecordButtonDisabled} 
          onClick={() => {
            setUpdateRecordModalVisible(true);
            setUpdateRecord(record);
          }}>编辑</Button>
        {record.status === 1 ? ( // 系统内置超级管理员账户不允许禁用
          <Popconfirm placement="topRight" title="确定要禁用该用户吗？" okText="确定" cancelText="取消" 
            okButtonProps={{ style: { backgroundColor: '#ff4d4f', borderColor: '#ff4d4f' } }}
            onConfirm={() => modifyRecordStatusHandle(record.id, 'disable')}>
            <Button color="danger" variant="link" icon={<DeleteOutlined />} disabled={modifyRecordStatusButtonDisabled || record.id === 1}>禁用</Button>
          </Popconfirm>
        ) : (
          <Popconfirm placement="topRight" title="确定要启用该用户吗？" okText="确定" cancelText="取消"
            okButtonProps={{ style: { backgroundColor: '#52c41a', borderColor: '#52c41a' } }} 
            onConfirm={() => { modifyRecordStatusHandle(record.id, 'enable'); }}>
            <Button color="success" variant="link" icon={<RestOutlined />} disabled={modifyRecordStatusButtonDisabled}>启用</Button>
          </Popconfirm>
        )}
      </>
    );
  };

  // 表格列定义，字段说明：
  // title：列标题
  // dataIndex：数据列对应的 JSON 字段
  // minWidth：最小宽度
  // render：渲染函数，用于自定义列的显示内容
  const tableColumns = [
    { title: '头像', dataIndex: 'avatar', render: (avatar) => <Avatar src={avatar} /> },
    { title: '中文名', dataIndex: 'cnName', minWidth: 80 },
    { title: '英文名', dataIndex: 'enName', minWidth: 80 },
    { title: '性别', dataIndex: 'gender', minWidth: 50, render: (gender) => generateGenderIcon(gender) },
    { title: '用户名', dataIndex: 'username', minWidth: 80 },
    { title: '邮箱', dataIndex: 'email' },
    { title: '手机号', dataIndex: 'phone' },
    { title: '部门', dataIndex: 'systemDepartments', minWidth: 100, render: (systemDepartments) => systemDepartments?.map((department) => department.name).join(',') },
    { title: '岗位', dataIndex: 'systemJobPositions', minWidth: 120, render: (systemJobPositions) => systemJobPositions?.map((jobPosition) => jobPosition.name).join(',') },
    { title: '角色名称', dataIndex: ['systemRole', 'name'], render: (name) => generateRoleTag(name) },
    { title: '状态', dataIndex: 'status', render: (status) => generateStatusTag(status) },
    { title: '创建时间', dataIndex: 'createdAt' },
    { title: '操作', key: 'action', fixed: 'right', render: (_, record) => actionButtonGroup(record) }
  ];

  // 定义筛选列表和数据限制，以下是字段说明（适用于筛选，添加，修改）：
  // label：页面显示名称
  // name：表单字段名称
  // placeholder：字段输入提示
  // type：字段类型，支持 input、inputPassword、textarea、select、treeSelect
  // rules：字段验证规则
  // search：是否允许搜索
  // tree：是否展示树形结构
  // multiple：是否允许多选
  // data：数据列表
  const filterRecordFields = [
    { label: '用户名', name: 'username', placeholder: '使用用户名进行搜索', type: 'input', rules: [{ message: '用户名长度不能超过30个字符', max: 30 }] },
    { label: '姓名', name: 'name', placeholder: '使用中文名或者英文名进行搜索', type: 'input', rules: [{ message: '姓名长度不能超过30个字符', max: 30 }] },
    { label: '邮箱', name: 'email', placeholder: '使用邮箱地址进行搜索', type: 'input', rules: [{ message: '邮箱长度不能超过30个字符', max: 30 }] },
    { label: '手机号', name: 'phone', placeholder: '使用手机号码进行搜索', type: 'input', rules: [{ message: '手机号长度不能超过15个字符', max: 15 }] },
    { label: '状态', name: 'status', placeholder: '选择用户状态进行搜索', type: 'select', search: false, tree: false, multiple: false, data: DATA_STATUS_MAP, rules: [] },
    { label: '性别', name: 'gender', placeholder: '选择性别进行搜索', type: 'select', search: false, tree: false, multiple: false, data: GENDER_MAP, rules: [] },
    { label: '部门', name: 'systemDepartment', placeholder: '选择部门进行搜索', type: 'treeSelect', search: true, tree: true, multiple: false, data: systemDepartmentList, rules: [] },
    { label: '岗位', name: 'systemJobPosition', placeholder: '选择岗位进行搜索', type: 'select', search: true, tree: false, multiple: false, data: systemJobPositionList, rules: [] },
    { label: '角色', name: 'systemRole', placeholder: '选择角色进行搜索', type: 'select', search: true, tree: false, multiple: false, data: systemRoleList, rules: [] }
  ];

  // 生成搜索栏表单项
  const generateFilterRecordFormItems = () => {
    return filterRecordFields
      .slice(0, expandFilterRecordItems ? filterRecordFields.length : PageConfig.defaultFilterExpandItemCount)
      .map(item => (
        <Col span={6} key={item.name}>
          {generateRecordFormItem(item)}
        </Col>
      ));
  };
  
  // 条件搜索请求封装，默认请求其实也属于搜索的一种类型
  const filterRecordListRequest = async (params) => {
    try {
      const res = await AxiosGet(PageConfig.tableDataListAPI, params);
      if (res?.code === 200) {
        setTableRecordList(res?.data?.list);
        setPageSize(res?.data?.pagination?.pageSize);
        setPageNumber(res?.data?.pagination?.pageNumber);
        setTotal(res?.data?.pagination?.total);
      } else {
        message.error(res?.message);
      }
    } catch (error) {
      message.error('服务异常，获取数据列表失败');
      console.log(error);
    }
  };

  // 手动搜索方法，需要先初始化页码，避免在搜索结果溢出搜索数据的页码的时候，导致请求参数中带了页码，无法请求到数据
  // 比如在第三页的时候搜索，但是结果只有两页，则会因为页码问题，显示没有数据
  const filterRecordListHandle = (params) => {
    setPageNumber(PageConfig.defaultPageNumber);
    setPageSize(PageConfig.defaultPageSize);
    setTotal(PageConfig.defaultTotal);
    setFilterRecordParams(params);
  };

  // 监听展示条件变化，然后请求数据
  useEffect(() => {
    const params = { ...filterRecordParams, pageSize, pageNumber, isPagination };
    filterRecordListRequest(params);
  }, [pageSize, pageNumber, filterRecordParams]);

  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 添加数据
  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 添加表单
  const [addRecordForm] = Form.useForm();

  // 添加数据弹窗
  const [addRecordModalVisible, setAddRecordModalVisible] = useState(false);

  // 表单基础字段
  const recordFormBasicFields = [
    { label: '用户名', name: 'username', placeholder: '请输入用户名', type: 'input', rules: [
      {required: true,message: '用户名不能为空'},
      {pattern: /^[a-zA-Z][a-zA-Z0-9_]*$/,message: '用户名只能以字母开头，且只能包含字母、数字和下划线'},
      {max: 30,message: '用户名长度不能超过30个字符'},
      {min: 3,message: '用户名长度不能小于3个字符'}
    ]},
    { label: '中文名', name: 'cnName', placeholder: '请输入中文名', type: 'input', rules: [
      {required: true,message: '中文名不能为空'},
      {max: 30,message: '中文名长度不能超过30个字符'},
      {min: 2,message: '中文名长度不能小于2个字符'},
      {pattern: /^[\u4E00-\u9FA5]+$/,message: '中文名只能包含中文'}
    ]},
    { label: '英文名', name: 'enName', placeholder: '请输入英文名', type: 'input', rules: [
      {required: true,message: '英文名不能为空'},
      {pattern: /^[a-zA-Z]+$/,message: '英文名只能包含字母'},
      {max: 30,message: '英文名长度不能超过30个字符'},
      {min: 2,message: '英文名长度不能小于2个字符'}
    ]},
    { label: '邮箱', name: 'email', placeholder: '请输入邮箱', type: 'input', rules: [
      {required: true,message: '邮箱不能为空'},
      {type: 'email',message: '邮箱格式不正确'}
    ]},
    { label: '手机号', name: 'phone', placeholder: '请输入手机号', type: 'input', rules: [
      {required: true,message: '手机号不能为空'},
      {pattern: /^1[3-9]\d{9}$/,message: '手机号格式不正确'}
    ]},
    { label: '隐藏手机号', name: 'hidePhone', type: 'select', search: false, tree: false, multiple: false, data: TRUE_FALSE_MAP, rules: [
      { required: true, message: '隐藏手机号不能为空' }
    ]},
    { label: '性别', name: 'gender', type: 'select', search: false, tree: false, multiple: false, data: GENDER_MAP, rules: [
      { required: true, message: '性别不能为空' }
    ]},
    { label: '部门', name: 'systemDepartments', type: 'treeSelect', search: true, tree: true, multiple: true, data: systemDepartmentList, rules: [
      { required: true, message: '部门不能为空' }
    ]},
    { label: '岗位', name: 'systemJobPositions', type: 'select', search: true, tree: false, multiple: true, data: systemJobPositionList, rules: [
      { required: true, message: '岗位不能为空' }
    ]},
    { label: '角色', name: 'systemRole', type: 'select', search: true, tree: false, multiple: false, data: systemRoleList, rules: [
      { required: true, message: '角色不能为空' }
    ]},
    { label: '个人介绍', name: 'description', placeholder: '请输入个人介绍', type: 'textarea' }
  ];

  // 定义添加数据的字段
  const addRecordFields = [
    // 将密码字段往前放
    ...recordFormBasicFields.slice(0, 1),
    { label: '密码', name: 'password', placeholder: '请输入密码', type: 'inputPassword', rules: [
      {required: true,message: '密码不能为空'},
      {max: 30,message: '密码长度不能超过30个字符'},
      {min: 8,message: '密码长度不能小于8个字符'},
      {pattern: /^[A-Za-z0-9@$!%*?&]+$/,message: '密码只能包含大小写字母、数字和特殊字符（@$!%*?&）'}
    ]},
    ...recordFormBasicFields.slice(1),
  ];

  // 生成添加数据表单项
  const generateAddRecordFormItems = () => {
    return addRecordFields?.map(item => generateRecordFormItem(item));
  };

  // 添加数据方法
  const addRecordHandle = async (data) => {
    try {
      const res = await AxiosPost(Apis.System.User.Add, data);
      if (res.code === 200) {
        message.success('添加成功');
        setAddRecordModalVisible(false);
        refreshCurrentPage();
      } else {
        message.error(res.message);
      }
    } catch (error) {
      message.error('服务异常，添加数据失败');
      console.error(error);
    }
  };

  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 批量导入
  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 批量添加数据弹窗
  const [multiAddRecordModalVisible, setMultiAddRecordModalVisible] = useState(false);  
  // 批量导入文件列表
  const [multiAddFileList, setMultiAddFileList] = useState([]);
  // 批量导入数据列表
  const [multiAddRecordList, setMultiAddRecordList] = useState([]);
  // 批量导入按钮是否禁用，没数据时候不能导入
  const [multiAddRecordSubmitButtonDisabled, setMultiAddRecordSubmitButtonDisabled] = useState(true);

  // 监听 mutiAddRecordList 的变化，如果列表不为空，则按钮可以点击
  useEffect(() => {
    setMultiAddRecordSubmitButtonDisabled(multiAddRecordList.length === 0);
  }, [multiAddRecordList]);

  // 选择文件后的操作，读取 Excel 文件中的数据
  const multiAddRecordProps = {
    name: 'file',
    // 通过下面两个参数共同限制只能上传一个文件
    multiple: false,
    maxCount: 1,
    fileList: multiAddFileList,
    beforeUpload(file) {
      // 由于不需要上传文件到服务器，所以需要阻止自动上传，并判断如果文件不是 xlsx 后缀，则报错
      if (!file.name.endsWith('.xlsx')) {
        message.error('请根据用户批量添加模板上传 Excel 文件');
      }
      return false;
    },
    // 文件列表变化时，读取文件，单文件不需要管 drop 事件，反正都会触发 change 事件
    onChange(info) { 
      // 获取文件列表，然后读取文件，如果文件列表为空，则清空列表
      setMultiAddFileList(info.fileList);
      if (info.fileList.length === 0) {
        setMultiAddRecordList([]);
        setMultiAddFileList([]);
        return;
      }
      // 使用 xlsx 读取文件
      const f = info.fileList[0]?.originFileObj;
      const reader = new FileReader();
      reader.readAsArrayBuffer(f);
      reader.onload = (e) => {
        const data = new Uint8Array(e.target.result);
        // 为了避免因为单元格格式问题导致和后端数据类型不一致的情况，所有字段都使用字符串上传
        const workbook = XLSX.read(data, { type: 'array', raw: false });
        // 获取第一个 sheet 的数据，其它的不读取
        const firstSheet = workbook.Sheets[workbook.SheetNames[0]];
        // 将 sheet 转换为 json 数据，并设置空单元格的默认值为空字符串
        const json = [XLSX.utils.sheet_to_json(firstSheet, { raw: false, defval: '' })];
        // 由于前 4 条数据都是模板数据，不需要，所以截取掉
        const jsonData = json[0].slice(4);
        setMultiAddRecordList(jsonData);
      };
    }
  };

  // 批量添加方法
  const multiAddRecordHandle = async () => {
    try {
      if (multiAddRecordList.length === 0) {
        message.error('未获取到导入数据，请检测文件格式是否正确');
        return;
      }
      const res = await AxiosPost(Apis.System.User.MultiAdd, multiAddRecordList);
      if (res.code === 200) {
        message.success('批量导入任务已创建，请稍后查看任务状态');
        setMultiAddRecordModalVisible(false);
      } else {
        message.error(res.message);
      }
    } catch (error) {
      message.error('服务异常，批量导入失败');
      console.error(error);
    }
    // 清理上传文件
    setMultiAddRecordModalVisible(false);
    setMultiAddFileList([]);
    setMultiAddRecordList([]);
  };

  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 批量导入历史
  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 批量添加历史任务记录弹窗
  const [multiAddRecordTaskHistoryModalVisible, setMultiAddRecordTaskHistoryModalVisible] = useState(false);

  // 批量导入任务表格列
  const multiAddRecordTaskHistoryColumns = [
    {title: '任务ID', dataIndex: 'taskId'},
    {title: '执行人', dataIndex: 'executor'},
    {title: '记录数', dataIndex: 'recordCount'},
    {title: '成功数', dataIndex: 'successCount'},
    {title: '失败数', dataIndex: 'failCount'},
    {title: '任务状态', dataIndex: 'status', render: (_, record) =>  generateTaskStatusTag(record.taskStatus)},
    {title: '任务开始时间', dataIndex: 'taskStartTime'},
    {title: '任务结束时间', dataIndex: 'taskEndTime'},
    {title: '操作', key: 'action', fixed: 'right', render: (_, record) => (<Space size="middle"><a><UnorderedListOutlined /> 查看明细</a></Space>)}
  ];

  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 批量操作
  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 批量操作的 ID 列表
  const [multiModifyRecordIds, setMultiModifyRecordIds] = useState('');

  // 批量操作的行选择
  const multiModifyRecordRowSelection = {
    selectedRowKeys: multiModifyRecordIds,
    onChange: (selectedRowKeys, selectedRows) => {
      setMultiModifyRecordIds(selectedRowKeys);
    },
    // 获取行选择框的属性，系统超级管理员账户不允许选择
    getCheckboxProps: (record) => ({ disabled: record.id === 1 })
  };

  // 批量操作的方法
  const multiModifyRecordHandle = async (data) => {
    try {
      const res = await AxiosPost(Apis.System.User.MultiModifyStatus, data);
      if (res.code === 200) {
        message.success('批量操作成功');
        setMultiModifyRecordIds([]);
        refreshCurrentPage();
      } else {
        message.error(res.message);
      }
    } catch (error) {
      message.error('服务异常，批量操作失败');
      console.error(error);
    }
  };

  // 批量操作菜单
  const multiModifyRecordMenuProps = {
    items: [
      {label: '批量禁用', key: 'disable', danger: true, disabled: multiModifyRecordIds.length === 0},
      {label: '批量启用', key: 'enable', danger: false, disabled: multiModifyRecordIds.length === 0}
    ],
    onClick: (key) => {
      if (multiModifyRecordIds.length > 0) {
        const data = {
          ids: multiModifyRecordIds,
          operate: key.key
        };
        multiModifyRecordHandle(data);
      }
    }
  };

  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 编辑用户
  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 编辑表单
  const [updateRecordForm] = Form.useForm();
  // 编辑弹窗
  const [updateRecordModalVisible, setUpdateRecordModalVisible] = useState(false);
  // 编辑数据
  const [updateRecord, setUpdateRecord] = useState(null);

  // 定义编辑数据的字段
  const updateRecordFields = [
    // 添加隐藏的 ID 字段
    {
      label: 'ID',
      name: 'id',
      type: 'input',
      value: updateRecord?.id,
      hidden: true
    },
    // 复用并扩展 recordFormBasicFields 的字段，增加回填数据
    ...recordFormBasicFields.map(field => ({
      ...field,
      value: field.name === 'systemDepartments' ? 
        updateRecord?.systemDepartments?.map(item => item.id) :
        field.name === 'systemJobPositions' ? 
        updateRecord?.systemJobPositions?.map(item => item.id) :
        field.name === 'systemRole' ? 
        updateRecord?.systemRole?.id :
        updateRecord?.[field.name]
    }))
  ];

  // 生成用户编辑表单
  const generateUpdateRecordFormItems = () => {
    return updateRecordFields?.map(item => generateRecordFormItem(item));
  };

  // 编辑用户方法
  const updateRecordHandle = async (data) => {
    try {
      const res = await AxiosPost(Apis.System.User.Update, data);
      if (res.code === 200) {
        message.success('编辑成功');
        setUpdateRecordModalVisible(false);
        refreshCurrentPage();
      } else {
        message.error(res.message);
      }
    } catch (error) {
      message.error('服务异常，编辑失败');
      console.error(error);
    }
  };

  // 监听 updateRecord 的变化，并在变化时重置表单
  useEffect(() => {
    if (updateRecord && updateRecordModalVisible) {
      updateRecordForm.setFieldsValue({
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
  }, [updateRecord, updateRecordModalVisible]);

  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 修改用户状态
  /////////////////////////////////////////////////////////////////////////////////////////////////////
  // 修改状态的方法
  const modifyRecordStatusHandle = async (recordId, operate) => {
    try {
      const res = await AxiosPost(Apis.System.User.ModifyStatus, { id: recordId, operate: operate });
      if (res.code === 200) {
        message.success('操作成功');
        refreshCurrentPage();
      } else {
        message.error(res.message);
      }
    } catch (error) {
      message.error('服务异常，修改状态失败');
      console.error(error);
    }
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
        {/* 搜索栏 */}
        <div className="admin-page-search">
          <Form form={filterRecordForm} labelCol={{ span: 6 }} wrapperCol={{ span: 18 }} colon={false} name="filterRecordForm" onFinish={filterRecordListHandle} autoComplete="off">
            <Row gutter={24}>
              {generateFilterRecordFormItems()}
              <Col span={24} key="x" style={{ marginTop: '10px', textAlign: 'right' }}>
                <Space>
                  <Button icon={<SearchOutlined />} htmlType="submit">条件搜索</Button>
                  <Button icon={<ClearOutlined />} onClick={() => filterRecordForm.resetFields()}>清理条件</Button>
                  {filterRecordFields.length > PageConfig.defaultFilterExpandItemCount && (
                    <a onClick={() => setExpandFilterRecordItems(!expandFilterRecordItems)}>
                      <DownOutlined rotate={expandFilterRecordItems ? 180 : 0} /> {expandFilterRecordItems ? '收起条件' : '展开更多'}
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
              <Button type="primary" icon={<UserAddOutlined />} onClick={() => setAddRecordModalVisible(true)} disabled={addRecordButtonDisabled}>添加{PageConfig.pageKeyword}</Button>
              <Button icon={<UploadOutlined />} onClick={() => setMultiAddRecordModalVisible(true)} disabled={multiAddRecordButtonDisabled}>批量导入</Button>
              <Dropdown menu={multiModifyRecordMenuProps} disabled={multiModifyRecordStatusButtonDisabled}>
                <Button><Space><DownOutlined /> 批量操作</Space></Button>
              </Dropdown>
            </Space>
            <Space style={{ float: 'right' }}>
              <Button icon={<DownloadOutlined />} onClick={() => { window.open('/template/用户批量添加模板.xlsx'); }}>模板下载</Button>
              <Button icon={<ClockCircleOutlined />} onClick={() => setMultiAddRecordTaskHistoryModalVisible(true)}>导入记录</Button>
            </Space>
          </div>
          <Table
            // 表格布局大小
            size="small"
            // 表格布局方式，支持 fixed、auto
            tableLayout="auto"
            // 表格行选择
            rowSelection={{type: 'checkbox', ...multiModifyRecordRowSelection}}
            // 表格列
            columns={tableColumns}
            // 表格展开信息
            expandable={{
              expandedRowRender: (record) => {
                const [name, cnName, enName, email] = record.creator.split(',');
                const items = [
                  {label: '用户创建者', children: `${cnName} / ${enName} (${name} / ${email})`},
                  {label: '个人介绍', children: record.description || '-'},
                  {label: '更新时间', children: record.updatedAt || '-'}, 
                  {label: '最后登录 IP', children: record.lastLoginIP || '-'},
                  {label: '最后登录时间', children: record.lastLoginTime || '-'}
                ];
                return <Descriptions column={1} items={items} />;
              },
              // 用于限制是否可以展开
              rowExpandable: (record) => record.name !== 'Not Expandable'
            }}
            dataSource={tableRecordList}
            // 行唯一标识
            rowKey="id"
            // 表格分页
            pagination={{
              pageSize: pageSize,
              current: pageNumber,
              total: total,
              showTotal: (total) => `总共 ${total} 条记录`,
              // 是否隐藏分页器，当只有一页时隐藏
              // hideOnSinglePage: true,
              // 是否允许修改显示数量
              showSizeChanger: true,
              // 是否显示快速跳转
              showQuickJumper: true,
              // 页码变化时触发
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
      <Modal title={'添加' + PageConfig.pageKeyword} open={addRecordModalVisible} onCancel={() => setAddRecordModalVisible(false)} width={400} maskClosable={false} footer={null}>
        <Form form={addRecordForm} labelCol={{ span: 6 }} wrapperCol={{ span: 18 }} colon={false} name="addRecordForm" onFinish={addRecordHandle} autoComplete="off">
          {generateAddRecordFormItems()}
          <Form.Item wrapperCol={{ span: 24 }}>
            <Button type="primary" htmlType="submit" block>添加{PageConfig.pageKeyword}</Button>
          </Form.Item>
        </Form>
      </Modal>

      {/* 用户编辑弹窗 */}
      <Modal title={'编辑' + PageConfig.pageKeyword} open={updateRecordModalVisible} onCancel={() => {
        setUpdateRecordModalVisible(false);
        updateRecordForm.resetFields(); // 关闭时重置表单，避免数据不更新
      }} width={400} maskClosable={false} footer={null}>
        <Form form={updateRecordForm} labelCol={{ span: 6 }} wrapperCol={{ span: 18 }} colon={false} name="updateRecordForm" onFinish={updateRecordHandle} autoComplete="off">
          {generateUpdateRecordFormItems()}
          <Form.Item wrapperCol={{ span: 24 }}>
            <Button type="primary" htmlType="submit" block>保存编辑</Button>
          </Form.Item>
        </Form>
      </Modal>

      {/* 批量导入弹窗 */}
      <Modal title={'批量导入' + PageConfig.pageKeyword} open={multiAddRecordModalVisible} onCancel={() => {
        setMultiAddRecordModalVisible(false);
        setMultiAddRecordList([]);
        setMultiAddFileList([]);
      }} width={800} maskClosable={false} footer={null}>
        <Dragger {...multiAddRecordProps}>
          <p className="ant-upload-drag-icon"><InboxOutlined /></p>
          <p className="ant-upload-text">点击或者拖拽批量导入模板文件到此区域完成创建</p>
          <p className="ant-upload-hint">为了数据安全，只支持单个文件模板导入，严格禁止上传公司数据或者其它违规文件。</p>
        </Dragger>
        <div style={{ marginTop: '10px' }}>
          <Button disabled={multiAddRecordSubmitButtonDisabled} type="primary" htmlType="submit" block onClick={multiAddRecordHandle}>批量导入</Button>
        </div>
      </Modal>

      {/* 批量导入记录 */}
      <Modal title={'批量导入任务'} open={multiAddRecordTaskHistoryModalVisible} onCancel={() => setMultiAddRecordTaskHistoryModalVisible(false)} width={1200} maskClosable={false} footer={null}>
        <Table
          // 表格布局大小
          size="small"
          // 表格布局方式，支持 fixed、auto
          tableLayout="auto"
          // 表格列
          columns={multiAddRecordTaskHistoryColumns}
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
