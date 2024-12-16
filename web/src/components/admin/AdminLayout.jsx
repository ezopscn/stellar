import { useEffect, useState } from 'react';
import { Avatar, Badge, Dropdown, Layout, Menu, App } from 'antd';
import { Logo } from '@/common/Image.jsx';
import { Outlet, useLocation, useNavigate } from 'react-router';
import { FooterText } from '@/common/Text.jsx';
import { TreeFindPath } from '@/utils/Path.jsx';
import { RouteRules } from '@/routes/RouteRules.jsx';
import { DynamicIcon } from '@/utils/IconLoad.jsx';
import { AxiosGet } from '@/utils/Request.jsx';
import { Apis } from '@/common/APIConfig.jsx';
import { jwtDecode } from 'jwt-decode';
import { ManOutlined, WomanOutlined, QuestionOutlined } from '@ant-design/icons';

const { Header, Content, Sider } = Layout;

// 生成菜单
const getItem = (label, key, icon, children) => {
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
  getItem('个人中心', '/me', <DynamicIcon iconName={'UserOutlined'} />),
  getItem('系统设置', '/system', <DynamicIcon iconName={'ClusterOutlined'} />, [
    getItem('组织架构', '/system/departments'),
    getItem('职位管理', '/system/jobpositions'),
    getItem('用户管理', '/system/users'),
    getItem('用户角色', '/system/roles'),
    getItem('系统菜单', '/system/menus'),
    getItem('系统接口', '/system/apis'),
    getItem('权限配置', '/system/permissions')
  ]),
  getItem('安全审计', '/securityaudit', <DynamicIcon iconName={'FileProtectOutlined'} />, [getItem('登录日志', '/securityaudit/login'), getItem('操作日志', '/securityaudit/operation')])
];

const AdminLayout = () => {
  // 消息提示
  const { message } = App.useApp();
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
    setOpenKeys(TreeFindPath(RouteRules, (data) => data.path === pathname));
    setSelectedKeys(pathname);
  }, [pathname]);

  // 用户注销
  const logoutHandler = async () => {
    const res = await AxiosGet(Apis.Public.Logout);
    if (res.code === 200) {
      localStorage.clear();
      message.success('注销成功');
      navigate('/login');
    } else {
      message.error('注销异常，' + res.message);
    }
  };

  // 从 Token 中解析出用户相关信息
  const token = localStorage.getItem('token');
  const userInfo = jwtDecode(token);

  // 下拉菜单
  const dropdownMenus = [
    {
      label: `${userInfo?.cnName} / ${userInfo?.enName}`,
      key: '0',
      disabled: true
    },
    {
      label: <a target="_blank">修改资料</a>,
      key: '1',
      onClick: () => {
        navigate('/me');
      }
    },
    {
      type: 'divider'
    },
    {
      label: '注销登录',
      key: '2',
      onClick: logoutHandler
    }
  ];

  return (
    <Layout>
      <Header className="admin-header">
        <div className="admin-left">
          <div className="admin-logo">
            <img className="admin-unselect" src={Logo} alt="" />
          </div>
        </div>
        <div className="admin-right">
          <Badge
            count={
              userInfo?.gender === 1 ? (
                <ManOutlined style={{ backgroundColor: '#165dff' }} />
              ) : userInfo?.gender === 2 ? (
                <WomanOutlined style={{ backgroundColor: '#ff4d4f' }} />
              ) : (
                <QuestionOutlined style={{ backgroundColor: '#999' }} />
              )
            }
          >
            <Dropdown
              menu={{
                items: dropdownMenus
              }}
            >
              <Avatar shape="circle" size={30} src={userInfo?.avatar} />
            </Dropdown>
          </Badge>
        </div>
      </Header>
      <Layout className="admin-main">
        <Sider className="admin-sider" theme="light" width={menuWidth} collapsedWidth={menuCollapsedWidth} collapsible collapsed={collapsed} onCollapse={(value) => setCollapsed(value)}>
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
