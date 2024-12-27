import { Input, TreeSelect, Select } from 'antd';

// 用户状态映射
const SYSTEM_USER_STATUS_MAP = [{ label: '启用', value: 1 }, { label: '禁用', value: 0 }];

// 用户性别映射
const SYSTEM_USER_GENDER_MAP = [{ label: '男', value: 1 }, { label: '女', value: 2 }, { label: '未知', value: 3 }];

// 组件映射
const FORM_ITEM_COMPONENT_MAP = {
  // 输入框
  input: (props) => <Input {...props} />, 
  // 密码输入框
  inputPassword: (props) => <Input.Password {...props} />,
  // 文本域
  textarea: (props) => <Input.TextArea {...props} />,
  // 树形选择框
  treeSelect: (props) => <TreeSelect {...props} treeDefaultExpandAll />, 
  // 下拉选择框
  select: (props) => <Select {...props} />,
};

// 是否映射
const TRUE_FALSE_MAP = [{ label: '是', value: 1 }, { label: '否', value: 0 }];

export { SYSTEM_USER_STATUS_MAP, SYSTEM_USER_GENDER_MAP, FORM_ITEM_COMPONENT_MAP, TRUE_FALSE_MAP };
