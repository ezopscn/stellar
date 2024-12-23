import { Helmet } from 'react-helmet';
import { TitleSuffix } from '@/common/Text.jsx';

const Query = () => {
  const title = '即时查询' + TitleSuffix;
  return (
    <>
      <Helmet>
        <title>{title}</title>
        <meta name="description" content={title} />
      </Helmet>
      <h1>Query</h1>
    </>
  );
};

export default Query;