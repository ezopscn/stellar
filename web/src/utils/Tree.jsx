// 递归生成 TreeSelect
const GenerateSelectTree = (dataList, parentId) => {
  const tree = [];
  for (const item of dataList) {
    if (item.parentId === parentId) {
      // 只保留需要的字段
      const node = {
        label: item.name,
        value: item.id,
        children: GenerateSelectTree(dataList, item.id)
      };

      // 如果没有子节点，则不包含 children 字段
      if (node.children.length === 0) {
        delete node.children;
      }

      tree.push(node);
    }
  }
  return tree;
}

// 递归生成 TreeNode
const GenerateTreeNode = (dataList, parentId) => {
  const tree = [];
  for (const item of dataList) {
    if (item.parentId === parentId) {
      const node = { ...item };
      node.children = GenerateTreeNode(dataList, item.id);

      if (node.children.length === 0) {
        delete node.children;
      }

      tree.push(node);
    }
  }
  return tree;
}

// 递归生成部门树
const GenerateSystemDepartmentTree = (dataList = [], parentId = null) => {
  return dataList
    .filter(item => item.parentId === parentId)
    .map(item => {
      const node = {
        ...item,
        title: item.name,
        key: item.id.toString()
      };
      
      const children = GenerateSystemDepartmentTree(dataList, item.id);
      if (children.length > 0) {
        node.children = children;
      }
      
      return node;
    });
};

// 转换树形数据 name-id 为 label-value 格式
const ConvertNameIdToLabelValueTree = (treeData) => {
  return treeData.map(node => ({
    label: node.name,
    value: node.id,
    children: node.children ? ConvertNameIdToLabelValueTree(node.children) : []
  }));
};

// 转换树形数据 name-id 为 title-key 格式
const ConvertNameIdToTitleKeyTree = (treeData) => {
  return treeData.map(node => ({
    key: node.id,
    title: node.name,
    children: node.children ? ConvertNameIdToTitleKeyTree(node.children) : []
  }));
};

// 获取树形结构所有节点的 key
const GetExpandedAllTreeKeys = (data) => {
  let keys = [];
  if (!Array.isArray(data)) return keys;
  data.forEach((item) => {
    if (item?.id) keys.push(item.id);
    if (item?.children?.length > 0) {
      keys = keys.concat(GetExpandedAllTreeKeys(item.children));
    }
  });
  return keys;
};

// 检查是否存在子集
const HasChildren = (treeData, id) => {
  const findNode = (nodes) => {
    for (const node of nodes) {
      if (node.key === id) {
        // 直接返回是否有子节点
        return Boolean(node.children?.length);
      }
      if (node.children) {
        const result = findNode(node.children);
        if (result !== undefined) return result;
      }
    }
    return undefined; // 明确返回 undefined 表示未找到节点
  };
  return findNode(treeData) || false; // 如果未找到节点，返回 false
};


export { GenerateSelectTree, GenerateTreeNode, GenerateSystemDepartmentTree, ConvertNameIdToLabelValueTree, ConvertNameIdToTitleKeyTree, GetExpandedAllTreeKeys, HasChildren };
