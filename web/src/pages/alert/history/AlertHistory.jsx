const AlertHistory = () => {
    return (
      <>
        <div className="admin-page-header admin-unselect">
          <div className="admin-page-title">历史告警 / HISTORY ALERT.</div>
          <div className="admin-page-desc">
            <ul>
              <li>告警历史是指告警在一定时间内被处理和解决的频率。</li>
              <li>告警历史越高，说明告警处理和解决的频率越高，告警的严重程度越高。</li>
            </ul>
          </div>
        </div>
        <div className="admin-page-main"></div>
      </>
    );
  };
  
  export default AlertHistory;
