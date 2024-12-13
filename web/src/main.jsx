import { createRoot } from 'react-dom/client';
import MainApp from '@/App.jsx';
// 由于 antd 组件的默认文案是英文，所以需要修改为中文
import { ConfigProvider } from 'antd';
// 中文支持
import zhCN from 'antd/es/locale/zh_CN';
// 时间选择器中文支持
import 'dayjs/locale/zh-cn';
// 小米字体
import 'misans/lib/Normal/MiSans-Regular.min.css';
// antd 样式覆盖和其它样式
import '@/assets/css/AntdRewrite.less';
import '@/assets/css/Admin.less';

createRoot(document.getElementById('root')).render(
  // autoInsertSpaceInButton：解决按钮的文本为两个汉字时中间自动补充空格的问题，新版本已经废弃
  // theme：修改默认主题，参考文档：https://ant-design.antgroup.com/docs/react/customize-theme-cn
  <ConfigProvider
    locale={zhCN}
    theme={{
      hashed: false, // 禁用 css hash
      token: {
        colorLink: '#003399', // 链接颜色
        colorPrimary: '#001529', // 主色调
        colorPrimaryHover: '#001529', // 主色调 hover
        borderRadius: 0, // 圆角
        fontFamily: 'MiSans, serif', // 文字字体
        fontSize: 13, // 默认字号
        controlHeight: 28 // 按钮和输入框等基础控件的高度
      }
    }}
  >
    <MainApp />
  </ConfigProvider>
);
