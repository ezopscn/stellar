import { Form } from 'antd';
import { FORM_ITEM_COMPONENT_MAP } from '@/common/GlobalConfig.jsx';

// 传入元素，生成表单项通用方法
const GenerateRecordFormItem = (item) => {
  const commonProps = {
    allowClear: true,
    placeholder: item.placeholder,
    hidden: item.hidden,
    value: item.value,
    disabled: item.disabled
  };
  const componentProps = {
    input: {},
    inputPassword: {},
    textarea: {},
    treeSelect: {
      treeData: item.data,
      showSearch: item.search,
      treeNodeFilterProp: 'label',
      allowClear: false,
      multiple: item.multiple
    },
    select: {
      options: item.data,
      showSearch: item.search,
      optionFilterProp: 'label',
      allowClear: false,
      mode: item.multiple ? 'multiple' : 'default'
    }
  };
  return (
    <Form.Item key={item.name} name={item.name} label={item.label} rules={item.rules} hidden={item.hidden}>
      {FORM_ITEM_COMPONENT_MAP[item.type]?.({
        ...commonProps,
        ...componentProps[item.type]
      })}
    </Form.Item>
  );
};

export { GenerateRecordFormItem };
