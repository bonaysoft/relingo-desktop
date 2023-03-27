<script lang="ts" setup>
import { ref, reactive, onMounted } from 'vue'
import { GetVocabulary, GetVocabularyList } from '../../wailsjs/go/relingo/Client';
import { FindNewWords } from '../../wailsjs/go/service/WordService'
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
                </tr>
            </thead>
            <tbody>
                <tr v-for="item in words" :key="item.name">
                    <td class="text-left">{{ item.word }}</td>
                </tr>
            </tbody>
        </v-table>
    </main>
</template>

<style scoped>

</style>
