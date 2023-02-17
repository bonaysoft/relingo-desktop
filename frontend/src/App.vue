<template>
  <v-app v-if="!ready">
    <Login />
  </v-app>
  <v-app v-else>
    <v-navigation-drawer permanent>
      <v-list-item :prepend-avatar="uInfo.avatar" :title="uInfo.email" nav>
        <template v-slot:append>
          <v-btn variant="text" icon="mdi-chevron-left"></v-btn>
        </template>
      </v-list-item>

      <v-divider></v-divider>

      <v-list density="compact" nav>
        <v-list-item prepend-icon="mdi-home-city" title="我的生词" value="home" to="home"></v-list-item>
        <v-list-item prepend-icon="mdi-account" title="全部单词" value="account" to="words"></v-list-item>
        <v-list-item prepend-icon="mdi-account-group-outline" title="已掌握单词" value="users" to="mastered"></v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-main>
      <v-container>
        <router-view></router-view>
      </v-container>
    </v-main>
  </v-app>
</template>

<script setup lang="ts">
import Login from "@/components/Login.vue";
import { ref, onMounted } from 'vue';
import { Ready, GetUserInfo } from '@/../wailsjs/go/relingo/Client';

const ready = ref(false)
const uInfo = ref<any>({})
onMounted(async () => {
  ready.value = await Ready()
  if(!ready.value){
    return
  }

  uInfo.value = await GetUserInfo()
})


</script>
