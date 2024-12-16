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

export { GenerateSelectTree, GenerateTreeNode };
