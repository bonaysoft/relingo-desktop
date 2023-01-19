<script lang="ts" setup>
import { ref, reactive, onMounted } from 'vue'
import { Greet } from '../../wailsjs/go/main/App'

const data = reactive({
  name: "",
  resultText: "Please enter your name below 👇",
})

const words = ref<any>([])
const refresh = () => {
  Greet(data.name).then(result => {
    words.value = result.sort((a: any, b: any) => b.Exposures - a.Exposures)
  })
}

onMounted(refresh)

</script>

<template>
  <main>
    <button class="btn" @click="refresh">刷新</button>

    <div v-for="w in words">
      {{ w.Source }} {{ w.Exposures }}
    </div>
  </main>
</template>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
