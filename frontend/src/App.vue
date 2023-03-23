
<script setup lang="ts">
import Login from "@/components/Login.vue";
import { ref, onMounted } from 'vue';
import { Ready, GetUserInfo } from '@/../wailsjs/go/relingo/Client';

const ready = ref(false)
const uInfo = ref<any>({})
onMounted(async () => {
  ready.value = await Ready()
  if (!ready.value) {
    return
  }

  uInfo.value = await GetUserInfo()
})
</script>

<template>
  <v-app v-if="!ready">
    <Login />
  </v-app>
  <v-app v-else>
    <v-navigation-drawer width="200" permanent>
      <v-list-item :prepend-avatar="uInfo.avatar" :title="uInfo.email" nav></v-list-item>
      <v-divider></v-divider>

      <v-list density="compact" nav>
        <v-list-group>
          <template v-slot:activator="{ props }">
            <v-list-item v-bind="props" prepend-icon="mdi-home-city" title="背单词"></v-list-item>
          </template>
          <v-list-item title="未掌握" value="unmastered" to="home"></v-list-item>
          <v-list-item title="已掌握" value="mastered" to="mastered"></v-list-item>
        </v-list-group>
        <v-list-item prepend-icon="mdi-book" title="单词本" value="words" to="words"></v-list-item>
        <!-- <v-list-item prepend-icon="mdi-cog" title="设置" value="settings" to="settings"></v-list-item> -->
      </v-list>
    </v-navigation-drawer>

    <v-main>
      <v-container>
        <router-view></router-view>
      </v-container>
    </v-main>
  </v-app>
</template>

