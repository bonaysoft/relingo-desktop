<script lang="ts" setup>
import { ref, reactive, onMounted } from 'vue'
import { FindNewWords } from '../../wailsjs/go/main/App'
import { GetUserInfo, GetVocabularyList, GetVocabulary } from '@/../wailsjs/go/relingo/Client';
import { model } from '../../wailsjs/go/models';

const words = ref<any>([])
const refresh = () => {

  GetUserInfo().then(uInfo => {
    console.log(uInfo)
  })

  GetVocabularyList().then(vocabularyList => {
    console.log(vocabularyList);

    const mastered = vocabularyList.find(el => el.type == 'mastered')
    console.log(mastered)

    GetVocabulary('relingo-A1', 'buildin').then(el => {
      console.log(el)
    })
  })

  FindNewWords().then(result => {
    words.value = result.sort((a: model.Word, b: model.Word) => b.exposures - a.exposures)
  })
}

onMounted(refresh)

</script>

<template>
  <main>
    <v-btn class="btn" @click="refresh">刷新</v-btn>

    <!-- <v-list-item v-for="item in words" :key="item.source" :title="item.source" :subtitle="item.exposures"></v-list-item> -->

    <v-table density="compact">
      <thead>
        <tr>
          <th class="text-left">
            单词
          </th>
          <th class="text-center">
            出现次数
          </th>
          <th class="text-left" style="width: 180px">
            上次出现
          </th>
          <th class="text-left" style="width: 180px">
            最初出现
          </th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="item in words" :key="item.name">
          <td class="text-left">{{ item.source }}</td>
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
