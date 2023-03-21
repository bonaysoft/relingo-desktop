<script lang="ts" setup>
import { ref, reactive, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router';
import { FindNewWords, SubmitVocabulary } from '../../wailsjs/go/main/App'
import { Play } from '../../wailsjs/go/youdao/Client'
import { main, model } from '../../wailsjs/go/models';

const dateTabs = ref<any[]>([
  { label: "今日生词", value: 'today' },
  { label: "昨日生词", value: 'yesterday' },
  { label: "本周生词", value: 'weekly' },
  { label: "全部生词", value: 'all' }
])
const total = ref<number>()
const page = ref({ pageNo: 1, pageSize: 6, total: 0 })
const words = ref<main.Word[]>([])
const currentTab = ref<any>('today')
const refresh = () => {
  const { pageNo, pageSize } = page.value
  FindNewWords(currentTab.value, pageNo, pageSize).then((result: main.ListResult) => {
    words.value = result.items
    total.value = result.total
    page.value.pageNo = pageNo
    page.value.total = Math.round(result.total / pageSize)
  })
}

const submitVocabulary = (words: string[]) => {
  SubmitVocabulary(words).then(refresh)
}
const play = async (word: string) => {
  const audio = new Audio(await Play(word))
  audio.play()
}

onMounted(refresh)
watch(currentTab, refresh)
</script>

<template>
  <main>
    <div>
      <v-tabs v-model="currentTab" bg-color="primary">
        <v-tab v-for="(v, idx) in dateTabs" :key="idx" :value="v.value">{{ v.label }}</v-tab>
        <span style="position: absolute;right: 10px;top: 13px"> {{ total }}</span>
      </v-tabs>
    </div>
    <div>
      <v-container fluid>
        <v-row dense>
          <v-col v-for="word in words" :key="word.id" :cols="6">
            <v-card>
              <v-card-item>
                <v-card-title>
                  <span>{{ word.name }}</span>
                  <span style="font-size: small; margin-left: 10px;">
                    <span class="my-2">{{ word.raw_object.phonetic }}</span>
                    <v-icon icon="mdi-volume-high" size="small" @click="play(word.name)"></v-icon>

                    <span style="float: right;">
                      <v-rating :model-value="word.raw_object.wordFrequency" color="amber" density="compact"
                        size="x-small" half-increments readonly></v-rating>
                    </span>
                  </span>
                </v-card-title>
                <v-card-subtitle>
                  <v-chip class="mr-2" size="x-small" v-for="tag in word.engra_data?.tags"> {{ tag }} </v-chip>
                </v-card-subtitle>
              </v-card-item>

              <v-card-text>
                <div class="my-1 text-subtitle-1">
                  {{ word.raw_object?.translations.map(el => el.target).join('；') }}
                </div>

                <div>{{ word.engra_data?.mnemonic }}</div>
              </v-card-text>

              <v-card-actions>
                <div class="counter">
                  <span>累计出现 {{ word.exposures }} 次</span>
                </div>
                <v-spacer></v-spacer>
                <v-btn size="small" color="surface-variant" variant="text" icon="mdi-check-all"
                  @click="submitVocabulary([word.name || ''])"></v-btn>
              </v-card-actions>
            </v-card>
          </v-col>
        </v-row>
      </v-container>
    </div>
    <div class="text-center my-5">
      <v-pagination rounded="circle" v-model="page.pageNo" :length="page.total" :total-visible="6"
        @update:model-value="refresh"></v-pagination>
    </div>
  </main>
</template>

<style scoped>
.counter {
  font-size: 12px;
  color: #c9c9c9;
  font-weight: 500;
  margin-left: 10px;
}
</style>
