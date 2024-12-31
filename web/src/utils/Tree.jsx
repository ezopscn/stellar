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
const GenerateDepartmentTree = (dataList = [], parentId = null) => {
  return dataList
    .filter(item => item.parentId === parentId)
    .map(item => {
      const node = {
        ...item,
        title: item.name,
        key: item.id.toString()
      };
      
      const children = GenerateDepartmentTree(dataList, item.id);
      if (children.length > 0) {
        node.children = children;
      }
      
      return node;
    });
};

export { GenerateSelectTree, GenerateTreeNode, GenerateDepartmentTree };
