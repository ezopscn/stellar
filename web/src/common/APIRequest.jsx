import { GenerateTreeNode, GenerateSelectTree } from './GenerateTree';
import { AxiosGet } from './Request';
import { message } from 'antd';

// 消息提示
const { message } = App.useApp();

// API 请求方法
const APIRequest = {
  // 获取普通的列表数据
  GetDataList: async (api, setter, tree = false) => {
    try {
      const res = await AxiosGet(api);
      if (res.code === 200) {
        if (tree) {
          const treeData = GenerateTreeNode(res.data, 0);
          setter(treeData);
        } else {
          setter(res.data);
        }
      } else {
        message.error(res.message);
      }
    } catch (error) {
      console.log(`后端服务异常，接口请求失败：${api}`);
      console.log(error);
    }
  },
  // 获取用于填充 Select 的数据
  GetSelectDataList: async (api, setter, tree = false) => {
    try {
      const res = await AxiosGet(api);
      if (res.code === 200) {
        if (tree) {
          const treeData = GenerateSelectTree(res.data, 0);
          setter(treeData);
        } else {
          setter(
            res.data.map((item) => ({
              label: item.name,
              value: item.id
            }))
          );
        }
      } else {
        message.error(res.message);
      }
    } catch (error) {
      console.log(`后端服务异常，接口请求失败：${api}`);
      console.log(error);
    }
  }
};
