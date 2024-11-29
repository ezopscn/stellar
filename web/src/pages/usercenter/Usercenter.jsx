import { Helmet } from 'react-helmet';
import { TitleSuffix } from '@/common/Text.jsx';

const Usercenter = () => {
  const title = '个人中心' + TitleSuffix;
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

export default Usercenter;