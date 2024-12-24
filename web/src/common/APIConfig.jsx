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
      Add: BackendURL + '/system/user/add', // 添加用户
      MutiAdd: BackendURL + '/system/user/muti-add', // 批量添加用户
      StatusModify: BackendURL + '/system/user/status-modify', // 修改用户状态
    },
    Menu: {
      Tree: BackendURL + '/system/menu/tree' // 菜单树
    },
    Role: {
      List: BackendURL + '/system/role/list' // 角色列表
    },
    JobPosition: {
      List: BackendURL + '/system/job-position/list' // 岗位列表
    },
    Department: {
      List: BackendURL + '/system/department/list' // 部门列表
    }
  }
};

export { BackendURL, Apis };
