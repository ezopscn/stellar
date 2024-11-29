// 路由匹配
import { useEffect } from 'react';
import { RouteRules } from '@/routes/RouteRules.jsx';
import { useLocation, useNavigate } from 'react-router';

// 路由匹配
const RouteMatch = (path, routes) => {
  for (const route of routes) {
    if (route.path === path) {
      return route;
    }
    if (route.children) {
      const childRoute = RouteMatch(path, route.children);
      if (childRoute) {
        return childRoute;
      }
    }
  }
  return null; // 如果没有匹配的路由，返回 null
};

// 路由守卫
const RouteGuard = ({ children }) => {
  // 获取当前路由信息
  const location = useLocation();
  // 路由跳转
  const navigator = useNavigate();
  // 监听路由变化，执行路由验证
  useEffect(() => {
    // 判断路由数据是否存在，不存在就返回 404
    if (RouteMatch(location.pathname, RouteRules) === null) {
      navigator('/error/404');
    }

    // 还可以进行登录认证等判断
  }, [location.pathname]);
  return children;
};

export default RouteGuard;
