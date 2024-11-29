import { Suspense } from 'react';
import RouteLoading from '@/routes/RouteLoading.jsx';

// 惰性加载，避免路由加载的时候所有页面都会被加载
const RouteLazyLoad = (Component) => (
  <Suspense fallback={<RouteLoading />}>
    <Component />
  </Suspense>
);

export default RouteLazyLoad;