import { Input, TreeSelect, Select } from 'antd';

// 数据状态映射
const DATA_STATUS_MAP = [{ label: '启用', value: 1 }, { label: '禁用', value: 0 }];

// 性别映射
const GENDER_MAP = [{ label: '男', value: 1 }, { label: '女', value: 2 }, { label: '未知', value: 3 }];

// 是否映射
const TRUE_FALSE_MAP = [{ label: '是', value: 1 }, { label: '否', value: 0 }];

// 组件映射
const FORM_ITEM_COMPONENT_MAP = {
  // 输入框
  input: (props) => <Input {...props} autoComplete="off" />, 
  // 密码输入框
  inputPassword: (props) => <Input.Password {...props} autoComplete="off" />,
  // 文本域
  textarea: (props) => <Input.TextArea {...props} autoComplete="off" />,
  // 树形选择框
  treeSelect: (props) => <TreeSelect {...props} treeDefaultExpandAll />, 
  // 下拉选择框
  select: (props) => <Select {...props} />,
};

// 节点角色映射
const NODE_ROLE_MAP = {
  'Leader': { text: 'Leader', color: 'red' },
  'Worker': { text: 'Worker', color: 'blue' },
  'WebServer': { text: 'WebServer', color: 'green' }
};

// 需要禁止选择的角色 ID 列表
const DISABLED_ROLE_IDS = [1];

export { 
  DATA_STATUS_MAP, 
  GENDER_MAP, 
  FORM_ITEM_COMPONENT_MAP, 
  TRUE_FALSE_MAP, 
  NODE_ROLE_MAP, 
  DISABLED_ROLE_IDS 
};
