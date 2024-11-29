import { useEffect, useState } from 'react';
import { Avatar, Badge, Dropdown, Layout, Menu, Select } from 'antd';
import { DefaultAvatar, Logo } from '@/common/Image.jsx';
import { Outlet, useLocation, useNavigate } from 'react-router';
import { FooterText } from '@/common/Text.jsx';
import { TreeFindPath } from '@/utils/Path.jsx';
import { RouteRules } from '@/routes/RouteRules.jsx';
import { DynamicIcon } from '@/utils/IconLoad.jsx';

const { Header, Content, Sider } = Layout;

// 生成菜单
function getItem(label, key, icon, children) {
  return {
    key,
    icon,
    children,
    label
  };
}

// 侧边菜单
const siderMenus = [
  getItem('工作空间', '/dashboard', <DynamicIcon iconName={'DesktopOutlined'} />),
  getItem('指标配置', '/metrics', <DynamicIcon iconName={'ConsoleSqlOutlined'} />),
  getItem('数据源', '/datasources', <DynamicIcon iconName={'ApiOutlined'} />),
  getItem('个人中心', '/usercenter', <DynamicIcon iconName={'UserOutlined'} />),
  getItem('系统设置', '/system', <DynamicIcon iconName={'ClusterOutlined'} />, [
    getItem('用户中心', '/system/users'),
    getItem('用户角色', '/system/roles'),
    getItem('系统菜单', '/system/menus'),
    getItem('系统接口', '/system/apis'),
    getItem('权限配置', '/system/permissions'),
  ]),
  getItem('安全审计', '/securityaudit', <DynamicIcon iconName={'FileProtectOutlined'} />, [
    getItem('登录日志', '/securityaudit/login'),
    getItem('操作日志', '/securityaudit/operation'),
  ])
];

// 下拉菜单
const dropdownMenus = [
  {
    label: 'DK / 吴彦祖',
    key: '0',
    disabled: true
  },
  {
    label: (
      <a target="_blank">
        消息中心<Badge size="small" count={5}></Badge>
      </a>
    ),
    key: '1'
  },
  {
    label: (
      <a target="_blank">
        个人资料
      </a>
    ),
    key: '2'
  },
  {
    type: 'divider'
  },
  {
    label: '注销登录',
    key: '3'
  }
];

const AdminLayout = () => {
  // 菜单跳转
  const navigate = useNavigate();
  // 菜单展开收起状态
  const [collapsed, setCollapsed] = useState(false);
  // 展开和收缩菜单宽度
  const menuWidth = 200;
  const menuCollapsedWidth = 60;

  // 获取当前的请求路径，并监听该路径是否改变，如果改变则修改页面菜单数据
  const { pathname } = useLocation(); // 当前页面
  const [openKeys, setOpenKeys] = useState([pathname]); // 展开菜单，父级菜单
  const [selectedKeys, setSelectedKeys] = useState([pathname]); // 选中菜单
  useEffect(() => {
    setOpenKeys(TreeFindPath(RouteRules, data => data.path === pathname));
    setSelectedKeys(pathname);
  }, [pathname]);

  return (
    <Layout>
      <Header className="admin-header">
        <div className="admin-left">
          <div className="admin-logo">
            <img className="admin-unselect" src={Logo} alt="" />
          </div>
        </div>
        <div className="admin-right">
          <Badge size="small" count={5}>
            <Dropdown menu={{
              items: dropdownMenus
            }}>
              <Avatar shape="circle" size={30}
                      src={DefaultAvatar} />
            </Dropdown>
          </Badge>
        </div>
      </Header>
      <Layout className="admin-main">
        <Sider className="admin-sider" theme="light" width={menuWidth} collapsedWidth={menuCollapsedWidth} collapsible
               collapsed={collapsed} onCollapse={(value) => setCollapsed(value)}>
          <Menu
            className="admin-menu"
            mode="inline"
            openKeys={openKeys}
            onOpenChange={(key) => {
              setOpenKeys(key); // 解决展开菜单问题
            }}
            selectedKeys={selectedKeys}
            items={siderMenus}
            onClick={({ key }) => {
              console.log(key);
              navigate(key);
            }}
          />
        </Sider>
        <Layout className="admin-body">
          <Content className="admin-content">
            <Outlet />
          </Content>
          <div className="admin-footer">
            <FooterText />
          </div>
        </Layout>
      </Layout>
    </Layout>
  );
};

export default AdminLayout;