import { BG404 } from '@/common/Image.jsx';

const NotFoundError = () => {
  return (
    <>
      <img className="admin-unselect" src={BG404} alt="" />
      <div className="admin-error-code admin-unselect">404</div>
    </>
  );
};

export default NotFoundError;