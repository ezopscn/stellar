import { Helmet } from 'react-helmet';
import { TitleSuffix } from '@/common/Text.jsx';

const Metrics = () => {
  const title = '监控指标' + TitleSuffix;
  return (
    <>
      <Helmet>
        <title>{title}</title>
        <meta name="description" content={title} />
      </Helmet>
      <h1>Metrics</h1>
    </>
  );
};

export default Metrics;