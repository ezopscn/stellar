import { Helmet } from 'react-helmet';
import { TitleSuffix } from '@/common/Text.jsx';

const SecurityauditLogin = () => {
  const title = '登录日志' + TitleSuffix;
  return (
    <>
      <Helmet>
        <title>{title}</title>
        <meta name="description" content={title} />
      </Helmet>
      <h1>Hello</h1>
    </>
  );
};

export default SecurityauditLogin;