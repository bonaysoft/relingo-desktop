<script lang="ts" setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { GetVocabulary, GetVocabularyList, MasteredWords } from '../../wailsjs/go/relingo/Client';
import { FindNewWords, SubmitVocabulary } from '../../wailsjs/go/service/WordService'
import { model } from '../../wailsjs/go/models';

const vocabularyTab = ref<any>([])
const masteredId = ref()
const glossaries = ref<any>([])
const words = ref<any>([])
const refresh = async () => {
    const masteredWords = await MasteredWords(masteredId.value)
    const { id, type } = vocabularyTab.value
    words.value = (await GetVocabulary(id, type)).map(el => ({ word: el, checked: false, mastered: masteredWords.indexOf(el) == -1 }))
}

const selectedWords = computed(() => {
    console.log(words.value.find((el: any) => el.checked == true));

    return words.value.filter((el: any) => el.checked == true).map((el: any) => el.word)
})

const submitVocabulary = (words: string[]) => {
    SubmitVocabulary(words).then(refresh)
}

onMounted(async () => {
    const ret = await GetVocabularyList()
    masteredId.value = ret.find((el: any) => el.type == 'mastered')?._id
    glossaries.value = ret.filter((el: any) => el.type == 'buildin')

    vocabularyTab.value = glossaries.value[0]
    try {
        refresh()
    } catch (error) {
        console.log(error);
    }
})

</script>

<template>
    <main>
        <v-tabs v-model="vocabularyTab" bg-color="primary" @update:model-value="refresh">
            <v-tab v-for="(v, idx) in glossaries" :key="idx" :value="v">{{ v.name }}</v-tab>
        </v-tabs>

        <v-table density="compact">
            <thead>
                <tr>
                    <th class="text-left">
                        单词
                    </th>
                    <th class="text-right">
                        操作
                    </th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="item in words" :key="item.name">
                    <td class="text-left">
                        <v-checkbox :label="item.word" v-model="item.checked"></v-checkbox>
                    </td>
                    <td class="text-right">
                        <v-btn v-show="!item.mastered" size="small" color="surface-variant ml-3" variant="text"
                            icon="mdi-check-all" @click="submitVocabulary([item.word || ''])"></v-btn>
                    </td>
                </tr>
            </tbody>
        </v-table>

        <div class="ft" v-show="selectedWords && selectedWords.length > 0">
            <v-btn @click="submitVocabulary(selectedWords)">一键掌握</v-btn>
        </div>
    </main>
</template>

<style scoped>
.ft {
    position: fixed;
    bottom: 10px;
    left: 50%;
}
</style>
