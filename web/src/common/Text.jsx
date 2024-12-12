import React, { useEffect, useState } from 'react';
import { AxiosGet } from '@/utils/Request.jsx';
import { Apis } from '@/common/APIConfig.jsx';
import { GithubOutlined } from '@ant-design/icons';

// 底部内容
const FooterText = () => {
  // 运行环境
  const runEnv = window.CONFIG.env;
  const runEnvText = 'Running Env: ' + runEnv;
  // 版本信息
  const [versionText, setVersionText] = useState('unknown');
  useEffect(() => {
    const getVersion = async () => {
      const version = localStorage.getItem('version');
      if (!version) {
        try {
          const res = await AxiosGet(Apis.Public.Version);
          if (res.code === 200) {
            const { systemVersion } = res.data;
            setVersionText(systemVersion);
            localStorage.setItem('version', systemVersion);
          } else {
            console.error('获取后端版本信息失败: ', res.message);
          }
        } catch (error) {
          console.error('获取后端版本信息失败: ', error);
        }
      } else {
        setVersionText(version);
      }
    };
    getVersion();
  }, []);

  return (
    <>
      <b>👻 STELLAR </b>© 2024 EZOPS.CN. Current Version: {versionText} / Latest Version:{' '}
      <a href="https://github.com/ezopscn/stellar/releases" target="_blank" rel="noreferrer">
        <GithubOutlined />
      </a>{' '}
      / {runEnvText}
    </>
  );
};

// Title
const TitleDesc = 'Stellar is a middleware tool that converts data warehouse data into Prometheus metrics';
const TitleSuffix = ' | ' + TitleDesc;

export { FooterText, TitleDesc, TitleSuffix };
