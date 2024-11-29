import { Outlet } from 'react-router';
import { Logo } from '@/common/Image.jsx';
import { FooterText } from '@/common/Text.jsx';

const LoginLayout = () => {
  return (
    <>
      <div className="admin-login-container">
        <div className="admin-login-box">
          <div className="admin-login-header">
            <img className="admin-unselect" src={Logo} alt="logo" />
          </div>
          <Outlet />
          <div className="admin-login-footer admin-unselect">
            <FooterText />
          </div>
        </div>
      </div>
    </>
  );
};

export default LoginLayout;