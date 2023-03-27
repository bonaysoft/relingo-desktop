<script lang="ts" setup>
import { onMounted, ref } from "vue";
import { Ready, GetUserInfo } from '@/../wailsjs/go/relingo/Client';
import { DownloadCert } from '@/../wailsjs/go/service/System';
const refresh = () => {
    location.reload();
}

const myInterval = ref()
const detectReady = async () => {
    const isReady = await Ready()
    if (isReady) {
        clearInterval(myInterval.value);
        refresh();
    }
}

onMounted(async () => {
    const isReady = await Ready()
    if (isReady) {
        return
    }

    myInterval.value = setInterval(detectReady, 1000)
})
</script>

<template>
    <div style="text-align: center;">
        <p>未检测到relingo请求，请先启用relingo插件后进行一次网页刷新。</p>
        <p>第一次使用请先获取证书并设置信任</p>

        <br />
        <v-btn @click="DownloadCert">获取证书</v-btn>
    </div>
</template>