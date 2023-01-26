<script lang="ts" setup>
import { ref, reactive, onMounted } from 'vue'
import { GetVocabulary, GetVocabularyList } from '../../wailsjs/go/relingo/Client';
import { FindNewWords } from '../../wailsjs/go/main/App'
import { model } from '../../wailsjs/go/models';

const vocabularyTab = ref()
const glossaries = ref<any>([])
const refresh = async () => {
    glossaries.value = (await GetVocabularyList()).filter(el => el.type == 'buildin')

    vocabularyTab.value = glossaries.value[0]
    onCurrentTabChange()
}

const words = ref<any>([])
const onCurrentTabChange = async () => {
    const { id, type } = vocabularyTab.value
    words.value = (await GetVocabulary(id, type)).map(el => ({ word: el }))
}

onMounted(refresh)

</script>

<template>
    <main>
        <v-tabs v-model="vocabularyTab" bg-color="primary" @update:model-value="onCurrentTabChange">
            <v-tab v-for="(v, idx) in glossaries" :key="idx" :value="v">{{ v.name }}</v-tab>
        </v-tabs>

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
                    <td class="text-left">{{ item.word }}</td>
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
