import { BG403 } from '@/common/Image.jsx';

const ForbiddenError = () => {
  return (
    <>
      <img className="admin-unselect" src={BG403} alt="" />
      <div className="admin-error-code admin-unselect">403</div>
    </>
  );
};

export default ForbiddenError;