// 通过子菜单 path 获取父级菜单列表
export function TreeFindPath(tree, func, path = []) {
  if (!tree) return [];
  for (const data of tree) {
    path.push(data.path);
    if (func(data)) return path;
    if (data.children) {
      const findChildren = TreeFindPath(data.children, func, path);
      if (findChildren.length) return findChildren;
    }
    path.pop();
  }
  return [];
}
