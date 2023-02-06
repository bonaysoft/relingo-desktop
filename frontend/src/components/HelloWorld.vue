<script lang="ts" setup>
import { ref, reactive, onMounted, watch } from 'vue'
import { FindNewWords, SubmitVocabulary } from '../../wailsjs/go/main/App'
import { model } from '../../wailsjs/go/models';

const words = ref<any>([])
const currentTab = ref<any>()
const refresh = () => {
  console.log(currentTab.value, 111);


  FindNewWords(currentTab.value).then(result => {
    words.value = result.sort((a: model.Word, b: model.Word) => b.exposures - a.exposures)
  })
}

const submitVocabulary = (words: string[]) => {
  SubmitVocabulary(words).then(refresh)
}

onMounted(refresh)
watch(currentTab, refresh)

</script>

<template>
  <main>
    <v-tabs bg-color="primary" v-model="currentTab">
      <v-tab value="today">今日生词</v-tab>
      <v-tab value="yesterday">昨日生词</v-tab>
      <v-tab value="week">本周生词</v-tab>
      <v-tab value="all">全部生词</v-tab>
    </v-tabs>
    {{ words.length }}

    <v-table density="compact">
      <thead>
        <tr>
          <th class="text-left">
            单词
          </th>
          <th class="text-center">
            累计出现次数
          </th>
          <th class="text-left" style="width: 180px">
            最近出现
          </th>
          <th class="text-left" style="width: 180px">
            最初出现
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in words" :key="item.name">
          <td class="text-left">{{ item.source }}
            <v-btn size="x-small" @click="submitVocabulary([item.source])">掌握</v-btn>
          </td>
          <td class="text-center">{{ item.exposures }}</td>
          <td class="text-left">{{ $dayjs(item.updated_at).format("YYYY-MM-DD HH:mm:ss") }}</td>
          <td class="text-left">{{ $dayjs(item.created_at).format("YYYY-MM-DD HH:mm:ss") }}</td>
        </tr>
      </tbody>
    </v-table>
  </main>
</template>

<style scoped>

</style>
