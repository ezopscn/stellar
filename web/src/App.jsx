import React from 'react';
import { App } from 'antd';
import { HashRouter } from 'react-router-dom';
import { Routes } from '@/routes/RouteRules.jsx';
import RouteGuard from '@/routes/RouteGuard.jsx';

const MainApp = () => {
  return (
    <App>
      <HashRouter>
        <RouteGuard>
          <Routes />
        </RouteGuard>
      </HashRouter>
    </App>
  );
};

export default MainApp;
