import { Outlet, useNavigate } from 'react-router';
import { Logo } from '@/common/Image.jsx';
import { FooterText } from '@/common/Text.jsx';
import { Button } from 'antd';
import { ArrowLeftOutlined } from '@ant-design/icons';

const ErrorLayout = () => {
  const navigate = useNavigate();
  return (
    <>
      <div className="admin-error">
        <div className="admin-error-header">
          <img className="admin-unselect" src={Logo} alt="logo" />
        </div>
        <div className="admin-error-body">
          <Outlet />
          <Button type="primary" icon={<ArrowLeftOutlined />} onClick={() => navigate('/')}>回到首页</Button>
        </div>
        <div className="admin-error-footer admin-unselect">
          <FooterText />
        </div>
      </div>
    </>
  );
};

export default ErrorLayout;