import { Tag } from 'antd';
import { ManOutlined, WomanOutlined, QuestionOutlined } from '@ant-design/icons';

// 生成性别图标
const GenerateGenderIcon = (gender) => {
  const genderIcons = {
    1: <ManOutlined style={{ color: '#165dff' }} />,
    2: <WomanOutlined style={{ color: '#ff4d4f' }} />,
    default: <QuestionOutlined style={{ color: '#999999' }} />
  };
  return genderIcons[gender] || genderIcons.default;
};

// 生成角色标签
const GenerateRoleTag = (name) => {
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
};

// 生成状态标签
const GenerateStatusTag = (status) => {
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
};

// 生成任务状态标签
const GenerateTaskStatusTag = (status) => {
  const statusMap = {
    1: { text: '执行中', color: 'gray' },
    2: { text: '执行成功', color: 'green' },
    3: { text: '执行失败', color: 'red' }
  };
  const { text, color } = statusMap[status] || {};
  return (
    <Tag bordered={false} color={color}>
      {text}
    </Tag>
  );
};

export { 
    GenerateGenderIcon,
    GenerateRoleTag, 
    GenerateStatusTag, 
    GenerateTaskStatusTag 
};
