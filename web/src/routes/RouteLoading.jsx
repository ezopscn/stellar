import { Spin } from 'antd';

// 加载中的时候显示的界面
const RouteLoading = () => {
  return (
    <div className="admin-loading">
      <Spin size="large" />
    </div>
  );
};

export default RouteLoading;
