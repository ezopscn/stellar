import { Helmet } from 'react-helmet';
import { TitleSuffix } from '@/common/Text.jsx';

const Datasource = () => {
  const title = '数据源' + TitleSuffix;
  return (
    <>
      <Helmet>
        <title>{title}</title>
        <meta name="description" content={title} />
      </Helmet>
      <h1>Datasources</h1>
    </>
  );
};

export default Datasource;