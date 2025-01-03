import { Helmet } from 'react-helmet';
import { TitleSuffix } from '@/common/Text.jsx';

const SecurityauditOperation = () => {
  const title = '操作日志' + TitleSuffix;
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

export default SecurityauditOperation;