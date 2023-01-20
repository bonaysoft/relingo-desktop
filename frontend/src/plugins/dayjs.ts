import dayjs from 'dayjs';
import relativeTime from "dayjs/plugin/relativeTime";
import "dayjs/locale/zh-cn";
import { App } from "vue";

dayjs.locale("zh");
dayjs.extend(relativeTime);

declare module "@vue/runtime-core" {
  interface ComponentCustomProperties {
    $dayjs(date?: dayjs.ConfigType): dayjs.Dayjs;
  }
}

export default {
  install: (app: App<Element>): void => {
    app.config.globalProperties.$dayjs = dayjs;
  },
};