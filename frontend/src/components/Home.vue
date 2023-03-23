<script lang="ts" setup>
import { ref, reactive, onMounted, watch } from 'vue'
import { useRoute } from 'vue-router';
import { FindNewWords, GetRootByName, SubmitVocabulary } from '../../wailsjs/go/main/App'
import { Play } from '../../wailsjs/go/youdao/Client'
import { main, model } from '../../wailsjs/go/models';


import "jsmind/style/jsmind.css";
import * as jsMind from "jsmind";

const dateTabs = ref<any[]>([
  { label: "今日生词", value: 'today' },
  { label: "昨日生词", value: 'yesterday' },
  { label: "本周生词", value: 'weekly' },
  { label: "全部生词", value: 'all' }
])
const drawer = ref<boolean>(false)
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

const openDrawer = async (name: string) => {
  const node = document.createElement('div')
  node.setAttribute('id', 'jsmind_container')
  node.setAttribute('style', 'height: 100%')
  document.getElementById("jsmind_container")?.remove()
  document.getElementById("jsdrawer")?.appendChild(node)

  const fillId = (item: any) => {
    if (item.children) {
      item.children = item.children.map(fillId)
    }
    // item.direction = 'left'
    // item.expanded = false
    if (item.name == name) {
      //   item.direction = 'right'
      //   item.expanded = true
      item['background-color'] = 'orange'
    }

    item.id = item.name
    item.topic = item.name
    if (item.meaning) {
      item.topic += "<br /><br />" + item.meaning
    }
    return item
  }
  try {
    const root = (await GetRootByName(name))
    let abf = root.children.map((el: any) => fillId(el))
    root.id = '-' + root.name + '-'
    root.topic = '-' + root.name + '-'
    if (root.meaning) {
      root.topic += '<br /><div style="text-align:center; font-size: 16px; margin-top:5px">' + root.meaning + '</div>'
    }

    root.isroot = true
    root.children = abf
    console.log(root);
    drawer.value = true
    if (!root) {
      return
    }

    var mind = {
      meta: {
        name: 'demo',
        author: 'hizzgdev@163.com',
        version: '0.2',
      },
      format: 'node_tree',
      data: root,
    };
    var options = {
      container: "jsmind_container",
      theme: "greensea",
      editable: false,
      support_html: true,
      view: {
        draggable: true,
        hide_scrollbars_when_draggable: true
      },
    };

    var jm = jsMind.show(options, mind);
  } catch (error) {

  }

}

onMounted(refresh)
watch(currentTab, () => {
  page.value.pageNo = 1;
  refresh()
})
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
                  <b>{{ word.name }}</b>
                  <span style="font-size: small; margin-left: 10px;">
                    <span class="my-2">/ {{ word.raw_object.phonetic[0] }} /</span>
                    <v-icon class="ml-1" icon="mdi-volume-high" size="small" @click="play(word.name)"></v-icon>

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

              <v-card-text style="height: 100px;">
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
                <v-btn size="small" icon="mdi-family-tree" @click="openDrawer(word.name)"></v-btn>
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

    <v-navigation-drawer v-model="drawer" location="bottom" width="600" temporary>
      <div id="jsdrawer">
        <div id="jsmind_container"></div>
      </div>
    </v-navigation-drawer>
  </main>
</template>

<style scoped>
.counter {
  font-size: 12px;
  color: #c9c9c9;
  font-weight: 500;
  margin-left: 10px;
}

#jsdrawer {
  height: 100%;
}
</style>
