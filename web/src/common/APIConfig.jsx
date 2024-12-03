// 后端地址
const BackendURL = window.CONFIG.backendUrl;

// 接口数据
const Apis = {
  RunEnv: window.CONFIG.env, // 运行环境
  BackendURL: BackendURL, // 后端基础接口
  Public: {
    Version: BackendURL + '/version', // 版本信息
    Login: BackendURL + '/login', // 用户登录
    Logout: BackendURL + '/logout', // 注销登录
    TokenVerification: BackendURL + '/token/verification' // Token 校验
  },
  Current: {
    User: {
      Menu: {
        Tree: BackendURL + '/current/user/menu/tree' // 当前用户菜单树
      }
    }
  },
  System: {
    User: {
      List: BackendURL + '/system/user/list', // 用户列表
      Info: BackendURL + '/system/user/detail', // 当前用户详情
      Specified: {
        DetailById: BackendURL + '/system/user/specified/detail/by/id/', // 指定 ID 用户详情
        DetailByUsername: BackendURL + '/system/user/specified/detail/by/username/', // 指定 Username 用户详情
        DetailByPhone: BackendURL + '/system/user/specified/detail/by/phone/', // 指定 Phone 用户详情
        DetailByEmail: BackendURL + '/system/user/specified/detail/by/email/' // 指定 Email 用户详情
      }
    },
    Menu: {
      Tree: BackendURL + '/system/menu/tree' // 菜单树
    }
  }
};

export { BackendURL, Apis };