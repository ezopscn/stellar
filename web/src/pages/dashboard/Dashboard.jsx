import { Helmet } from 'react-helmet';
import { TitleSuffix } from '@/common/Text.jsx';

const Dashboard = () => {
  const title = '工作空间' + TitleSuffix;
  return (
    <>
      <Helmet>
        <title>{title}</title>
        <meta name="description" content={title} />
      </Helmet>
      <h1>Dashboard</h1>
    </>
  );
};

export default Dashboard;