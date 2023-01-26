<script lang="ts" setup>
import { ref, reactive, onMounted, computed } from 'vue'
import { GetVocabularyList, GetVocabulary } from '@/../wailsjs/go/relingo/Client';

const words = ref<any>([])
const masteredCount = computed(() => {return words.value.length})

const refresh = () => {
  GetVocabularyList().then(vocabularyList => {
    const mastered = vocabularyList.find(el => el.type == 'mastered')
    words.value = mastered?.words?.map(word => ({ name: word }))
  })

}

onMounted(refresh)

</script>

<template>
  <main>
    已掌握总数{{ masteredCount }}

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
          <td class="text-left">{{ item.name }}</td>
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
