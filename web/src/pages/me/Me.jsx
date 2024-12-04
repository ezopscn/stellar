import { Helmet } from 'react-helmet';
import { TitleSuffix } from '@/common/Text.jsx';

const Me = () => {
  const title = '个人中心' + TitleSuffix;
  return (
    <>
      <Helmet>
        <title>{title}</title>
        <meta name="description" content={title} />
      </Helmet>
      <h1>Me</h1>
    </>
  );
};

export default Me;