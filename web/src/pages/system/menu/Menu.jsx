const Menu = () => {
  return (
    <>
      <div className="admin-page-header admin-unselect">
        <div className="admin-page-title">系统菜单 / MENU MANAGEMENT.</div>
        <div className="admin-page-desc">
          <ul>
            <li>出于数据安全考虑，系统强制使用禁用用户替代删除用户。</li>
            <li>对于某些特殊的用户，例如老板或者高管，我们建议隐藏其联系方式，保护个人隐私。</li>
          </ul>
        </div>
      </div>
      <div className="admin-page-main"></div>
    </>
  );
};

export default Menu;