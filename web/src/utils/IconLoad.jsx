import {
  ApiOutlined,
  ConsoleSqlOutlined,
  DesktopOutlined,
  SettingOutlined,
  UsergroupAddOutlined,
  InsuranceOutlined,
  UserOutlined,
  FileProtectOutlined,
  ClusterOutlined,
  AlertOutlined,
  MailOutlined,
  ProjectOutlined
} from '@ant-design/icons';

// 图标字符串映射
// eslint-disable-next-line react-refresh/only-export-components
export const IconMap = {
  DesktopOutlined: DesktopOutlined,
  ConsoleSqlOutlined: ConsoleSqlOutlined,
  ApiOutlined: ApiOutlined,
  SettingOutlined: SettingOutlined,
  UsergroupAddOutlined: UsergroupAddOutlined,
  InsuranceOutlined: InsuranceOutlined,
  UserOutlined: UserOutlined,
  FileProtectOutlined: FileProtectOutlined,
  ClusterOutlined: ClusterOutlined,
  AlertOutlined: AlertOutlined,
  MailOutlined: MailOutlined,
  ProjectOutlined: ProjectOutlined
};

// 生成 Icon
// 用法：<DynamicIcon iconName={'DesktopOutlined'} />
// eslint-disable-next-line react/prop-types
export const DynamicIcon = ({ iconName }) => {
  const IconComponent = IconMap[iconName];
  return IconComponent ? <IconComponent /> : null;
};
