// 路由匹配
import { useEffect } from 'react';
import { RouteRules } from '@/routes/RouteRules.jsx';
import { useLocation, useNavigate } from 'react-router';
import { GetToken } from '@/utils/Token.jsx';
import { AxiosGet } from '@/utils/Request.jsx';
import { Apis } from '@/common/APIConfig.jsx';

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

// 路由守卫，认证拦截
const RouteGuard = ({ children }) => {
  const location = useLocation();
  const navigator = useNavigate();

  useEffect(() => {
    const matchRoute = RouteMatch(location.pathname, RouteRules);
    if (!matchRoute) {
      navigator('/error/404'); // 验证路由是否存在
    } else {
      if (matchRoute.auth && !GetToken()) {
        // 需要登录但是未登录
        navigator('/login');
      } else if (matchRoute.auth) {
        // 需要登录且已登录，则校验 Token 是否过期
        AxiosGet(Apis.Public.TokenVerification).then((res) => {
          if (res.code !== 200) {
            navigator('/login');
          }
        });
      } else if (location.pathname === '/login' && GetToken()) {
        // 登录页面还有 Token，则清理 localStorage
        localStorage.clear();
      }
      // 验证路由是否在用户权限范围内，不再则返回 403 页面
    }
  }, [location.pathname]);

  return children;
};

export default RouteGuard;
